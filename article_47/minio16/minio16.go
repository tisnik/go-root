// Seriál "Programovací jazyk Go"
//
// Čtyřicátá sedmá část
//     Další možnosti nabízené projektem MinIO
//     https://www.root.cz/clanky/dalsi-moznosti-nabizene-projektem-minio/
//
// Demonstrační příklad číslo 16:
//     Použití metody GetBucketLifecycle

package main

import (
	"flag"
	"fmt"
	"github.com/minio/minio-go/v6"
	"log"
)

func printBucketLifecycle(minioClient *minio.Client, bucket string) {
	fmt.Printf("Lifecycle for bucket: %s\n", bucket)
	lifecycle, err := minioClient.GetBucketLifecycle(bucket)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(lifecycle)
	fmt.Println()
}

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

	printBucketLifecycle(minioClient, "foo")
	printBucketLifecycle(minioClient, "readonly")
	printBucketLifecycle(minioClient, "writeonly")
	printBucketLifecycle(minioClient, "readwrite")
}
