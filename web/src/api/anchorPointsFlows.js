import service from '@/utils/request'

// @Tags AnchorPointsFlows
// @Summary 创建anchorPointsFlows表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AnchorPointsFlows true "创建anchorPointsFlows表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /anchorPointsFlows/createAnchorPointsFlows [post]
export const createAnchorPointsFlows = (data) => {
  return service({
    url: '/anchorPointsFlows/createAnchorPointsFlows',
    method: 'post',
    data
  })
}

// @Tags AnchorPointsFlows
// @Summary 删除anchorPointsFlows表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AnchorPointsFlows true "删除anchorPointsFlows表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /anchorPointsFlows/deleteAnchorPointsFlows [delete]
export const deleteAnchorPointsFlows = (data) => {
  return service({
    url: '/anchorPointsFlows/deleteAnchorPointsFlows',
    method: 'delete',
    data
  })
}

// @Tags AnchorPointsFlows
// @Summary 批量删除anchorPointsFlows表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除anchorPointsFlows表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /anchorPointsFlows/deleteAnchorPointsFlows [delete]
export const deleteAnchorPointsFlowsByIds = (data) => {
  return service({
    url: '/anchorPointsFlows/deleteAnchorPointsFlowsByIds',
    method: 'delete',
    data
  })
}

// @Tags AnchorPointsFlows
// @Summary 更新anchorPointsFlows表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AnchorPointsFlows true "更新anchorPointsFlows表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /anchorPointsFlows/updateAnchorPointsFlows [put]
export const updateAnchorPointsFlows = (data) => {
  return service({
    url: '/anchorPointsFlows/updateAnchorPointsFlows',
    method: 'put',
    data
  })
}

// @Tags AnchorPointsFlows
// @Summary 用id查询anchorPointsFlows表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.AnchorPointsFlows true "用id查询anchorPointsFlows表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /anchorPointsFlows/findAnchorPointsFlows [get]
export const findAnchorPointsFlows = (params) => {
  return service({
    url: '/anchorPointsFlows/findAnchorPointsFlows',
    method: 'get',
    params
  })
}

// @Tags AnchorPointsFlows
// @Summary 分页获取anchorPointsFlows表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取anchorPointsFlows表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /anchorPointsFlows/getAnchorPointsFlowsList [get]
export const getAnchorPointsFlowsList = (params) => {
  return service({
    url: '/anchorPointsFlows/getAnchorPointsFlowsList',
    method: 'get',
    params
  })
}
