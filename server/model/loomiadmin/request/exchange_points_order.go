package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/loomiadmin"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type ExchangePointsOrderSearch struct{
    loomiadmin.ExchangePointsOrder
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}
