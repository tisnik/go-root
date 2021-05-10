// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Čtyřicátá šestá část
//     Projekt MinIO: jedna z nejužitečnějších aplikací naprogramovaných v Go
//     https://www.root.cz/clanky/projekt-minio-jedna-z-nejuzitecnejsich-aplikaci-naprogramovanych-v-go/#k19
//
// Demonstrační příklad číslo 9:
//     Poslání dat do Minia s jejich uložením do objektu

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

	length, err := minioClient.FPutObject("foo", "minio9.go", "minio9.go", minio.PutObjectOptions{
		ContentType: "text/plain;charset=UTF-8",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Successfully uploaded bytes: ", length)

	printObject(minioClient, "foo", "minio9.go")
}
