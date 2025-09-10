package pkg

import (
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
			input:    uint64(18446744073709551615),
			expected: []byte{0xCF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF},
		},
		{
			name:     "int8",
			input:    int8(-128),
			expected: []byte{0xD0, 0x80},
		},
		{
			name:     "int16",
			input:    int16(-300),
			expected: []byte{0xD1, 0xFE, 0xD4},
		},
		{
			name:     "int32",
			input:    int32(-70000),
			expected: []byte{0xD2, 0xFF, 0xFE, 0xEE, 0x90},
		},
		{
			name:     "int64",
			input:    int64(-123456789),
			expected: []byte{0xD3, 0xFF, 0xFF, 0xFF, 0xFF, 0xF8, 0xA4, 0x32, 0xEB},
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
