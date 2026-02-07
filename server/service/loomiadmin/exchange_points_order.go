package loomiadmin

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/loomiadmin"
	loomiadminReq "github.com/flipped-aurora/gin-vue-admin/server/model/loomiadmin/request"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ExchangePointsOrderService struct {
}

const TopicAuditPassed = "exchange_audit_passed"

type AuditPassedMsg struct {
	ExchangeOrderID uint `json:"exchange_order_id"`
}

// CreateExchangePointsOrder 创建exchangePointsOrder表记录
// Author [piexlmax](https://github.com/piexlmax)
func (exchangePointsOrderService *ExchangePointsOrderService) CreateExchangePointsOrder(exchangePointsOrder *loomiadmin.ExchangePointsOrder) (err error) {
	err = global.MustGetGlobalDBByDBName("loomi").Create(exchangePointsOrder).Error
	return err
}

// DeleteExchangePointsOrder 删除exchangePointsOrder表记录
// Author [piexlmax](https://github.com/piexlmax)
func (exchangePointsOrderService *ExchangePointsOrderService) DeleteExchangePointsOrder(exchangePointsOrder loomiadmin.ExchangePointsOrder) (err error) {
	err = global.MustGetGlobalDBByDBName("loomi").Delete(&exchangePointsOrder).Error
	return err
}

// DeleteExchangePointsOrderByIds 批量删除exchangePointsOrder表记录
// Author [piexlmax](https://github.com/piexlmax)
func (exchangePointsOrderService *ExchangePointsOrderService) DeleteExchangePointsOrderByIds(ids request.IdsReq) (err error) {
	err = global.MustGetGlobalDBByDBName("loomi").Delete(&[]loomiadmin.ExchangePointsOrder{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateExchangePointsOrder 更新exchangePointsOrder表记录
// Author [piexlmax](https://github.com/piexlmax)
func (exchangePointsOrderService *ExchangePointsOrderService) UpdateExchangePointsOrder(req loomiadmin.ExchangePointsOrder) (err error) {
	db := global.MustGetGlobalDBByDBName("loomi")
	var order loomiadmin.ExchangePointsOrder
	if err := db.First(&order, req.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("记录未找到")
		}
		return err
	}

	// 状态检查
	reqStatus := req.Status
	if reqStatus != loomiadmin.ExchangeStatusWaitAudit &&
		reqStatus != loomiadmin.ExchangeStatusFail &&
		reqStatus != loomiadmin.ExchangeStatusAuditReject {
		return errors.New("同意、打款失败、打款拒绝之外的状态不可在此修改")
	}

	if order.Status != loomiadmin.ExchangeStatusPending { // Only Pending can be audited
		return errors.New("当前申请状态不允许操作")
	}

	err = db.Transaction(func(tx *gorm.DB) error {
		// 更新字段
		order.Status = req.Status
		order.Desc = req.Desc
		order.OperatorId = req.OperatorId
		order.UpdatedAt = time.Now()

		if err := tx.Save(&order).Error; err != nil {
			return err
		}

		// 同步修改金流状态
		if order.PointsFlowId > 0 {
			if err := tx.Model(&loomiadmin.AnchorPointsFlows{}).Where("id = ?", order.PointsFlowId).Update("status", int8(reqStatus)).Error; err != nil {
				return err
			}
		}

		return nil
	})

	if err == nil && reqStatus == loomiadmin.ExchangeStatusWaitAudit {
		exchangePointsOrderService.SendAuditPassedMsg(order.ID)
	}

	return err
}

func (exchangePointsOrderService *ExchangePointsOrderService) SendAuditPassedMsg(orderID uint) {
	global.GVA_LOG.Info("send audit passed msg", zap.Uint("order_id", orderID))
	msg := AuditPassedMsg{ExchangeOrderID: orderID}
	val, _ := json.Marshal(msg)
	if err := global.GVA_REDIS.Publish(context.Background(), TopicAuditPassed, string(val)).Err(); err != nil {
		global.GVA_LOG.Error("send audit passed msg failed", zap.Error(err), zap.Uint("order_id", orderID))
	}
}

// GetExchangePointsOrder 根据id获取exchangePointsOrder表记录
// Author [piexlmax](https://github.com/piexlmax)
func (exchangePointsOrderService *ExchangePointsOrderService) GetExchangePointsOrder(id uint) (exchangePointsOrder loomiadmin.ExchangePointsOrder, err error) {
	err = global.MustGetGlobalDBByDBName("loomi").Where("id = ?", id).First(&exchangePointsOrder).Error
	return
}

// GetExchangePointsOrderInfoList 分页获取exchangePointsOrder表记录
// Author [piexlmax](https://github.com/piexlmax)
func (exchangePointsOrderService *ExchangePointsOrderService) GetExchangePointsOrderInfoList(info loomiadminReq.ExchangePointsOrderSearch) (list []loomiadmin.ExchangePointsOrder, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("loomi").Model(&loomiadmin.ExchangePointsOrder{})
	var exchangePointsOrders []loomiadmin.ExchangePointsOrder
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&exchangePointsOrders).Error
	return exchangePointsOrders, total, err
}
