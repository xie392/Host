package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xie392/restful-api/apps/host"
	"github.com/xie392/restful-api/pkg/utils"
	"net/http"
)

func (h *Handler) createHost(c *gin.Context) {
	ins := host.NewHost()

	// 将HTTP协议里面 解析出来用户的请求参数
	// read c.Request.Body
	// json unmarshal
	// 用户传递过来的参数进行解析, 实现了一个json 的unmarshal
	if err := c.Bind(ins); err != nil {
		fmt.Println("Invalid request body", err)
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request body")
		return
	}

	// 进行接口调用, 返回 肯定有成功或者失败
	ins, err := h.svc.CreateHost(c.Request.Context(), ins)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Internal server error")
		return
	}

	// 成功, 把对象实例返回给HTTP API调用
	utils.SuccessResponse(c, ins)
}

//
//func (h *Handler) queryHost(c *gin.Context) {
//	// 从http请求的query string 中获取参数
//	req := host.NewQueryHostFromHTTP(c.Request)
//
//	// 进行接口调用, 返回 肯定有成功或者失败
//	set, err := h.svc.QueryHost(c.Request.Context(), req)
//	if err != nil {
//		response.Failed(c.Writer, err)
//		return
//	}
//
//	response.Success(c.Writer, set)
//}
//
//func (h *Handler) describeHost(c *gin.Context) {
//	// 从http请求的query string 中获取参数
//	req := host.NewDescribeHostRequestWithId(c.Param("id"))
//
//	// 进行接口调用, 返回 肯定有成功或者失败
//	set, err := h.svc.DescribeHost(c.Request.Context(), req)
//	if err != nil {
//		response.Failed(c.Writer, err)
//		return
//	}
//
//	response.Success(c.Writer, set)
//}
//
//func (h *Handler) putHost(c *gin.Context) {
//	// 从http请求的query string 中获取参数
//	req := host.NewPutUpdateHostRequest(c.Param("id"))
//
//	// 解析Body里面的数据
//	if err := c.Bind(req.Host); err != nil {
//		response.Failed(c.Writer, err)
//		return
//	}
//	req.Id = c.Param("id")
//
//	// 进行接口调用, 返回 肯定有成功或者失败
//	set, err := h.svc.UpdateHost(c.Request.Context(), req)
//	if err != nil {
//		response.Failed(c.Writer, err)
//		return
//	}
//
//	response.Success(c.Writer, set)
//}
//
//func (h *Handler) patchHost(c *gin.Context) {
//	// 从http请求的query string 中获取参数, PATH/QUERY
//	req := host.NewPatchUpdateHostRequest(c.Param("id"))
//
//	// 解析Body里面的数据
//	if err := c.Bind(req.Host); err != nil {
//		response.Failed(c.Writer, err)
//		return
//	}
//	req.Id = c.Param("id")
//
//	// 进行接口调用, 返回 肯定有成功或者失败
//	set, err := h.svc.UpdateHost(c.Request.Context(), req)
//	if err != nil {
//		response.Failed(c.Writer, err)
//		return
//	}
//
//	response.Success(c.Writer, set)
//}
