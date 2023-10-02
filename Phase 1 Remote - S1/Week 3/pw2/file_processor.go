package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"sync"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run file_processor.go input.csv output.csv")
		return
	}

	inputFilename := os.Args[1]
	outputFilename := os.Args[2]

	// Buka file input
	inputFile, err := os.Open(inputFilename)
	if err != nil {
		fmt.Printf("Error opening input file: %v\n", err)
		return
	}
	defer inputFile.Close()

	// Buka file output
	outputFile, err := os.Create(outputFilename)
	if err != nil {
		fmt.Printf("Error creating output file: %v\n", err)
		return
	}
	defer outputFile.Close()

	// Membaca file CSV input
	reader := csv.NewReader(inputFile)
	writer := csv.NewWriter(outputFile)

	// Membaca baris judul (header)
	header, err := reader.Read()
	if err != nil {
		fmt.Printf("Error reading CSV header: %v\n", err)
		return
	}

	// Menulis judul (header) ke file output
	err = writer.Write(header)
	if err != nil {
		fmt.Printf("Error writing CSV header: %v\n", err)
		return
	}

	var wg sync.WaitGroup

	// Membaca dan memproses baris-baris data secara konkuren
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		wg.Add(1)

		go func(record []string) {
			defer wg.Done()

			// Melakukan pemrosesan data pada record (misalnya, mengubah nama menjadi huruf besar dan menambahkan "Mr.")
			record[0] = strings.ToUpper(record[0])
			record[2] = "Mr. " + record[2]

			// Menulis record yang telah diproses ke file output
			err := writer.Write(record)
			if err != nil {
				fmt.Printf("Error writing CSV record: %v\n", err)
				return
			}
		}(record)
	}

	wg.Wait()

	writer.Flush()
	if err := writer.Error(); err != nil {
		fmt.Printf("Error flushing writer: %v\n", err)
		return
	}

	fmt.Println("File processing complete. Check output.csv for the results.")
}
