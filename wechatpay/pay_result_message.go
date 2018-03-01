package wechatpay

import (
	"encoding/xml"
	"fmt"
	"strings"
)

/* ================================================================================
* PayResultMessage数据域结构
* qq group: 582452342
* email   : 2091938785@qq.com
* author  : 美丽的地球啊 - mliu
  <xml>
  <return_code><![CDATA[SUCCESS]]></return_code>
  <return_msg><![CDATA[OK]]></return_msg>
  </xml>
* ================================================================================ */
type PayResultMessage struct {
	XMLName    xml.Name `xml:"xml"`
	ReturnCode string   `form:"return_code" json:"return_code" xml:"return_code"` //SUCCESS/FAIL SUCCESS表示商户接收通知成功并校验成功
	ReturnMsg  string   `form:"return_msg" json:"return_msg" xml:"return_msg"`    //返回信息，如非空，为错误原因：签名失败 参数格式校验错误
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 对象转成Xml
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *PayResultMessage) ToXml() string {
	mapData := s.ToMap()
	results := make([]string, 0)

	results = append(results, "<xml>")

	for k, v := range mapData {
		if len(v) > 0 {
			value := fmt.Sprintf("<![CDATA[%s]]>", v)
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
func (s *PayResultMessage) ToMap() map[string]string {
	mapData := map[string]string{
		"return_code": s.ReturnCode,
		"return_msg":  s.ReturnMsg,
	}

	return mapData
}
