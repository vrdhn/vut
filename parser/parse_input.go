package parser

//
//import (
//	"errors"
//	"fmt"
//	"strconv"
//	"strings"
//)
//
//// ParseValue parses input string which can be either an integer or a percentage string (e.g. "75%").
//// If input is integer, returns that integer.
//// If input ends with '%', returns percentage of max.
//// Returns error if input is invalid.
//func ParseValue(max int, input string) (int, error) {
//	input = strings.TrimSpace(input)
//
//	if strings.HasSuffix(input, "%") {
//		pctStr := strings.TrimSuffix(input, "%")
//		pct, err := strconv.ParseFloat(pctStr, 64)
//		if err != nil {
//			return 0, errors.New("invalid percentage value")
//		}
//		if pct < 0 || pct > 100 {
//			return 0, errors.New("percentage out of range 0-100")
//		}
//		return int(float64(max) * pct / 100.0), nil
//	}
//
//	// Otherwise, try parsing as integer
//	val, err := strconv.Atoi(input)
//	if err != nil {
//		return 0, errors.New("invalid integer value")
//	}
//	if val < 0 || val > max {
//		return 0, errors.New("value out of range")
//	}
//
//	return val, nil
//}
//
