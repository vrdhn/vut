package tools

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Json is a lightweight wrapper for parsed JSON data.
type Json struct {
	data interface{}
}

// ParseJson parses a JSON string and returns a *Json wrapper.
func parseJson(in string) (Json, error) {
	var v interface{}
	err := json.Unmarshal([]byte(in), &v)
	if err != nil {
		return Json{nil}, fmt.Errorf("failed to parse JSON: %w", err)
	}
	return Json{data: v}, nil
}

// GetJson navigates through the JSON using keys or indexes.
// Returns a new *Json and error if any lookup fails.
func (j Json) GetJson(args ...any) (*Json, error) {
	cur := j.data
	for _, arg := range args {
		switch a := arg.(type) {
		case string:
			m, ok := cur.(map[string]interface{})
			if !ok {
				return nil, errors.New("expected object for key lookup")
			}
			val, exists := m[a]
			if !exists {
				return nil, fmt.Errorf("key not found: %s", a)
			}
			cur = val
		case int:
			arr, ok := cur.([]interface{})
			if !ok {
				return nil, errors.New("expected array for index lookup")
			}
			if a < 0 || a >= len(arr) {
				return nil, fmt.Errorf("index out of range: %d", a)
			}
			cur = arr[a]
		default:
			return nil, fmt.Errorf("unsupported argument type: %T", a)
		}
	}
	return &Json{data: cur}, nil
}

// GetString retrieves a string value from JSON.
func (j Json) GetString(args ...any) (string, error) {
	val, err := j.GetJson(args...)
	if err != nil {
		return "", err
	}
	s, ok := val.data.(string)
	if !ok {
		return "", errors.New("value is not a string")
	}
	return s, nil
}

// GetInt retrieves an integer (from float64 JSON numbers).
func (j Json) GetInt(args ...any) (int, error) {
	val, err := j.GetJson(args...)
	if err != nil {
		return 0, err
	}
	switch n := val.data.(type) {
	case float64:
		return int(n), nil
	case int:
		return n, nil
	default:
		return 0, errors.New("value is not a number")
	}
}

// GetFloat retrieves a float64 value.
func (j *Json) GetFloat(args ...any) (float64, error) {
	val, err := j.GetJson(args...)
	if err != nil {
		return 0, err
	}
	f, ok := val.data.(float64)
	if !ok {
		return 0, errors.New("value is not a float")
	}
	return f, nil
}
