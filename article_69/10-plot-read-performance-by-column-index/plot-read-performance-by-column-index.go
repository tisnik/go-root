// Seriál "Programovací jazyk Go"
//
// Šedesátá devátá část
//    Komunikace se sloupcovými databázemi z jazyka Go: Parquet soubory (dokončení)
//    https://www.root.cz/clanky/komunikace-se-sloupcovymi-databazemi-z-jazyka-go-parquet-soubory-dokonceni/

// This tool is able to read all records stored in selected Parquet file.
// Currently, only records with the structure `Record` is read correctly. Name
// of input Parquet file needs to be selected from command line.
package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/xitongsys/parquet-go-source/local"
	"github.com/xitongsys/parquet-go/reader"
	"github.com/xitongsys/parquet-go/source"
)

// maximum block size for reading Parquet files by blocks
const maxBlockSize = 100

// Record represents one record stored in Parquet file
type Record struct {
	ID      uint64 `parquet:"name=id, type=UINT_64, encoding=PLAIN"`
	Name    string `parquet:"name=name, type=UTF8, encoding=PLAIN_DICTIONARY"`
	Surname string `parquet:"name=surname, type=UTF8, encoding=PLAIN"`
	Email   string `parquet:"name=email, type=UTF8, encoding=PLAIN"`
	Active  bool   `parquet:"name=active, type=BOOLEAN"`
	Color   string `parquet:"name=color, type=UTF8, encoding=PLAIN_DICTIONARY"`
}

const activeColumnIndex = 4

// closeReader tries to close the given Parquet file reader
func closeReader(reader source.ParquetFile) {
	err := reader.Close()
	if err != nil {
		log.Println("close reader:", err)
	}
}

func readParquetFile(fileName string, blockSize int, readers int) {
	// construct the file reader and try to open the Parquet file for
	// reading
	fileReader, err := local.NewLocalFileReader(fileName)

	if err != nil {
		log.Fatal("Can't open file", err)
		return
	}

	// fileReader needs to be closed properly
	defer closeReader(fileReader)

	// initializa Parquet file reader
	parquetColumnReader, err := reader.NewParquetColumnReader(fileReader, int64(readers))

	if err != nil {
		log.Fatal("Can't create parquet column reader", err)
		return
	}

	// parquetReader needs to be stopped
	defer parquetColumnReader.ReadStop()

	readColumnData(parquetColumnReader, activeColumnIndex, blockSize)

}

func readColumnData(parquetReader *reader.ParquetReader, columnIndex int, blockSize int) {
	valuesCount := int(parquetReader.GetNumRows())

	activeCount := 0
	inactiveCount := 0

	readValues := 0

	// try to read and display all records
	for readValues < valuesCount {
		// try to read value
		data, _, _, err := parquetReader.ReadColumnByIndex(int64(columnIndex), int64(blockSize))
		if err != nil {
			log.Println("Read error", err)
			continue
		} else {
			readValues += len(data)
		}
		for _, active := range data {
			if active.(bool) {
				activeCount++
			} else {
				inactiveCount++
			}
		}
	}
	log.Println("Read", readValues, "values", "active", activeCount, "inactive", inactiveCount)
}

func main() {
	var readers []int = []int{1, 8, 16, 32}

	for _, numReaders := range readers {
		// create and open new CSV file
		csvFileName := fmt.Sprintf("durations-%d-readers.csv", numReaders)

		csvFile, err := os.Create(csvFileName)
		if err != nil {
			log.Fatal("Create CSV file", err)
		}

		defer csvFile.Close()

		// initialize CSV writer
		csvWriter := csv.NewWriter(csvFile)
		defer csvWriter.Flush()

		csvWriter.Write([]string{"Block size", "Time to read"})

		for blockSize := 1; blockSize <= maxBlockSize; blockSize++ {
			t1 := time.Now()

			readParquetFile("1000000records_compression_none.parquet", blockSize, numReaders)

			// compute and print duration
			since := time.Since(t1)
			log.Printf("Block size: %d  Readers: %d  Duration: %d\n", blockSize, numReaders, since)

			// write duration into CSV file
			csvWriter.Write([]string{strconv.Itoa(blockSize), strconv.Itoa(int(since))})
		}
	}
}
