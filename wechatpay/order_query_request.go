package wechatpay

import (
	"encoding/xml"
	"fmt"
	"strings"
)

/* ================================================================================
 * OrderQueryRequest数据域结构
 * qq group: 582452342
 * email   : 2091938785@qq.com
 * author  : 美丽的地球啊 - mliu
 * ================================================================================ */
type OrderQueryRequest struct {
	XMLName       xml.Name `form:"-" json:"-" xml:"xml"`
	AppId         string   `form:"appid" json:"appid" xml:"appid"`                            //必填 微信开放平台审核通过的应用APPID
	MchId         string   `form:"mch_id" json:"mch_id" xml:"mch_id"`                         //必填 微信支付分配的商户号
	OutTradeNo    string   `form:"out_trade_no" json:"out_trade_no" xml:"out_trade_no"`       //必填--二选一 商户系统内部的订单号，当没提供transaction_id时需要传这个
	TransactionId string   `form:"transaction_id" json:"transaction_id" xml:"transaction_id"` //必填--二选一 微信的订单号，优先使用
	NonceStr      string   `form:"nonce_str" json:"nonce_str" xml:"nonce_str"`                //必填 随机字符串，不长于32位
	Sign          string   `form:"sign" json:"sign" xml:"sign"`                               //必填 签名
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 对象转成Xml
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *OrderQueryRequest) ToXml() string {
	mapData := s.ToMap()
	results := make([]string, 0)

	results = append(results, "<xml>")

	for k, v := range mapData {
		if len(v) > 0 {
			value := fmt.Sprintf("%s", v)
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
func (s *OrderQueryRequest) ToMap() map[string]string {
	mapData := map[string]string{
		"appid":     s.AppId,
		"mch_id":    s.MchId,
		"nonce_str": s.NonceStr,
		"sign":      s.Sign,
	}

	if len(s.TransactionId) > 0 {
		mapData["transaction_id"] = s.TransactionId
	} else if len(s.OutTradeNo) > 0 {
		mapData["out_trade_no"] = s.OutTradeNo
	}

	return mapData
}
