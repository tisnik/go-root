// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Šedesátá osmá část
//    Komunikace se sloupcovými databázemi z jazyka Go: Parquet soubory
//    https://www.root.cz/clanky/komunikace-se-sloupcovymi-databazemi-z-jazyka-go-parquet-soubory/

// This tool is able to read all records stored in selected Parquet file.
// Currently, only records with the structure `Record` is read correctly. Name
// of input Parquet file needs to be selected from command line.
package main

import (
	"log"

	"github.com/xitongsys/parquet-go-source/local"
	"github.com/xitongsys/parquet-go/reader"
	"github.com/xitongsys/parquet-go/source"
)

const defaultInputFile = "flat.parquet"

// Record represents one record stored in Parquet file
type Record struct {
	Id      int64  `parquet:"name=id, type=INT64"`
	Name    string `parquet:"name=name, type=UTF8, encoding=PLAIN"`
	Surname string `parquet:"name=surname, type=UTF8, encoding=PLAIN"`
	Active  bool   `parquet:"name=active, type=BOOLEAN"`
	Remark  string `parquet:"name=remark, type=UTF8, encoding=PLAIN"`
}

// closeReader tries to close the given Parquet file reader
func closeReader(reader source.ParquetFile) {
	err := reader.Close()
	if err != nil {
		log.Println("close reader:", err)
	}
}

func displayContentOfParquetFile(fileName string) {
	const parallelNumber = 4

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

	displayRecords(parquetReader)
}

// displayRecords function lists all records from Parquet file
func displayRecords(parquetReader *reader.ParquetReader) {
	recordCount := int(parquetReader.GetNumRows())

	// try to read and display all records
	for i := 0; i < recordCount; i++ {
		record := make([]Record, 1)

		// try to read record
		err := parquetReader.Read(&record)
		if err != nil {
			log.Println("Read error", err)
		} else {
			// and display it
			log.Println(record)
		}
	}
}

func main() {
	displayContentOfParquetFile(defaultInputFile)
}
