package loomiadmin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ExchangePointsOrderRouter struct {
}

// InitExchangePointsOrderRouter 初始化 exchangePointsOrder表 路由信息
func (s *ExchangePointsOrderRouter) InitExchangePointsOrderRouter(Router *gin.RouterGroup) {
	exchangePointsOrderRouter := Router.Group("exchangePointsOrder").Use(middleware.OperationRecord())
	exchangePointsOrderRouterWithoutRecord := Router.Group("exchangePointsOrder")
	var exchangePointsOrderApi = v1.ApiGroupApp.LoomiadminApiGroup.ExchangePointsOrderApi
	{
		exchangePointsOrderRouter.POST("createExchangePointsOrder", exchangePointsOrderApi.CreateExchangePointsOrder)   // 新建exchangePointsOrder表
		exchangePointsOrderRouter.DELETE("deleteExchangePointsOrder", exchangePointsOrderApi.DeleteExchangePointsOrder) // 删除exchangePointsOrder表
		exchangePointsOrderRouter.DELETE("deleteExchangePointsOrderByIds", exchangePointsOrderApi.DeleteExchangePointsOrderByIds) // 批量删除exchangePointsOrder表
		exchangePointsOrderRouter.PUT("updateExchangePointsOrder", exchangePointsOrderApi.UpdateExchangePointsOrder)    // 更新exchangePointsOrder表
	}
	{
		exchangePointsOrderRouterWithoutRecord.GET("findExchangePointsOrder", exchangePointsOrderApi.FindExchangePointsOrder)        // 根据ID获取exchangePointsOrder表
		exchangePointsOrderRouterWithoutRecord.GET("getExchangePointsOrderList", exchangePointsOrderApi.GetExchangePointsOrderList)  // 获取exchangePointsOrder表列表
	}
}
