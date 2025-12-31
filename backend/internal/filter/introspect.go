package filter

import (
	"database/sql"
	"reflect"
	"strings"
	"time"
)

func BuildSchemaFromStruct(model any) map[string]FieldSchema {
	t := reflect.TypeOf(model)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	schema := map[string]FieldSchema{}

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)

		jsonTag := strings.Split(f.Tag.Get("json"), ",")[0]
		if jsonTag == "" || jsonTag == "-" {
			continue
		}

		ui := f.Tag.Get("ui")
		gormTag := f.Tag.Get("gorm")

		column := jsonTag
		for _, part := range strings.Split(gormTag, ";") {
			if strings.HasPrefix(part, "column:") {
				column = strings.TrimPrefix(part, "column:")
			}
		}

		fieldType, ops := inferTypeAndOps(f.Type)

		schema[jsonTag] = FieldSchema{
			JSONKey:    jsonTag,
			DBColumn:   column,
			Type:       fieldType,
			Operators:  ops,
			Sortable:   strings.Contains(ui, "sortable"),
			Filterable: strings.Contains(ui, "filterable"),
			Editable:   strings.Contains(ui, "editable"),
			Visible:    strings.Contains(ui, "visible"),
			Selection:  extractSelection(ui),
			TimeFormat: f.Tag.Get("time_format"),
		}
	}

	return schema
}

func inferTypeAndOps(t reflect.Type) (FieldType, []Operator) {
	// Check for time.Time
	switch t {
	case reflect.TypeOf(time.Time{}):
		return DateTime, []Operator{Gt, Gte, Lt, Lte, Between, IsNull}
	}

	// Check for sql.Null types
	switch t {
	case reflect.TypeOf(sql.NullString{}):
		return String, []Operator{Eq, Ne, Contains, StartsWith, EndsWith, In, IsNull}
	case reflect.TypeOf(sql.NullInt64{}), reflect.TypeOf(sql.NullInt32{}), reflect.TypeOf(sql.NullInt16{}):
		return Number, []Operator{Eq, Ne, Gt, Gte, Lt, Lte, Between, In, IsNull}
	case reflect.TypeOf(sql.NullFloat64{}):
		return Number, []Operator{Eq, Ne, Gt, Gte, Lt, Lte, Between, In, IsNull}
	case reflect.TypeOf(sql.NullBool{}):
		return Boolean, []Operator{Eq, IsNull}
	case reflect.TypeOf(sql.NullTime{}):
		return DateTime, []Operator{Gt, Gte, Lt, Lte, Between, IsNull}
	case reflect.TypeOf(sql.NullByte{}):
		return Number, []Operator{Eq, Ne, Gt, Gte, Lt, Lte, Between, In, IsNull}
	}

	// Check for gorm datatypes by type name string matching
	typeName := t.String()
	switch {
	// Date and Time types
	case strings.Contains(typeName, "datatypes.Date"):
		return DateTime, []Operator{Gt, Gte, Lt, Lte, Between, IsNull}
	case strings.Contains(typeName, "datatypes.Time"):
		return DateTime, []Operator{Gt, Gte, Lt, Lte, Between, IsNull}

	// Null types from datatypes package
	case strings.Contains(typeName, "datatypes.NullString"):
		return String, []Operator{Eq, Ne, Contains, StartsWith, EndsWith, In, IsNull}
	case strings.Contains(typeName, "datatypes.NullInt64"),
		strings.Contains(typeName, "datatypes.NullInt32"),
		strings.Contains(typeName, "datatypes.NullInt16"):
		return Number, []Operator{Eq, Ne, Gt, Gte, Lt, Lte, Between, In, IsNull}
	case strings.Contains(typeName, "datatypes.NullFloat64"):
		return Number, []Operator{Eq, Ne, Gt, Gte, Lt, Lte, Between, In, IsNull}
	case strings.Contains(typeName, "datatypes.NullBool"):
		return Boolean, []Operator{Eq, IsNull}
	case strings.Contains(typeName, "datatypes.NullTime"):
		return DateTime, []Operator{Gt, Gte, Lt, Lte, Between, IsNull}
	case strings.Contains(typeName, "datatypes.NullByte"):
		return Number, []Operator{Eq, Ne, Gt, Gte, Lt, Lte, Between, In, IsNull}
	case strings.Contains(typeName, "datatypes.Null["):
		// Generic Null[T] type - treat as string for now
		return String, []Operator{Eq, Ne, Contains, IsNull}

	// JSON types
	case strings.Contains(typeName, "datatypes.JSON"):
		return String, []Operator{Contains, IsNull}
	case strings.Contains(typeName, "datatypes.JSONMap"):
		return String, []Operator{Contains, IsNull}
	case strings.Contains(typeName, "datatypes.JSONType"):
		return String, []Operator{Contains, IsNull}
	case strings.Contains(typeName, "datatypes.JSONSlice"):
		return String, []Operator{Contains, IsNull}

	// UUID types
	case strings.Contains(typeName, "datatypes.UUID"):
		return String, []Operator{Eq, Ne, In, IsNull}
	case strings.Contains(typeName, "datatypes.BinUUID"):
		return String, []Operator{Eq, Ne, In, IsNull}

	// URL type
	case strings.Contains(typeName, "datatypes.URL"):
		return String, []Operator{Eq, Ne, Contains, StartsWith, EndsWith, IsNull}

	// GORM schema types
	case strings.Contains(typeName, "schema.DeletedAt"):
		return DateTime, []Operator{Gt, Gte, Lt, Lte, Between, IsNull}
	}

	// Check for basic kinds
	switch t.Kind() {
	case reflect.String:
		return String, []Operator{Eq, Ne, Contains, StartsWith, EndsWith, In, IsNull}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return Number, []Operator{Eq, Ne, Gt, Gte, Lt, Lte, Between, In, IsNull}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return Number, []Operator{Eq, Ne, Gt, Gte, Lt, Lte, Between, In, IsNull}
	case reflect.Float32, reflect.Float64:
		return Number, []Operator{Eq, Ne, Gt, Gte, Lt, Lte, Between, In, IsNull}
	case reflect.Bool:
		return Boolean, []Operator{Eq, IsNull}
	}

	return String, []Operator{Eq}
}

func extractSelection(ui string) string {
	for _, part := range strings.Split(ui, ";") {
		if strings.HasPrefix(part, "selection:") {
			return strings.TrimPrefix(part, "selection:")
		}
	}
	return ""
}
