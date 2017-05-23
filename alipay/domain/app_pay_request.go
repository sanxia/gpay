package domain

import (
	"github.com/sanxia/glib"
)

/* ================================================================================
 * AppPayRequest数据域结构
 * alipay.trade.app.pay  app支付接口2.0
 * alipay.trade.create 商户通过该接口进行交易的创建下单
 * alipay.trade.close 统一收单交易关闭接口
 * alipay.trade.cancel 统一收单交易撤销接口
 * author: mliu
 * ================================================================================ */
type AppPayRequest struct {
	AppId      string                `form:"app_id" json:"app_id"`           //支付宝分配给开发者的应用ID,2014072300007148
	BizContent *AppPayRequestContent `form:"biz_content" json:"biz_content"` //业务请求参数的集合，最大长度不限，除公共参数外所有请求参数都必须放在这个参数中传递，具体参照各产品快速接入文档
	Method     string                `form:"method" json:"method"`           //接口名称,alipay.trade.app.pay
	Format     string                `form:"format" json:"format"`           //仅支持JSON,JSON
	Charset    string                `form:"charset" json:"charset"`         //请求使用的编码格式，如utf-8,gbk,gb2312等,utf-8
	Sign       string                `form:"sign" json:"sign"`               //商户请求参数的签名串
	SignType   string                `form:"sign_type" json:"sign_type"`     //商户生成签名字符串所使用的签名算法类型，目前支持RSA2和RSA，推荐使用RSA2,RSA2
	NotifyUrl  string                `form:"notify_url" json:"notify_url"`   //支付宝服务器主动通知商户服务器里指定的页面http/https路径。建议商户使用https
	Timestamp  string                `form:"timestamp" json:"timestamp"`     //发送请求的时间，格式"yyyy-MM-dd HH:mm:ss",2014-07-24 03:07:50
	Version    string                `form:"version" json:"version"`         //调用的接口版本，固定为：1.0
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 对象转成字典
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *AppPayRequest) ToMap() map[string]string {
	mapData := map[string]string{
		"app_id":     s.AppId,
		"method":     s.Method,
		"format":     s.Format,
		"charset":    s.Charset,
		"sign":       s.Sign,
		"sign_type":  s.SignType,
		"notify_url": s.NotifyUrl,
		"timestamp":  s.Timestamp,
		"version":    s.Version,
	}

	bizcontent := ""
	if s.BizContent != nil {
		if jsonString, err := glib.ToJson(s.BizContent); err == nil {
			bizcontent = jsonString
		}
	}

	mapData["biz_content"] = bizcontent

	return mapData
}
