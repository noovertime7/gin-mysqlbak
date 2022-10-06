package pkg

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/noovertime7/gin-mysqlbak/agent/pkg/trace"
	"github.com/noovertime7/gin-mysqlbak/agent/proto/bak"
	"github.com/noovertime7/gin-mysqlbak/agent/proto/bakhistory"
	"github.com/noovertime7/gin-mysqlbak/agent/proto/esbak"
	"github.com/noovertime7/gin-mysqlbak/agent/proto/host"
	"github.com/noovertime7/gin-mysqlbak/agent/proto/task"
	"github.com/noovertime7/gin-mysqlbak/agent/repository"
	"github.com/noovertime7/gin-mysqlbak/conf"
	"log"
)

var s micro.Service

var JaegerAddr = conf.GetStringConf("jaeger", "addr")

var AgentService *repository.AgentService

func GetHostService(serviceName string) (host.HostService, string, error) {
	//配置jaeger连接
	jaegerTracer, closer, err := trace.NewJaegerTracer(serviceName, JaegerAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer closer.Close()
	addr, err := AgentService.GetServiceAddr(context.Background(), serviceName)
	if err != nil {
		return nil, "", err
	}
	if conf.GetBoolConf("jaeger", "enable") {
		s = micro.NewService()
		micro.WrapClient(
			opentracing.NewClientWrapper(jaegerTracer))
		s.Init()
		service := host.NewHostService(fmt.Sprintf("%s", serviceName), s.Client())
		return service, addr, nil
	}
	s = micro.NewService()
	s.Init()
	service := host.NewHostService(fmt.Sprintf("%s", serviceName), s.Client())
	return service, addr, nil
}

func GetTaskService(serviceName string) (task.TaskService, string, error) {
	// 配置jaeger连接
	jaegerTracer, closer, err := trace.NewJaegerTracer(serviceName, JaegerAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer closer.Close()
	// 获取服务地址
	addr, err := AgentService.GetServiceAddr(context.Background(), serviceName)
	if err != nil {
		return nil, "", err
	}
	if conf.GetBoolConf("jaeger", "enable") {
		s = micro.NewService(
			micro.WrapClient(
				opentracing.NewClientWrapper(jaegerTracer)),
		)
		s.Init()
		service := task.NewTaskService(fmt.Sprintf("%s", serviceName), s.Client())
		return service, addr, nil
	}
	s = micro.NewService()
	s.Init()
	service := task.NewTaskService(fmt.Sprintf("%s", serviceName), s.Client())
	return service, addr, nil
}

func GetHistoryService(serviceName string) (bakhistory.HistoryService, string, error) {
	// 配置jaeger连接
	jaegerTracer, closer, err := trace.NewJaegerTracer(serviceName, JaegerAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer closer.Close()
	addr, err := AgentService.GetServiceAddr(context.Background(), serviceName)
	if err != nil {
		return nil, "", err
	}
	// 配置jaeger连接
	if conf.GetBoolConf("jaeger", "enable") {
		s := micro.NewService(
			micro.WrapClient(
				opentracing.NewClientWrapper(jaegerTracer),
			))

		s.Init()
		ser := bakhistory.NewHistoryService(serviceName, s.Client())
		return ser, addr, nil
	}
	s := micro.NewService(
		micro.WrapClient(
			opentracing.NewClientWrapper(jaegerTracer),
		))
	s.Init()
	ser := bakhistory.NewHistoryService(serviceName, s.Client())
	return ser, addr, nil
}

// GetESTaskService 获取es备份的task服务
func GetESTaskService(serviceName string) (esbak.EsService, string, error) {
	// 配置jaeger连接
	jaegerTracer, closer, err := trace.NewJaegerTracer(serviceName, JaegerAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer closer.Close()
	addr, err := AgentService.GetServiceAddr(context.Background(), serviceName)
	if err != nil {
		return nil, "", err
	}
	// 配置jaeger连接
	if conf.GetBoolConf("jaeger", "enable") {
		s := micro.NewService(
			micro.WrapClient(
				opentracing.NewClientWrapper(jaegerTracer),
			))
		s.Init()
		ser := esbak.NewEsService(serviceName, s.Client())
		return ser, addr, nil
	}
	s := micro.NewService(
		micro.WrapClient(
			opentracing.NewClientWrapper(jaegerTracer),
		))
	s.Init()
	ser := esbak.NewEsService(serviceName, s.Client())
	return ser, addr, nil
}

func GetESBakService(serviceName string) (esbak.EsBakService, string, error) {
	// 配置jaeger连接
	jaegerTracer, closer, err := trace.NewJaegerTracer(serviceName, JaegerAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer closer.Close()
	addr, err := AgentService.GetServiceAddr(context.Background(), serviceName)
	if err != nil {
		return nil, "", err
	}
	// 配置jaeger连接
	if conf.GetBoolConf("jaeger", "enable") {
		s := micro.NewService(
			micro.WrapClient(
				opentracing.NewClientWrapper(jaegerTracer),
			))
		s.Init()
		ser := esbak.NewEsBakService(serviceName, s.Client())
		return ser, addr, nil
	}
	s := micro.NewService(
		micro.WrapClient(
			opentracing.NewClientWrapper(jaegerTracer),
		))
	s.Init()
	ser := esbak.NewEsBakService(serviceName, s.Client())
	return ser, addr, nil
}

func GetESHistoryService(serviceName string) (esbak.EsHistoryService, string, error) {
	// 配置jaeger连接
	jaegerTracer, closer, err := trace.NewJaegerTracer(serviceName, JaegerAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer closer.Close()
	addr, err := AgentService.GetServiceAddr(context.Background(), serviceName)
	if err != nil {
		return nil, "", err
	}
	// 配置jaeger连接
	if conf.GetBoolConf("jaeger", "enable") {
		s = micro.NewService(
			micro.WrapClient(
				opentracing.NewClientWrapper(jaegerTracer)),
		)
		s.Init()
		ser := esbak.NewEsHistoryService(fmt.Sprintf("%s", serviceName), s.Client())
		return ser, addr, nil
	}
	s = micro.NewService()
	s.Init()
	ser := esbak.NewEsHistoryService(fmt.Sprintf("%s", serviceName), s.Client())
	return ser, addr, nil
}

func GetBakService(serviceName string) (bak.BakService, string, error) {
	// 配置jaeger连接
	jaegerTracer, closer, err := trace.NewJaegerTracer(serviceName, JaegerAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer closer.Close()
	addr, err := AgentService.GetServiceAddr(context.Background(), serviceName)
	if err != nil {
		return nil, "", err
	}
	// 配置jaeger连接
	if conf.GetBoolConf("jaeger", "enable") {
		s = micro.NewService(
			micro.WrapClient(
				opentracing.NewClientWrapper(jaegerTracer)),
		)
		s.Init()
		ser := bak.NewBakService(fmt.Sprintf("%s", serviceName), s.Client())
		return ser, addr, nil
	}
	s = micro.NewService()
	s.Init()
	ser := bak.NewBakService(fmt.Sprintf("%s", serviceName), s.Client())
	return ser, addr, nil
}
