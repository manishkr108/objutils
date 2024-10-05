package objutils

import (
	"reflect"
	"testing"
)

func TestStructKeyShortener(t *testing.T) {
	tests := []struct {
		input       interface{}
		expected    map[string]interface{}
		expectError bool
	}{
		{
			input: map[string]interface{}{
				"keyOne": 1,
				"keyTwo": 2,
			},
			expected: map[string]interface{}{
				"key":  1,
				"keyT": 2,
			},
			expectError: false,
		},
		{
			input: struct {
				FieldOne int
				FieldTwo int
			}{FieldOne: 1, FieldTwo: 2},
			expected: map[string]interface{}{
				"FieldOne": 1,
				"FieldTwo": 2,
			},
			expectError: false,
		},
		{
			input:       123, // Invalid input
			expectError: true,
		},
	}

	for _, test := range tests {
		result, err := StructKeyShortener(test.input, map[string]string{"keyOne": "key", "keyTwo": "keyT"})
		if (err != nil) != test.expectError {
			t.Errorf("expected error: %v, got: %v", test.expectError, err)
			continue
		}
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("expected: %v, got: %v", test.expected, result)
		}
	}
}

func TestMaxValue(t *testing.T) {
	tests := []struct {
		input    map[string]interface{}
		expected interface{}
	}{
		{
			input: map[string]interface{}{
				"a": 1,
				"b": 2,
				"c": 3,
			},
			expected: 3.0,
		},
		{
			input: map[string]interface{}{
				"a": 1.5,
				"b": 2.5,
				"c": 3.5,
			},
			expected: 3.5,
		},
		{
			input: map[string]interface{}{
				"a": 10,
				"b": nil,
				"c": 5,
			},
			expected: 10.0,
		},
	}

	for _, test := range tests {
		result, err := MaxValue(test.input)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
			continue
		}
		if result != test.expected {
			t.Errorf("expected: %v, got: %v", test.expected, result)
		}
	}
}

func TestGetAllKeys(t *testing.T) {
	tests := []struct {
		input    map[string]interface{}
		expected []string
	}{
		{
			input: map[string]interface{}{
				"a": 1,
				"b": 2,
				"c": 3,
			},
			expected: []string{"a", "b", "c"},
		},
	}

	for _, test := range tests {
		result := GetAllKeys(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("expected: %v, got: %v", test.expected, result)
		}
	}
}

func TestGetAllValues(t *testing.T) {
	tests := []struct {
		input    map[string]interface{}
		expected []interface{}
	}{
		{
			input: map[string]interface{}{
				"a": 1,
				"b": 2,
				"c": 3,
			},
			expected: []interface{}{1, 2, 3},
		},
	}

	for _, test := range tests {
		result := GetAllValues(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("expected: %v, got: %v", test.expected, result)
		}
	}
}

func TestCombineMaps(t *testing.T) {
	m1 := map[string]interface{}{"a": 1, "b": 2}
	m2 := map[string]interface{}{"b": 3, "c": 4}
	expected := map[string]interface{}{"a": 1, "b": 3, "c": 4} // last map's value takes precedence

	result := CombineMaps(m1, m2)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected: %v, got: %v", expected, result)
	}
}

func TestGetValuesByKey(t *testing.T) {
	input := map[string]interface{}{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	tests := []struct {
		keys     []string
		expected []interface{}
	}{
		{keys: []string{"a", "c"}, expected: []interface{}{1, 3}},
		{keys: []string{"b"}, expected: []interface{}{2}},
	}

	for _, test := range tests {
		result := GetValuesByKey(input, test.keys...)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("expected: %v, got: %v", test.expected, result)
		}
	}
}

func TestSortAndPreserveDuplicates(t *testing.T) {
	input := []string{"b", "c", "e", "d", "t", "a", "b", "t"}
	expected := []string{"a", "b", "b", "c", "d", "e", "t", "t"}

	result := SortAndPreserveDuplicates(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected: %v, got: %v", expected, result)
	}
}

func TestSortMapByKey(t *testing.T) {
	input := map[string]interface{}{
		"b": 1,
		"d": 2,
		"a": 6,
		"c": 5,
		"e": 5,
	}

	expected := map[string]interface{}{
		"a": 6, // Latest value for 'a'
		"b": 1,
		"c": 5,
		"d": 2, // Latest value for 'd'
		"e": 5,
	}

	sorted, err := SortMapByKey(input)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if !reflect.DeepEqual(sorted, expected) {
		t.Errorf("expected: %v, got: %v", expected, sorted)
	}
}

// TestSortJSONKeys checks that the SortJSONKeys function correctly handles duplicate keys.
func TestSortJSONKeys(t *testing.T) {
	// Input JSON string with duplicate keys
	jsonInput := `{"b":1,"d":2,"d":4,"a":6,"c":5,"e":5,"a":2}`

	expected := []KeyValue{
		{"a", "6"},
		{"a", "2"},
		{"b", "1"},
		{"c", "5"},
		{"d", "2"},
		{"d", "4"},
		{"e", "5"},
	}

	sortedKeyValues, err := SortJSONKeys(jsonInput)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// Compare the sorted key-value pairs with the expected result.
	for i, kv := range sortedKeyValues {
		if kv.Key != expected[i].Key || kv.Value != expected[i].Value {
			t.Errorf("expected: %+v, got: %+v", expected[i], kv)
		}
	}
}
