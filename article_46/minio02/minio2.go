// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Čtyřicátá šestá část
//     Projekt MinIO: jedna z nejužitečnějších aplikací naprogramovaných v Go
//     https://www.root.cz/clanky/projekt-minio-jedna-z-nejuzitecnejsich-aplikaci-naprogramovanych-v-go/#k19
//
// Demonstrační příklad číslo 2:
//     Zadání parametrů připojení z příkazového řádku

package main

import (
	"flag"
	"github.com/minio/minio-go/v6"
	"log"
)

func main() {
	var endpoint = flag.String("endpoint", "127.0.0.1:9000", "MinIO service endpoint")
	var accessKeyID = flag.String("accessKeyID", "", "Access key ID for MinIO")
	var secretAccessKey = flag.String("secretAccessKey", "", "Secret access key for MinIO")
	var useSSL = flag.Bool("useSSL", false, "Use SSL for communication with MinIO")
	flag.Parse()

	// initialize minio client object
	minioClient, err := minio.New(*endpoint, *accessKeyID, *secretAccessKey, *useSSL)
	if err != nil {
		log.Fatalln(err)
	}

	// everything seems to be ok
	log.Printf("%#v\n", minioClient)
}
