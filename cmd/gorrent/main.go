package main

import (
	"fmt"
	"log"

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


	byteInt := "i67e"
	fmt.Printf("Raw Integer: %s\n\n", byteInt)

	rawBytes := []byte(byteInt)

	decodedInt, err := bencode.DecodeInteger(rawBytes)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Decoded Integer: %d\n", *decodedInt)
}
