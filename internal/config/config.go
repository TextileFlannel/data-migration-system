// Package config holds runtime configuration (connection strings, paths).
package config

// Config carries connection settings for the CLI commands.
type Config struct {
	MongoURI    string // MongoDB connection URI
	MongoDB     string // MongoDB database name
	PostgresURL string // PostgreSQL connection URL
}
