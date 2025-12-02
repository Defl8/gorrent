package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/Defl8/gorrent/internal/bencode"
)

func main() {
	// byteString := "16:thisshoulddecode"
	// _, _ = fmt.Printf("Original byte string: %s\n\n", byteString)
	//
	// rawBytes := []byte(byteString)
	//
	// decodedString, err := bencode.DecodeByteString(rawBytes)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// fmt.Printf("Decoded Info:\nLength: %d\nContents: %s\n\n", decodedString.Length, decodedString.Contents)

	// byteInt := "i67e"
	// fmt.Printf("Raw Integer: %s\n\n", byteInt)
	//
	// rawBytes := []byte(byteInt)
	//
	// decodedInt, err := bencode.DecodeInteger(rawBytes)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// fmt.Printf("Decoded Integer: %d\n", *decodedInt)

	byteList := "li67e4:teste"
	fmt.Printf("Raw List: %s\n\n", byteList)

	bytes := []byte(byteList)

	decodedList, byteCount, err := bencode.DecodeList(bytes)
	if err != nil {
		log.Fatal(err)
	}

	parts := make([]string, 0, len(decodedList.Elements))
	for _, elem := range decodedList.Elements {
		switch v := elem.(type){
		case *int64:
			parts = append(parts, fmt.Sprintf("%d", *v))
		case *bencode.BencodeByteString:
			parts = append(parts, v.Contents)
		}
	}

	stringsList := strings.Join(parts, ", ")

	fmt.Printf("Decoded List Contents: %s\nBytes Consumed: %d\n", stringsList, byteCount)
}
