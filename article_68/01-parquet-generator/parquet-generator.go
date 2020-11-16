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
	Id      int64  `parquet:"name=id, type=INT64"`
	Name    string `parquet:"name=name, type=UTF8, encoding=PLAIN"`
	Surname string `parquet:"name=surname, type=UTF8, encoding=PLAIN"`
	Active  bool   `parquet:"name=active, type=BOOLEAN"`
	Remark  string `parquet:"name=remark, type=UTF8, encoding=PLAIN"`
}

func writeRecord(pw *writer.ParquetWriter) {
	// create report structure to be stored in Parquet file
	record := Record{
		Id:      int64(1),
		Name:    "Pepa",
		Surname: "Vyskoč",
		Active:  false,
		Remark:  "foo bar baz",
	}

	// write the record structure into Parquet file
	err := pw.Write(record)
	if err != nil {
		log.Println("Write into Parquet error", err)
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

func main() {
	w, err := os.Create("flat.parquet")
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
	pw.CompressionType = parquet.CompressionCodec_SNAPPY

	defer stopWrite(pw)

	writeRecord(pw)

	log.Println("Write Finished")
}
