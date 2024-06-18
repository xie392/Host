package host

import "context"

type Service interface {
	CreateHost(context.Context, *Host) (*Host, error)                  // 录入主机信息
	QueryHost(context.Context, *QueryHostRequest) (*ListHost, error)   // 查询主机列表
	DescribeHost(context.Context, *DescribeHostRequest) (*Host, error) // 查询主机详情
	UpdateHost(context.Context, *UpdateHostRequest) (*Host, error)     // 更新主机信息
	DeleteHost(context.Context, *DeleteHostRequest) (*Host, error)     // 删除主机信息
}
