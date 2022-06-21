package public

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"github.com/noovertime7/mysqlbak/pkg/log"
	"io"
	"os"
)

func GenSaltPassword(salt, password string) string {
	s1 := sha256.New()
	s1.Write([]byte(password))
	str1 := fmt.Sprintf("%x", s1.Sum(nil))
	s2 := sha256.New()
	s2.Write([]byte(str1 + salt))
	str2 := fmt.Sprintf("%x", s2.Sum(nil))
	return str2
}

//MD5 md5加密
func MD5(s string) string {
	h := md5.New()
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func GetFileSize(fileanme string) int {
	fileInfo, err := os.Stat(fileanme)
	if err != nil {
		return 0
	}
	tmp := int(fileInfo.Size()) / 1024
	return tmp
}

//创建文件夹
func CreateDir(path string) {
	_exist, _err := HasDir(path)
	if _err != nil {
		log.Logger.Errorf("获取文件夹异常 -> %v\n", _err)
		return
	}
	if _exist {
		return
	} else {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Logger.Errorf("创建目录异常 -> %v\n", err)
		} else {
			log.Logger.Infof("创建文件夹%s成功!", path)
		}
	}
}

func HasDir(path string) (bool, error) {
	_, _err := os.Stat(path)
	if _err == nil {
		return true, nil
	}
	if os.IsNotExist(_err) {
		return false, nil
	}
	return false, _err
}
