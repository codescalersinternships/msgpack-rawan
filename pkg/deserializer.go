package pkg

import "math"

func Deserialize(bytes []byte) (interface{}, int, error) {
	var result interface{}

	switch bytes[0] {

	case 0xC0: //nil
		result = nil
		return result, 1, nil
	case 0xC3: //bool true
		result = true
		return result, 1, nil
	case 0xC2: // bool false
		result = false
		return result, 1, nil
	case 0xCC: // uint8
		result = uint8(bytes[1])
		return result, 2, nil
	case 0xCD: // uint16
		result = uint16(bytes[1])<<8 | uint16(bytes[2])
		return result, 3, nil
	case 0xCE: // uint32
		result = uint32(bytes[1])<<24 | uint32(bytes[2])<<16 | uint32(bytes[3])<<8 | uint32(bytes[4])
		return result, 5, nil
	case 0xCF: // uint64
		result = uint64(bytes[1])<<56 | uint64(bytes[2])<<48 | uint64(bytes[3])<<40 | uint64(bytes[4])<<32 | uint64(bytes[5])<<24 | uint64(bytes[6])<<16 | uint64(bytes[7])<<8 | uint64(bytes[8])
		return result, 9, nil
	case 0xD0: // int8
		result = int8(bytes[1])
		return result, 2, nil
	case 0xD1: // int16
		result = int16(bytes[1])<<8 | int16(bytes[2])
		return result, 3, nil
	case 0xD2: // int32
		result = int32(bytes[1])<<24 | int32(bytes[2])<<16 | int32(bytes[3])<<8 | int32(bytes[4])
		return result, 5, nil
	case 0xD3: // int64
		result = int64(bytes[1])<<56 | int64(bytes[2])<<48 | int64(bytes[3])<<40 | int64(bytes[4])<<32 | int64(bytes[5])<<24 | int64(bytes[6])<<16 | int64(bytes[7])<<8 | int64(bytes[8])
		return result, 9, nil
	case 0xCA: // float32
		result = math.Float32frombits(uint32(bytes[1])<<24 | uint32(bytes[2])<<16 | uint32(bytes[3])<<8 | uint32(bytes[4]))
		return result, 5, nil
	case 0xCB: // float64
		result = math.Float64frombits(uint64(bytes[1])<<56 | uint64(bytes[2])<<48 | uint64(bytes[3])<<40 | uint64(bytes[4])<<32 | uint64(bytes[5])<<24 | uint64(bytes[6])<<16 | uint64(bytes[7])<<8 | uint64(bytes[8]))
		return result, 9, nil
	//fixstr
	case 0xA0, 0xA1, 0xA2, 0xA3, 0xA4, 0xA5, 0xA6, 0xA7, 0xA8, 0xA9, 0xAA, 0xAB, 0xAC, 0xAD, 0xAE, 0xAF, 0xB0, 0xB1, 0xB2, 0xB3, 0xB4, 0xB5, 0xB6, 0xB7, 0xB8, 0xB9, 0xBA, 0xBB, 0xBC, 0xBD, 0xBE, 0xBF:
		strLen := int(bytes[0] & 0x1F) // lower 5 bits
		result = string(bytes[1 : 1+strLen])
		return result, int(bytes[0]&0x1F) + 1, nil // lower 5 bits
	case 0xD9: // str8
		strLen := int(bytes[1])
		result = string(bytes[2 : 2+strLen])
		return result, strLen + 2, nil
	case 0xDA: // str16
		strLen := int(bytes[1])<<8 | int(bytes[2])
		result = string(bytes[3 : 3+strLen])
		return result, strLen + 3, nil
	case 0xDB: // str32
		strLen := int(bytes[1])<<24 | int(bytes[2])<<16 | int(bytes[3])<<8 | int(bytes[4])
		result = string(bytes[5 : 5+strLen])
		return result, strLen + 5, nil
	case 0xC4: // bin8
		binLen := int(bytes[1])
		result = bytes[2 : 2+binLen]
		return result, binLen + 2, nil
	case 0xC5: // bin16
		binLen := int(bytes[1])<<8 | int(bytes[2])
		result = bytes[3 : 3+binLen]
		return result, binLen + 3, nil
	case 0xC6: // bin32
		binLen := int(bytes[1])<<24 | int(bytes[2])<<16 | int(bytes[3])<<8 | int(bytes[4])
		result = bytes[5 : 5+binLen]
		return result, binLen + 5, nil
	//fixarray
	case 0x90, 0x91, 0x92, 0x93, 0x94, 0x95, 0x96, 0x97, 0x98, 0x99, 0x9A, 0x9B, 0x9C, 0x9D, 0x9E, 0x9F:
		arrayLen := int(bytes[0] & 0x0F) //get lower 4 bits
		arr := make([]any, 0, arrayLen)
		rest := bytes[1:]
		for i := 0; i < arrayLen; i++ {
			elem, n, _ := Deserialize(rest)
			arr = append(arr, elem)
			rest = rest[n:]
		}
		result = arr
		return result, len(bytes) - len(rest), nil
	case 0xDC: // array16
		arrayLen := int(bytes[1])<<8 | int(bytes[2])
		arr := make([]any, 0, arrayLen)

		rest := bytes[3:]
		for i := 0; i < arrayLen; i++ {
			elem, n, _ := Deserialize(rest)
			arr = append(arr, elem)
			rest = rest[n:]
		}
		result = arr
		return result, len(bytes) - len(rest), nil
	case 0xDD: // array32
		arrayLen := int(bytes[1])<<24 | (int(bytes[2]) << 16) | (int(bytes[3]) << 8) | int(bytes[4])
		arr := make([]any, 0, arrayLen)

		rest := bytes[5:]
		for i := 0; i < arrayLen; i++ {
			elem, n, _ := Deserialize(rest)
			arr = append(arr, elem)
			rest = rest[n:]
		}
		result = arr
		return result, len(bytes) - len(rest), nil
	//fixmap
	case 0x80, 0x81, 0x82, 0x83, 0x84, 0x85, 0x86, 0x87, 0x88, 0x89, 0x8A, 0x8B, 0x8C, 0x8D, 0x8E, 0x8F:
		mapLen := int(bytes[0] & 0x0F) //get lower 4 bits
		temp_map := make(map[any]any, mapLen)

		rest := bytes[1:]
		consumed := 1
		for i := 0; i < mapLen; i++ {
			key, n, _ := Deserialize(rest)
			consumed += n
			rest = rest[n:]

			val, n, _ := Deserialize(rest)
			consumed += n
			rest = rest[n:]
			temp_map[key] = val
		}
		result = temp_map
		return result, len(bytes) - len(rest), nil
	default:
		if bytes[0] < 128 { // positive fixint
			result = int(bytes[0])
			return result, 1, nil
		} else if bytes[0] >= 0xE0 { // negative fixint
			result = int(int8(bytes[0]))
			return result, 1, nil
		}
	}
	return result, 0, nil
}
