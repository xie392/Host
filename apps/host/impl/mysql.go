package impl

import (
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/xie392/restful-api/apps"
	"github.com/xie392/restful-api/apps/host"
	"github.com/xie392/restful-api/conf"
	"gorm.io/gorm"
)

var impl = &HostServiceImpl{}

func NewHostService() host.Service {
	// Host service 服务的子 Loggger
	// 封装的Zap让其满足 Logger接口
	return &HostServiceImpl{
		l:  zap.L().Named("Host"),
		db: conf.C().MySQL.GetDB(),
	}
}

type HostServiceImpl struct {
	l  logger.Logger
	db *gorm.DB
}

func (i *HostServiceImpl) Config() {
	i.l = zap.L().Named("Host")
	i.db = conf.C().MySQL.GetDB()
}

// Name 服务服务的名
func (i *HostServiceImpl) Name() string {
	return host.AppName
}

// _ import app 自动执行注册逻辑
func init() {
	//  对象注册到ioc容器
	apps.RegistryImpl(impl)
}
