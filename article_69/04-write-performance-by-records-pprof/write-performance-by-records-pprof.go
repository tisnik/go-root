// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Šedesátá devátá část
//    Komunikace se sloupcovými databázemi z jazyka Go: Parquet soubory (dokončení)
//    https://www.root.cz/clanky/komunikace-se-sloupcovymi-databazemi-z-jazyka-go-parquet-soubory-dokonceni/

package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	"runtime/pprof"

	"github.com/bxcodec/faker/v3"
	"github.com/xitongsys/parquet-go/parquet"
	"github.com/xitongsys/parquet-go/writer"
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

func generateColor() string {
	var colors []string = []string{
		"black",
		"blue",
		"red",
		"magenta",
		"green",
		"cyan",
		"yellow",
		"white",
	}
	return colors[rand.Int()%len(colors)]
}

func writeRecords(pw *writer.ParquetWriter, n int) {
	// create report structure to be stored in Parquet file
	record := Record{}

	for i := 0; i < n; i++ {
		record.ID = uint64(i)
		record.Name = faker.FirstName()
		record.Surname = faker.LastName()
		record.Email = faker.Email()
		record.Active = i%2 == 0
		record.Color = generateColor()

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
	t1 := time.Now()

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

	// compute and print duration
	t2 := time.Now()
	since := time.Since(t1)
	log.Println("Start time: ", t1)
	log.Println("End time:   ", t2)
	log.Println("Duration:   ", since)
}

func main() {
	f, err := os.Create("write-performance.prof")
	if err != nil {
		log.Fatalf("failed to create profiler output file: %v", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalf("failed to close profiler file: %v", err)
		}
	}()

	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatalf("failed to start profle: %v", err)
	}
	defer pprof.StopCPUProfile()

	createAndWriteIntoParquetFile("1000000records_compression_none.parquet", 1000, parquet.CompressionCodec_UNCOMPRESSED)
	createAndWriteIntoParquetFile("1000000records_compression_snappy.parquet", 1000000, parquet.CompressionCodec_SNAPPY)
	createAndWriteIntoParquetFile("1000000records_compression_gzip.parquet", 1000000, parquet.CompressionCodec_GZIP)
}
