// Seriál "Programovací jazyk Go"
//
// Čtyřicátá šestá část
//
// Demonstrační příklad číslo 4:
//     Výpis všech objektů ve zvoleném bucketu

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

func listObjects(minioClient *minio.Client, bucket string) {
	fmt.Println("List of objects for bucket:", bucket)

	done := make(chan struct{})
	defer close(done)

	objects := minioClient.ListObjects(bucket, "", false, done)
	for object := range objects {
		if object.Err != nil {
			log.Println(object.Err)
			return
		}
		fmt.Println(object)
	}
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
	listObjects(minioClient, "foo")
}
