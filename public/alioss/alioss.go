package alioss

import (
	"github.com/noovertime7/mysqlbak/pkg/log"
	"strings"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func AliOssUploadFile(filename, Endpoint, Accesskey, Secretkey, BucketName, Directory string) error {
	// 创建OSSClient实例。
	// yourEndpoint填写Bucket对应的Endpoint，以华东1（杭州）为例，填写为https://oss-cn-hangzhou.aliyuncs.com。其它Region请按实际情况填写。
	// 阿里云账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM用户进行API访问或日常运维，请登录RAM控制台创建RAM用户。

	client, err := oss.New(Endpoint, Accesskey, Secretkey)
	if err != nil {
		log.Logger.Error("Error:", err)
		return err
	}

	// 填写存储空间名称，例如examplebucket。
	bucket, err := client.Bucket(BucketName)
	if err != nil {
		log.Logger.Error("Error:", err)
		return err
	}

	// 依次填写Object的完整路径（例如exampledir/exampleobject.txt）和本地文件的完整路径（例如D:\\localpath\\examplefile.txt）。
	file := strings.Split(filename, "/")[len(strings.Split(filename, "/"))-1] //需要处理一下拿到文件名
	err = bucket.PutObjectFromFile(Directory+file, filename)
	if err != nil {
		log.Logger.Error("Error:", err)
		return err
	}
	log.Logger.Info("阿里云对象存储OSS上传成功")
	return nil
}
