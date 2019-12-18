package main

import (
	"flag"
	"fmt"
	"github.com/minio/minio-go/v6"
	"log"
)

func printBucketPolicy(minioClient *minio.Client, bucket string) {
	fmt.Printf("Policy for bucket: %s\n", bucket)
	policy, err := minioClient.GetBucketPolicy(bucket)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(policy)
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

	policyFoo := `{"Version": "2012-10-17","Statement": [{"Effect":"Allow","Principal":{"AWS": ["*"]},"Resource": ["arn:aws:s3:::foo/*"],"Sid": "", "Action":["s3:PutObject"]}]}`
	err = minioClient.SetBucketPolicy("foo", policyFoo)
	if err != nil {
		fmt.Println(err)
		return
	}

	policyReadOnly := `{"Version": "2012-10-17","Statement": [{"Effect":"Allow","Principal":{"AWS": ["*"]},"Resource": ["arn:aws:s3:::readonly/*"],"Sid": "", "Action":["s3:GetObject","s3:PutObject"]}]}`
	err = minioClient.SetBucketPolicy("readonly", policyReadOnly)
	if err != nil {
		fmt.Println(err)
		return
	}

	printBucketPolicy(minioClient, "foo")
	printBucketPolicy(minioClient, "readonly")
	printBucketPolicy(minioClient, "writeonly")
	printBucketPolicy(minioClient, "readwrite")
}
