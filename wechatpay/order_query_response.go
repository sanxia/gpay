package wechatpay

import (
	"encoding/xml"
)

/* ================================================================================
* OrderQueryResponse数据域结构
* qq group: 582452342
* email   : 2091938785@qq.com
* author  : 美丽的地球啊 - mliu
  <xml>
  <return_code><![CDATA[SUCCESS]]></return_code>
  <return_msg><![CDATA[OK]]></return_msg>
  <appid><![CDATA[wx2421b1c4370ec43b]]></appid>
  <mch_id><![CDATA[10000100]]></mch_id>
  <device_info><![CDATA[1000]]></device_info>
  <nonce_str><![CDATA[TN55wO9Pba5yENl8]]></nonce_str>
  <sign><![CDATA[BDF0099C15FF7BC6B1585FBB110AB635]]></sign>
  <result_code><![CDATA[SUCCESS]]></result_code>
  <openid><![CDATA[oUpF8uN95-Ptaags6E_roPHg7AG0]]></openid>
  <is_subscribe><![CDATA[Y]]></is_subscribe>
  <trade_type><![CDATA[APP]]></trade_type>
  <bank_type><![CDATA[CCB_DEBIT]]></bank_type>
  <total_fee>1</total_fee>
  <fee_type><![CDATA[CNY]]></fee_type>
  <transaction_id><![CDATA[1008450740201411110005820873]]></transaction_id>
  <out_trade_no><![CDATA[1415757673]]></out_trade_no>
  <attach><![CDATA[订单额外描述]]></attach>
  <time_end><![CDATA[20141111170043]]></time_end>
  <trade_state><![CDATA[SUCCESS]]></trade_state>
  </xml>
* ================================================================================ */
type OrderQueryResponse struct {
	XMLName        xml.Name `form:"-" json:"-" xml:"xml"`
	ReturnCode     string   `form:"return_code" json:"return_code" xml:"return_code"`                //必填 返回状态码 SUCCESS
	ReturnMsg      string   `form:"return_msg" json:"return_msg" xml:"return_msg"`                   //必填 返回信息
	AppId          string   `form:"appid" json:"appid" xml:"appid"`                                  //必填 应用APPID
	MchId          string   `form:"mch_id" json:"mch_id" xml:"mch_id"`                               //必填 商户号
	OutTradeNo     string   `form:"out_trade_no" json:"out_trade_no" xml:"out_trade_no"`             //必填 商户系统内部订单号，要求32个字符内，只能是数字、大小写字母_-|*@ ，且在同一个商户号下唯一
	TransactionId  string   `form:"transaction_id" json:"transaction_id" xml:"transaction_id"`       //必填 微信支付订单号
	DeviceInfo     string   `form:"device_info" json:"device_info" xml:"device_info"`                //设备号
	NonceStr       string   `form:"nonce_str" json:"nonce_str" xml:"nonce_str"`                      //必填 随机字符串
	Attach         string   `form:"attach" json:"attach" xml:"attach"`                               //商家数据包，原样返回
	OpenId         string   `form:"open_id" json:"open_id" xml:"open_id"`                            //必填 用户在商户appid下的唯一标识
	CashFee        int      `form:"cash_fee" json:"cash_fee" xml:"cash_fee"`                         //必填 现金支付金额订单现金支付金额
	TotalFee       int      `form:"total_fee" json:"total_fee" xml:"total_fee"`                      //必填 订单总金额，单位为分
	CouponCount    int      `form:"coupon_count" json:"coupon_count" xml:"coupon_count"`             //代金券或立减优惠使用数量
	CouponFee      int      `form:"coupon_fee" json:"coupon_fee" xml:"coupon_fee"`                   //代金券或立减优惠金额<=订单总金额，订单总金额-代金券或立减优惠金额=现金支付金额
	IsSubscribe    string   `form:"is_subscribe" json:"is_subscribe" xml:"is_subscribe"`             //用户是否关注公众账号，Y-关注，N-未关注，仅在公众账号类型支付有效
	TradeState     string   `form:"trade_state" json:"trade_state" xml:"trade_state"`                //必填 SUCCESS—支付成功 REFUND—转入退款 NOTPAY—未支付 CLOSED—已关闭 REVOKED—已撤销（刷卡支付）USERPAYING--用户支付中 PAYERROR--支付失败(其他原因，如银行返回失败)
	TradeStateDesc string   `form:"trade_state_desc" json:"trade_state_desc" xml:"trade_state_desc"` //必填 对当前查询订单状态的描述和下一步操作的指引
	TradeType      string   `form:"trade_type" json:"trade_type" xml:"trade_type"`                   //必填 调用接口提交的交易类型
	BankType       string   `form:"bank_type" json:"bank_type" xml:"bank_type"`                      //必填 银行类型，采用字符串类型的银行标识
	CashFeeType    string   `form:"cash_fee_type" json:"cash_fee_type" xml:"cash_fee_type"`          //现金支付货币类型
	FeeType        string   `form:"fee_type" json:"fee_type" xml:"fee_type"`                         //符合ISO 4217标准的三位字母代码，默认人民币：CNY
	ResultCode     string   `form:"result_code" json:"result_code" xml:"result_code"`                //必填 业务结果
	ErrCode        string   `form:"err_code" json:"err_code" xml:"err_code"`                         //错误代码
	ErrCodeDes     string   `form:"err_code_des" json:"err_code_des" xml:"err_code_des"`             //错误代码描述
	TimeEnd        string   `form:"time_end" json:"time_end" xml:"time_end"`                         //必填 支付完成时间，格式为yyyyMMddHHmmss，如2009年12月25日9点10分10秒表示为20091225091010
	Sign           string   `form:"sign" json:"sign" xml:"sign"`                                     //必填 签名
}
