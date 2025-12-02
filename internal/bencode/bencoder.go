package bencode

import (
	"errors"
	"strconv"
)

type BencodeType interface {
	decode(data []byte, currentPos int) int
}

func Decode(data []byte) {
	for i := range data {
		switch data[i] {
		case 'd':
		}
	}

}

type BencodeByteString struct {
	Length   int64
	Contents string
}

func DecodeByteString(data []byte) (*BencodeByteString, error) {
	colonIndex := -1
	for i, b := range data {
		if b == ':' {
			colonIndex = i
			break
		}
	}

	if colonIndex == -1 {
		return nil, errors.New("No colon found in byte string.")
	}

	length, err := strconv.Atoi(string(data[:colonIndex]))
	if err != nil {
		return nil, err
	}

	contentStart := colonIndex + 1

	if len(data) < contentStart+length {
		return nil, errors.New("Byte string length greater than available data")
	}

	contents := string(data[contentStart : contentStart+length])

	return &BencodeByteString{
		Length:   int64(length),
		Contents: contents,
	}, nil
}

func DecodeInteger(data []byte) (*int64, error) {
	if len(data) < 3 {
		return nil, errors.New("Integer data too short.")
	}

	if data[0] != 'i' {
		return nil, errors.New("Expected 'i' at beginning of integer data.")
	}

	endIndex := -1
	for i := 1; i < len(data); i++ {
		if data[i] == 'e' {
			endIndex = i
			break
		}
	}

	if endIndex == -1 {
		return nil, errors.New("Expected 'e' at the end of integer.")
	}

	intBytes := data[1:endIndex]
	if len(intBytes) == 0 {
		return nil, errors.New("Empty integer value.")
	}

	if len(intBytes) > 1 && intBytes[0] == 0 {
		return nil, errors.New("Invalid leading zero found in integer.")
	}

	if len(intBytes) > 2 && string(intBytes) == "-0" {
		return nil, errors.New("Negative zero is invalid.")
	}

	value, err := strconv.ParseInt(string(intBytes), 10, 64)
	if err != nil {
		return nil, err
	}

	return &value, nil
}
