package wechatpay

/* ================================================================================
* UnifiedOrderResult数据域结构
* qq group: 582452342
* email   : 2091938785@qq.com
* author  : 美丽的地球啊 - mliu
* ================================================================================ */
type UnifiedOrderResult struct {
	AppId     string `form:"appid" json:"appid" xml:"appid"`             //必填 微信开放平台审核通过的应用APPID
	PartnerId string `form:"partnerid" json:"partnerid" xml:"partnerid"` //必填 微信支付分配的商户号
	PrepayId  string `form:"prepayid" json:"prepayid" xml:"prepayid"`    //必填 微信返回的支付交易会话ID
	Package   string `form:"package" json:"package" xml:"package"`       //必填 暂填写固定值Sign=WXPay
	NonceStr  string `form:"noncestr" json:"noncestr" xml:"noncestr"`    //必填 随机字符串
	Timestamp string `form:"timestamp" json:"timestamp" xml:"timestamp"` //必填 时间戳 1970的秒数
	Sign      string `form:"sign" json:"sign" xml:"sign"`                //必填 签名
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 对象转成字典
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *UnifiedOrderResult) ToMap() map[string]string {
	mapData := map[string]string{
		"appid":     s.AppId,
		"partnerid": s.PartnerId,
		"prepayid":  s.PrepayId,
		"package":   s.Package,
		"noncestr":  s.NonceStr,
		"timestamp": s.Timestamp,
		"sign":      s.Sign,
	}

	return mapData
}
