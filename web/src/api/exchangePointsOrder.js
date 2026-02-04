import service from '@/utils/request'

// @Tags ExchangePointsOrder
// @Summary 创建exchangePointsOrder表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExchangePointsOrder true "创建exchangePointsOrder表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /exchangePointsOrder/createExchangePointsOrder [post]
export const createExchangePointsOrder = (data) => {
  return service({
    url: '/exchangePointsOrder/createExchangePointsOrder',
    method: 'post',
    data
  })
}

// @Tags ExchangePointsOrder
// @Summary 删除exchangePointsOrder表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExchangePointsOrder true "删除exchangePointsOrder表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /exchangePointsOrder/deleteExchangePointsOrder [delete]
export const deleteExchangePointsOrder = (data) => {
  return service({
    url: '/exchangePointsOrder/deleteExchangePointsOrder',
    method: 'delete',
    data
  })
}

// @Tags ExchangePointsOrder
// @Summary 批量删除exchangePointsOrder表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除exchangePointsOrder表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /exchangePointsOrder/deleteExchangePointsOrder [delete]
export const deleteExchangePointsOrderByIds = (data) => {
  return service({
    url: '/exchangePointsOrder/deleteExchangePointsOrderByIds',
    method: 'delete',
    data
  })
}

// @Tags ExchangePointsOrder
// @Summary 更新exchangePointsOrder表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExchangePointsOrder true "更新exchangePointsOrder表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /exchangePointsOrder/updateExchangePointsOrder [put]
export const updateExchangePointsOrder = (data) => {
  return service({
    url: '/exchangePointsOrder/updateExchangePointsOrder',
    method: 'put',
    data
  })
}

// @Tags ExchangePointsOrder
// @Summary 用id查询exchangePointsOrder表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ExchangePointsOrder true "用id查询exchangePointsOrder表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /exchangePointsOrder/findExchangePointsOrder [get]
export const findExchangePointsOrder = (params) => {
  return service({
    url: '/exchangePointsOrder/findExchangePointsOrder',
    method: 'get',
    params
  })
}

// @Tags ExchangePointsOrder
// @Summary 分页获取exchangePointsOrder表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取exchangePointsOrder表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /exchangePointsOrder/getExchangePointsOrderList [get]
export const getExchangePointsOrderList = (params) => {
  return service({
    url: '/exchangePointsOrder/getExchangePointsOrderList',
    method: 'get',
    params
  })
}
