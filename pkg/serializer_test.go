package pkg

import (
	"bytes"
	"reflect"
	"testing"
)

func TestSerialize(t *testing.T) {
	testcases := []struct {
		name     string
		input    interface{}
		expected []byte
	}{
		{
			name:     "nil",
			input:    nil,
			expected: []byte{0xC0},
		},
		{
			name:     "boolean true",
			input:    true,
			expected: []byte{0xC3},
		},
		{
			name:     "boolean false",
			input:    false,
			expected: []byte{0xC2},
		},
		{
			name:     "uint8",
			input:    uint8(255),
			expected: []byte{0xCC, 0xFF},
		},
		{
			name:     "uint16",
			input:    uint16(65535),
			expected: []byte{0xCD, 0xFF, 0xFF},
		},
		{
			name:     "uint32",
			input:    uint32(4294967295),
			expected: []byte{0xCE, 0xFF, 0xFF, 0xFF, 0xFF},
		},
		{
			name:     "uint64",
			input:    uint64(5000000000),
			expected: []byte{0xCF, 0x00, 0x00, 0x00, 0x01, 0x2A, 0x05, 0xF2, 0x00},
		},
		{
			name:     "positive fixint",
			input:    100,
			expected: []byte{0x64},
		},
		{
			name:     "negative fixint",
			input:    -20,
			expected: []byte{0xEC},
		},
		{
			name:     "int8",
			input:    -128,
			expected: []byte{0xD0, 0x80},
		},
		{
			name:     "int16",
			input:    -300,
			expected: []byte{0xD1, 0xFE, 0xD4},
		},
		{
			name:     "int32",
			input:    -70000,
			expected: []byte{0xD2, 0xFF, 0xFE, 0xEE, 0x90},
		},
		{
			name:     "int64",
			input:    -5000000000,
			expected: []byte{0xD3, 0xFF, 0xFF, 0xFF, 0xFE, 0xD5, 0xFA, 0x0E, 0x00},
		},
		{
			name:     "float32",
			input:    float32(3.14),
			expected: []byte{0xCA, 0x40, 0x48, 0xF5, 0xC3},
		},
		{
			name:     "float64",
			input:    float64(3.141592653589793),
			expected: []byte{0xCB, 0x40, 0x09, 0x21, 0xFB, 0x54, 0x44, 0x2D, 0x18},
		},
		{
			name:     "fix string",
			input:    "hello",
			expected: []byte{0xA5, 'h', 'e', 'l', 'l', 'o'},
		},
		{
			name:     "str8",
			input:    string(make([]byte, 100)),
			expected: append([]byte{0xD9, 100}, make([]byte, 100)...),
		},
		{
			name:     "str16",
			input:    string(make([]byte, 7000)),
			expected: append([]byte{0xDA, 0x1B, 0x58}, make([]byte, 7000)...),
		},
		{
			name:     "str32",
			input:    string(make([]byte, 70000)),
			expected: append([]byte{0xDB, 0x00, 0x01, 0x11, 0x70}, make([]byte, 70000)...),
		},
		{
			name:     "bin8",
			input:    []byte{0x01, 0x02, 0x03, 0x04, 0x05},
			expected: []byte{0xC4, 0x05, 0x01, 0x02, 0x03, 0x04, 0x05},
		},
		{
			name:     "bin16",
			input:    make([]byte, 300),
			expected: append([]byte{0xC5, 0x01, 0x2C}, make([]byte, 300)...),
		},
		{
			name:     "bin32",
			input:    make([]byte, 70000),
			expected: append([]byte{0xC6, 0x00, 0x01, 0x11, 0x70}, make([]byte, 70000)...),
		},
		{
			name:     "fix array",
			input:    []any{1, 2, 3},
			expected: []byte{0x93, 0x01, 0x02, 0x03},
		},
		{
			name:     "array16",
			input:    make([]any, 200),
			expected: append([]byte{0xDC, 0x00, 0xC8}, bytes.Repeat([]byte{0xC0}, 200)...),
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := Serialize(tc.input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, result)
			}
		})
	}

}
