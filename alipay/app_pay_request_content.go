package alipay

/* ================================================================================
 * AppPayRequestContent数据域结构
 * alipay.trade.app.pay  app支付接口2.0
 * alipay.trade.create 商户通过该接口进行交易的创建下单
 * alipay.trade.close 统一收单交易关闭接口
 * alipay.trade.cancel 统一收单交易撤销接口
 * qq group: 582452342
 * email   : 2091938785@qq.com
 * author  : 美丽的地球啊 - mliu
 * ================================================================================ */
type AppPayRequestContent struct {
	SellerId           string `form:"seller_id" json:"seller_id"`                       //非必填，收款支付宝用户ID。 如果该值为空，则默认为商户签约账号对应的支付宝用户ID，如：2088102147948060
	OutTradeNo         string `form:"out_trade_no" json:"out_trade_no"`                 //必填，商户网站唯一订单号
	Subject            string `form:"subject" json:"subject"`                           //必填，商品的标题/交易标题/订单标题/订单关键字等。
	Body               string `form:"body" json:"body"`                                 //非必填，对一笔交易的具体描述信息。如果是多种商品，请将商品描述字符串累加传给body。
	GoodsType          string `form:"goods_type" json:"goods_type"`                     //非必填，商品主类型 :0-虚拟类商品,1-实物类商品
	ProductCode        string `form:"product_code" json:"product_code"`                 //必填，销售产品码，商家和支付宝签约的产品码，为固定值QUICK_MSECURITY_PAY
	EnablePayChannels  string `form:"enable_pay_channels" json:"enable_pay_channels"`   //非必填，可用渠道，用户只能在指定渠道范围内支付当有多个渠道时用“,”分隔
	DisablePayChannels string `form:"disable_pay_channels" json:"disable_pay_channels"` //非必填，可用渠道，用户只能在指定渠道范围内支付当有多个渠道时用“,”分隔, 注：与disable_pay_channels互斥
	PassbackParams     string `form:"passback_params" json:"passback_params"`           //非必填，公用回传参数，如果请求时传递了该参数，则返回给商户时会回传该参数。支付宝只会在同步返回（包括跳转回商户网站）和异步通知时将该参数原样返回。本参数必须进行UrlEncode之后才可以发送给支付宝。
	PromoParams        string `form:"promo_params" json:"promo_params"`                 //非必填，优惠参数,仅与支付宝协商后可用
	ExtendParams       string `form:"extend_params" json:"extend_params"`               //非必填，业务扩展参数
	StoreId            string `form:"store_id" json:"store_id"`                         //非必填，商户门店编号
	TimeExpire         string `form:"time_expire" json:"time_expire"`                   //绝对超时时间，格式为yyyy-MM-dd HH:mm。
	TimeoutExpress     string `form:"timeout_express" json:"timeout_express"`           //非必填，该笔订单允许的最晚付款时间，逾期将关闭交易。取值范围：1m～15d。m-分钟，h-小时，d-天，1c-当天（1c-当天的情况下，无论交易何时创建，都在0点关闭）。 该参数数值不接受小数点， 如 1.5h，可转换为 90m。
	TotalAmount        string `form:"total_amount" json:"total_amount"`                 //必填，订单总金额，单位为元，精确到小数点后两位，取值范围[0.01,100000000]
}
