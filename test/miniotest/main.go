package main

import "github.com/noovertime7/gin-mysqlbak/public/minio"

func main() {
	client := minio.NewClient("10.20.110.61:9000", "minioadmin", "minioadmin", "mysqlbak", "ceshi", "G:\\gin-mysql\\public\\minio\\test.txt")
	client.UploadFile()
}
