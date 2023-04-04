package receiver

import (
	"chatmono/models"
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/exce"
	"github.com/xgpc/dsg/frame"
	"github.com/xgpc/dsg/models/cond"
)

type receiverOrderListReq struct {
	ProfitsharingStatus int `gorm:"column:profitsharing_status" json:"profitsharing_status"` // 分账状态:[0所有, 1 未分账, 2 已分帐]
	PayStatus           int `gorm:"column:pay_status" json:"pay_status"`                     // 支付状态:[0所有, 1未支付, 2已取消或超时, 3支付完成, 4已退款]
	DeliveryStatus      int `gorm:"column:delivery_status" json:"delivery_status"`           // 配送状态:[0所有 1未接单, 2已接单/已发货, 3已送达, 4已取消]
}

// OrderList 我的分账订单
// @Summary      我的分账订单
// @Description  我的分账订单
// @Accept       json
// @Produce      json
// @param        root  body  receiverOrderListReq  true  "参数"
// @param        page               query    int     true  "页数"
// @param        page_size           query    int     true  "页展示条数"
// @Tags         分销管理
// @Success      200  {object}  render.Response
// @Router       /api/receiver/order/list [Post]
func OrderList(ctx iris.Context) {
	p := frame.NewBase(ctx)
	var param receiverOrderListReq
	p.Init(&param)

	var list []models.ProductOrder
	var total int64

	tx := p.DB().Model(&models.ProductOrder{}).
		Preload("Product").
		Where(models.ProductOrderColumns.ReferrerID, p.MyId()).
		Where(models.ProductOrderColumns.PayStatus, models.PayStatusOK)
	if param.ProfitsharingStatus != 0 {
		tx = tx.Where(models.ProductOrderColumns.ProfitsharingStatus, param.ProfitsharingStatus)
	}

	if param.PayStatus != 0 {
		tx = tx.Where(models.ProductOrderColumns.PayStatus, param.PayStatus)
	}

	if param.DeliveryStatus != 0 {
		tx = tx.Where(models.ProductOrderColumns.DeliveryStatus, param.DeliveryStatus)
	}

	err := tx.Count(&total).Scopes(cond.PageByQuery(ctx)).Find(&list).Error
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	p.SuccessWithList(list, total)

}
