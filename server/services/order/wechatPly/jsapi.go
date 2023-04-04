// Package wechatPly Package orderPay
// @Author:        asus
// @Description:   $
// @File:          jsapi
// @Data:          2021/12/2510:25
//
package wechatPly

import (
	"context"
	"errors"
	"fmt"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/verifiers"
	"github.com/wechatpay-apiv3/wechatpay-go/core/downloader"
	"github.com/wechatpay-apiv3/wechatpay-go/core/notify"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments"
	jsapisvc "github.com/wechatpay-apiv3/wechatpay-go/services/payments/jsapi"
	"github.com/wechatpay-apiv3/wechatpay-go/services/refunddomestic"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
	"net/http"
)

type JsApiPay struct {
	AppId       string
	MchID       string
	mchAPIv3Key string
	notifyUrl   string
	ctx         context.Context
	client      *core.Client
	svc         *jsapisvc.JsapiApiService
	refundSvc   *refunddomestic.RefundsApiService
}

//NewJsapi 初始化
func NewJsapi(mchID, mchCertificateSerialNumber, mchAPIv3Key, mchPrivateKeyPath, notifyUrl, appid string) (*JsApiPay, error) {
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath(mchPrivateKeyPath)
	if err != nil {
		return nil, errors.New("商户密钥获取失败:" + err.Error())
	}

	ctx := context.Background()
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(mchID, mchCertificateSerialNumber, mchPrivateKey, mchAPIv3Key),
	}

	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		return nil, errors.New("初始化微信JSAPI支付配置失败:" + err.Error())
	}

	return &JsApiPay{
		AppId:       appid,
		MchID:       mchID,
		mchAPIv3Key: mchAPIv3Key,
		notifyUrl:   notifyUrl,
		ctx:         ctx,
		client:      client,
		svc:         &jsapisvc.JsapiApiService{Client: client},
		refundSvc:   &refunddomestic.RefundsApiService{Client: client},
	}, nil
}

//PayUnifiedOrder 下单获取支付凭证
func (client *JsApiPay) PayUnifiedOrder(order *OrderRequest) (*jsapisvc.PrepayWithRequestPaymentResponse, error) {
	if order.OpenId == "" {
		return nil, errors.New("用户OpenId不能为空")
	}
	if order.Desc == "" {
		return nil, errors.New("请填写商品名称")
	}

	prepayResp, _, err := client.svc.PrepayWithRequestPayment(client.ctx, jsapisvc.PrepayRequest{
		Appid:       &client.AppId,
		Mchid:       &client.MchID,
		Description: &order.Desc,
		OutTradeNo:  &order.OrderSn,
		TimeExpire:  order.GetTimeExpire(),
		Attach:      &order.Attach,
		NotifyUrl:   &client.notifyUrl,
		Amount: &jsapisvc.Amount{
			Total: order.GetInt64Price(),
		},
		Payer: &jsapisvc.Payer{
			Openid: &order.OpenId,
		},
		SettleInfo: &jsapisvc.SettleInfo{
			ProfitSharing: core.Bool(order.ProfitSharing),
		},
	})

	if err != nil {
		return nil, err
	}
	return prepayResp, nil
}

// NewParseNotifyRequest 解密和验签
func (client *JsApiPay) NewParseNotifyRequest(request *http.Request) (*PaySuccess, error) {

	certVisitor := downloader.MgrInstance().GetCertificateVisitor(client.MchID)
	handler := notify.NewNotifyHandler(client.mchAPIv3Key, verifiers.NewSHA256WithRSAVerifier(certVisitor))

	transaction := new(payments.Transaction)
	notifyReq, err := handler.ParseNotifyRequest(client.ctx, request, transaction)
	//_, err := handler.ParseNotifyRequest(client.ctx, request, transaction)
	// 如果验签未通过，或者解密失败
	if err != nil {
		return nil, errors.New("验签或解密失败:" + err.Error())
	}

	fmt.Printf("支付解析内容：%+v", notifyReq)
	return &PaySuccess{Transaction: transaction}, nil
}

// QueryOrder 订单查询
func (client *JsApiPay) QueryOrder(orderSn string) (*PaySuccess, error) {
	resp, _, err := client.svc.QueryOrderByOutTradeNo(client.ctx,
		jsapisvc.QueryOrderByOutTradeNoRequest{
			OutTradeNo: &orderSn,
			Mchid:      &client.MchID,
		},
	)
	// 如果验签未通过，或者解密失败
	if err != nil {
		return nil, errors.New("订单查询微信接口请求失败" + err.Error())
	}

	return &PaySuccess{Transaction: resp}, nil
}

// CancelOrder 订单关闭
func (client *JsApiPay) CancelOrder(orderSn string) (bool, error) {
	_, err := client.svc.CloseOrder(client.ctx, jsapisvc.CloseOrderRequest{
		OutTradeNo: &orderSn,
		Mchid:      &client.MchID,
	})

	if err != nil {
		return false, err
	}

	return true, nil
}

// Refund 订单退款
func (client *JsApiPay) Refund(orderSn, outRefundSn, reason string, totalPrice, refundPrice uint32) (*refunddomestic.Refund, error) {

	resp, result, err := client.refundSvc.Create(client.ctx,
		refunddomestic.CreateRequest{
			OutTradeNo:  &orderSn,
			OutRefundNo: &outRefundSn,
			Reason:      &reason,
			Amount: &refunddomestic.AmountReq{
				Currency: core.String("CNY"),
				Refund:   core.Int64(int64(refundPrice)),
				Total:    core.Int64(int64(totalPrice)),
			},
		},
	)

	if err != nil {
		//TODO 日志记录
		return nil, errors.New(fmt.Sprintf("退款失败原因:%s,状态:%d", result.Response.Body, result.Response.StatusCode))
	}

	return resp, nil
}

// RefundQuery 订单退款查询
func (client *JsApiPay) RefundQuery(refundSn string) (*refunddomestic.Refund, error) {
	resp, result, err := client.refundSvc.QueryByOutRefundNo(client.ctx,
		refunddomestic.QueryByOutRefundNoRequest{
			OutRefundNo: &refundSn,
		},
	)

	if err != nil {
		//TODO 日志记录
		return nil, errors.New(fmt.Sprintf("退款失败原因:%s,状态:%d", result.Response.Body, result.Response.StatusCode))
	}

	return resp, nil
}
