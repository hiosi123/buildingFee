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
	structType := reflect.TypeOf(s)

	if structType.Kind() != reflect.Struct {
		return "", errors.New("input must be struct")
	}

	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		values = append(values, field.Tag.Get("db"))
	}

	// Join the values into a comma-separated string
	valuesStr := strings.Join(values, ", ")

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

	for i := 1; i < numFields; i++ {
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
		quotedColumns[i] = fmt.Sprintf("`%s`", col)
	}
	return strings.Join(quotedColumns, ", ")
}

func placeholders(n int) string {
	return strings.Repeat("?, ", n-1) + "?"
}

func UpdateQueryStr(tableName string, s any, condition string) (string, []any, error) {
	value := reflect.ValueOf(s)
	typ := reflect.TypeOf(s)
	var conditionValue any

	if value.Kind() != reflect.Struct {
		return "", nil, errors.New("input must be struct")
	}

	numFields := typ.NumField()
	columns := make([]string, 0, numFields)
	values := make([]any, 0, numFields)

	for i := 0; i < numFields; i++ {
		field := typ.Field(i)
		column := field.Tag.Get("db")
		fieldValue := value.Field(i).Interface()

		if column == condition {
			conditionValue = fieldValue
			continue
		}

		// Check if the field has a zero value
		if reflect.DeepEqual(fieldValue, reflect.Zero(field.Type).Interface()) {
			// Skip this field, it has a nil value
			continue
		}

		columns = append(columns, column)
		values = append(values, fieldValue)

	}

	setExpressions := make([]string, len(columns))
	for i, column := range columns {
		setExpressions[i] = fmt.Sprintf("%s = ?", column)
	}

	query := fmt.Sprintf("UPDATE %s SET %s WHERE %s = ?", tableName, strings.Join(setExpressions, ", "), condition)

	return query, append(values, conditionValue), nil
}
