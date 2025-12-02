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

func DecodeByteString(data []byte) (*ByteString, error) {
	colonIndex := -1
	for i, b := range data {
		if b == ':' {
			colonIndex = i
			break;
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

	if len(data) < contentStart + length {
		return nil, errors.New("Byte string length greater than available data")
	}

	contents := string(data[contentStart:contentStart+length])

	return &ByteString{
		Length:   int64(length),
		Contents: contents,
	}, nil
}

type ByteString struct {
	Length   int64
	Contents string
}
