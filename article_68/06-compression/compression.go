// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Šedesátá osmá část
//    Komunikace se sloupcovými databázemi z jazyka Go: Parquet soubory
//    https://www.root.cz/clanky/komunikace-se-sloupcovymi-databazemi-z-jazyka-go-parquet-soubory/

package main

import (
	"log"
	"os"

	"github.com/xitongsys/parquet-go/parquet"
	"github.com/xitongsys/parquet-go/writer"
)

const defaultOutputFile = "flat.parquet"

// Record represents one record stored in Parquet file
type Record struct {
	Active bool `parquet:"name=active, type=BOOLEAN"`
}

func writeRecords(pw *writer.ParquetWriter, n int) {
	// create report structure to be stored in Parquet file
	record := Record{
		Active: false,
	}

	for i := 0; i < n; i++ {
		record.Active = i%2 == 0
		// write the record structure into Parquet file
		err := pw.Write(record)
		if err != nil {
			log.Println("Write into Parquet error", err)
		}
	}
}

// stopWrite function stop writing into Parquet file
func stopWrite(pw *writer.ParquetWriter) {
	err := pw.WriteStop()

	// most write errors are caught at this time
	if err != nil {
		log.Println("WriteStop error", err)
	}
}

func createAndWriteIntoParquetFile(filename string, records int, compression parquet.CompressionCodec) {
	w, err := os.Create(filename)
	if err != nil {
		log.Println("Can't create local file", err)
		return
	}

	defer w.Close()

	// initialize Parquet file writer
	pw, err := writer.NewParquetWriterFromWriter(w, new(Record), 1)
	if err != nil {
		log.Println("Can't create parquet writer", err)
		return
	}

	pw.RowGroupSize = 128 * 1024 * 1024 //128M
	pw.CompressionType = compression

	defer stopWrite(pw)

	writeRecords(pw, records)

	log.Println("Write Finished")
}

func main() {
	createAndWriteIntoParquetFile("1000000records_compression_none.parquet", 1000000, parquet.CompressionCodec_UNCOMPRESSED)
	createAndWriteIntoParquetFile("1000000records_compression_snappy.parquet", 1000000, parquet.CompressionCodec_SNAPPY)
	createAndWriteIntoParquetFile("1000000records_compression_gzip.parquet", 1000000, parquet.CompressionCodec_GZIP)
}
