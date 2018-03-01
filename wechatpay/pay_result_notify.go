package wechatpay

import (
	"encoding/xml"
)

/* ================================================================================
* PayResultNotify数据域结构
* qq group: 582452342
* email   : 2091938785@qq.com
* author  : 美丽的地球啊 - mliu
  <xml>
  <appid><![CDATA[wx2421b1c4370ec43b]]></appid>
  <attach><![CDATA[支付测试]]></attach>
  <bank_type><![CDATA[CFT]]></bank_type>
  <fee_type><![CDATA[CNY]]></fee_type>
  <is_subscribe><![CDATA[Y]]></is_subscribe>
  <mch_id><![CDATA[10000100]]></mch_id>
  <nonce_str><![CDATA[5d2b6c2a8db53831f7eda20af46e531c]]></nonce_str>
  <openid><![CDATA[oUpF8uMEb4qRXf22hE3X68TekukE]]></openid>
  <out_trade_no><![CDATA[1409811653]]></out_trade_no>
  <result_code><![CDATA[SUCCESS]]></result_code>
  <return_code><![CDATA[SUCCESS]]></return_code>
  <sign><![CDATA[B552ED6B279343CB493C5DD0D78AB241]]></sign>
  <sub_mch_id><![CDATA[10000100]]></sub_mch_id>
  <time_end><![CDATA[20140903131540]]></time_end>
  <total_fee>1</total_fee>
  <trade_type><![CDATA[JSAPI]]></trade_type>
  <transaction_id><![CDATA[1004400740201409030005092168]]></transaction_id>
  </xml>
* ================================================================================ */
type PayResultNotify struct {
	XMLName       xml.Name `form:"-" json:"-" xml:"xml"`
	ReturnCode    string   `form:"return_code" json:"return_code" xml:"return_code"`          //必填 返回状态码 SUCCESS
	ReturnMsg     string   `form:"return_msg" json:"return_msg" xml:"return_msg"`             //返回信息
	AppId         string   `form:"appid" json:"appid" xml:"appid"`                            //必填 应用APPID
	MchId         string   `form:"mch_id" json:"mch_id" xml:"mch_id"`                         //必填 商户号
	OutTradeNo    string   `form:"out_trade_no" json:"out_trade_no" xml:"out_trade_no"`       //必填 商户系统内部订单号，要求32个字符内，只能是数字、大小写字母_-|*@ ，且在同一个商户号下唯一
	TransactionId string   `form:"transaction_id" json:"transaction_id" xml:"transaction_id"` //必填 微信支付订单号
	OpenId        string   `form:"open_id" json:"open_id" xml:"open_id"`                      //必填 用户在商户appid下的唯一标识
	DeviceInfo    string   `form:"device_info" json:"device_info" xml:"device_info"`          //设备号
	NonceStr      string   `form:"nonce_str" json:"nonce_str" xml:"nonce_str"`                //必填 随机字符串
	Attach        string   `form:"attach" json:"attach" xml:"attach"`                         //商家数据包，原样返回
	CouponCount   int      `form:"coupon_count" json:"coupon_count" xml:"coupon_count"`       //代金券或立减优惠使用数量
	CouponFee     int      `form:"coupon_fee" json:"coupon_fee" xml:"coupon_fee"`             //代金券或立减优惠金额<=订单总金额，订单总金额-代金券或立减优惠金额=现金支付金额
	CashFee       int      `form:"cash_fee" json:"cash_fee" xml:"cash_fee"`                   //必填 现金支付金额订单现金支付金额
	TotalFee      int      `form:"total_fee" json:"total_fee" xml:"total_fee"`                //必填 订单总金额，单位为分
	IsSubscribe   string   `form:"is_subscribe" json:"is_subscribe" xml:"is_subscribe"`       //用户是否关注公众账号，Y-关注，N-未关注，仅在公众账号类型支付有效
	TradeType     string   `form:"trade_type" json:"trade_type" xml:"trade_type"`             //必填 交易类型
	BankType      string   `form:"bank_type" json:"bank_type" xml:"bank_type"`                //必填 银行类型，采用字符串类型的银行标识
	CashfeeType   string   `form:"cash_fee_type" json:"cash_fee_type" xml:"cash_fee_type"`    //货币类型，符合ISO4217标准的三位字母代码，默认人民币：CNY
	FeeType       string   `form:"fee_type" json:"fee_type" xml:"fee_type"`                   //货币类型，符合ISO4217标准的三位字母代码，默认人民币：CNY
	ResultCode    string   `form:"result_code" json:"result_code" xml:"result_code"`          //必填 业务结果
	ErrCode       string   `form:"err_code" json:"err_code" xml:"err_code"`                   //错误代码
	ErrCodeDes    string   `form:"err_code_des" json:"err_code_des" xml:"err_code_des"`       //错误代码描述
	TimeEnd       string   `form:"time_end" json:"time_end" xml:"time_end"`                   //必填 支付完成时间，格式为yyyyMMddHHmmss，如2009年12月25日9点10分10秒表示为20091225091010
	Sign          string   `form:"sign" json:"sign" xml:"sign"`                               //必填 签名
}
