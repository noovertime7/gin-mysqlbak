package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	"github.com/noovertime7/gin-mysqlbak/agent/proto/bakhistory"
	"github.com/noovertime7/mysqlbak/pkg/log"
)

const (
	address = "localhost:39800"
)

func main() {
	//conn, err := grpc.Dial(address, grpc.WithInsecure()) //建立客户端和服务器之间的链接
	//if err != nil {
	//	log.Fatalf("did not connect %v", err)
	//}
	//defer conn.Close()
	service := micro.NewService()
	service.Init()
	var ops client.CallOption = func(options *client.CallOptions) {
		options.Address = []string{"127.0.0.1:30000"}
	}
	h := bakhistory.NewHistoryService("test5.local", service.Client())
	historyListInput := &bakhistory.HistoryListInput{
		Info:     "",
		PageNo:   1,
		PageSize: 10,
		Sort:     "",
	}

	for i := 0; i < 10; i++ {
		data, err := h.GetHistoryList(context.Background(), historyListInput, ops)
		if err != nil {
			log.Logger.Error("agent获取历史记录列表失败", err)
			return
		}
		fmt.Println(data)
	}

}
