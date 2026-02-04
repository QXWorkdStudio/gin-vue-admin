package loomiadmin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/loomiadmin"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    loomiadminReq "github.com/flipped-aurora/gin-vue-admin/server/model/loomiadmin/request"
)

type ExchangePointsOrderService struct {
}

// CreateExchangePointsOrder 创建exchangePointsOrder表记录
// Author [piexlmax](https://github.com/piexlmax)
func (exchangePointsOrderService *ExchangePointsOrderService) CreateExchangePointsOrder(exchangePointsOrder *loomiadmin.ExchangePointsOrder) (err error) {
	err = global.MustGetGlobalDBByDBName("loomi").Create(exchangePointsOrder).Error
	return err
}

// DeleteExchangePointsOrder 删除exchangePointsOrder表记录
// Author [piexlmax](https://github.com/piexlmax)
func (exchangePointsOrderService *ExchangePointsOrderService)DeleteExchangePointsOrder(exchangePointsOrder loomiadmin.ExchangePointsOrder) (err error) {
	err = global.MustGetGlobalDBByDBName("loomi").Delete(&exchangePointsOrder).Error
	return err
}

// DeleteExchangePointsOrderByIds 批量删除exchangePointsOrder表记录
// Author [piexlmax](https://github.com/piexlmax)
func (exchangePointsOrderService *ExchangePointsOrderService)DeleteExchangePointsOrderByIds(ids request.IdsReq) (err error) {
	err = global.MustGetGlobalDBByDBName("loomi").Delete(&[]loomiadmin.ExchangePointsOrder{},"id in ?",ids.Ids).Error
	return err
}

// UpdateExchangePointsOrder 更新exchangePointsOrder表记录
// Author [piexlmax](https://github.com/piexlmax)
func (exchangePointsOrderService *ExchangePointsOrderService)UpdateExchangePointsOrder(exchangePointsOrder loomiadmin.ExchangePointsOrder) (err error) {
	err = global.MustGetGlobalDBByDBName("loomi").Save(&exchangePointsOrder).Error
	return err
}

// GetExchangePointsOrder 根据id获取exchangePointsOrder表记录
// Author [piexlmax](https://github.com/piexlmax)
func (exchangePointsOrderService *ExchangePointsOrderService)GetExchangePointsOrder(id uint) (exchangePointsOrder loomiadmin.ExchangePointsOrder, err error) {
	err = global.MustGetGlobalDBByDBName("loomi").Where("id = ?", id).First(&exchangePointsOrder).Error
	return
}

// GetExchangePointsOrderInfoList 分页获取exchangePointsOrder表记录
// Author [piexlmax](https://github.com/piexlmax)
func (exchangePointsOrderService *ExchangePointsOrderService)GetExchangePointsOrderInfoList(info loomiadminReq.ExchangePointsOrderSearch) (list []loomiadmin.ExchangePointsOrder, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.MustGetGlobalDBByDBName("loomi").Model(&loomiadmin.ExchangePointsOrder{})
    var exchangePointsOrders []loomiadmin.ExchangePointsOrder
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&exchangePointsOrders).Error
	return  exchangePointsOrders, total, err
}
