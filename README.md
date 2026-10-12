# Intelligent Data Migration System

Интеллектуальная миграция данных **MongoDB → PostgreSQL**: обученная модель решает,
как разложить бесхемные документы по реляционным таблицам.

Полная концепция и архитектура — в [DESIGN.md](DESIGN.md).

## Раскладка

- `cmd/dms/` — точка входа CLI (только разбор команд/флагов).
- `internal/` — вся логика (профилирование, признаки, модель, план, миграция, проверка).
- `ml/` — Python: генерация синтетики и обучение модели.

## Сборка и запуск

```sh
go build ./...
go run ./cmd/dms            # список команд
```

## Команды (пока заглушки)

| Команда             | Что делает                                     |
|---------------------|------------------------------------------------|
| `dms build-dataset` | собрать `train.csv` из документов + меток (Go) |
| `dms profile`       | профилировать MongoDB → `profile.json`         |
| `dms plan`          | `profile.json` + модель → `plan.json`          |
| `dms migrate`       | создать схему в Postgres и перенести данные    |
| `dms verify`        | проверить целостность после миграции           |

## ML (Python)

```sh
pip install -r ml/requirements.txt
python ml/generate.py      # синтетика: документы + метки
go run ./cmd/dms build-dataset   # → train.csv (признаки считает Go)
python ml/train.py         # → models/model.txt
python ml/evaluate.py      # метрики, сравнение с baseline
```
