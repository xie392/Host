package impl_test

import (
	"context"
	"fmt"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/stretchr/testify/assert"
	"github.com/xie392/restful-api/apps/host"
	"github.com/xie392/restful-api/apps/host/impl"
	"github.com/xie392/restful-api/conf"
	"testing"
)

var (
	// 定义对象是满足该接口的实例
	service host.Service
)

func TestCreate(t *testing.T) {
	should := assert.New(t)

	// 创建一个Host实例
	ins := host.NewHost()

	// 设置Host的属性
	ins.Id = "test-02"
	ins.Region = "广州"
	ins.Type = "small"
	ins.Name = "接口测试主机"
	ins.ResourceID = "test-02"
	ins.CPU = 1
	ins.Memory = 2048

	ins, err := service.CreateHost(context.Background(), ins)
	if should.NoError(err) {
		fmt.Println(ins)
	}
}

//
//func TestQuery(t *testing.T) {
//	should := assert.New(t)
//
//	req := host.NewQueryHostRequest()
//	req.Keywords = "接口测试"
//	set, err := service.QueryHost(context.Background(), req)
//	if should.NoError(err) {
//		for i := range set.Items {
//			fmt.Println(set.Items[i].Id)
//		}
//	}
//}
//
//func TestDescribe(t *testing.T) {
//	should := assert.New(t)
//
//	req := host.NewDescribeHostRequestWithId("ins-09")
//	ins, err := service.DescribeHost(context.Background(), req)
//	if should.NoError(err) {
//		fmt.Println(ins.Id)
//	}
//}
//
//func TestUpdate(t *testing.T) {
//	should := assert.New(t)
//
//	req := host.NewPutUpdateHostRequest("ins-09")
//	req.Name = "更新测试02"
//	req.Region = "rg 02"
//	req.Type = "small"
//	req.CPU = 1
//	req.Memory = 2048
//	req.Description = "测试更新"
//	ins, err := service.UpdateHost(context.Background(), req)
//	if should.NoError(err) {
//		fmt.Println(ins.Id)
//	}
//}
//
//
//func TestPatch(t *testing.T) {
//	should := assert.New(t)
//
//	req := host.NewPatchUpdateHostRequest("ins-09")
//	req.Description = "Patch更新模式测试"
//	ins, err := service.UpdateHost(context.Background(), req)
//	if should.NoError(err) {
//		fmt.Println(ins.Id)
//	}
//}

func TestInit(t *testing.T) {
	// 测试用例的配置文件
	//err := conf.LoadConfigFromEnv()
	err := conf.LoadConfigFromToml("../../../etc/config.toml")
	if err != nil {
		panic(err)
	}

	// 需要初始化全局Logger,
	// 为什么不设计为默认打印, 因为性能
	fmt.Println(zap.DevelopmentSetup())

	// host service 的具体实现
	service = impl.NewHostService()

	// 测试用例
	TestCreate(t)
}
