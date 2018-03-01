package alipay

/* ================================================================================
 * AppPayResponse数据域结构
 * qq group: 582452342
 * email   : 2091938785@qq.com
 * author  : 美丽的地球啊 - mliu
 * ================================================================================ */
type AppPayResponse struct {
	TradeNo     string `form:"trade_no" json:"trade_no"`         //该交易在支付宝系统中的交易流水号。
	OutTradeNo  string `form:"out_trade_no" json:"out_trade_no"` //商户网站唯一订单号
	SellerId    string `form:"seller_id" json:"seller_id"`       //收款支付宝账号对应的支付宝唯一用户号。  以2088开头的纯16位数字
	TotalAmount string `form:"total_amount" json:"total_amount"` //该笔订单的资金总额，单位为RMB-Yuan。取值范围为[0.01，100000000.00]，精确到小数点后两位。
}
