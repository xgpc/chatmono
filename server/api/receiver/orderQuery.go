package receiver

import (
	"chatmono/models"
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/exce"
	"github.com/xgpc/dsg/frame"
	"github.com/xgpc/dsg/models/cond"
)

type receiverOrderQueryReq struct {
	UserID              uint32 `json:"user_id"`                                                 // 查询指定用户(推荐人UserID), 填0为查询所有
	ProfitsharingStatus int    `gorm:"column:profitsharing_status" json:"profitsharing_status"` // 分账状态:[0所有, 1 未分账, 2 已分帐]
	PayStatus           int    `gorm:"column:pay_status" json:"pay_status"`                     // 支付状态:[0所有, 1未支付, 2已取消或超时, 3支付完成, 4已退款]
	DeliveryStatus      int    `gorm:"column:delivery_status" json:"delivery_status"`           // 配送状态:[0所有 1未接单, 2已接单/已发货, 3已送达, 4已取消]
}

// ReceiverQuery 查询分销单
// @Summary      查询分销单
// @Description  查询分销单(一般只存在两种情况 已分帐or未分账, 但是特殊情况已经退款不在考虑范围内, 因此存在已退款已分帐这种需要人为处理)
// @Accept       json
// @Produce      json
// @param        root  body  receiverOrderQueryReq  true  "参数"
// @param        page               query    int     true  "页数"
// @param        page_size           query    int     true  "页展示条数"
// @Tags         分销管理
// @Success      200  {object}  render.Response
// @Router       /api/receiver/order/query [Post]
func ReceiverQuery(ctx iris.Context) {
	p := frame.NewBase(ctx)
	var param receiverOrderQueryReq
	p.Init(&param)
	var list []models.ProductOrder
	var total int64

	tx := p.DB().Model(&models.ProductOrder{}).Preload("Product").
		Where(models.ProductOrderColumns.PayStatus, models.PayStatusOK)

	if param.UserID != 0 {
		tx = tx.Where(models.ProductOrderColumns.ReferrerID, param.UserID)
	}

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
