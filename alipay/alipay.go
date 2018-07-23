package alipay

import (
	"errors"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"
	"time"
)

import (
	"github.com/sanxia/glib"
)

/* ================================================================================
 * 支付宝支付签名工具模块
 * qq group: 582452342
 * email   : 2091938785@qq.com
 * author  : 美丽的地球啊 - mliu
 * ================================================================================ */
type AlipayClient struct {
	appId           string //商户app id
	appPrivate      string //商户app私匙（单行数据，不带-----BEGIN ... KEY-----）
	alipayPublicKey string //阿里支付公匙（单行数据，不带-----BEGIN ... KEY-----）
	sellerId        string //商户支付宝收款账号
	gatewayUrl      string //阿里支付网关地址
	notifyUrl       string //异步通知地址
	timeoutExpress  string //订单过期时间字符串(10m,24h,1d)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 创建Alipay客户端
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func NewAlipayClient(appId, appPrivate, alipayPublicKey string) *AlipayClient {
	alipayClient := new(AlipayClient)
	alipayClient.appId = appId
	alipayClient.appPrivate = appPrivate
	alipayClient.alipayPublicKey = alipayPublicKey
	return alipayClient
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 设置卖家支付宝id（不设置则已申请支付时绑定的支付宝为默认值）
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *AlipayClient) SetSellerId(sellerId string) {
	s.sellerId = sellerId
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 设置网关地址
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *AlipayClient) SetGatewayUrl(gatewayUrl string) {
	s.gatewayUrl = gatewayUrl
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 设置通知地址
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *AlipayClient) SetNotifyUrl(notifyUrl string) {
	s.notifyUrl = notifyUrl
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 设置订单支付过期时间（15m,24h,1d）
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *AlipayClient) SetTimeoutExpress(timeoutExpress string) {
	s.timeoutExpress = timeoutExpress
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取订单字符串给APP支付客户端
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *AlipayClient) GetOrderString(
	outTradeNo, subject, body string,
	amount float64,
	passbackParams string,
	creationDate time.Time) (string, error) {

	appPayRequest := new(AppPayRequest)
	appPayRequest.AppId = s.appId
	appPayRequest.Method = "alipay.trade.app.pay"
	appPayRequest.Format = "json"
	appPayRequest.Charset = "utf-8"
	appPayRequest.SignType = "RSA2"
	appPayRequest.NotifyUrl = s.notifyUrl
	appPayRequest.Timestamp = glib.TimeToString(creationDate)

	//业务内容
	appPayRequestContent := new(AppPayRequestContent)
	appPayRequestContent.SellerId = s.sellerId
	appPayRequestContent.OutTradeNo = outTradeNo
	appPayRequestContent.Subject = subject
	appPayRequestContent.Body = body
	appPayRequestContent.PassbackParams = passbackParams
	appPayRequestContent.TotalAmount = fmt.Sprintf("%.2f", amount)
	appPayRequestContent.ProductCode = "QUICK_MSECURITY_PAY"

	timeoutExpress := "24h"
	if len(s.timeoutExpress) > 0 {
		timeoutExpress = s.timeoutExpress
	}
	appPayRequestContent.TimeoutExpress = timeoutExpress

	appPayRequest.BizContent = appPayRequestContent
	appPayRequest.Version = "1.0"

	//请求对象转字典参数
	paramMap := appPayRequest.ToMap()

	//签名
	sign, err := s.Sign(paramMap)
	if err != nil {
		return "", err
	}

	//base64格式的签名附加到字典参数
	paramMap["sign"] = sign

	//字典kv用&链接起来，v需要url编码
	orderString := glib.JoinMapToString(paramMap, []string{}, true)

	return orderString, nil
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取预创建支付二维码地址
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *AlipayClient) GetOrderQrCode(
	outTradeNo, subject, body string,
	amount float64) (string, error) {

	if len(outTradeNo) == 0 || len(subject) == 0 || amount <= 0.0 {
		return "", errors.New("GetOrderQrCode Args Error")
	}

	preCreateyRequestContent := new(PreCreateyRequestContent)
	preCreateyRequestContent.OutTradeNo = outTradeNo
	preCreateyRequestContent.Subject = subject
	preCreateyRequestContent.Body = body

	timeoutExpress := "24h"
	if len(s.timeoutExpress) > 0 {
		timeoutExpress = s.timeoutExpress
	}
	preCreateyRequestContent.TimeoutExpress = timeoutExpress

	//预创建请求
	preCreateRequest := new(PreCreateRequest)
	preCreateRequest.AppId = s.appId
	preCreateRequest.BizContent = preCreateyRequestContent
	preCreateRequest.Method = "alipay.trade.precreate"
	preCreateRequest.Format = "json"
	preCreateRequest.Charset = "utf-8"
	preCreateRequest.SignType = "RSA2"
	preCreateRequest.NotifyUrl = s.notifyUrl
	preCreateRequest.Timestamp = glib.TimeToString(time.Now())
	preCreateRequest.Version = "1.0"

	//请求对象转字典参数
	paramMap := preCreateRequest.ToMap()

	//签名
	sign, err := s.Sign(paramMap)
	if err != nil {
		return "", err
	}

	//base64格式的签名附加到字典参数
	paramMap["sign"] = sign

	//字典kv用&链接起来，v需要url编码
	orderString := glib.JoinMapToString(paramMap, []string{}, true)

	gatewayUrl := "https://openapi.alipay.com/gateway.do"
	if len(s.gatewayUrl) > 0 {
		gatewayUrl = s.gatewayUrl
	}

	//发起post请求
	respData, err := glib.HttpPost(gatewayUrl, orderString)
	if err != nil {
		return "", err
	}

	log.Printf("UnifiedOrder raw resp: %s", respData)

	var preCreateResponse PreCreateResponse
	if err := glib.FromJson(respData, &preCreateResponse); err != nil {
		return "", err
	}

	qrCode := ""
	if preCreateResponse.AlipayTradePreCreateResponse != nil {
		qrCode = preCreateResponse.AlipayTradePreCreateResponse.QrCode
	}

	log.Printf("UnifiedOrder qrCode: %s", qrCode)

	return qrCode, nil
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 同步验签
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *AlipayClient) ReturnVerify(
	returnResultResp *AppPayReturnResultResponse) (bool, error) {
	err := errors.New("ReturnVerify SignError")
	signString := s.GetReturnResultSignString(returnResultResp.RawResultString)
	sign := returnResultResp.Result.Sign

	if len(signString) == 0 || len(sign) == 0 {
		return false, err
	}

	return glib.Sha256WithRsaVerify(signString, sign, s.alipayPublicKey)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 异步验签
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *AlipayClient) NotifyVerify(dataParams map[string]string) (bool, error) {
	err := errors.New("NotifyVerify SignError")
	if len(dataParams) == 0 {
		return false, err
	}

	outTradeNo, isOutTradeNo := dataParams["out_trade_no"]
	sign, isSign := dataParams["sign"]

	if !isOutTradeNo || !isSign || len(outTradeNo) == 0 || len(sign) == 0 {
		return false, err
	}

	//待签名字符串
	signString := glib.JoinMapToString(dataParams, []string{"sign", "sign_type"}, false)

	return glib.Sha256WithRsaVerify(signString, sign, s.alipayPublicKey)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取签名
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *AlipayClient) Sign(params map[string]string) (string, error) {
	//待签名字符串
	waitingSignString := glib.JoinMapToString(params, []string{"sign"}, false)
	sign, err := glib.Sha256WithRsa(waitingSignString, s.appPrivate)

	return sign, err
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 从同步结果原始字符串获取待签名的字符串
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *AlipayClient) GetReturnResultSignString(returnResultString string) string {
	signString := ""
	patern := `"alipay_trade_app_pay_response":(.*[\}]),`
	if reg, err := regexp.Compile(patern); err == nil {
		results := reg.FindStringSubmatch(returnResultString)
		if len(results) > 0 {
			signString = fmt.Sprintf("%s", results[1])
		}
	}

	return signString
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 从异步通知原始字符串获取通知响应
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *AlipayClient) GetNotifyResponse(rawData string) *AppPayNotifyResponse {
	dataMap := make(map[string]string, 0)
	notifyResponse := new(AppPayNotifyResponse)

	if len(rawData) == 0 {
		return notifyResponse
	}

	//请求参数转换成字典
	datas := strings.Split(rawData, "&")
	for _, data := range datas {
		kvPairs := strings.Split(data, "=")
		key := kvPairs[0]
		value, _ := url.QueryUnescape(kvPairs[1])

		dataMap[key] = value
	}

	//通知数据
	if appId, isOk := dataMap["app_id"]; isOk {
		notifyResponse.AppId = appId
	}

	if subject, isOk := dataMap["subject"]; isOk {
		notifyResponse.Subject = subject
	}

	if body, isOk := dataMap["body"]; isOk {
		notifyResponse.Body = body
	}

	if tradeNo, isOk := dataMap["trade_no"]; isOk {
		notifyResponse.TradeNo = tradeNo
	}

	if outTradeNo, isOk := dataMap["out_trade_no"]; isOk {
		notifyResponse.OutTradeNo = outTradeNo
	}

	if outBizNo, isOk := dataMap["out_biz_no"]; isOk {
		notifyResponse.OutBizNo = outBizNo
	}

	if buyerId, isOk := dataMap["buyer_id"]; isOk {
		notifyResponse.BuyerId = buyerId
	}

	if buyerLogonId, isOk := dataMap["buyer_logon_id"]; isOk {
		notifyResponse.BuyerLogonId = buyerLogonId
	}

	if sellerId, isOk := dataMap["seller_id"]; isOk {
		notifyResponse.SellerId = sellerId
	}

	if sellerEmail, isOk := dataMap["seller_email"]; isOk {
		notifyResponse.SellerEmail = sellerEmail
	}

	if tradeStatus, isOk := dataMap["trade_status"]; isOk {
		notifyResponse.TradeStatus = tradeStatus
	}

	if totalAmountString, isOk := dataMap["total_amount"]; isOk {
		notifyResponse.TotalAmount = totalAmountString
	}

	if receiptAmount, isOk := dataMap["receipt_amount"]; isOk {
		notifyResponse.ReceiptAmount = receiptAmount
	}

	if invoiceAmount, isOk := dataMap["invoice_amount"]; isOk {
		notifyResponse.InvoiceAmount = invoiceAmount
	}

	if buyerPayAmount, isOk := dataMap["buyer_pay_amount"]; isOk {
		notifyResponse.BuyerPayAmount = buyerPayAmount
	}

	if pointAmount, isOk := dataMap["point_amount"]; isOk {
		notifyResponse.PointAmount = pointAmount
	}

	if refundFee, isOk := dataMap["refund_fee"]; isOk {
		notifyResponse.RefundFee = refundFee
	}

	if fundBillList, isOk := dataMap["fund_bill_list"]; isOk {
		notifyResponse.FundBillList = fundBillList
	}

	if voucherDetailList, isOk := dataMap["voucher_detail_list"]; isOk {
		notifyResponse.VoucherDetailList = voucherDetailList
	}

	if passbackParams, isOk := dataMap["passback_params"]; isOk {
		notifyResponse.PassbackParams = passbackParams
	}

	if charset, isOk := dataMap["charset"]; isOk {
		notifyResponse.Charset = charset
	}

	if sign, isOk := dataMap["sign"]; isOk {
		notifyResponse.Sign = sign
	}

	if signType, isOk := dataMap["sign_type"]; isOk {
		notifyResponse.SignType = signType
	}

	if notifyId, isOk := dataMap["notify_id"]; isOk {
		notifyResponse.NotifyId = notifyId
	}

	if notifyType, isOk := dataMap["notify_type"]; isOk {
		notifyResponse.NotifyType = notifyType
	}

	if notifyTime, isOk := dataMap["notify_time"]; isOk {
		notifyResponse.NotifyTime = notifyTime
	}

	if gmtCreate, isOk := dataMap["gmt_create"]; isOk {
		notifyResponse.GmtCreate = gmtCreate
	}

	if gmtPayment, isOk := dataMap["gmt_payment"]; isOk {
		notifyResponse.GmtPayment = gmtPayment
	}

	if gmtRefund, isOk := dataMap["gmt_refund"]; isOk {
		notifyResponse.GmtRefund = gmtRefund
	}

	if gmtClose, isOk := dataMap["gmt_close"]; isOk {
		notifyResponse.GmtClose = gmtClose
	}

	if version, isOk := dataMap["version"]; isOk {
		notifyResponse.Version = version
	}

	return notifyResponse
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取同步验签编码对应的消息描述
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func ReturnCodeToMsg(code string) string {
	msg := "未知错误"
	messageMap := map[string]string{
		"4000": "订单支付失败",
		"5000": "重复请求",
		"6001": "用户中途取消",
		"6002": "网络连接出错",
		"6004": "支付结果未知",
		"8000": "正在处理中",
		"9000": "操作成功",
	}

	if _msg, isOk := messageMap[code]; isOk {
		msg = _msg
	}

	return msg
}
