package alipay

/* ================================================================================
 * AppPayReturnResponse数据域结构
 * qq group: 582452342
 * email   : 2091938785@qq.com
 * author  : 美丽的地球啊 - mliu
{
   "memo" : "xxxxx",
   "result" : "{
       \"alipay_trade_app_pay_response\":{
           \"code\":\"10000\",
           \"msg\":\"Success\",
           \"app_id\":\"2014072300007148\",
           \"out_trade_no\":\"081622560194853\",
           \"trade_no\":\"2016081621001004400236957647\",
           \"total_amount\":\"0.01\",
           \"seller_id\":\"2088702849871851\",
           \"charset\":\"utf-8\",
           \"timestamp\":\"2016-10-11 17:43:36\"
       },
       \"sign\":\"NGfStJf3i3ooWBuCDIQSumOpaGBcQz+aoAqyGh3W6EqA/gmyPYwLJ2REFijY9XPTApI9YglZyMw+ZMhd3kb0mh4RAXMrb6mekX4Zu8Nf6geOwIa9kLOnw0IMCjxi4abDIfXhxrXyj********\",
       \"sign_type\":\"RSA2\"
   }",
   "resultStatus" : "9000"
}
* 4000: 订单支付失败
* 5000: 重复请求
* 6001: 用户中途取消
* 6002: 网络连接出错
* 6004: 支付结果未知（有可能已经支付成功），请查询商户订单列表中订单的支付状态
* 8000: 正在处理中，支付结果未知（有可能已经支付成功），请查询商户订单列表中订单的支付状态
* 9000: 订单支付成功
* 10000: 接口调用成功，调用结果请参考具体的API文档所对应的业务返回参数
* 20000: 服务不可用
* 20001: 授权权限不足
* 40001: 缺少必选参数
* 40002: 非法的参数
* 40004: 业务处理失败
* 40006: 权限不足
* ================================================================================ */
type AppPayReturnResultResponse struct {
	Memo            string              `form:"memo" json:"memo"`
	Result          *AppPayReturnResult `form:"result" json:"result"`                       //同步结果对象
	RawResultString string              `form:"raw_result_string" json:"raw_result_string"` //原始同步结果字符串
	ResultStatus    string              `form:"result_status" json:"result_status"`         //同步结果状态（9000: 订单支付成功）
}

type AppPayReturnResult struct {
	AppPayTradeAppPayResponse *AppPayTradeAppPayResponse `form:"alipay_trade_app_pay_response" json:"alipay_trade_app_pay_response"`
	Sign                      string                     `form:"sign" json:"sign"`           //签名字符串
	SignType                  string                     `form:"sign_type" json:"sign_type"` //签名类型
}

type AppPayTradeAppPayResponse struct {
	AppId       string `form:"app_id" json:"app_id"`             //支付宝分配给开发者的应用Id
	TradeNo     string `form:"trade_no" json:"trade_no"`         //该交易在支付宝系统中的交易流水号。最长64位。
	OutTradeNo  string `form:"out_trade_no" json:"out_trade_no"` //商户网站唯一订单号
	SellerId    string `form:"seller_id" json:"seller_id"`       //收款支付宝账号对应的支付宝唯一用户号。以2088开头的纯16位数字
	TotalAmount string `form:"total_amount" json:"total_amount"` //该笔订单的资金总额，单位为RMB-Yuan。取值范围为[0.01,100000000.00]，精确到小数点后两位。
	Code        string `form:"code" json:"code"`                 //结果码
	Msg         string `form:"msg" json:"msg"`                   //处理结果的描述，信息来自于code返回结果的描述
	Charset     string `form:"charset" json:"charset"`           //编码格式
	Timestamp   string `form:"timestamp" json:"timestamp"`       //2016-10-11 17:43:36
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 对象转成字典
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *AppPayTradeAppPayResponse) ToMap() map[string]string {
	mapData := map[string]string{
		"app_id":       s.AppId,
		"trade_no":     s.TradeNo,
		"out_trade_no": s.OutTradeNo,
		"seller_id":    s.SellerId,
		"total_amount": s.TotalAmount,
		"code":         s.Code,
		"msg":          s.Msg,
		"charset":      s.Charset,
		"timestamp":    s.Timestamp,
	}
	return mapData
}
