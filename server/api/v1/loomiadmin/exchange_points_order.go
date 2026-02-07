package loomiadmin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/loomiadmin"
	loomiadminReq "github.com/flipped-aurora/gin-vue-admin/server/model/loomiadmin/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ExchangePointsOrderApi struct {
}

var exchangePointsOrderService = service.ServiceGroupApp.LoomiadminServiceGroup.ExchangePointsOrderService

// CreateExchangePointsOrder 创建exchangePointsOrder表
// @Tags ExchangePointsOrder
// @Summary 创建exchangePointsOrder表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param x-token header string false "x-token"
// @Param data body loomiadmin.ExchangePointsOrder true "创建exchangePointsOrder表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /exchangePointsOrder/createExchangePointsOrder [post]
func (exchangePointsOrderApi *ExchangePointsOrderApi) CreateExchangePointsOrder(c *gin.Context) {
	var exchangePointsOrder loomiadmin.ExchangePointsOrder
	err := c.ShouldBindJSON(&exchangePointsOrder)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := exchangePointsOrderService.CreateExchangePointsOrder(&exchangePointsOrder); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteExchangePointsOrder 删除exchangePointsOrder表
// @Tags ExchangePointsOrder
// @Summary 删除exchangePointsOrder表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param x-token header string false "x-token"
// @Param data body loomiadmin.ExchangePointsOrder true "删除exchangePointsOrder表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /exchangePointsOrder/deleteExchangePointsOrder [delete]
func (exchangePointsOrderApi *ExchangePointsOrderApi) DeleteExchangePointsOrder(c *gin.Context) {
	var exchangePointsOrder loomiadmin.ExchangePointsOrder
	err := c.ShouldBindJSON(&exchangePointsOrder)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := exchangePointsOrderService.DeleteExchangePointsOrder(exchangePointsOrder); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteExchangePointsOrderByIds 批量删除exchangePointsOrder表
// @Tags ExchangePointsOrder
// @Summary 批量删除exchangePointsOrder表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param x-token header string false "x-token"
// @Param data body request.IdsReq true "批量删除exchangePointsOrder表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /exchangePointsOrder/deleteExchangePointsOrderByIds [delete]
func (exchangePointsOrderApi *ExchangePointsOrderApi) DeleteExchangePointsOrderByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := exchangePointsOrderService.DeleteExchangePointsOrderByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateExchangePointsOrder 更新exchangePointsOrder表
// @Tags ExchangePointsOrder
// @Summary 更新exchangePointsOrder表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param x-token header string false "x-token"
// @Param data body loomiadmin.ExchangePointsOrder true "更新exchangePointsOrder表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /exchangePointsOrder/updateExchangePointsOrder [put]
func (exchangePointsOrderApi *ExchangePointsOrderApi) UpdateExchangePointsOrder(c *gin.Context) {
	var exchangePointsOrder loomiadmin.ExchangePointsOrder
	err := c.ShouldBindJSON(&exchangePointsOrder)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 获取当前登录用户ID
	userID := utils.GetUserID(c)
	opID := int(userID)
	exchangePointsOrder.OperatorId = &opID

	if err := exchangePointsOrderService.UpdateExchangePointsOrder(exchangePointsOrder); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindExchangePointsOrder 用id查询exchangePointsOrder表
// @Tags ExchangePointsOrder
// @Summary 用id查询exchangePointsOrder表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param x-token header string false "x-token"
// @Param data query loomiadmin.ExchangePointsOrder true "用id查询exchangePointsOrder表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /exchangePointsOrder/findExchangePointsOrder [get]
func (exchangePointsOrderApi *ExchangePointsOrderApi) FindExchangePointsOrder(c *gin.Context) {
	var exchangePointsOrder loomiadmin.ExchangePointsOrder
	err := c.ShouldBindQuery(&exchangePointsOrder)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reexchangePointsOrder, err := exchangePointsOrderService.GetExchangePointsOrder(exchangePointsOrder.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reexchangePointsOrder": reexchangePointsOrder}, c)
	}
}

// GetExchangePointsOrderList 分页获取exchangePointsOrder表列表
// @Tags ExchangePointsOrder
// @Summary 分页获取exchangePointsOrder表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param x-token header string false "x-token"
// @Param data query loomiadminReq.ExchangePointsOrderSearch true "分页获取exchangePointsOrder表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /exchangePointsOrder/getExchangePointsOrderList [get]
func (exchangePointsOrderApi *ExchangePointsOrderApi) GetExchangePointsOrderList(c *gin.Context) {
	var pageInfo loomiadminReq.ExchangePointsOrderSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := exchangePointsOrderService.GetExchangePointsOrderInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
