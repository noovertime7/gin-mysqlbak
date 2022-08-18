package pkg

import (
	"context"
	"errors"
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/noovertime7/gin-mysqlbak/agent/agentdto"
	"github.com/noovertime7/gin-mysqlbak/agent/pkg/trace"
	"github.com/noovertime7/gin-mysqlbak/agent/proto/bak"
	"github.com/noovertime7/gin-mysqlbak/agent/proto/bakhistory"
	"github.com/noovertime7/gin-mysqlbak/agent/proto/host"
	"github.com/noovertime7/gin-mysqlbak/agent/proto/task"
	"github.com/noovertime7/gin-mysqlbak/conf"
	"log"
)

var s micro.Service
var reg = etcd.NewRegistry(registry.Addrs(conf.GetStringConf("etcd", "addr")))

var JaegerAddr = conf.GetStringConf("jaeger", "addr")

// log wrapper logs every time a request is made
type hystrixWrapper struct {
	client.Client
}

// Call 熔断器的使用，超过1秒熔断
func (h *hystrixWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	//md, _ := metadata.FromContext(ctx)
	name := req.Service() + "." + req.Endpoint()
	config := hystrix.CommandConfig{
		//超时时间
		Timeout: 2000,
		//请求阈值，有20个请求才会进行错误计算
		RequestVolumeThreshold: 5,
		//过多长时间熔断器，再次开启
		SleepWindow: 5000,
		//错误百分比
		ErrorPercentThreshold: 20,
	}
	hystrix.ConfigureCommand(name, config)
	return hystrix.Do(name, func() error {
		return h.Client.Call(ctx, req, rsp)
	}, func(err error) error {
		if err != nil {
			return errors.New("请求Agent服务超时")
		}
		return nil
	})
}

func NewHystrixWrapper(c client.Client) client.Client {
	return &hystrixWrapper{Client: c}
}

func GetHostService(serviceName string) interface{} {
	// 配置jaeger连接
	//jaegerTracer, closer, err := trace.NewJaegerTracer(serviceName, JaegerAddr)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer closer.Close()
	// 配置jaeger连接
	if conf.GetBoolConf("jaeger", "enable") {
		s = micro.NewService()
		//	micro.WrapClient(
		//		NewHystrixWrapper,
		//		opentracing.NewClientWrapper(jaegerTracer)),
		//)
		s.Init()
		return host.NewHostService(fmt.Sprintf("%s", serviceName), s.Client())
	}
	s = micro.NewService(
		micro.Registry(reg),
		micro.WrapClient(
			NewHystrixWrapper,
		))
	s.Init()
	return host.NewHostService(fmt.Sprintf("%s", serviceName), s.Client())
}

func GetTaskService(serviceName string) interface{} {
	// 配置jaeger连接
	jaegerTracer, closer, err := trace.NewJaegerTracer(serviceName, JaegerAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer closer.Close()
	// 配置jaeger连接
	if conf.GetBoolConf("jaeger", "enable") {
		s = micro.NewService(
			micro.Registry(reg),
			micro.WrapClient(
				NewHystrixWrapper,
				opentracing.NewClientWrapper(jaegerTracer)),
		)
		s.Init()
		return task.NewTaskService(fmt.Sprintf("%s", serviceName), s.Client())
	}
	s = micro.NewService(
		micro.Registry(reg),
		micro.WrapClient(
			NewHystrixWrapper,
		))
	s.Init()
	return task.NewTaskService(fmt.Sprintf("%s", serviceName), s.Client())
}

func GetHistoryService(serviceName string) interface{} {
	service := micro.NewService()
	service.Init()
	return bakhistory.NewHistoryService(serviceName, service.Client())
	//// 配置jaeger连接
	//jaegerTracer, closer, err := trace.NewJaegerTracer(serviceName, JaegerAddr)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer closer.Close()
	//// 配置jaeger连接
	//if conf.GetBoolConf("jaeger", "enable") {
	//	s = micro.NewService(
	//		micro.WrapClient(
	//			NewHystrixWrapper,
	//			opentracing.NewClientWrapper(jaegerTracer)),
	//	)
	//	s.Init()
	//	return bakhistory.NewHistoryService(fmt.Sprintf("%s", serviceName), s.Client())
	//}
	//s = micro.NewService(
	//	micro.Registry(reg),
	//	micro.Address("127.0.0.1:39800"),
	//	micro.WrapClient(
	//		NewHystrixWrapper,
	//	))
	//s.Init()
	//fmt.Println(s.Client().Options().CallOptions.Address)
	//return bakhistory.NewHistoryService(fmt.Sprintf("%s", serviceName), s.Client())
}

func GetBakService(serviceName string) interface{} {
	// 配置jaeger连接
	jaegerTracer, closer, err := trace.NewJaegerTracer(serviceName, JaegerAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer closer.Close()
	// 配置jaeger连接
	if conf.GetBoolConf("jaeger", "enable") {
		s = micro.NewService(
			micro.Registry(reg),
			micro.WrapClient(
				NewHystrixWrapper,
				opentracing.NewClientWrapper(jaegerTracer)),
		)
		s.Init()
		return bak.NewBakService(fmt.Sprintf("%s", serviceName), s.Client())
	}
	s = micro.NewService(
		micro.Registry(reg),
		micro.WrapClient(
			NewHystrixWrapper,
		))
	s.Init()
	return bak.NewBakService(fmt.Sprintf("%s", serviceName), s.Client())
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
