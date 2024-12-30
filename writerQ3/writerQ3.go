package main

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"os"
)

func main() {
	file, err := os.Create("test.txt.gz")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := gzip.NewWriter(file)
	defer writer.Close()
	
	multiWriter := io.MultiWriter(writer, os.Stdout)

	encoder := json.NewEncoder(multiWriter)
	encoder.SetIndent("", "    ")
	encoder.Encode(map[string]string{
		"Name": "Alice",
		"Age":  "25",
	})

	writer.Flush()
}
