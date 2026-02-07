package loomiadmin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/loomiadmin"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    loomiadminReq "github.com/flipped-aurora/gin-vue-admin/server/model/loomiadmin/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type AnchorPointsFlowsApi struct {
}

var anchorPointsFlowsService = service.ServiceGroupApp.LoomiadminServiceGroup.AnchorPointsFlowsService


// CreateAnchorPointsFlows 创建anchorPointsFlows表
// @Tags AnchorPointsFlows
// @Summary 创建anchorPointsFlows表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body loomiadmin.AnchorPointsFlows true "创建anchorPointsFlows表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /anchorPointsFlows/createAnchorPointsFlows [post]
func (anchorPointsFlowsApi *AnchorPointsFlowsApi) CreateAnchorPointsFlows(c *gin.Context) {
	var anchorPointsFlows loomiadmin.AnchorPointsFlows
	err := c.ShouldBindJSON(&anchorPointsFlows)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := anchorPointsFlowsService.CreateAnchorPointsFlows(&anchorPointsFlows); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteAnchorPointsFlows 删除anchorPointsFlows表
// @Tags AnchorPointsFlows
// @Summary 删除anchorPointsFlows表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body loomiadmin.AnchorPointsFlows true "删除anchorPointsFlows表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /anchorPointsFlows/deleteAnchorPointsFlows [delete]
func (anchorPointsFlowsApi *AnchorPointsFlowsApi) DeleteAnchorPointsFlows(c *gin.Context) {
	var anchorPointsFlows loomiadmin.AnchorPointsFlows
	err := c.ShouldBindJSON(&anchorPointsFlows)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := anchorPointsFlowsService.DeleteAnchorPointsFlows(anchorPointsFlows); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteAnchorPointsFlowsByIds 批量删除anchorPointsFlows表
// @Tags AnchorPointsFlows
// @Summary 批量删除anchorPointsFlows表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除anchorPointsFlows表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /anchorPointsFlows/deleteAnchorPointsFlowsByIds [delete]
func (anchorPointsFlowsApi *AnchorPointsFlowsApi) DeleteAnchorPointsFlowsByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := anchorPointsFlowsService.DeleteAnchorPointsFlowsByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateAnchorPointsFlows 更新anchorPointsFlows表
// @Tags AnchorPointsFlows
// @Summary 更新anchorPointsFlows表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body loomiadmin.AnchorPointsFlows true "更新anchorPointsFlows表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /anchorPointsFlows/updateAnchorPointsFlows [put]
func (anchorPointsFlowsApi *AnchorPointsFlowsApi) UpdateAnchorPointsFlows(c *gin.Context) {
	var anchorPointsFlows loomiadmin.AnchorPointsFlows
	err := c.ShouldBindJSON(&anchorPointsFlows)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := anchorPointsFlowsService.UpdateAnchorPointsFlows(anchorPointsFlows); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindAnchorPointsFlows 用id查询anchorPointsFlows表
// @Tags AnchorPointsFlows
// @Summary 用id查询anchorPointsFlows表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query loomiadmin.AnchorPointsFlows true "用id查询anchorPointsFlows表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /anchorPointsFlows/findAnchorPointsFlows [get]
func (anchorPointsFlowsApi *AnchorPointsFlowsApi) FindAnchorPointsFlows(c *gin.Context) {
	var anchorPointsFlows loomiadmin.AnchorPointsFlows
	err := c.ShouldBindQuery(&anchorPointsFlows)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reanchorPointsFlows, err := anchorPointsFlowsService.GetAnchorPointsFlows(anchorPointsFlows.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reanchorPointsFlows": reanchorPointsFlows}, c)
	}
}

// GetAnchorPointsFlowsList 分页获取anchorPointsFlows表列表
// @Tags AnchorPointsFlows
// @Summary 分页获取anchorPointsFlows表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query loomiadminReq.AnchorPointsFlowsSearch true "分页获取anchorPointsFlows表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /anchorPointsFlows/getAnchorPointsFlowsList [get]
func (anchorPointsFlowsApi *AnchorPointsFlowsApi) GetAnchorPointsFlowsList(c *gin.Context) {
	var pageInfo loomiadminReq.AnchorPointsFlowsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := anchorPointsFlowsService.GetAnchorPointsFlowsInfoList(pageInfo); err != nil {
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
