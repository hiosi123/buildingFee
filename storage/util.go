package storage

import (
	"fmt"
	"reflect"
	"strings"
)

func GetPureFieldStr(myInstance struct{}) string {
	// Use reflection to iterate over struct fields and build a comma-separated string
	var values []string
	structValue := reflect.ValueOf(myInstance)
	for i := 0; i < structValue.NumField(); i++ {
		field := structValue.Field(i)

		// Convert the field value to a string and add it to the values slice
		values = append(values, fmt.Sprintf("'%v'", field.Interface()))
	}

	// Join the values into a comma-separated string
	valuesStr := strings.Join(values, ",")

	return valuesStr
}
