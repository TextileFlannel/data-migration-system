// Command dms is the CLI for the intelligent data migration system
// (MongoDB -> PostgreSQL).
//
// Это только точка входа: разбор команд и флагов + вызовы в internal/*.
// Вся логика живёт в пакетах internal/* (profiler, features, dataset, model,
// planner, migrator, verify). Пока они — заглушки.
//
// Подкоманды:
//
//	build-dataset  собрать train.csv из документов + меток (dev)
//	profile        профилировать коллекции MongoDB -> profile.json
//	plan           profile.json + модель -> план миграции
//	migrate        создать схему в Postgres и перенести данные
//	verify         проверить целостность после миграции
package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"

	"intelligent-data-migration-system/internal/config"
	"intelligent-data-migration-system/internal/dataset"
	"intelligent-data-migration-system/internal/migrator"
	"intelligent-data-migration-system/internal/model"
	"intelligent-data-migration-system/internal/planner"
	"intelligent-data-migration-system/internal/profiler"
	"intelligent-data-migration-system/internal/schema"
	"intelligent-data-migration-system/internal/verify"
)

func main() {
	err := run(os.Args[1:])
	switch {
	case err == nil:
		return
	case errors.Is(err, flag.ErrHelp):
		return
	default:
		fmt.Fprintln(os.Stderr, "dms: "+err.Error())
		os.Exit(1)
	}
}

// command is one CLI subcommand.
type command struct {
	name    string
	summary string
	run     func(ctx context.Context, args []string) error
}

func commands() []command {
	return []command{
		{"build-dataset", "собрать train.csv из документов + меток (dev)", cmdBuildDataset},
		{"profile", "профилировать коллекции MongoDB -> profile.json", cmdProfile},
		{"plan", "profile.json + модель -> план миграции", cmdPlan},
		{"migrate", "создать схему в Postgres и перенести данные", cmdMigrate},
		{"verify", "проверить целостность после миграции", cmdVerify},
	}
}

func run(args []string) error {
	if len(args) == 0 {
		usage()
		return nil
	}
	switch args[0] {
	case "-h", "--help", "help":
		usage()
		return nil
	}
	for _, c := range commands() {
		if c.name == args[0] {
			return c.run(context.Background(), args[1:])
		}
	}
	usage()
	return fmt.Errorf("неизвестная команда %q", args[0])
}

func usage() {
	fmt.Println("dms — интеллектуальная миграция MongoDB -> PostgreSQL")
	fmt.Println()
	fmt.Println("Использование:")
	fmt.Println("  dms <команда> [флаги]")
	fmt.Println()
	fmt.Println("Команды:")
	for _, c := range commands() {
		fmt.Printf("  %-14s %s\n", c.name, c.summary)
	}
}

// --- Обработчики команд: только разбор флагов и вызов internal/* ---

// cmdBuildDataset: документы + метки -> признаки (Go) -> train.csv.
func cmdBuildDataset(ctx context.Context, args []string) error {
	fs := flag.NewFlagSet("build-dataset", flag.ContinueOnError)
	docs := fs.String("docs", "testdata/docs.jsonl", "JSONL с документами от генератора")
	labels := fs.String("labels", "testdata/labels.json", "файл меток (поле -> решение)")
	out := fs.String("out", "train.csv", "обучающий CSV на выходе")
	if err := fs.Parse(args); err != nil {
		return err
	}
	return dataset.Build(ctx, dataset.Options{DocsPath: *docs, LabelsPath: *labels, OutCSV: *out})
}

// cmdProfile: профилирование MongoDB -> profile.json.
func cmdProfile(ctx context.Context, args []string) error {
	fs := flag.NewFlagSet("profile", flag.ContinueOnError)
	mongoURI := fs.String("mongo", "mongodb://localhost:27017", "MongoDB connection URI")
	db := fs.String("db", "", "имя базы MongoDB")
	out := fs.String("out", "profile.json", "файл профиля на выходе")
	if err := fs.Parse(args); err != nil {
		return err
	}
	prof, err := profiler.New(config.Config{MongoURI: *mongoURI, MongoDB: *db}).Profile(ctx)
	if err != nil {
		return err
	}
	if err := writeJSON(*out, prof); err != nil {
		return err
	}
	fmt.Printf("профиль сохранён: %s\n", *out)
	return nil
}

// cmdPlan: profile.json + модель -> plan.json.
func cmdPlan(ctx context.Context, args []string) error {
	fs := flag.NewFlagSet("plan", flag.ContinueOnError)
	in := fs.String("in", "profile.json", "входной профиль")
	modelPath := fs.String("model", "models/model.txt", "файл обученной модели")
	out := fs.String("out", "plan.json", "план миграции на выходе")
	if err := fs.Parse(args); err != nil {
		return err
	}
	var prof schema.Profile
	if err := readJSON(*in, &prof); err != nil {
		return err
	}
	m, err := model.Load(*modelPath)
	if err != nil {
		return err
	}
	plan, err := planner.New(m).Plan(&prof)
	if err != nil {
		return err
	}
	return writeJSON(*out, plan)
}

// cmdMigrate: выполнить plan.json (создать схему в Postgres и перенести данные).
func cmdMigrate(ctx context.Context, args []string) error {
	fs := flag.NewFlagSet("migrate", flag.ContinueOnError)
	planPath := fs.String("plan", "plan.json", "план миграции")
	mongoURI := fs.String("mongo", "mongodb://localhost:27017", "MongoDB connection URI")
	db := fs.String("db", "", "имя базы MongoDB")
	pg := fs.String("postgres", "postgres://localhost:5432/postgres", "PostgreSQL URL")
	if err := fs.Parse(args); err != nil {
		return err
	}
	var plan schema.MappingPlan
	if err := readJSON(*planPath, &plan); err != nil {
		return err
	}
	cfg := config.Config{MongoURI: *mongoURI, MongoDB: *db, PostgresURL: *pg}
	return migrator.New(cfg).Migrate(ctx, &plan)
}

// cmdVerify: проверить целостность после миграции.
func cmdVerify(ctx context.Context, args []string) error {
	fs := flag.NewFlagSet("verify", flag.ContinueOnError)
	planPath := fs.String("plan", "plan.json", "план миграции")
	mongoURI := fs.String("mongo", "mongodb://localhost:27017", "MongoDB connection URI")
	db := fs.String("db", "", "имя базы MongoDB")
	pg := fs.String("postgres", "postgres://localhost:5432/postgres", "PostgreSQL URL")
	if err := fs.Parse(args); err != nil {
		return err
	}
	var plan schema.MappingPlan
	if err := readJSON(*planPath, &plan); err != nil {
		return err
	}
	cfg := config.Config{MongoURI: *mongoURI, MongoDB: *db, PostgresURL: *pg}
	return verify.New(cfg).Verify(ctx, &plan)
}

// --- Вспомогательное чтение/запись JSON ---

func readJSON(path string, v any) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return json.NewDecoder(f).Decode(v)
}

func writeJSON(path string, v any) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	return enc.Encode(v)
}
