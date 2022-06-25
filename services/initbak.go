package services

import (
	"fmt"
	"github.com/e421083458/golang_common/lib"
	"github.com/noovertime7/mysqlbak/pkg/log"
	"net/http"
)

func InitBak() {
	var url string
	addrlenth := len(lib.GetStringConf("base.http.addr"))
	if addrlenth <= 6 {
		url = fmt.Sprintf("http://127.0.0.1%s/public/initbak", lib.GetStringConf("base.http.addr"))
		if _, err := http.Get(url); err != nil {
			log.Logger.Error(err)
			return
		}
		return
	} else if addrlenth <= 21 && addrlenth > 6 {
		url = fmt.Sprintf("http://%s/public/initbak", lib.GetStringConf("base.http.addr"))
		if _, err := http.Get(url); err != nil {
			log.Logger.Error(err)
			return
		}
		return
	}
	log.Logger.Fatalf("您输入的IP地址 %s 格式有误，请重新输入", lib.GetStringConf("base.http.addr"))
}
