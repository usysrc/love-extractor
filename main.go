package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: love-extractor <input-exe> <output-love>")
		return
	}

	// Read the executable
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	// Look for the ZIP header (which is the beginning of a .love file)
	zipHeader := []byte{0x50, 0x4B, 0x03, 0x04}
	pos := bytes.Index(data, zipHeader)

	if pos == -1 {
		fmt.Println("No .love file found appended to the executable.")
		return
	}

	// Extract the .love data
	loveData := data[pos:]

	// Write the .love data to a new file
	err = os.WriteFile(os.Args[2], loveData, 0644)
	if err != nil {
		fmt.Println("Error writing .love file:", err)
		return
	}

	fmt.Println(".love file successfully extracted!")
}
