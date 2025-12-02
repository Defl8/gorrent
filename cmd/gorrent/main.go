package main

import (
	"fmt"
	"log"

	"github.com/Defl8/gorrent/internal/bencode"
)


func main() {
	byteString := "16:thisshoulddecode"
	_, _ = fmt.Printf("Original byte string: %s\n\n", byteString)

	rawBytes := []byte(byteString)

	decodedString, err := bencode.DecodeByteString(rawBytes)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Decoded Info:\nLength: %d\nContents: %s\n", decodedString.Length, decodedString.Contents)
}
