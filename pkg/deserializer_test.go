package pkg

import (
	"bytes"
	"reflect"
	"testing"
)

func TestDeserialize(t *testing.T) {
	testcases := []struct {
		name     string
		input    []byte
		expected interface{}
	}{
		{
			name:     "nil",
			input:    []byte{0xC0},
			expected: nil,
		},
		{
			name:     "boolean true",
			input:    []byte{0xC3},
			expected: true,
		},
		{
			name:     "boolean false",
			input:    []byte{0xC2},
			expected: false,
		},
		{
			name:     "uint8",
			input:    []byte{0xCC, 0xFF},
			expected: uint8(255),
		},
		{
			name:     "uint16",
			input:    []byte{0xCD, 0xFF, 0xFF},
			expected: uint16(65535),
		},
		{
			name:     "uint32",
			input:    []byte{0xCE, 0xFF, 0xFF, 0xFF, 0xFF},
			expected: uint32(4294967295),
		},
		{
			name:     "uint64",
			input:    []byte{0xCF, 0x00, 0x00, 0x00, 0x01, 0x2A, 0x05, 0xF2, 0x00},
			expected: uint64(5000000000),
		},
		{
			name:     "float32",
			input:    []byte{0xCA, 0x40, 0x48, 0xF5, 0xC3},
			expected: float32(3.14),
		},
		{
			name:     "float64",
			input:    []byte{0xCB, 0x40, 0x09, 0x21, 0xFB, 0x54, 0x44, 0x2D, 0x18},
			expected: float64(3.141592653589793),
		},
		{
			name:     "fix string",
			input:    []byte{0xA5, 'h', 'e', 'l', 'l', 'o'},
			expected: "hello",
		},
		{
			name:     "str8",
			input:    append([]byte{0xD9, 100}, make([]byte, 100)...),
			expected: string(make([]byte, 100)),
		},
		{
			name:     "str16",
			input:    append([]byte{0xDA, 0x1B, 0x58}, make([]byte, 7000)...),
			expected: string(make([]byte, 7000)),
		},
		{
			name:     "str32",
			input:    append([]byte{0xDB, 0x00, 0x01, 0x11, 0x70}, make([]byte, 70000)...),
			expected: string(make([]byte, 70000)),
		},
		{
			name:     "bin8",
			input:    []byte{0xC4, 0x05, 0x01, 0x02, 0x03, 0x04, 0x05},
			expected: []byte{0x01, 0x02, 0x03, 0x04, 0x05},
		},
		{
			name:     "bin16",
			input:    append([]byte{0xC5, 0x01, 0x2C}, make([]byte, 300)...),
			expected: make([]byte, 300),
		},
		{
			name:     "bin32",
			input:    append([]byte{0xC6, 0x00, 0x01, 0x11, 0x70}, make([]byte, 70000)...),
			expected: make([]byte, 70000),
		},
		{
			name:     "empty array",
			input:    []byte{0x90},
			expected: []any{},
		},
		{
			name:     "fix array",
			input:    []byte{0x93, 0x01, 0x02, 0x03},
			expected: []any{1, 2, 3},
		},
		{
			name:     "array16",
			input:    append([]byte{0xDC, 0x00, 0xC8}, bytes.Repeat([]byte{0xC0}, 200)...),
			expected: make([]any, 200),
		},
		{
			name:     "array32",
			input:    append([]byte{0xDD, 0x00, 0x01, 0x11, 0x70}, bytes.Repeat([]byte{0xC0}, 70000)...),
			expected: make([]any, 70000),
		},
		{
			name:     "nested array",
			input:    []byte{0x93, 0x01, 0x92, 0x02, 0x03, 0x04},
			expected: []any{1, []any{2, 3}, 4},
		},
		{
			name:     "empty map",
			input:    []byte{0x80},
			expected: map[any]any{},
		},
		{
			name:  "fix map",
			input: []byte{0x82, 0xA1, 'a', 0x01, 0xA1, 'b', 0xC3},
			expected: map[any]any{
				"a": 1,
				"b": true,
			},
		},
		{
			name:     "positive fixint",
			input:    []byte{0x64},
			expected: 100,
		},
		{
			name:     "negative fixint",
			input:    []byte{0xEC},
			expected: -20,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			result, _, err := Deserialize(tc.input)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, result)
				t.Errorf("Type of expected: %T, Type of got: %T", tc.expected, result)
			}
		})
	}
}
