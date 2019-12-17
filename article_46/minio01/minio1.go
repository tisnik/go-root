// Seriál "Programovací jazyk Go"
//
// Čtyřicátá šestá část
//
// Demonstrační příklad číslo 1:
//     Kostra aplikace, která provede inicializaci klienta služby Minio

package main

import (
	"github.com/minio/minio-go/v6"
	"log"
)

func main() {
	endpoint := "127.0.0.1:9000"

	// it is needed to change the following two keys
	accessKeyID := "3V8WMANF061SGOIVR7AA"
	secretAccessKey := "AHTM6+74n1Z8DZRZ4V7o83QcnYRnTEVblVb8sIlE"

	useSSL := true

	// initialize minio client object
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalln(err)
	}

	// everything seems to be ok
	log.Printf("%#v\n", minioClient)
}
