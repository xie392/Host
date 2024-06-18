package impl

import (
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/xie392/restful-api/apps/host"
)

// 接口实现的静态方法
//var _ host.Service = (*HostServiceImpl)(nil)

func NewHostService() host.Service {
	// Host service 服务的子 Loggger
	// 封装的Zap让其满足 Logger接口
	return &HostServiceImpl{l: zap.L().Named("Host")}
}

type HostServiceImpl struct {
	l logger.Logger
}
