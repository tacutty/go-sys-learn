package main

import (
	"io"
	"os"
	"path/filepath"
	"fmt"

	"github.com/edsrzf/mmap-go"
)

func main() {
	var testData = []byte("0123456789ABCDEF")
	var testPath = filepath.Join(os.TempDir(), "testdata")
	err := os.WriteFile(testPath, testData, 0644)
	if err != nil {
		panic(err)
	}

	// メモリにマッピング
	f, err := os.OpenFile(testPath, os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	m, err := mmap.Map(f, mmap.COPY, 0)
	if err != nil {
		panic(err)
	}

	defer m.Unmap()

	// マッピングしたメモリを書き換え
	m[9] = 'X'
	m.Flush()

	fileData, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	fmt.Printf("original: %s\n", testData)
	fmt.Printf("mmap: %s\n", m)
	fmt.Printf("file: %s\n", fileData)
}
