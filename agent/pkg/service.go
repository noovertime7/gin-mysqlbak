package pkg

import (
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/noovertime7/gin-mysqlbak/agent/agentdto"
	"github.com/noovertime7/gin-mysqlbak/agent/proto"
)

var s micro.Service
var reg = etcd.NewRegistry(registry.Addrs("127.0.0.1:2379"))

func init() {
	s = micro.NewService(
		micro.Registry(reg))
	s.Init()
}

func GetMicroService(serviceName string) interface{} {
	return proto.NewHostService(fmt.Sprintf("%s", serviceName), s.Client())
}

func GetServiceList() agentdto.AgentOutPut {
	services, _ := reg.ListServices()
	var AgentOutPutItems []agentdto.AgentOutPutItem
	for _, s := range services {
		nodes := s.Nodes
		for _, node := range nodes {
			item := agentdto.AgentOutPutItem{
				Name:      s.Name,
				Address:   node.Address,
				ServiceID: node.Id,
			}
			AgentOutPutItems = append(AgentOutPutItems, item)
		}
	}
	total := len(AgentOutPutItems)
	return agentdto.AgentOutPut{Total: total, AgentOutPutItem: AgentOutPutItems}
}
