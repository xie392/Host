package apps

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xie392/restful-api/apps/host"
)

// IOC 容器 管理所有的服务的实例
// 1. HostService的实例必须注册过 HostService才会有具体的实例, 服务启动时注
//2. HTTP 暴露模块, 依赖Ioc里面的HostService

var (
	// HostService 使用Interface{} + 断言进行抽象
	HostService host.Service
	// 维护当前所有的服务
	implApps = map[string]ImplService{}
	// 维护当前所有的Gin服务
	ginApps = map[string]GinService{}
)

func RegistryImpl(svc ImplService) {
	// 服务实例注册 svcs map当中
	if _, ok := implApps[svc.Name()]; ok {
		panic(fmt.Sprintf("service %s has registried", svc.Name()))
	}

	implApps[svc.Name()] = svc
	// 更加对象满足的接口来注册具体的服
	if v, ok := svc.(host.Service); ok {
		HostService = v
	}
}

// 如果指定了具体类 就导致没增加一种类多一个Get方法
//func GetHostImpl(name string) host.Service

// GetImpl Get 一个Impl服务的实例：implApps
// 返回一个对象, 任何类型都可使用 interface{} 由使用方进行断言
func GetImpl(name string) interface{} {
	for k, v := range implApps {
		if k == name {
			return v
		}
	}

	return nil
}

func RegistryGin(svc GinService) {
	// 服务实例注册到 svcs map当中
	if _, ok := ginApps[svc.Name()]; ok {
		panic(fmt.Sprintf("service %s has registried", svc.Name()))
	}

	ginApps[svc.Name()] = svc
}

// InitImpl 用户初始注册到Ioc容器里面的所有服
func InitImpl() {
	for _, v := range implApps {
		v.Config()
	}
}

// LoadedGinApps 已经加载完成的Gin App由Ioc管理
func LoadedGinApps() (names []string) {
	for k := range ginApps {
		names = append(names, k)
	}
	return
}

// InitGin 用户初始 注册到Ioc容器里面的所有服务
func InitGin(r gin.IRouter) {
	// 先初始化好所有对
	for _, v := range ginApps {
		v.Config()
	}

	// 完成Http Handler的注册
	for _, v := range ginApps {
		v.Registry(r)
	}
}

type ImplService interface {
	Config()
	Name() string
}

// GinService 注册Gin编写的Handler
// 比如 编写了Http服务A, 只需要实现Registry方法, 就能把Handler注册给Root Router
type GinService interface {
	Registry(r gin.IRouter)
	Config()
	Name() string
}
