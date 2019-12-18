from minio import Minio
from minio.error import ResponseError

endpoint = "127.0.0.1:9000"
accessKeyID = "3V8WMANF061SGOIVR7AA"
secretAccessKey = "AHTM6+74n1Z8DZRZ4V7o83QcnYRnTEVblVb8sIlE"
useSSL = True

minioClient = Minio(endpoint,
                    access_key=accessKeyID,
                    secret_key=secretAccessKey,
                    secure=useSSL)

print(minioClient)
