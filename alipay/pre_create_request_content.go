package alipay

/* ================================================================================
 * PreCreateyRequestContent数据域结构
 * qq group: 582452342
 * email   : 2091938785@qq.com
 * author  : 美丽的地球啊 - mliu
 * ================================================================================ */
type PreCreateyRequestContent struct {
	SellerId             string                   `form:"seller_id" json:"seller_id"`                         //卖家支付宝用户ID。 如果该值为空，则默认为商户签约账号对应的支付宝用户ID
	OutTradeNo           string                   `form:"out_trade_no" json:"out_trade_no"`                   //必填 商户订单号,64个字符以内、只能包含字母、数字、下划线；需保证在商户端不重复
	Subject              string                   `form:"subject" json:"subject"`                             //必填 订单标题
	Body                 string                   `form:"body " json:"body"`                                  //对交易或商品的描述
	BuyerLogonId         string                   `form:"buyer_logon_id" json:"buyer_logon_id"`               //买家支付宝账号
	OperatorId           string                   `form:"operator_id " json:"operator_id"`                    //商户操作员编号
	StoreId              string                   `form:"store_id  " json:"store_id"`                         //商户门店编号
	AlipayStoreId        string                   `form:"alipay_store_id " json:"alipay_store_id"`            //支付宝店铺的门店ID
	TerminalId           string                   `form:"terminal_id   " json:"terminal_id"`                  //商户机具终端编号
	GoodsDetail          []*PreCreateyGoodsDetail `form:"terminal_id   " json:"terminal_id"`                  //订单包含的商品列表信息.Json格式. 其它说明详见：“商品明细说明”
	UndiscountableAmount float64                  `form:"undiscountable_amount" json:"undiscountable_amount"` //不可打折金额. 不参与优惠计算的金额，单位为元，精确到小数点后两位，取值范围[0.01,100000000] 如果该值未传入，但传入了【订单总金额】,【打折金额】，则该值默认为【订单总金额】-【打折金额】
	DiscountableAmount   float64                  `form:"discountable_amount" json:"discountable_amount"`     //可打折金额. 参与优惠计算的金额，单位为元，精确到小数点后两位，取值范围[0.01,100000000] 如果该值未传入，但传入了【订单总金额】，【不可打折金额】则该值默认为【订单总金额】-【不可打折金额】
	TotalAmount          float64                  `form:"total_amount" json:"total_amount"`                   //必填 订单总金额，单位为元，精确到小数点后两位，取值范围[0.01,100000000] 如果同时传入了【打折金额】，【不可打折金额】，【订单总金额】三者，则必须满足如下条件：【订单总金额】=【打折金额】+【不可打折金额】
	TimeoutExpress       string                   `form:"timeout_express" json:"timeout_express"`             //该笔订单允许的最晚付款时间，逾期将关闭交易。取值范围：1m～15d。m-分钟，h-小时，d-天，1c-当天（1c-当天的情况下，无论交易何时创建，都在0点关闭）。 该参数数值不接受小数点， 如 1.5h，可转换为 90m。
}

type PreCreateyGoodsDetail struct {
	GoodsId       string  `form:" goods_id " json:" goods_id"`               //商品的编号
	AlipayGoodsId string  `form:" alipay_goods_id " json:" alipay_goods_id"` //支付宝定义的统一商品编号
	GoodsName     string  `form:"goods_name " json:"goods_name"`             //商品名称
	GoodsCategory string  `form:"goods_category " json:" goods_category"`    //商品类目
	Body          string  `form:"body " json:" body"`                        //商品描述信息
	ShowUrl       string  `form:"show_url " json:" show_url"`                //商品的展示地址
	Quantity      int     `form:"quantity " json:"quantity"`                 //商品数量
	Price         float64 `form:"price " json:"price"`                       //商品单价，单位为元
}
