// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Šedesátá devátá část
//    Komunikace se sloupcovými databázemi z jazyka Go: Parquet soubory (dokončení)
//    https://www.root.cz/clanky/komunikace-se-sloupcovymi-databazemi-z-jazyka-go-parquet-soubory-dokonceni/

// This tool is able to read all records stored in selected Parquet file.
// Currently, only records with the structure `Record` is read correctly. Name
// of input Parquet file needs to be selected from command line.
package main

import (
	"log"
	"time"

	"github.com/xitongsys/parquet-go-source/local"
	"github.com/xitongsys/parquet-go/reader"
	"github.com/xitongsys/parquet-go/source"
)

// Record represents one record stored in Parquet file
type Record struct {
	ID      uint64 `parquet:"name=id, type=UINT_64, encoding=PLAIN"`
	Name    string `parquet:"name=name, type=UTF8, encoding=PLAIN_DICTIONARY"`
	Surname string `parquet:"name=surname, type=UTF8, encoding=PLAIN"`
	Email   string `parquet:"name=email, type=UTF8, encoding=PLAIN"`
	Active  bool   `parquet:"name=active, type=BOOLEAN"`
	Color   string `parquet:"name=color, type=UTF8, encoding=PLAIN_DICTIONARY"`
}

// closeReader tries to close the given Parquet file reader
func closeReader(reader source.ParquetFile) {
	err := reader.Close()
	if err != nil {
		log.Println("close reader:", err)
	}
}

func readParquetFile(fileName string) {
	t1 := time.Now()

	const parallelNumber = 1

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
	parquetReader, err := reader.NewParquetReader(fileReader, new(Record),
		parallelNumber)

	if err != nil {
		log.Fatal("Can't create parquet reader", err)
		return
	}

	// parquetReader needs to be stopped
	defer parquetReader.ReadStop()

	readRecords(parquetReader)

	// compute and print duration
	t2 := time.Now()
	since := time.Since(t1)
	log.Println("Start time: ", t1)
	log.Println("End time:   ", t2)
	log.Println("Duration:   ", since)
}

func readRecords(parquetReader *reader.ParquetReader) {
	recordCount := int(parquetReader.GetNumRows())
	log.Println("Records to read", recordCount)

	record := make([]Record, 1)
	records := 0

	// try to read and display all records
	for i := 0; i < recordCount; i++ {
		// try to read record
		err := parquetReader.Read(&record)
		if err != nil {
			log.Println("Read error", err)
			continue
		} else {
			records++
		}
	}
	log.Println("Read", records, "records")
}

func main() {
	readParquetFile("1000000records_compression_none.parquet")
	readParquetFile("1000000records_compression_snappy.parquet")
	readParquetFile("1000000records_compression_gzip.parquet")
}
