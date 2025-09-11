package pkg

import (
	"math"
	"reflect"
)

func handleInt(objType int64, result []byte) []byte {
	if objType >= 0 && objType < 128 { // positive fixint
		result = append(result, byte(objType))

	} else if objType >= -32 && objType < 0 { // negative fixint
		result = append(result, byte(int8(objType)))

	} else if objType >= math.MinInt8 && objType <= math.MaxInt8 { // int 8
		result = append(result, 0xD0)
		result = append(result, byte(int8(objType)))

	} else if objType >= math.MinInt16 && objType <= math.MaxInt16 { // int 16
		result = append(result, 0xD1)
		result = append(result, byte(int16(objType)>>8), byte(int16(objType)))

	} else if objType >= math.MinInt32 && objType <= math.MaxInt32 { // int 32
		result = append(result, 0xD2)
		result = append(result, byte(int32(objType)>>24), byte(int32(objType)>>16), byte(int32(objType)>>8), byte(int32(objType)))

	} else { // int 64
		result = append(result, 0xD3)
		result = append(result, byte(int64(objType)>>56), byte(int64(objType)>>48), byte(int64(objType)>>40), byte(int64(objType)>>32), byte(int64(objType)>>24), byte(int64(objType)>>16), byte(int64(objType)>>8), byte(int64(objType)))
	}
	return result
}

func handleUint(objType uint64, result []byte) []byte {
	if objType < (1 << 7) { // positive fixint
		result = append(result, byte(objType))
	} else if objType < (1 << 8) { // uint 8
		result = append(result, 0xCC)
		result = append(result, byte(objType))
	} else if objType < (1 << 16) { // uint 16
		result = append(result, 0xCD)
		result = append(result, byte(objType>>8), byte(objType))
	} else if objType < (1 << 32) { // uint 32
		result = append(result, 0xCE)
		result = append(result, byte(objType>>24), byte(objType>>16), byte(objType>>8), byte(objType))
	} else { // uint 64
		result = append(result, 0xCF)
		result = append(result, byte(objType>>56), byte(objType>>48), byte(objType>>40), byte(objType>>32), byte(objType>>24), byte(objType>>16), byte(objType>>8), byte(objType))
	}
	return result
}

func Serialize(object interface{}) ([]byte, error) {
	var result []byte

	switch objType := object.(type) {

	case nil:
		result = append(result, 0xC0)

	case bool:
		if objType {
			result = append(result, 0xC3)
		} else {
			result = append(result, 0xC2)
		}

	case uint, uint8, uint16, uint32, uint64:
		result = handleUint(reflect.ValueOf(objType).Uint(), result)

	case int, int8, int16, int32, int64:
		result = handleInt(reflect.ValueOf(objType).Int(), result)
	case float32:
		result = append(result, 0xCA)
		floatBits := math.Float32bits(objType)
		result = append(result, byte(floatBits>>24), byte(floatBits>>16), byte(floatBits>>8), byte(floatBits))
	case float64:
		result = append(result, 0xCB)
		floatBits := math.Float64bits(objType)
		result = append(result, byte(floatBits>>56), byte(floatBits>>48), byte(floatBits>>40), byte(floatBits>>32), byte(floatBits>>24), byte(floatBits>>16), byte(floatBits>>8), byte(floatBits))
	case string:
		strLen := len(objType)
		if strLen < 32 {
			result = append(result, 0xA0|byte(strLen))

		} else if strLen < int(math.Pow(2, 8)) {
			result = append(result, 0xD9)
			result = append(result, byte(strLen))

		} else if strLen < int(math.Pow(2, 16)) {
			result = append(result, 0xDA)
			result = append(result, byte(strLen>>8), byte(strLen))

		} else if strLen < int(math.Pow(2, 32)) {
			result = append(result, 0xDB)
			result = append(result, byte(strLen>>24), byte(strLen>>16), byte(strLen>>8), byte(strLen))
		}
		result = append(result, []byte(objType)...)
	case []byte:
		byteLen := len(objType)
		if byteLen < int(math.Pow(2, 8)) {
			result = append(result, 0xC4)
			result = append(result, byte(byteLen))

		} else if byteLen < int(math.Pow(2, 16)) {
			result = append(result, 0xC5)
			result = append(result, byte(byteLen>>8), byte(byteLen))

		} else if byteLen < int(math.Pow(2, 32)) {
			result = append(result, 0xC6)
			result = append(result, byte(byteLen>>24), byte(byteLen>>16), byte(byteLen>>8), byte(byteLen))
		}
		result = append(result, []byte(objType)...)
	case []any:
		arrLen := len(objType)
		if arrLen < 16 {
			result = append(result, 0x90|byte(arrLen))

		} else if arrLen < int(math.Pow(2, 16)) {
			result = append(result, 0xDC)
			result = append(result, byte(arrLen>>8), byte(arrLen))

		} else if arrLen < int(math.Pow(2, 32)) {
			result = append(result, 0xDD)
			result = append(result, byte(arrLen>>24), byte(arrLen>>16), byte(arrLen>>8), byte(arrLen))
		}

		for _, element := range objType {
			serializedElement, err := Serialize(element)
			if err != nil {
				return nil, err
			}
			result = append(result, serializedElement...)
		}
	}
	return result, nil
}
