package receiver

import (
	"chatmono/models"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/exce"
	"github.com/xgpc/dsg/frame"
)

type receiverOrderTotalReq struct {
}

type receiverOrderTotalRsp struct {
	Total      float64 `json:"total"`      // 总金额
	Received   float64 `json:"received"`   // 已分销金额
	Unreceived float64 `json:"unreceived"` // 未分销金额
}

// Total 我的分销金汇总
// @Summary      我的分销金汇总
// @Description  我的分销金汇总(未到账金额包含未付款)
// @Accept       json
// @Produce      json
// @param        root  body  receiverOrderTotalReq  true  "参数"
// @Tags         分销管理
// @Success      200  {object}  render.Response
// @Router       /api/receiver/order/total [Post]
func Total(ctx iris.Context) {
	p := frame.NewBase(ctx)
	var param receiverOrderTotalReq
	p.Init(&param)

	var info receiverOrderTotalRsp

	err := p.DB().
		Model(&models.ProductOrder{}).
		Where(models.ProductOrderColumns.ReferrerID, p.MyId()).
		Where(models.ProductOrderColumns.PayStatus, models.PayStatusOK).
		Pluck(
			fmt.Sprintf("if(SUM(%s) is null , 0.0, SUM(%s))",
				models.ProductOrderColumns.ProfitsharingMoney,
				models.ProductOrderColumns.ProfitsharingMoney,
			), &info.Total).Error
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	err = p.DB().
		Model(&models.ProductOrder{}).
		Where(models.ProductOrderColumns.ReferrerID, p.MyId()).
		Where(models.ProductOrderColumns.ProfitsharingStatus, models.ProfitsharingStatusTrue).
		Pluck(
			fmt.Sprintf("if(SUM(%s) is null , 0.0, SUM(%s))",
				models.ProductOrderColumns.ProfitsharingMoney,
				models.ProductOrderColumns.ProfitsharingMoney,
			), &info.Received).Error
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	err = p.DB().
		Model(&models.ProductOrder{}).
		Where(models.ProductOrderColumns.ReferrerID, p.MyId()).
		Where(models.ProductOrderColumns.ProfitsharingStatus, models.ProfitsharingStatusFalse).
		Pluck(
			fmt.Sprintf("if(SUM(%s) is null , 0.0, SUM(%s))",
				models.ProductOrderColumns.ProfitsharingMoney,
				models.ProductOrderColumns.ProfitsharingMoney,
			), &info.Unreceived).Error
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	p.SuccessWithData(info)
}
