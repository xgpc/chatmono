package order

import (
	"chatmono/models"
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/exce"
	"github.com/xgpc/dsg/frame"
	"github.com/xgpc/dsg/models/cond"
)

type orderQueryReq struct {
	ProductID      uint32 `gorm:"column:product_id" json:"product_id"`           // 商品ID 0所有
	UserID         uint32 `json:"user_id"`                                       // 下单人ID 0所有
	PayStatus      int    `gorm:"column:pay_status" json:"pay_status"`           // 支付状态:[0所有, 1未支付, 2已取消或超时, 3支付完成, 4退款]
	DeliveryStatus int    `gorm:"column:delivery_status" json:"delivery_status"` // 配送状态:[0所有, 1未接单, 2已接单/已发货, 3已送达, 4已取消]
	RefundStatus   int    `gorm:"column:refund_status" json:"refund_status"`     // 退款状态:[0所有, 1 正常, 2申请退款, 3驳回退款, 4同意退款]
	OrderType      int    `gorm:"column:order_type" json:"order_type"`           // 订单类型:[0所有, 1在线支付, 2积分支付]
}

type ProductOrderRes struct {
	models.ProductOrder
}

// Query 订单查询
// @Summary      订单查询
// @Description  订单查询
// @Accept       json
// @Produce      json
// @param        root  body  orderQueryReq  true  "参数"
// @param        page               query    int     true  "页数"
// @param        page_size           query    int     true  "页展示条数"
// @Tags         订单管理
// @Success      200  {object}  render.Response{data=models.ProductOrder}
// @Router       /api/order/query [Post]
func Query(ctx iris.Context) {
	p := frame.NewBase(ctx)
	var param orderQueryReq
	p.Init(&param)
	var list []ProductOrderRes
	var total int64

	//m := &models.Product{}

	tx := p.DB().Debug().Model(&ProductOrderRes{}).Preload("Product")

	if param.ProductID != 0 {
		tx = tx.Where(models.ProductOrderColumns.ProductID, param.ProductID)
	}
	if param.UserID != 0 {
		tx = tx.Where(models.ProductOrderColumns.CreatedUserID, param.UserID)
	}
	if param.PayStatus != 0 {
		tx = tx.Where(models.ProductOrderColumns.PayStatus, param.PayStatus)
	}
	if param.DeliveryStatus != 0 {
		tx = tx.Where(models.ProductOrderColumns.DeliveryStatus, param.DeliveryStatus)
	}
	if param.RefundStatus != 0 {
		tx = tx.Where(models.ProductOrderColumns.RefundStatus, param.RefundStatus)
	}

	if param.OrderType != 0 {
		tx = tx.Where(models.ProductOrderColumns.OrderType, param.OrderType)
	}

	err := tx.Count(&total).Scopes(cond.PageByQuery(ctx)).Find(&list).Error
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	p.SuccessWithList(list, total)
}
