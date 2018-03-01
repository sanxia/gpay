package alipay

/* ================================================================================
 * PreCreateResponse数据域结构
 * qq group: 582452342
 * email   : 2091938785@qq.com
 * author  : 美丽的地球啊 - mliu
 * ================================================================================ */
type PreCreateResponse struct {
	AlipayTradePreCreateResponse *AlipayTradePreCreateResponse `form:"alipay_trade_precreate_response" json:"alipay_trade_precreate_response"` //当前预下单请求生成的二维码码串，可以用二维码生成工具根据该码串值生成对应的二维码
	Sign                         string                        `form:"sign" json:"sign"`                                                       //签名
}

type AlipayTradePreCreateResponse struct {
	Code       string `form:"code" json:"code"`                 //网关返回码
	Msg        string `form:"msg" json:"msg"`                   //网关返回码描述
	SubCode    string `form:"sub_code" json:"sub_code"`         //业务返回码
	SubMSg     string `form:"sub_msg" json:"sub_msg"`           //业务返回码描述
	OutTradeNo string `form:"out_trade_no" json:"out_trade_no"` //商户的订单号
	QrCode     string `form:"qr_code" json:"qr_code"`           //当前预下单请求生成的二维码码串，可以用二维码生成工具根据该码串值生成对应的二维码
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 对象转成字典
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *AlipayTradePreCreateResponse) ToMap() map[string]string {
	mapData := map[string]string{
		"code":         s.Code,
		"msg":          s.Msg,
		"sub_code":     s.SubCode,
		"sub_msg":      s.SubMSg,
		"out_trade_no": s.OutTradeNo,
		"qr_code":      s.QrCode,
	}

	return mapData
}
