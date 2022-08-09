package minio

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
	"path/filepath"
)

type client struct {
	client         *minio.Client
	bucketName     string
	targetFilePath string
	Dir            string
}

func NewClient(endpoint, accessKeyID, secretAccessKey, bucketName, dir, filepath string) *client {
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: false,
	})
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	return &client{
		client:         minioClient,
		bucketName:     bucketName, //目标bucket
		targetFilePath: filepath,
		Dir:            dir,
	}
}

//checkBucket 检测目标bucket是否存在，不存在就创建一个
func (c *client) checkBucket() {
	isExists, err := c.client.BucketExists(context.Background(), c.bucketName)
	if err != nil {
		log.Println("check bucket exist error ")
		return
	}
	if !isExists {
		err2 := c.client.MakeBucket(context.Background(), c.bucketName, minio.MakeBucketOptions{Region: "cn-north-1", ObjectLocking: false})
		if err2 != nil {
			log.Println("MakeBucket error ")
			fmt.Println(err2)
			return
		}
		log.Printf("minio创建bucket %s\n", c.bucketName)
	}
}

func (c *client) UploadFile() error {
	c.checkBucket()
	_, filename := filepath.Split(c.targetFilePath)
	dirname := "/" + c.Dir + "/" + filename
	_, err := c.client.FPutObject(context.Background(), c.bucketName, dirname, c.targetFilePath, minio.PutObjectOptions{})
	if err != nil {
		return err
	}
	return nil
}
