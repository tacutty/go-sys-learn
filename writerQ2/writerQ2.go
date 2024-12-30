package main

import (
	"encoding/csv"
	"os"
)

func main() {
	data := [][]string{
		{"ID", "Name", "Age"},
		{"1", "Alice", "30"},
		{"2", "Bob", "25"},
		{"3", "Charlie", "35"},
	}
	file, err := os.Create("test.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	if err := writer.WriteAll(data); err != nil {
		panic(err)
	}
	writer.Flush()
}
