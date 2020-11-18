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

func readParquetFile(fileName string, blockSize int) {
	t1 := time.Now()

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
	parquetReader, err := reader.NewParquetReader(fileReader, new(Record), 1)

	if err != nil {
		log.Fatal("Can't create parquet reader", err)
		return
	}

	// parquetReader needs to be stopped
	defer parquetReader.ReadStop()

	readRecords(parquetReader, blockSize)

	// compute and print duration
	since := time.Since(t1)
	log.Printf("Block size: %d  Duration: %d", blockSize, since)
}

func readRecords(parquetReader *reader.ParquetReader, blockSize int) {
	recordCount := int(parquetReader.GetNumRows())
	log.Println("Records to read", recordCount)

	records := make([]Record, blockSize)
	readRecords := 0

	// try to read and display all records
	for readRecords < recordCount {
		// try to read record
		err := parquetReader.Read(&records)
		if err != nil {
			log.Println("Read error", err)
			continue
		} else {
			readRecords += len(records)
		}
	}
	// log.Println("Read", readRecords, "records")
}

func main() {
	readParquetFile("1000000records_compression_none.parquet", 1)
	readParquetFile("1000000records_compression_none.parquet", 100)
	readParquetFile("1000000records_compression_none.parquet", 1000)
	//readParquetFile("1000000records_compression_snappy.parquet", 1)
	//readParquetFile("1000000records_compression_gzip.parquet", 1)
}
