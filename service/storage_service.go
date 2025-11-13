package service

import (
	"context"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// create connection for minIO

type StorageService struct {
	Client     *minio.Client
	BucketName string
}

func NewStorageService() *StorageService {

	endpoints := "localhost:9000"
	accessKey := "minioadmin"
	secretKey := "minioadmin"
	bucketName := "uploads"

	// initilizing minio client

	minioClient, err := minio.New(endpoints, &minio.Options{

		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: false,
	})
	if err != nil {
		log.Fatal("could not initilize minio client")
	}

	//checking if bucket exist or not ? create if not
	ctx := context.Background()
	exist, err := minioClient.BucketExists(ctx, bucketName)
	if err != nil {
		log.Fatal("Error while checking bucket")
	}

	if !exist {
		err := minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
		if err != nil {
			log.Fatal("error while creating bucket")
		}
		log.Println("Bucket is Created ! ")
	} else {
		log.Println("Bucket is already exist")

	}

	return &StorageService{
		Client:     minioClient,
		BucketName: bucketName,
	}

}

//  Generate the temporary link for user to upload the files
