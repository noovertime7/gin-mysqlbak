package main

import (
	"fmt"
	"github.com/noovertime7/gin-mysqlbak/public/ding"
	"github.com/noovertime7/mysqlbak/pkg/log"
)

func main() {
	content := fmt.Sprintf(
		"# %s/%s备份状态\n - 测试  - 测试", "127.0.0.1", "mysql")
	webhook := ding.Webhook{AtAll: true, Secret: "SEC72586e3f7ff6db4b2ad24eac905f308a9ddb0b1b9809af31e5623a14abb424b2", AccessToken: "77f579efbefeefc316b55d3caea1ba1963db2f1319aa7520cbfd9626de073fdc"}
	if err := webhook.SendMarkDown(content, "备份状态"); err != nil {
		log.Logger.Error("钉钉消息发送失败", err)
		return
	}
}
