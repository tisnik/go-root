// Seriál "Programovací jazyk Go"
//
// Čtyřicátá sedmá část
//
// Demonstrační příklad číslo 12:
//     Použití metody ComposeObject

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

func composeObject(minioClient *minio.Client, bucket string, from1 string, from2 string, to string) {
	src1 := minio.NewSourceInfo(bucket, from1, nil)
	src2 := minio.NewSourceInfo(bucket, from2, nil)
	sources := []minio.SourceInfo{src1, src2}

	dst, err := minio.NewDestinationInfo(bucket, to, nil, nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = minioClient.ComposeObject(dst, sources)
	if err != nil {
		log.Fatalln(err)
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

	printObject(minioClient, "foo", "minio9.go")
	printObject(minioClient, "foo", "minio10.go")
	composeObject(minioClient, "foo", "minio9.go", "minio10.go", "minio9_10.go")
	printObject(minioClient, "foo", "minio9_10.go")
}
