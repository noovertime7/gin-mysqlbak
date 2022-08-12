package pkg

import (
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-plugins/wrapper/breaker/hystrix/v2"
	"github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/noovertime7/gin-mysqlbak/agent/agentdto"
	"github.com/noovertime7/gin-mysqlbak/agent/pkg/trace"
	"github.com/noovertime7/gin-mysqlbak/agent/proto"
	"log"
)

var s micro.Service
var reg = etcd.NewRegistry(registry.Addrs("127.0.0.1:2379"))

const JaegerAddr = "127.0.0.1:6831"

func init() {

}

func GetMicroService(serviceName string) interface{} {
	// 配置jaeger连接
	jaegerTracer, closer, err := trace.NewJaegerTracer(serviceName, JaegerAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer closer.Close()
	// 配置jaeger连接
	s = micro.NewService(
		micro.Registry(reg),
		micro.WrapClient(
			hystrix.NewClientWrapper(),
			opentracing.NewClientWrapper(jaegerTracer)),
	)
	s.Init()
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
