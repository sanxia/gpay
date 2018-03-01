package wechatpay

import (
	"encoding/xml"
)

/* ================================================================================
* UnifiedOrderResponse数据域结构
* qq group: 582452342
* email   : 2091938785@qq.com
* author  : 美丽的地球啊 - mliu
  <xml>
  <return_code><![CDATA[SUCCESS]]></return_code>
  <return_msg><![CDATA[OK]]></return_msg>
  <appid><![CDATA[wx2421b1c4370ec43b]]></appid>
  <mch_id><![CDATA[10000100]]></mch_id>
  <nonce_str><![CDATA[IITRi8Iabbblz1Jc]]></nonce_str>
  <sign><![CDATA[7921E432F65EB8ED0CE9755F0E86D72F]]></sign>
  <result_code><![CDATA[SUCCESS]]></result_code>
  <prepay_id><![CDATA[wx201411101639507cbf6ffd8b0779950874]]></prepay_id>
  <trade_type><![CDATA[APP]]></trade_type>
  </xml>
* ================================================================================ */
type UnifiedOrderResponse struct {
	XMLName    xml.Name `form:"-" json:"-" xml:"xml"`
	ReturnCode string   `form:"return_code" json:"return_code" xml:"return_code"`    //必填 返回状态码 SUCCESS/FAIL此字段是通信标识，非交易标识，交易是否成功需要查看result_code来判断
	ReturnMsg  string   `form:"return_msg" json:"return_msg" xml:"return_msg"`       //返回信息，如非空，为错误原因 签名失败 参数格式校验错误
	AppId      string   `form:"appid" json:"appid" xml:"appid"`                      //必填 应用APPID 在return_code为SUCCESS的时候有返回
	MchId      string   `form:"mch_id" json:"mch_id" xml:"mch_id"`                   //必填 商户号 在return_code为SUCCESS的时候有返回
	DeviceInfo string   `form:"device_info" json:"device_info" xml:"device_info"`    //设备号 在return_code为SUCCESS的时候有返回
	NonceStr   string   `form:"nonce_str" json:"nonce_str" xml:"nonce_str"`          //必填 随机字符串 在return_code为SUCCESS的时候有返回
	PrepayId   string   `form:"prepay_id" json:"prepay_id" xml:"prepay_id"`          //必填 预支付交易会话标识 在return_code和result_code都为SUCCESS的时候有返回
	CodeUrl    string   `form:"code_url" json:"code_url" xml:"code_url"`             //trade_type为NATIVE时有返回，用于生成二维码，展示给用户进行扫码支付
	TradeType  string   `form:"trade_type" json:"trade_type" xml:"trade_type"`       //必填 交易类型 在return_code和result_code都为SUCCESS的时候有返回
	ResultCode string   `form:"result_code" json:"result_code" xml:"result_code"`    //必填 业务结果 在return_code为SUCCESS的时候有返回
	ErrCode    string   `form:"err_code" json:"err_code" xml:"err_code"`             //错误代码 在return_code为SUCCESS的时候有返回
	ErrCodeDes string   `form:"err_code_des" json:"err_code_des" xml:"err_code_des"` //错误代码描述 在return_code为SUCCESS的时候有返回
	Sign       string   `form:"sign" json:"sign" xml:"sign"`                         //必填 签名 在return_code为SUCCESS的时候有返回
}
