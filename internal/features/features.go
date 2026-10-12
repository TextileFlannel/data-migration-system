// Package features turns a field profile into a numeric feature vector.
//
// Это ЕДИНСТВЕННАЯ реализация признаков — используется и при сборке обучающего
// набора (dms build-dataset), и на инференсе (dms plan). Поэтому модель всегда
// видит одинаковые числа (нет train/serve skew).
package features

import "intelligent-data-migration-system/internal/schema"

// Names lists the feature columns in a fixed order. Этот же порядок — в train.csv.
var Names = []string{
	"depth",
	"presence_ratio",
	"is_array",
	"avg_array_len",
	"max_array_len",
	"frac_string",
	"frac_int",
	"frac_long",
	"frac_double",
	"frac_bool",
	"frac_date",
	"frac_objectid",
	"frac_object",
	"frac_array",
	"frac_null",
	"type_entropy",
	"distinct_ratio",
	"avg_len",
	"max_len",
	"nested_fields",
}

// Extract computes the feature vector for a single field profile.
// Длина результата совпадает с len(Names) и выровнена по порядку Names.
func Extract(fp schema.FieldProfile) []float64 {
	// TODO: посчитать признаки по fp (доли типов, энтропия, длины и т.д.).
	return make([]float64, len(Names))
}
