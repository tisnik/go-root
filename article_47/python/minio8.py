import argparse

from minio import Minio
from minio.error import ResponseError


def print_object(minioClient, bucket, object_name):
    try:
        data = minioClient.get_object(bucket, object_name)
        for line in data.readlines():
            print(line.decode('utf-8')[:-1])
    except ResponseError as err:
        print(err)


parser = argparse.ArgumentParser()
parser.add_argument("--endpoint", default="127.0.0.1:9000",
                    help="MinIO service endpoint")
parser.add_argument("--accessKeyID", default="",
                    help="Access key ID for MinIO")
parser.add_argument("--secretAccessKey", default="",
                    help="Secret access key for MinIO")
parser.add_argument("--enable-ssl", dest="useSSL", action="store_true",
                    help="Use SSL for communication with MinIO")
parser.add_argument("--disable-ssl", dest="useSSL", action="store_false",
                    help="Don't SSL for communication with MinIO")
args = parser.parse_args()

minioClient = Minio(args.endpoint,
                    access_key=args.accessKeyID,
                    secret_key=args.secretAccessKey,
                    secure=args.useSSL)

print_object(minioClient, "foo", "t.go")
