package main

import (
	"testing"
)

func TestCamelToSnake(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"thisIsCamelCase", "this_is_camel_case"},
		{"simpleTest", "simple_test"},
		{"CamelCaseToSnakeCase", "camel_case_to_snake_case"},
		{"already_snake_case", "already_snake_case"},
		{"Single", "single"},
		{"", ""},
	}

	for _, test := range tests {
		result := CamelToSnake(test.input)
		if result != test.expected {
			t.Errorf("For input %q, expected %q but got %q", test.input, test.expected, result)
		}
	}
}
