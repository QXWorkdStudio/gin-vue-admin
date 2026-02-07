package loomiadmin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/loomiadmin"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    loomiadminReq "github.com/flipped-aurora/gin-vue-admin/server/model/loomiadmin/request"
)

type AnchorPointsFlowsService struct {
}

// CreateAnchorPointsFlows 创建anchorPointsFlows表记录
// Author [piexlmax](https://github.com/piexlmax)
func (anchorPointsFlowsService *AnchorPointsFlowsService) CreateAnchorPointsFlows(anchorPointsFlows *loomiadmin.AnchorPointsFlows) (err error) {
	err = global.MustGetGlobalDBByDBName("loomi").Create(anchorPointsFlows).Error
	return err
}

// DeleteAnchorPointsFlows 删除anchorPointsFlows表记录
// Author [piexlmax](https://github.com/piexlmax)
func (anchorPointsFlowsService *AnchorPointsFlowsService)DeleteAnchorPointsFlows(anchorPointsFlows loomiadmin.AnchorPointsFlows) (err error) {
	err = global.MustGetGlobalDBByDBName("loomi").Delete(&anchorPointsFlows).Error
	return err
}

// DeleteAnchorPointsFlowsByIds 批量删除anchorPointsFlows表记录
// Author [piexlmax](https://github.com/piexlmax)
func (anchorPointsFlowsService *AnchorPointsFlowsService)DeleteAnchorPointsFlowsByIds(ids request.IdsReq) (err error) {
	err = global.MustGetGlobalDBByDBName("loomi").Delete(&[]loomiadmin.AnchorPointsFlows{},"id in ?",ids.Ids).Error
	return err
}

// UpdateAnchorPointsFlows 更新anchorPointsFlows表记录
// Author [piexlmax](https://github.com/piexlmax)
func (anchorPointsFlowsService *AnchorPointsFlowsService)UpdateAnchorPointsFlows(anchorPointsFlows loomiadmin.AnchorPointsFlows) (err error) {
	err = global.MustGetGlobalDBByDBName("loomi").Save(&anchorPointsFlows).Error
	return err
}

// GetAnchorPointsFlows 根据id获取anchorPointsFlows表记录
// Author [piexlmax](https://github.com/piexlmax)
func (anchorPointsFlowsService *AnchorPointsFlowsService)GetAnchorPointsFlows(id uint) (anchorPointsFlows loomiadmin.AnchorPointsFlows, err error) {
	err = global.MustGetGlobalDBByDBName("loomi").Where("id = ?", id).First(&anchorPointsFlows).Error
	return
}

// GetAnchorPointsFlowsInfoList 分页获取anchorPointsFlows表记录
// Author [piexlmax](https://github.com/piexlmax)
func (anchorPointsFlowsService *AnchorPointsFlowsService)GetAnchorPointsFlowsInfoList(info loomiadminReq.AnchorPointsFlowsSearch) (list []loomiadmin.AnchorPointsFlows, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.MustGetGlobalDBByDBName("loomi").Model(&loomiadmin.AnchorPointsFlows{})
    var anchorPointsFlowss []loomiadmin.AnchorPointsFlows
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
	
	err = db.Find(&anchorPointsFlowss).Error
	return  anchorPointsFlowss, total, err
}
