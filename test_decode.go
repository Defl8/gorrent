package main

import (
	"fmt"
	"github.com/Defl8/gorrent/internal/bencode"
)

func main() {
	// Corrected test data with proper length prefixes
	data := []byte("d8:announce31:http://tracker.example.com:808013:announce-listll31:http://tracker.example.com:8080el30:http://backup.tracker.com:8080ee7:comment22:Test torrent with data10:created by13:test-client-113:creation datei1234567890e4:infod6:lengthi1048576e4:name9:test.file12:piece lengthi262144e6:pieces20:fakehash123456789012ee")
	
	result, consumed, err := bencode.DecodeDict(data)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		fmt.Printf("Bytes consumed before error: %d\n", consumed)
		return
	}
	
	fmt.Printf("Success! Decoded dictionary with %d pairs\n", len(result.Pairs))
	fmt.Printf("Bytes consumed: %d out of %d total\n", consumed, len(data))
}
