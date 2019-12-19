// Seriál "Programovací jazyk Go"
//
// Čtyřicátá sedmá část
//
// Demonstrační příklad číslo 11:
//     Použití metody StatObject

package main

import (
	"flag"
	"fmt"
	"github.com/minio/minio-go/v6"
	"log"
)

func listBuckets(minioClient *minio.Client) {
	fmt.Println("List of buckets:")

	buckets, err := minioClient.ListBuckets()
	if err != nil {
		log.Fatalln(err)
		return
	}
	for i, bucket := range buckets {
		fmt.Printf("%d\t%+v\n", i, bucket)
	}
}

func getObjectInfo(minioClient *minio.Client, bucket string, name string) {
	fmt.Printf("Info for object %s in bucket %s\n", name, bucket)

	info, err := minioClient.StatObject(bucket, name, minio.StatObjectOptions{})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Key:          %s\nSize:         %d\nTag:          %s\nContent-type: %s\n", info.Key, info.Size, info.ETag, info.ContentType)
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

	listBuckets(minioClient)
	fmt.Println()
	getObjectInfo(minioClient, "foo", "minio9.go")
	fmt.Println()
	getObjectInfo(minioClient, "foo", "logos.jpg")
	fmt.Println()
	getObjectInfo(minioClient, "foo", "something_else")
}
