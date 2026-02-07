package loomiadmin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AnchorPointsFlowsRouter struct {
}

// InitAnchorPointsFlowsRouter 初始化 anchorPointsFlows表 路由信息
func (s *AnchorPointsFlowsRouter) InitAnchorPointsFlowsRouter(Router *gin.RouterGroup) {
	anchorPointsFlowsRouter := Router.Group("anchorPointsFlows").Use(middleware.OperationRecord())
	anchorPointsFlowsRouterWithoutRecord := Router.Group("anchorPointsFlows")
	var anchorPointsFlowsApi = v1.ApiGroupApp.LoomiadminApiGroup.AnchorPointsFlowsApi
	{
		anchorPointsFlowsRouter.POST("createAnchorPointsFlows", anchorPointsFlowsApi.CreateAnchorPointsFlows)   // 新建anchorPointsFlows表
		anchorPointsFlowsRouter.DELETE("deleteAnchorPointsFlows", anchorPointsFlowsApi.DeleteAnchorPointsFlows) // 删除anchorPointsFlows表
		anchorPointsFlowsRouter.DELETE("deleteAnchorPointsFlowsByIds", anchorPointsFlowsApi.DeleteAnchorPointsFlowsByIds) // 批量删除anchorPointsFlows表
		anchorPointsFlowsRouter.PUT("updateAnchorPointsFlows", anchorPointsFlowsApi.UpdateAnchorPointsFlows)    // 更新anchorPointsFlows表
	}
	{
		anchorPointsFlowsRouterWithoutRecord.GET("findAnchorPointsFlows", anchorPointsFlowsApi.FindAnchorPointsFlows)        // 根据ID获取anchorPointsFlows表
		anchorPointsFlowsRouterWithoutRecord.GET("getAnchorPointsFlowsList", anchorPointsFlowsApi.GetAnchorPointsFlowsList)  // 获取anchorPointsFlows表列表
	}
}
