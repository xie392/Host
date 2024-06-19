package impl

import (
	"context"
	"github.com/xie392/restful-api/apps/host"
	"github.com/xie392/restful-api/conf"
	"gorm.io/gorm"
)

func (i *HostServiceImpl) CreateHost(ctx context.Context, ins *host.Host) (*host.Host, error) {
	// 生成id
	//ins.Id = utils.GenerateId(8)
	//ins.ResourceID = ins.Id

	// 添加创建时间
	//ins.CreateAt = time.Now().Unix()
	//ins.UpdateAt = ins.CreateAt

	// 插入资源和描述
	insertResource := func(tx *gorm.DB) error {
		return tx.Create(ins.Resource).Error
	}

	// 插入描述
	insertDescribe := func(tx *gorm.DB) error {
		return tx.Create(ins.Describe).Error
	}

	// 使用TransferFunds进行事务管理
	err := conf.TransferFunds(i.db, insertResource, insertDescribe)
	if err != nil {
		i.l.Error("create host failed", err)
		return nil, err
	}
	return ins, nil
}

func (i *HostServiceImpl) QueryHost(ctx context.Context, req *host.QueryHostRequest) (*host.ListHost, error) {
	return nil, nil
}

func (i *HostServiceImpl) DescribeHost(ctx context.Context, req *host.DescribeHostRequest) (*host.Host, error) {
	return nil, nil
}

func (i *HostServiceImpl) UpdateHost(ctx context.Context, req *host.UpdateHostRequest) (*host.Host, error) {
	return nil, nil
}

func (i *HostServiceImpl) DeleteHost(ctx context.Context, req *host.DeleteHostRequest) (*host.Host, error) {
	return nil, nil
}
