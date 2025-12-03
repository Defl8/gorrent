package bencode

import (
	"errors"
	"slices"
	"strconv"
)

// TODO: update once Encode functions are in place
type BencodeType any

type BencodeByteString struct {
	Length   int64
	Contents string
}

type BencodeInteger struct {
	Value int64
}

type BencodeList struct {
	Elements []BencodeType
}

type BencodeDict struct {
	Pairs map[BencodeByteString]BencodeType
}

func DecodeElement(data []byte) (BencodeType, int, error) {
	if len(data) == 0 {
		return nil, 0, errors.New("No data to decode.")
	}

	switch data[0] {
	case 'i':
		bCodeInt, bytesConsumed, err := DecodeInteger(data)
		if err != nil {
			return nil, 0, err
		}
		return bCodeInt, bytesConsumed, nil

	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		bCodeByteString, bytesConsumed, err := DecodeByteString(data)
		if err != nil {
			return nil, 0, err
		}
		return bCodeByteString, bytesConsumed, nil

	case 'l':
		bCodeList, bytesConsumed, err := DecodeList(data)
		if err != nil {
			return nil, 0, err
		}
		return bCodeList, bytesConsumed, nil

	case 'd':
		bCodeDict, bytesConsumed, err := DecodeDict(data)
		if err != nil {
			return nil, 0, err
		}
		return bCodeDict, bytesConsumed, nil

	default:
		return nil, 0, errors.New("Data passed could not be decoded.")
	}
}

func DecodeByteString(data []byte) (*BencodeByteString, int, error) {
	colonIndex := -1
	for i, b := range data {
		if b == ':' {
			colonIndex = i
			break
		}
	}

	if colonIndex == -1 {
		return nil, 0, errors.New("No colon found in byte string.")
	}

	length, err := strconv.Atoi(string(data[:colonIndex]))
	if err != nil {
		return nil, 0, err
	}

	contentStart := colonIndex + 1
	contentEnd := contentStart + length

	if len(data) < contentEnd {
		return nil, 0, errors.New("Byte string length greater than available data")
	}

	contents := string(data[contentStart:contentEnd])

	return &BencodeByteString{
		Length:   int64(length),
		Contents: contents,
	}, contentEnd, nil
}

func DecodeInteger(data []byte) (*BencodeInteger, int, error) {
	if len(data) < 3 {
		return nil, 0, errors.New("Integer data too short.")
	}

	if data[0] != 'i' {
		return nil, 0, errors.New("Expected 'i' at beginning of integer data.")
	}

	endIndex := -1
	for i := 1; i < len(data); i++ {
		if data[i] == 'e' {
			endIndex = i
			break
		}
	}

	if endIndex == -1 {
		return nil, 0, errors.New("Expected 'e' at the end of integer.")
	}

	intBytes := data[1:endIndex]
	if len(intBytes) == 0 {
		return nil, 0, errors.New("Empty integer value.")
	}

	if len(intBytes) > 1 && intBytes[0] == 0 {
		return nil, 0, errors.New("Invalid leading zero found in integer.")
	}

	if len(intBytes) > 2 && string(intBytes) == "-0" {
		return nil, 0, errors.New("Negative zero is invalid.")
	}

	value, err := strconv.ParseInt(string(intBytes), 10, 64)
	if err != nil {
		return nil, 0, err
	}
	bCodeInt := BencodeInteger{Value: value}
	return &bCodeInt, len(data[:endIndex+1]), nil
}

func DecodeList(data []byte) (*BencodeList, int, error) {
	if data[0] != 'l' {
		return nil, 0, errors.New("List missing starting delimiter 'l'")
	}

	if len(data) < 2 {
		return nil, 0, errors.New("List data too short to decode.")
	}

	elements := BencodeList{Elements: []BencodeType{}}

	index := 1 // start after the first char since we know that the first char is 'l'

	for index < len(data) {
		if data[index] == 'e' { // if the next char after 'l' is 'e', list empty
			return &elements, index + 1, nil
		}

		elem, bytesConsumed, err := DecodeElement(data[index:])
		if err != nil {
			return nil, 0, err
		}

		elements.Elements = append(elements.Elements, elem)

		index += bytesConsumed // don't need to look at bytes that have already been decoded.
	}

	return nil, 0, errors.New("List missing end delimiter 'e'")
}

func DecodeDict(data []byte) (*BencodeDict, int, error) {
	// d7:meaningi42e4:wiki7:bencodee

	if data[0] != 'd' {
		return nil, 0, errors.New("Dict missing starting delimiter 'd'")
	}

	if len(data) < 2 {
		return nil, 0, errors.New("Dict data too short to decode.")
	}

	bcodeDict := BencodeDict{Pairs: make(map[BencodeByteString]BencodeType)}
	index := 1
	hasKey := false

	// Make sure there are no duplicat keys
	keyMemo := []BencodeByteString{}

	var key BencodeByteString
	for index < len(data) {
		if data[index] == 'e' {
			return &bcodeDict, index + 1, nil
		}

		if hasKey {
			val, bytesConsumed, err := DecodeElement(data[index:])
			if err != nil {
				return nil, 0, err
			}

			// Complete dict pair in struct
			bcodeDict.Pairs[key] = val

			hasKey = false
			index += bytesConsumed
			continue
		}

		if !hasKey {
			keyByteString, bytesConsumed, err := DecodeByteString(data[index:])
			if err != nil {
				return nil, 0, err
			}

			key = *keyByteString

			// Error if dict has duplicate keys
			if slices.Contains(keyMemo, key) {
				return nil, 0, errors.New("Bencode dictionary contains duplicate keys.")
			}

			// Memoize if the key is not a dupe
			keyMemo = append(keyMemo, key)


			hasKey = true
			index += bytesConsumed
		}
	}
	return nil, 0, errors.New("Dict missing end delimiter 'e'")
}

func (d BencodeDict) Get(contents string) (BencodeType, error) {
	byteString := BencodeByteString{Length: int64(len(contents)), Contents: contents}

	bCodeType, ok := d.Pairs[byteString]

	if !ok {
		return nil, errors.New("Key Value pair does not exist in Bencode Dictionary.")
	}

	return bCodeType, nil
}
