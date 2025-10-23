package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"reflect"
	"strconv"
	"strings"
)

// ParseCSVIntoStructs parses CSV rows into a slice of structs based on `csv:"index"` tags.
func ParseCSVIntoStructs[T any](s string) ([]T, error) {
	reader := csv.NewReader(strings.NewReader(s))
	var results []T
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("error reading CSV: %w", err)
		}

		// Create new struct instance
		var item T
		v := reflect.ValueOf(&item).Elem()
		t := v.Type()

		// Iterate over struct fields
		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			tag := t.Field(i).Tag.Get("csv")
			if tag == "" {
				continue
			}

			// Parse column index from tag
			idx, err := strconv.Atoi(tag)
			if err != nil || idx < 0 || idx >= len(record) {
				continue
			}

			val := record[idx]

			// Convert value into field type
			switch field.Kind() {
			case reflect.String:
				field.SetString(val)
			case reflect.Int, reflect.Int64, reflect.Int32:
				if iv, err := strconv.ParseInt(val, 10, 64); err == nil {
					field.SetInt(iv)
				}
			case reflect.Float64, reflect.Float32:
				if fv, err := strconv.ParseFloat(val, 64); err == nil {
					field.SetFloat(fv)
				}
			}
		}

		results = append(results, item)
	}

	return results, nil
}
