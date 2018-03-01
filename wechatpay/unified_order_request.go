package wechatpay

import (
	"encoding/xml"
	"fmt"
	"strings"
)

import (
	"github.com/sanxia/glib"
)

/* ================================================================================
 * UnifiedOrderRequest数据域结构
 * qq group: 582452342
 * email   : 2091938785@qq.com
 * author  : 美丽的地球啊 - mliu
 * ================================================================================ */
type UnifiedOrderRequest struct {
	XMLName        xml.Name            `form:"-" json:"-" xml:"xml"`
	AppId          string              `form:"appid" json:"appid" xml:"appid"`                                  //必填 微信开放平台审核通过的应用APPID
	MchId          string              `form:"mch_id" json:"mch_id" xml:"mch_id"`                               //必填 微信支付分配的商户号
	DeviceInfo     string              `form:"device_info" json:"device_info" xml:"device_info"`                //终端设备号(门店号或收银设备ID)，默认请传"WEB"
	NonceStr       string              `form:"nonce_str" json:"nonce_str" xml:"nonce_str"`                      //必填 随机字符串，不长于32位
	Body           string              `form:"body" json:"body" xml:"body"`                                     //必填 商品描述交易字段格式根据不同的应用场景按照以下格式：APP——需传入应用市场上的APP名字-实际商品名称，天天爱消除-游戏充值。
	Detail         *UnifiedOrderDetail `form:"detail" json:"detail" xml:"detail"`                               //商品详细列表，使用Json格式，传输签名前请务必使用CDATA标签将JSON文本串保护起来。
	Attach         string              `form:"attach" json:"attach" xml:"attach"`                               //附加数据，在查询API和支付通知中原样返回，该字段主要用于商户携带订单的自定义数据
	OutTradeNo     string              `form:"out_trade_no" json:"out_trade_no" xml:"out_trade_no"`             //必填 商户系统内部订单号，要求32个字符内，只能是数字、大小写字母_-|*@ ，且在同一个商户号下唯一
	FeeType        string              `form:"fee_type" json:"fee_type" xml:"fee_type"`                         //符合ISO 4217标准的三位字母代码，默认人民币：CNY
	TotalFee       int                 `form:"total_fee" json:"total_fee" xml:"total_fee"`                      //必填 订单总金额，单位为分
	SpbillCreateIp string              `form:"spbill_create_ip" json:"spbill_create_ip" xml:"spbill_create_ip"` //必填 用户端实际ip
	TimeStart      string              `form:"time_start" json:"time_start" xml:"time_start"`                   //订单生成时间，格式为yyyyMMddHHmmss，如2009年12月25日9点10分10秒表示为20091225091010
	TimeExpire     string              `form:"time_expire" json:"time_expire" xml:"time_expire"`                //订单失效时间，格式为yyyyMMddHHmmss，如2009年12月27日9点10分10秒表示为20091227091010，注意：最短失效时间间隔必须大于5分钟
	GoodsTag       string              `form:"goods_tag" json:"goods_tag" xml:"goods_tag"`                      //商品标记，代金券或立减优惠功能的参数
	NotifyUrl      string              `form:"notify_url" json:"notify_url" xml:"notify_url"`                   //必填 接收微信支付异步通知回调地址，通知url必须为直接可访问的url，不能携带参数
	TradeType      string              `form:"trade_type" json:"trade_type" xml:"trade_type"`                   //必填 支付类型(APP)
	LimitPay       string              `form:"limit_pay" json:"limit_pay" xml:"limit_pay"`                      //no_credit--指定不能使用信用卡支付
	Sign           string              `form:"sign" json:"sign" xml:"sign"`                                     //必填 签名
	SignType       string              `form:"sign_type" json:"sign_type" xml:"sign_type"`                      //签名类型，目前支持HMAC-SHA256和MD5，默认为MD5
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 统一下单商品详细项列表
 * 使用Json格式，传输签名前请务必使用CDATA标签将JSON文本串保护起来
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
type UnifiedOrderDetail struct {
	GoodsDetail []*UnifiedOrderDetailItem `form:"detail" json:"detail"`
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 统一下单商品详细项
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
type UnifiedOrderDetailItem struct {
	GoodsId       string `form:"goods_id" json:"goods_id"`             //必填 商品的编号
	WxpayGoodsId  string `form:"wxpay_goods_id" json:"wxpay_goods_id"` //可选 微信支付定义的统一商品编号
	GoodsName     string `form:"goods_name" json:"goods_name"`         //必填 商品名称
	GoodsCategory string `form:"goods_category" json:"goods_category"` //可选 商品类目ID
	Quantity      int    `form:"quantity" json:"quantity"`             //必填 商品数量
	Price         int    `form:"price" json:"price"`                   //必填 商品单价，单位为分
	Body          string `form:"body" json:"body"`                     //可选 商品描述信息
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 对象转成Xml
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *UnifiedOrderRequest) ToXml() string {
	mapData := s.ToMap()
	results := make([]string, 0)

	results = append(results, "<xml>")

	for k, v := range mapData {
		if len(v) > 0 {
			value := fmt.Sprintf("%s", v)
			if k == "detail" {
				value = fmt.Sprintf("<![CDATA[%s]]>", v)
			}
			line := fmt.Sprintf("<%s>%s</%s>", k, value, k)
			results = append(results, line)
		}
	}

	results = append(results, "</xml>")

	return strings.Join(results, "")
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 对象转成字典
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *UnifiedOrderRequest) ToMap() map[string]string {
	mapData := map[string]string{
		"appid":            s.AppId,
		"mch_id":           s.MchId,
		"device_info":      s.DeviceInfo,
		"nonce_str":        s.NonceStr,
		"body":             s.Body,
		"detail":           "",
		"attach":           s.Attach,
		"out_trade_no":     s.OutTradeNo,
		"fee_type":         s.FeeType,
		"total_fee":        fmt.Sprintf("%d", s.TotalFee),
		"spbill_create_ip": s.SpbillCreateIp,
		"time_start":       s.TimeStart,
		"time_expire":      s.TimeExpire,
		"goods_tag":        s.GoodsTag,
		"notify_url":       s.NotifyUrl,
		"trade_type":       s.TradeType,
		"limit_pay":        s.LimitPay,
		"sign":             s.Sign,
		"sign_type":        s.SignType,
	}

	if s.Detail != nil {
		if jsongString, err := glib.ToJson(s.Detail); err == nil {
			mapData["detail"] = jsongString
		}
	}

	return mapData
}
