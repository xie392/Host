package http

import (
	"github.com/gin-gonic/gin"
	"github.com/xie392/restful-api/apps"
	"github.com/xie392/restful-api/apps/host"
)

// 面向接口, 真正Service的实现, 在服务实例化的时候传递进行
// 也就是(CLI)  Start时候
var handler = &Handler{}

// Handler 通过写一个实例类, 把内部的接口通过HTTP协议暴露出去
type Handler struct {
	svc host.Service
}

func NewHostHttpHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Config() {
	h.svc = apps.HostService
}

func (h *Handler) Registry(r gin.IRouter) {
	r.POST("/hosts", h.createHost)
	//r.GET("/hosts", h.queryHost)
	//r.GET("/hosts/:id", h.describeHost)
	//r.PUT("/hosts/:id", h.putHost)
	//r.PATCH("/hosts/:id", h.patchHost)
}
