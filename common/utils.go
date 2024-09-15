package common

import (
	"fmt"
	"strconv"
	"strings"
)

// IntListContains return true if list of int contains int
func IntListContains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// StringListContains return true if list of string contains string
func StringListContains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// StringToInt convert string to int, or return 0 on error
func StringToInt(strVal string) int {
	intVal, err := strconv.Atoi(strVal)
	if err != nil {
		return 0
	}
	return intVal
}

// StringToLong convert string to int64, or return 0 on error
func StringToLong(str string) int64 {
	long, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}
	return long
}

// StringToFloat64OrNil convert string to float64, or return 0 on error
func StringToFloat64OrNil(str string) (*float64, error) {
	//Some float number may be seperated by comma, instead of dot
	str = strings.Replace(str, ",", ".", -1)

	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return nil, fmt.Errorf("error on string to float64 convert: %v\n", err)
	}
	return &f, nil
}

// InterfaceToFloat64OrNil convert interface to float64, or return nil on error
func InterfaceToFloat64OrNil(str interface{}) (*float64, error) {
	if str == nil || &str == nil {
		return nil, fmt.Errorf("error on string to float64 convert: nil value\n")
	}
	f, err := strconv.ParseFloat(str.(string), 64)
	if err != nil {
		return nil, fmt.Errorf("error on string to float64 convert: %v\n", err)
	}
	return &f, nil
}

// StringToBool convert string to bool, or return false on error
func StringToBool(strVal string) bool {
	boolVal, err := strconv.ParseBool(strVal)
	if err != nil {
		return false
	}
	return boolVal
}

// ExcludeListElementsFromMap - exclude elements from map by key list
func ExcludeListElementsFromMap(m map[string]interface{}, ignoreKeyList []string) map[string]interface{} {
	res := map[string]interface{}{}
	for key, val := range m {
		if StringListContains(ignoreKeyList, key) {
			continue
		}
		res[key] = val
	}
	return res
}
