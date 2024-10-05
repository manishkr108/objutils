package objutils

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"sort"
)

type KeyValue struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

// StructKeyShortener reduces struct or map keys to shorter forms based on a provided mapping.
func StructKeyShortener(input interface{}, mapping map[string]string) (map[string]interface{}, error) {
	result := make(map[string]interface{})

	val := reflect.ValueOf(input)
	if val.Kind() != reflect.Map && val.Kind() != reflect.Struct {
		return nil, fmt.Errorf("input must be either a map or struct")
	}

	if val.Kind() == reflect.Map {
		for _, key := range val.MapKeys() {
			newKey := key.String()
			if mappedKey, exists := mapping[newKey]; exists {
				newKey = mappedKey
			}
			result[newKey] = val.MapIndex(key).Interface()
		}
	} else {
		for i := 0; i < val.NumField(); i++ {
			field := val.Type().Field(i)
			newKey := field.Name
			if mappedKey, exists := mapping[field.Name]; exists {
				newKey = mappedKey
			}
			result[newKey] = val.Field(i).Interface()
		}
	}

	return result, nil
}

// MaxValue returns the maximum value from a map of numeric values.
func MaxValue(m map[string]interface{}) (interface{}, error) {
	var max interface{}

	for _, value := range m {
		if value == nil {
			continue // skip nil values
		}
		val := reflect.ValueOf(value)

		var currentMax float64
		switch val.Kind() {
		case reflect.Int:
			currentMax = float64(val.Int())
		case reflect.Float64:
			currentMax = val.Float()
		default:
			return nil, fmt.Errorf("value must be numeric")
		}

		if max == nil || currentMax > reflect.ValueOf(max).Float() {
			max = currentMax
		}
	}

	return max, nil
}

// GetAllKeys returns all keys from a map.
func GetAllKeys(m map[string]interface{}) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys) // Sort keys alphabetically
	return keys
}

// GetAllValues returns all values from a map.
func GetAllValues(m map[string]interface{}) []interface{} {
	values := make([]interface{}, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

// CombineMaps combines multiple maps into a single map.
func CombineMaps(maps ...map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}

// GetValuesByKey retrieves values by key from a map.
func GetValuesByKey(m map[string]interface{}, keys ...string) []interface{} {
	values := make([]interface{}, 0, len(keys))
	for _, k := range keys {
		if v, exists := m[k]; exists {
			values = append(values, v)
		}
	}
	return values
}

// SortAndPreserveDuplicates sorts a slice of strings and preserves duplicates.
func SortAndPreserveDuplicates(keys []string) []string {
	sort.Strings(keys)
	return keys
}

// SortMapByKey sorts a map by its keys in ascending order
func SortMapByKey(input map[string]interface{}) (map[string]interface{}, error) {
	if input == nil {
		return nil, fmt.Errorf("input must be a non-nil map")
	}

	// Extract keys
	keys := make([]string, 0, len(input))
	for key := range input {
		keys = append(keys, key)
	}

	// Sort keys
	sort.Strings(keys)

	// Create a new map with sorted keys
	sortedMap := make(map[string]interface{})
	for _, key := range keys {
		sortedMap[key] = input[key]
	}

	return sortedMap, nil
}

// SortJSONKeys sorts the keys of the provided JSON string in ascending order.
// SortJSONKeys sorts key-value pairs while preserving duplicates.
func SortJSONKeys(jsonInput string) ([]KeyValue, error) {
	// Parse the JSON input to handle multiple instances of the same key
	keyValuePairs, err := parseDuplicateJSONKeys(jsonInput)
	if err != nil {
		return nil, errors.New("invalid JSON input")
	}

	// Sort key-value pairs by the key while preserving duplicates.
	sort.SliceStable(keyValuePairs, func(i, j int) bool {
		return keyValuePairs[i].Key < keyValuePairs[j].Key
	})

	return keyValuePairs, nil
}

// parseDuplicateJSONKeys converts a JSON object with duplicate keys into an array of key-value pairs
func parseDuplicateJSONKeys(jsonInput string) ([]KeyValue, error) {
	// Use a regex to extract the keys and values as key-value pairs
	re := regexp.MustCompile(`"(\w+)":(\d+)`)
	matches := re.FindAllStringSubmatch(jsonInput, -1)

	// Create a slice to hold key-value pairs
	var keyValuePairs []KeyValue
	for _, match := range matches {
		keyValuePairs = append(keyValuePairs, KeyValue{Key: match[1], Value: match[2]})
	}

	return keyValuePairs, nil
}
