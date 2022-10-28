package pkg

import (
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
)

func DecryptDir(key, dir, fileType string) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}
	if len(files) == 0 {
		return errors.New("Directory is empty")
	}
	for _, file := range files {
		filePath := dir + "/" + file.Name()
		data, err := Decrypt(key, filePath, fileType)
		if err != nil {
			fmt.Printf("decrypt %v err:%v continue\n", file.Name(), err)
			continue
		}
		fmt.Printf("decrypt %v success, decrypt file : %v\n", file.Name(), data)
	}
	return nil
}
