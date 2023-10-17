package storage

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func GetPureFieldStr(s any) (string, error) {
	// Use reflection to iterate over struct fields and build a comma-separated string
	var values []string
	structValue := reflect.ValueOf(s)

	if structValue.Kind() != reflect.Struct {
		return "", errors.New("input must be struct")
	}

	for i := 0; i < structValue.NumField(); i++ {
		field := structValue.Field(i)

		// Convert the field value to a string and add it to the values slice
		values = append(values, fmt.Sprintf("'%v'", field.Interface()))
	}

	// Join the values into a comma-separated string
	valuesStr := strings.Join(values, ",")

	return valuesStr, nil
}

func InsertQueryStr(s any) (string, []any, error) {
	value := reflect.ValueOf(s)
	typ := reflect.TypeOf(s)

	if value.Kind() != reflect.Struct {
		return "", nil, errors.New("input must be struct")
	}

	numFields := typ.NumField()
	columns := make([]string, 0, numFields)
	values := make([]any, 0, numFields)

	for i := 0; i < numFields; i++ {
		field := typ.Field(i)
		columns = append(columns, field.Tag.Get("db"))
		values = append(values, value.Field(i).Interface())
	}

	query := fmt.Sprintf("(%s) values (%s)",
		quoteColumns(columns),
		placeholders(len(columns)),
	)

	return query, values, nil
}

func quoteColumns(columns []string) string {
	quotedColumns := make([]string, len(columns))
	for i, col := range columns {
		quotedColumns[i] = fmt.Sprintf(`"%s"`, col)
	}
	return strings.Join(quotedColumns, ", ")
}

func placeholders(n int) string {
	return strings.Repeat("?, ", n-1) + "?"
}
