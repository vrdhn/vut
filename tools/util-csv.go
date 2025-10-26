package tools

import (
	"encoding/csv"
	"io"
	"reflect"
	"strconv"
	"strings"
)

func csvParser[T any](seperator rune) func(string) ([]T, error) {
	return func(in string) ([]T, error) {
		return parseCSV[T](in, seperator)
	}
}

// ParseCSVToStruct parses CSV data (no headers) into a slice of structs.
// It uses reflection to map fields in order.
func parseCSV[T any](data string, separator rune) ([]T, error) {
	reader := csv.NewReader(strings.NewReader(data))
	reader.Comma = separator
	reader.TrimLeadingSpace = true

	var results []T
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		var item T
		v := reflect.ValueOf(&item).Elem()
		//t := v.Type()

		for i := 0; i < len(record) && i < v.NumField(); i++ {
			field := v.Field(i)
			if !field.CanSet() {
				continue
			}

			val := record[i]
			switch field.Kind() {
			case reflect.String:
				field.SetString(val)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				if n, err := strconv.ParseInt(val, 10, 64); err == nil {
					field.SetInt(n)
				}
			case reflect.Float32, reflect.Float64:
				if f, err := strconv.ParseFloat(val, 64); err == nil {
					field.SetFloat(f)
				}
			case reflect.Bool:
				if b, err := strconv.ParseBool(val); err == nil {
					field.SetBool(b)
				}
				// You could extend this to handle time.Time, enums, etc.
			}
		}

		results = append(results, item)
	}

	return results, nil
}
