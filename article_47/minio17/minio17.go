// Seriál "Programovací jazyk Go"
//
// Čtyřicátá sedmá část
//
// Demonstrační příklad číslo 17:
//     Použití metody SetBucketLifecycle

package main

import (
	"flag"
	"github.com/minio/minio-go/v6"
	"log"
	"strconv"
)

func createBucket(minioClient *minio.Client, bucket string) {
	err := minioClient.MakeBucket(bucket, "us-east-1")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("New bucket has been created")
}

// <?xml version="1.0" encoding="UTF-8"?>
func setExpiration(minioClient *minio.Client, bucket string, days int) {
	rules := `
<LifecycleConfiguration>
    <Rule>
        <ID>expire-bucket</ID>
	<Prefix></Prefix>
	<Status>Enabled</Status>
	<Expiration>
	    <Days>` + strconv.Itoa(days) + `</Days>
	</Expiration>
    </Rule>
</LifecycleConfiguration>`

	println(rules)
	err := minioClient.SetBucketLifecycle(bucket, rules)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Expiration has been configured")
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

	//createBucket(minioClient, "temporary")
	setExpiration(minioClient, "temporary", 1)
}
