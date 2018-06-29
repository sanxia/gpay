package gpay

import (
	"time"
)

/* ================================================================================
 * TradeInfo数据域结构
 * api: alipay.open.auth.token.app.query
 * qq group: 582452342
 * email   : 2091938785@qq.com
 * author  : 美丽的地球啊 - mliu
 * ================================================================================ */
type TradeInfo struct {
	OutTradeNo   string    `form:"out_trade_no" json:"out_trade_no"`   //外部唯一交易号
	Subject      string    `form:"subject" json:"subject"`             //主题
	Body         string    `form:"body" json:"body"`                   //内容体
	Amount       float64   `form:"amount" json:"amount"`               //金额
	CreationDate time.Time `form:"creation_date" json:"creation_date"` //创建日期
	Extend       string    `form:"extend" json:"extend"`               //扩展数据
}
