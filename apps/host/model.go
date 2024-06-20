package host

import (
	"github.com/xie392/restful-api/pkg/utils"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

type ListHost struct {
	Total int     `json:"total"`
	Items []*Host `json:"items"`
}

type Host struct {
	*Resource // 资源公共属性部分
	*Describe // 资源独有属性部分
}

type Vendor int

const (
	PRIVATE_IDC Vendor = iota // 枚举的默认值
	ALIYUN                    // 阿里云
	QCLOUD                    // 腾讯云
	TXYUN                     // 天翼云
)

type Resource struct {
	Id          string `json:"id" gorm:"primaryKey" binding:"required"` // 全局唯一Id
	Vendor      Vendor `json:"vendor"`                                  // 厂商
	Region      string `json:"region" binding:"required"`               // 地域
	CreateAt    int64  `json:"create_at"  gorm:"autoCreateTime"`        // 创建时间
	ExpireAt    int64  `json:"expire_at" `                              // 过期时间
	Type        string `json:"type"  binding:"required"`                // 规格
	Name        string `json:"name"  binding:"required"`                // 名称
	Description string `json:"description"`                             // 描述
	Status      string `json:"status"`                                  // 服务商中的状态
	Tags        string `json:"tags"`                                    // 标签
	UpdateAt    int64  `json:"update_at"  gorm:"autoCreateTime"`        // 更新时间
	SyncAt      int64  `json:"sync_at"`                                 // 同步时间
	Account     string `json:"account"`                                 // 资源的所属账号
	PublicIP    string `json:"public_ip"`                               // 公网IP
	PrivateIP   string `json:"private_ip"`                              // 内网IP
}

type Describe struct {
	ResourceID   string `json:"resource_id" gorm:"primaryKey" binding:"required"` // 资源ID
	CPU          int    `json:"cpu" binding:"required"`                           // 核数
	Memory       int    `json:"memory" binding:"required"`                        // 内存
	GPUAmount    int    `json:"gpu_amount"`                                       // GPU数量
	GPUSpec      string `json:"gpu_spec"`                                         // GPU类型
	OSType       string `json:"os_type"`                                          // 操作系统类型，分为Windows和Linux
	OSName       string `json:"os_name"`                                          // 操作系统名称
	SerialNumber string `json:"serial_number"`                                    // 序列号
}

type QueryHostRequest struct {
	PageSize   int    `json:"page_size"`
	PageNumber int    `json:"page_number"`
	Keywords   string `json:"keywords"`
}

func NewQueryHostRequest() *QueryHostRequest {
	return &QueryHostRequest{
		PageSize:   20,
		PageNumber: 1,
	}
}

func NewQueryHostFromHTTP(r *http.Request) *QueryHostRequest {
	req := NewQueryHostRequest()
	qs := r.URL.Query()
	pss := qs.Get("page_size")
	if pss != "" {
		req.PageSize, _ = strconv.Atoi(pss)
	}

	pns := qs.Get("page_number")
	if pns != "" {
		req.PageNumber, _ = strconv.Atoi(pns)
	}

	req.Keywords = qs.Get("keywords")
	return req
}

type UpdateHostRequest struct {
	*Describe
}

type DeleteHostRequest struct {
	Id string `json:"id"`
}

type DescribeHostRequest struct {
	Id string `json:"id"`
}

func NewHost() *Host {
	id := utils.GenerateId(8)
	createAt := time.Now().Unix()
	return &Host{
		Resource: &Resource{
			Id:       id,
			CreateAt: createAt,
			UpdateAt: createAt,
		},
		Describe: &Describe{
			ResourceID: id,
		},
	}
}

//func NewHostFrom(r *http.Request) *Host {
//	req := NewQueryHostRequest()
//	qs := r.URL.Query()
//	pss := qs.Get("page_size")
//	if pss != "" {
//		req.PageSize, _ = strconv.Atoi(pss)
//	}
//
//	pns := qs.Get("page_number")
//	if pns != "" {
//		req.PageNumber, _ = strconv.Atoi(pns)
//	}
//
//	req.Keywords = qs.Get("keywords")
//	return req
//req := NewHost()
//id := utils.GenerateId(8)
//createAt := time.Now().Unix()
//}

func (Resource) TableName() string {
	return "resource"
}

func (Describe) TableName() string {
	return "host"
}

func AutoMigrateResource(db *gorm.DB) error {
	if err := db.AutoMigrate(&Resource{}); err != nil {
		return err
	}
	return nil
}

func AutoMigrateDescribe(db *gorm.DB) error {
	if err := db.AutoMigrate(&Describe{}); err != nil {
		return err
	}
	return nil
}
