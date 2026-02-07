// 自动生成模板ExchangePointsOrder
package loomiadmin

import "github.com/flipped-aurora/gin-vue-admin/server/global"

const (
	ExchangeStatusPending     = 1  // 待处理
	ExchangeStatusWaitAudit   = 2  // 待打款
	ExchangeStatusFail        = 10 // 打款失败
	ExchangeStatusAuditReject = 11 // 后台拒绝
	ExchangeStatusSuccess     = 21 // 成功
)

// exchangePointsOrder表 结构体  ExchangePointsOrder
type ExchangePointsOrder struct {
	global.GVA_MODEL
	AccountName   string  `json:"accountName" form:"accountName" gorm:"column:account_name;comment:用户名;size:50;"`                           //用户名
	AccountNumber string  `json:"accountNumber" form:"accountNumber" gorm:"column:account_number;comment:银行账号、钱包、其他账户标识;size:30;"`          //银行账号、钱包、其他账户标识
	AddPoints     *int    `json:"addPoints" form:"addPoints" gorm:"column:add_points;comment:变动的积分;size:10;"`                               //变动的积分
	Amount        string  `json:"amount" form:"amount" gorm:"column:amount;comment:兑换金额;size:255;"`                                         //兑换金额
	CountryName   string  `json:"countryName" form:"countryName" gorm:"column:country_name;comment:国家;size:64;"`                            //国家
	Currency      string  `json:"currency" form:"currency" gorm:"column:currency;comment:币种;size:3;"`                                       //币种
	Desc          string  `json:"desc" form:"desc" gorm:"column:desc;comment:描述;size:255;"`                                                 //描述
	ExtInfo       *string `json:"extInfo" form:"extInfo" gorm:"column:ext_info;comment:额外信息汇总；;"`                                           //额外信息汇总；
	MainChannel   string  `json:"mainChannel" form:"mainChannel" gorm:"column:main_channel;comment:主渠道;size:50;"`                           //主渠道
	MerchantId    string  `json:"merchantId" form:"merchantId" gorm:"column:merchant_id;comment:商户id;"`                                     //商户id
	OperatorId    *int    `json:"operatorId" form:"operatorId" gorm:"column:operator_id;comment:审核人id;size:19;"`                            //审核人id
	OrderNo       string  `json:"orderNo" form:"orderNo" gorm:"column:order_no;comment:商户订单号;size:64;"`                                     //商户订单号
	PointsAfter   *int    `json:"pointsAfter" form:"pointsAfter" gorm:"column:points_after;comment:修改后积分;size:10;"`                         //修改后积分
	PointsBefore  *int    `json:"pointsBefore" form:"pointsBefore" gorm:"column:points_before;comment:修改前积分余额;size:10;"`                    //修改前积分余额
	PointsFlowId  int     `json:"pointsFlowId" form:"pointsFlowId" gorm:"column:points_flow_id;comment:积分流水id;size:19;"`                    //积分流水id
	ProductId     *int    `json:"productId" form:"productId" gorm:"column:product_id;comment:产品id;size:10;"`                                //产品id
	Status        int     `json:"status" form:"status" gorm:"column:status;comment:对应订单状态 提现状态 1:待处理 2:待打款 10:打款失败 11:后台拒绝 21:成功;size:10;"` //对应订单状态 提现状态 1:待处理 2:待打款 10:打款失败 11:后台拒绝 21:成功
	SubChannel    string  `json:"subChannel" form:"subChannel" gorm:"column:sub_channel;comment:子渠道; bank_code;size:50;"`                   //子渠道; bank_code
	Uid           string  `json:"uid" form:"uid" gorm:"column:uid;comment:用户uid;size:50;"`                                                  //用户uid
}

// TableName exchangePointsOrder表 ExchangePointsOrder自定义表名 exchange_points_order
func (ExchangePointsOrder) TableName() string {
	return "exchange_points_order"
}
