package order

import (
	"chatmono/models"
	"chatmono/services/order"
	"github.com/kataras/iris/v12"
	"github.com/shopspring/decimal"
	"github.com/xgpc/dsg/exce"
	"github.com/xgpc/dsg/frame"
	"time"
)

type orderPayReq struct {
	ProductID  uint32  `json:"product_id"` // 商品ID
	Num        int     `json:"num"`        // 购买数量
	OpenID     string  `json:"open_id"`
	Price      float64 `gorm:"column:price" json:"price"`           // 总金额  = num * 商品单价
	AddressID  uint32  `gorm:"column:address_id" json:"address_id"` // 配送地址ID
	ReferrerID uint32  `json:"referrer_id"`                         // 推荐人ID
}

type orderPayRsp struct {
	ProductOrder *models.ProductOrder `json:"product_order"`
	WxPay        *order.ResOrderInfo  `json:"wx_pay"`
}

// Pay 下单
// @Summary      下单
// @Description  下单
// @Accept       json
// @Produce      json
// @param        root  body  orderPayReq  true  "参数"
// @Tags         订单管理
// @Success      200  {object}  render.Response{data=orderPayRsp}
// @Router       /api/order/pay [Post]
func Pay(ctx iris.Context) {
	p := frame.NewBase(ctx)
	var param orderPayReq
	p.Init(&param)

	// 获取商品
	productMgr := models.ProductMgr(p.DB())
	product, err := productMgr.GetFromID(param.ProductID)
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	// 判断金额
	d1 := decimal.NewFromFloat(product.Price).Mul(decimal.NewFromFloat(float64(param.Num)))
	Price, _ := d1.Float64()
	if param.Price != Price {
		exce.ThrowSys(exce.CodeRequestError, "商品金额不对应, 请检查后重新下单")
	}

	var info models.ProductOrder
	info.Price = param.Price
	info.ProductNum = param.Num
	info.ProductID = param.ProductID
	info.AddressID = param.AddressID
	info.ReferrerID = param.ReferrerID
	info.CreatedUserID = p.MyId()

	// 分账标识
	var ProfitSharing bool
	// 测试阶段 默认为true
	ProfitSharing = true

	// 计算分账金额
	if param.ReferrerID != 0 {
		d2 := decimal.NewFromFloat(product.RebateRate).Mul(decimal.NewFromFloat(float64(param.Num)))
		RebateRate, _ := d2.Float64()
		info.ProfitsharingMoney = RebateRate
		ProfitSharing = true
	}

	// 下单
	orderRes := order.Handle.Pay(p.MyId(), Price, param.OpenID, product.ProductName, ProfitSharing)
	if orderRes == nil {
		exce.ThrowSys(exce.CodeRequestError, "微信支付调用失败")
	}

	info.OrderSn = orderRes.OrderSn
	// 该数据改为支付成功时填充
	//info.TransactionID = *paySuccess.TransactionId
	info.CreatedAt = time.Now().Unix()

	// fix: 没有给默认值
	info.PayStatus = models.PayStatusWait
	info.DeliveryStatus = models.DeliveryStatusDefault
	info.RefundStatus = models.RefundStatusDefault
	info.ProfitsharingStatus = models.ProfitsharingStatusFalse
	info.OrderType = models.OrderTypeMoney

	err = p.DB().Model(&info).Create(&info).Error
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	res := orderPayRsp{}
	res.ProductOrder = &info
	res.WxPay = orderRes

	p.SuccessWithData(res)
}
