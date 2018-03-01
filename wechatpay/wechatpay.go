package wechatpay

import (
	"fmt"
	"log"
	"strings"
	"time"
)

import (
	"github.com/sanxia/glib"
)

/* ================================================================================
 * 微信支付签名工具模块
 * qq group: 582452342
 * email   : 2091938785@qq.com
 * author  : 美丽的地球啊 - mliu
 * ================================================================================ */
type WechatpayClient struct {
	appId           string
	apiSecret       string
	partnerId       string
	feeType         string
	unifiedOrderUrl string
	notifyUrl       string
	timeoutExpress  int
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 创建Wechatpay客户端
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func NewWechatpayClient(appId, partnerId, apiSecret string) *WechatpayClient {
	wechatpayClient := new(WechatpayClient)
	wechatpayClient.appId = appId
	wechatpayClient.partnerId = partnerId
	wechatpayClient.apiSecret = apiSecret
	return wechatpayClient
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 设置货币类型(CNY:人民币)
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *WechatpayClient) SetFeeType(feeType string) {
	s.feeType = feeType
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 设置统一下单地址
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *WechatpayClient) SetUnifiedOrderUrl(unifiedOrderUrl string) {
	s.unifiedOrderUrl = unifiedOrderUrl
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 设置通知地址
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *WechatpayClient) SetNotifyUrl(notifyUrl string) {
	s.notifyUrl = notifyUrl
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 设置订单支付过期时间，单位小时（24）
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *WechatpayClient) SetTimeoutExpress(timeoutExpress int) {
	s.timeoutExpress = timeoutExpress
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 统一下单
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *WechatpayClient) UnifiedOrder(
	outTradeNo, body string,
	amount float64,
	ip string,
	args ...*UnifiedOrderDetail) (*UnifiedOrderResponse, error) {

	//获取xml请求字符串
	xmlString := s.GetUnifiedOrderXml(outTradeNo, body, amount, ip)

	//发起post请求
	unifiedOrderUrl := "https://api.mch.weixin.qq.com/pay/unifiedorder"
	if len(s.unifiedOrderUrl) > 0 {
		unifiedOrderUrl = s.unifiedOrderUrl
	}
	respData, err := glib.HttpPost(unifiedOrderUrl, xmlString)
	if err != nil {
		return nil, err
	}

	log.Printf("UnifiedOrder raw resp: %s", respData)

	//解析Xml数据为对象
	unifiedOrderResponse := new(UnifiedOrderResponse)
	if err := glib.FromXml(respData, unifiedOrderResponse); err != nil {
		return nil, err
	}

	return unifiedOrderResponse, nil
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取统一下单Xml字符串
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *WechatpayClient) GetUnifiedOrderXml(
	outTradeNo, body string,
	amount float64,
	ip string,
	args ...*UnifiedOrderDetail) string {

	var detail *UnifiedOrderDetail
	if len(args) == 1 {
		detail = args[0]
	}

	//过期时间默认24小时
	timeoutExpress := 24
	if s.timeoutExpress > 0 {
		timeoutExpress = s.timeoutExpress
	}

	nowDate := time.Now()
	expiredate := nowDate.Add(time.Duration(timeoutExpress) * time.Hour)
	timeStartString := glib.TimeToString(nowDate, "20060102150405")
	timeExpireString := glib.TimeToString(expiredate, "20060102150405")
	log.Printf("timeExpireString: %s", timeExpireString)

	unifiedOrderRequest := new(UnifiedOrderRequest)
	unifiedOrderRequest.AppId = s.appId
	unifiedOrderRequest.MchId = s.partnerId
	unifiedOrderRequest.DeviceInfo = ""
	unifiedOrderRequest.NonceStr = glib.Guid()
	unifiedOrderRequest.Body = body
	unifiedOrderRequest.Detail = detail
	unifiedOrderRequest.Attach = ""
	unifiedOrderRequest.OutTradeNo = outTradeNo

	feeType := "CNY"
	if len(s.feeType) > 0 {
		feeType = s.feeType
	}
	unifiedOrderRequest.FeeType = feeType

	unifiedOrderRequest.TotalFee = int(amount * 100.0) //元转成分
	unifiedOrderRequest.SpbillCreateIp = ip
	unifiedOrderRequest.TimeStart = timeStartString
	unifiedOrderRequest.TimeExpire = timeExpireString
	unifiedOrderRequest.GoodsTag = ""
	unifiedOrderRequest.NotifyUrl = s.notifyUrl
	unifiedOrderRequest.TradeType = "APP"
	unifiedOrderRequest.LimitPay = ""
	unifiedOrderRequest.SignType = "MD5"

	//签名统一下单请求签名
	sign := s.UnifiedOrderRequestSign(unifiedOrderRequest)
	unifiedOrderRequest.Sign = sign
	log.Printf("UnifiedOrder sign: %s", sign)

	//获取xml请求字符串
	xmlString := unifiedOrderRequest.ToXml()
	log.Printf("UnifiedOrder xmlString: %s", xmlString)

	return xmlString
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取Api统一下单请求签名
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *WechatpayClient) UnifiedOrderRequestSign(
	unifiedOrderRequest *UnifiedOrderRequest) string {
	params := unifiedOrderRequest.ToMap()

	//待签名字符串
	waitingSignString := glib.JoinMapToString(params, []string{"sign"}, false)
	waitingSignString = fmt.Sprintf("%s&key=%s", waitingSignString, s.apiSecret)
	sign := strings.ToUpper(glib.Md5(waitingSignString))
	return sign
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取Api统一下单结果签名
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *WechatpayClient) UnifiedOrderResultSign(
	unifiedOrderResult *UnifiedOrderResult) string {
	params := unifiedOrderResult.ToMap()

	//待签名字符串
	waitingSignString := glib.JoinMapToString(params, []string{"sign"}, false)
	waitingSignString = fmt.Sprintf("%s&key=%s", waitingSignString, s.apiSecret)
	sign := strings.ToUpper(glib.Md5(waitingSignString))
	return sign
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取统一下单响应结果
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *WechatpayClient) GetUnifiedOrderResult(
	unifiedOrderResponse *UnifiedOrderResponse) *UnifiedOrderResult {

	unifiedOrderResult := new(UnifiedOrderResult)
	unifiedOrderResult.AppId = s.appId
	unifiedOrderResult.PartnerId = s.partnerId
	unifiedOrderResult.PrepayId = unifiedOrderResponse.PrepayId
	unifiedOrderResult.Package = "Sign=WXPay"
	unifiedOrderResult.NonceStr = glib.Guid()
	unifiedOrderResult.Timestamp = fmt.Sprintf("%d", glib.UnixTimestamp())

	//签名统一下单响应结果签名
	sign := s.UnifiedOrderResultSign(unifiedOrderResult)
	unifiedOrderResult.Sign = sign
	log.Printf("UnifiedOrderResult sign: %s", sign)

	return unifiedOrderResult
}
