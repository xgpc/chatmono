package order

import (
	"chatmono/models"
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/exce"
	"github.com/xgpc/dsg/frame"
	"github.com/xgpc/dsg/models/cond"
)

type orderListReq struct {
	ProductID      uint32 `gorm:"column:product_id" json:"product_id"`           // 商品ID 0所有
	PayStatus      int    `gorm:"column:pay_status" json:"pay_status"`           // 支付状态:[0所有 1未支付, 2已取消或超时, 3支付完成, 4退款]
	DeliveryStatus int    `gorm:"column:delivery_status" json:"delivery_status"` // 配送状态:[0所有 1未接单, 2已接单/已发货, 3已送达, 4已取消]
	RefundStatus   int    `gorm:"column:refund_status" json:"refund_status"`     // 退款状态:[1 正常, 2申请退款, 3驳回退款, 4同意退款]
	OrderType      int    `gorm:"column:order_type" json:"order_type"`           // 订单类型:[0所有, 1在线支付, 2积分支付]
}

// List 我的订单
// @Summary      我的订单
// @Description  我的订单
// @Accept       json
// @Produce      json
// @param        root  body  orderListReq  true  "参数"
// @param        page               query    int     true  "页数"
// @param        page_size           query    int     true  "页展示条数"
// @Tags         订单管理
// @Success      200  {object}  render.Response
// @Router       /api/order/list [Post]
func List(ctx iris.Context) {
	p := frame.NewBase(ctx)
	var param orderListReq
	p.Init(&param)

	var list []ProductOrderRes
	var total int64

	tx := p.DB().Debug().Model(&ProductOrderRes{}).Preload("Product").
		Where(models.ProductOrderColumns.CreatedUserID, p.MyId())

	if param.ProductID != 0 {
		tx = tx.Where(models.ProductOrderColumns.ProductID, param.ProductID)
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
