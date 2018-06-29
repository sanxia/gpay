# gpay
golang payment for alipay, wechatpay

Alipay, Wechatpay Golang Client
===========================

AlipayClient Api
---------------------------
* GetOrderString
* GetOrderQrCode
* ReturnVerify
* NotifyVerify
---------------------------

    import (
        "time"
    )
    import (
        "github.com/sanxia/gpay/alipay"
    )

    var alipayClient *alipay.AlipayClient

    func init(){
        appId := "you app id"
        appPrivate := "you app single line private"
        alipayPublicKey := "alipay single line public"
        alipayClient = alipay.NewAlipayClient(appId, appPrivate, alipayPublicKey)
    }

    func AlipayOrder() (string, error){
        outTradeNo := "test-1234567890"
        subject := "mall produce test123"
        body := "mall product info test123"
        amount := 0.01
        passbackParams := ""
        creationDate := time.Now()
        orderString, err := alipayClient.GetOrderString(
            outTradeNo,
            subject,
            body,
            amount,
            passbackParams,
            creationDate)

        return orderString, nil
    }

    func AlipayQrOrder() (string, error){
        outTradeNo := "test-1234567890"
        subject := "mall produce test123"
        body := "mall product info test123"
        amount := 0.01
        orderQrCode, err := alipayClient.GetOrderQrCode(
            outTradeNo,
            subject,
            body,
            amount)

        return orderQrCode, nil
    }


WechatpayClient Api
---------------------------
* UnifiedOrder
---------------------------
    import (
        "time"
    )
    import (
        "github.com/sanxia/gpay/wechatpay"
    )

    var wechatpayClient *wechatpay.WechatpayClient

    func init(){
        appId := "you app id"
        partnerId := "you partner id"
        apiSecret := "you api secret"
        wechatpayClient = wechatpay.NewWechatpayClient(appId, partnerId, apiSecret)
    }

    func WechatpayUnifiedOrder() (*wechatpay.UnifiedOrderResponse, error){
        outTradeNo := "test-1234567890"
        body := "mall product info test123"
        amount := 1
        attach := ""
        ip := "127.0.0.1"

        goodsDetail := &wechatpay.UnifiedOrderDetailItem{
            GoodsId: "test123",
            GoodsName: "test product",
            Quantity: 1,
            Price: amount,
        }
        unifiedOrderDetail := &wechatpay
        .UnifiedOrderDetail{
            GoodsDetail: []*UnifiedOrderDetailItem{
                goodsDetail
            },
        }

        return wechatpayClient.UnifiedOrder(
            outTradeNo,
            body,
            amount,
            attach,
            ip,
            unifiedOrderDetail)
    }
