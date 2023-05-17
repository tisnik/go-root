// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Čtyřicátá šestá část
//     Projekt MinIO: jedna z nejužitečnějších aplikací naprogramovaných v Go
//     https://www.root.cz/clanky/projekt-minio-jedna-z-nejuzitecnejsich-aplikaci-naprogramovanych-v-go/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů ze čtyřicáté šesté části:
//    https://github.com/tisnik/go-root/blob/master/article_46/README.md
//
// Demonstrační příklad číslo 7:
//    Přečtení obsahu objektu s textem

package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/minio/minio-go/v6"
	"log"
)

func printObject(minioClient *minio.Client, bucket string, objectName string) {
	object, err := minioClient.GetObject(bucket, objectName, minio.GetObjectOptions{})
	if err != nil {
		log.Fatalln(err)
	}
	defer object.Close()

	scanner := bufio.NewScanner(object)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
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

	printObject(minioClient, "foo", "t.go")
}
