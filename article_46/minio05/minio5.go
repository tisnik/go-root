// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Čtyřicátá šestá část
//     Projekt MinIO: jedna z nejužitečnějších aplikací naprogramovaných v Go
//     https://www.root.cz/clanky/projekt-minio-jedna-z-nejuzitecnejsich-aplikaci-naprogramovanych-v-go/#k19
//
// Demonstrační příklad číslo 5:
//     Podrobnější informace o objektech ve zvoleném bucketu

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

func listObjects(minioClient *minio.Client, bucket string, prefix string) {
	fmt.Println("List of objects for bucket:", bucket)

	done := make(chan struct{})
	defer close(done)

	objects := minioClient.ListObjects(bucket, prefix, false, done)
	for object := range objects {
		if object.Err != nil {
			log.Println(object.Err)
			return
		}
		fmt.Printf("Key: %s,  Size: %d,  Tag: %s\n", object.Key, object.Size, object.ETag)
	}
}

func main() {
	var endpoint = flag.String("endpoint", "127.0.0.1:9000", "MinIO service endpoint")
	var accessKeyID = flag.String("accessKeyID", "", "Access key ID for MinIO")
	var secretAccessKey = flag.String("secretAccessKey", "", "Secret access key for MinIO")
	var useSSL = flag.Bool("useSSL", false, "Use SSL for communication with MinIO")
	var objectPrefix = flag.String("prefix", "", "Prefix for objects to be listed")
	flag.Parse()

	// initialize minio client object
	minioClient, err := minio.New(*endpoint, *accessKeyID, *secretAccessKey, *useSSL)
	if err != nil {
		log.Fatalln(err)
	}

	// everything seems to be ok
	log.Printf("%#v\n", minioClient)

	listBuckets(minioClient)
	listObjects(minioClient, "foo", *objectPrefix)
}
