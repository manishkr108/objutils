package main

import (
	"encoding/json"
	"fmt"
	"log"
	"objutils/objutils" // Make sure to adjust the import path according to your project structure
)

func main() {
	// Example usage of StructKeyShortener
	inputMap := map[string]interface{}{
		"keyOne": 1,
		"keyTwo": 2,
	}
	mapping := map[string]string{
		"keyOne": "key",
		"keyTwo": "keyT",
	}

	shortened, err := objutils.StructKeyShortener(inputMap, mapping)
	if err != nil {
		log.Fatalf("Error in StructKeyShortener: %v", err)
	}
	fmt.Println("Shortened Map:", shortened)

	// Example usage of MaxValue
	numericMap := map[string]interface{}{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	maxValue, err := objutils.MaxValue(numericMap)
	if err != nil {
		log.Fatalf("Error in MaxValue: %v", err)
	}
	fmt.Println("Max Value:", maxValue)

	// Example usage of GetAllKeys
	keys := objutils.GetAllKeys(numericMap)
	fmt.Println("All Keys:", keys)

	// Example usage of GetAllValues
	values := objutils.GetAllValues(numericMap)
	fmt.Println("All Values:", values)

	// Example usage of CombineMaps
	m1 := map[string]interface{}{"a": 1, "b": 2}
	m2 := map[string]interface{}{"b": 3, "c": 4}
	combined := objutils.CombineMaps(m1, m2)
	fmt.Println("Combined Map:", combined)

	// Example usage of GetValuesByKey
	input := map[string]interface{}{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	valuesByKey := objutils.GetValuesByKey(input, "a", "c")
	fmt.Println("Values by Key:", valuesByKey)
	// Test SortAndPreserveDuplicates
	sortInput := []string{"b", "c", "e", "d", "t", "a", "b", "t"}
	sorted := objutils.SortAndPreserveDuplicates(sortInput)
	fmt.Println("Sorted with Duplicates Preserved:", sorted)

	// Original JSON input with duplicate keys
	jsonInput := `{"b":1,"d":2,"d":4,"a":6,"c":5,"e":5,"a":2}`

	// Call the library function SortJSONKeys
	sortedKeyValues, err := objutils.SortJSONKeys(jsonInput)
	if err != nil {
		log.Fatalf("Error sorting JSON keys: %v", err)
	}

	// Convert sorted key-values to a JSON array for printing
	resultJSON, err := json.MarshalIndent(sortedKeyValues, "", "  ")
	if err != nil {
		log.Fatalf("Error marshalling result to JSON: %v", err)
	}

	// Print the result
	fmt.Println("Sorted JSON with duplicate keys:")
	fmt.Println(string(resultJSON))
}
