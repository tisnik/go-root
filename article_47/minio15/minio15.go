package main

import (
	"flag"
	"fmt"
	"github.com/minio/minio-go/v6"
	"log"
)

func createBucket(minioClient *minio.Client, bucket string) {
	err := minioClient.MakeBucket(bucket, "us-east-1")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("New bucket has been created")
}

func enableVersioning(minioClient *minio.Client, bucket string) {
	err := minioClient.EnableVersioning(bucket)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Versioning has been enabled")
}

func uploadFile(minioClient *minio.Client, bucket string, filename string) {
	length, err := minioClient.FPutObject(bucket, filename, filename, minio.PutObjectOptions{
		ContentType: "text/plain;charset=UTF-8",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Successfully uploaded bytes: ", length)
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

	createBucket(minioClient, "versioned")
	enableVersioning(minioClient, "versioned")

	for i := 1; i < 10; i++ {
		uploadFile(minioClient, "versioned", "minio13.go")
	}
}
