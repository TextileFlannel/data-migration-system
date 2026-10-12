// Package schema defines the core types shared across the migration pipeline:
// the MongoDB-side profile, the per-field decision classes, and the resulting
// relational (PostgreSQL) mapping plan.
package schema

// BSONType is a MongoDB/BSON value type observed while profiling.
type BSONType string

const (
	BSONString   BSONType = "string"
	BSONInt      BSONType = "int"
	BSONLong     BSONType = "long"
	BSONDouble   BSONType = "double"
	BSONBool     BSONType = "bool"
	BSONDate     BSONType = "date"
	BSONObjectID BSONType = "objectId"
	BSONObject   BSONType = "object"
	BSONArray    BSONType = "array"
	BSONNull     BSONType = "null"
)

// FieldProfile is the aggregated statistics for one field path in a collection.
// Его строит profiler, а потребляет features.
type FieldProfile struct {
	Path          string             `json:"path"`           // напр. "address.city"
	Depth         int                `json:"depth"`          // глубина вложенности
	PresenceRatio float64            `json:"presence_ratio"` // доля документов, где поле есть
	IsArray       bool               `json:"is_array"`
	AvgArrayLen   float64            `json:"avg_array_len"`
	MaxArrayLen   int                `json:"max_array_len"`
	TypeCounts    map[BSONType]int64 `json:"type_counts"` // сколько раз встречался каждый тип
	DistinctRatio float64            `json:"distinct_ratio"`
	AvgLen        float64            `json:"avg_len"` // средняя длина строки
	MaxLen        int                `json:"max_len"`
	NumericMin    float64            `json:"numeric_min"`
	NumericMax    float64            `json:"numeric_max"`
	NestedFields  int                `json:"nested_fields"` // число подполей для объектов
}

// CollectionProfile is the profile of one MongoDB collection.
type CollectionProfile struct {
	Name     string         `json:"name"`
	DocCount int64          `json:"doc_count"`
	Fields   []FieldProfile `json:"fields"`
}

// Profile is the profile of a whole MongoDB database.
type Profile struct {
	Database    string              `json:"database"`
	Collections []CollectionProfile `json:"collections"`
}

// Decision is the structural mapping class predicted for a field (выход модели, задача A).
type Decision string

const (
	DecisionScalarColumn Decision = "scalar_column"
	DecisionNewTable1to1 Decision = "new_table_1to1"
	DecisionNewTable1toN Decision = "new_table_1toN"
	DecisionArrayColumn  Decision = "array_column"
	DecisionFlatten      Decision = "flatten"
	DecisionJSONB        Decision = "jsonb"
	DecisionForeignKey   Decision = "foreign_key"
	DecisionSkip         Decision = "skip"
)

// PostgresType is the target column type for a scalar column (выход модели, задача B).
type PostgresType string

const (
	PGSmallint    PostgresType = "smallint"
	PGInteger     PostgresType = "integer"
	PGBigint      PostgresType = "bigint"
	PGNumeric     PostgresType = "numeric"
	PGBoolean     PostgresType = "boolean"
	PGText        PostgresType = "text"
	PGTimestamptz PostgresType = "timestamptz"
	PGUUID        PostgresType = "uuid"
	PGJSONB       PostgresType = "jsonb"
)

// Column is one column in a planned relational table.
type Column struct {
	Name       string       `json:"name"`
	Type       PostgresType `json:"type"`
	Nullable   bool         `json:"nullable"`
	SourcePath string       `json:"source_path"` // откуда поле в документе
}

// Table is one planned relational table.
type Table struct {
	Name       string   `json:"name"`
	Columns    []Column `json:"columns"`
	PrimaryKey []string `json:"primary_key,omitempty"`
}

// ForeignKey is a planned relationship between two tables.
type ForeignKey struct {
	Table      string   `json:"table"`
	Columns    []string `json:"columns"`
	RefTable   string   `json:"ref_table"`
	RefColumns []string `json:"ref_columns"`
}

// MappingPlan is the full document->relational plan: его строит planner,
// просматривает/правит пользователь, выполняет migrator.
type MappingPlan struct {
	SourceDatabase string       `json:"source_database"`
	Tables         []Table      `json:"tables"`
	ForeignKeys    []ForeignKey `json:"foreign_keys"`
}
