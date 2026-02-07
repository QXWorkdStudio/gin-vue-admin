// 自动生成模板AnchorPointsFlows
package loomiadmin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/datatypes"
)

// anchorPointsFlows表 结构体  AnchorPointsFlows
type AnchorPointsFlows struct {
	global.GVA_MODEL
	AnchorUid       string         `json:"anchorUid" form:"anchorUid" gorm:"column:anchor_uid;comment:;size:50;"`                    //anchorUid字段
	ChannelId       string         `json:"channelId" form:"channelId" gorm:"column:channel_id;comment:;size:100;"`                   //channelId字段
	Coin            *int           `json:"coin" form:"coin" gorm:"column:coin;comment:用户消费金币;size:19;"`                              //用户消费金币
	EventId         string         `json:"eventId" form:"eventId" gorm:"column:event_id;comment:;size:100;"`                         //eventId字段
	Ext             datatypes.JSON `json:"ext" form:"ext" gorm:"column:ext;comment:额外字段;" swaggertype:"string"`                      //额外字段
	GuildId         string         `json:"guildId" form:"guildId" gorm:"column:guild_id;comment:;size:50;"`                          //guildId字段
	Num             *int           `json:"num" form:"num" gorm:"column:num;comment:;size:19;"`                                       //num字段
	PayUid          string         `json:"payUid" form:"payUid" gorm:"column:pay_uid;comment:;size:50;"`                             //payUid字段
	PointsAmount    *int           `json:"pointsAmount" form:"pointsAmount" gorm:"column:points_amount;comment:积分收益数;size:19;"`      //积分收益数
	PointsChg       *int           `json:"pointsChg" form:"pointsChg" gorm:"column:points_chg;comment:;size:19;"`                    //pointsChg字段
	PointsChgAfter  *int           `json:"pointsChgAfter" form:"pointsChgAfter" gorm:"column:points_chg_after;comment:;size:19;"`    //pointsChgAfter字段
	PointsChgBefore *int           `json:"pointsChgBefore" form:"pointsChgBefore" gorm:"column:points_chg_before;comment:;size:19;"` //pointsChgBefore字段
	PointsFlowId    int            `json:"pointsFlowId" form:"pointsFlowId" gorm:"column:points_flow_id;comment:积分流水id;size:19;"`    //积分流水id
	Ratio           *int           `json:"ratio" form:"ratio" gorm:"column:ratio;comment:兑换比例 10=10%;size:19;"`                      //兑换比例 10=10%
	Source          string         `json:"source" form:"source" gorm:"column:source;comment:;size:50;"`                              //source字段
	SourceType      string         `json:"sourceType" form:"sourceType" gorm:"column:source_type;comment:;size:50;"`                 //sourceType字段
	Status          int8           `json:"status" form:"status" gorm:"column:status;comment:;"`                                      //status字段
	SubSource       string         `json:"subSource" form:"subSource" gorm:"column:sub_source;comment:;size:50;"`                    //subSource字段
}

// TableName anchorPointsFlows表 AnchorPointsFlows自定义表名 anchor_points_flows
func (AnchorPointsFlows) TableName() string {
	return "anchor_points_flows"
}
