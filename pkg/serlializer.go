package pkg

import "math"

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
	case uint8:
		result = append(result, 0xCC)
		result = append(result, byte(objType))
	case uint16:
		result = append(result, 0xCD)
		result = append(result, byte(objType>>8), byte(objType))
	case uint32:
		result = append(result, 0xCE)
		result = append(result, byte(objType>>24), byte(objType>>16), byte(objType>>8), byte(objType))
	case uint64:
		result = append(result, 0xCF)
		result = append(result, byte(objType>>56), byte(objType>>48), byte(objType>>40), byte(objType>>32), byte(objType>>24), byte(objType>>16), byte(objType>>8), byte(objType))
	case int8:
		result = append(result, 0xD0)
		result = append(result, byte(objType))
	case int16:
		result = append(result, 0xD1)
		result = append(result, byte(objType>>8), byte(objType))
	case int32:
		result = append(result, 0xD2)
		result = append(result, byte(objType>>24), byte(objType>>16), byte(objType>>8), byte(objType))
	case int64:
		result = append(result, 0xD3)
		result = append(result, byte(objType>>56), byte(objType>>48), byte(objType>>40), byte(objType>>32), byte(objType>>24), byte(objType>>16), byte(objType>>8), byte(objType))
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
			result = append(result, []byte(objType)...)
		} else if strLen < 2^8 {
			result = append(result, 0xD9)
			result = append(result, byte(strLen))
			result = append(result, []byte(objType)...)
		} else if strLen < 2^16 {
			result = append(result, 0xDA)
			result = append(result, byte(strLen>>8), byte(strLen))
			result = append(result, []byte(objType)...)
		} else if strLen < 2^32 {
			result = append(result, 0xDB)
			result = append(result, byte(strLen>>24), byte(strLen>>16), byte(strLen>>8), byte(strLen))
			result = append(result, []byte(objType)...)
		}
	case []byte:
		byteLen := len(objType)
		if byteLen < 2^8 {
			result = append(result, 0xC4)
			result = append(result, byte(byteLen))
			result = append(result, []byte(objType)...)
		} else if byteLen < 2^16 {
			result = append(result, 0xC5)
			result = append(result, byte(byteLen>>8), byte(byteLen))
			result = append(result, []byte(objType)...)
		} else if byteLen < 2^32 {
			result = append(result, 0xC6)
			result = append(result, byte(byteLen>>24), byte(byteLen>>16), byte(byteLen>>8), byte(byteLen))
			result = append(result, []byte(objType)...)
		}
	}
	return result, nil
}
