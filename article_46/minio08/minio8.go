// Seriál "Programovací jazyk Go"
//
// Čtyřicátá šestá část
//
// Demonstrační příklad číslo 8:
//     Alternativní způsob přístupu k obsahu objektu

package main

import (
	"flag"
	"fmt"
	"github.com/minio/minio-go/v6"
	"io/ioutil"
	"log"
)

func printObject(minioClient *minio.Client, bucket string, objectName string) {
	object, err := minioClient.GetObject(bucket, objectName, minio.GetObjectOptions{})
	if err != nil {
		log.Fatalln(err)
	}
	defer object.Close()

	bytes, err := ioutil.ReadAll(object)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(bytes))
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

	printObject(minioClient, "foo", "t.go")
}
