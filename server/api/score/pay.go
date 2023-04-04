package score

import (
	"chatmono/models"
	"chatmono/services/order"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/exce"
	"github.com/xgpc/dsg/frame"
	"gorm.io/gorm"
	"time"
)

type orderPayReq struct {
	ProductID uint32 `json:"product_id"`                          // 商品ID
	Num       int    `json:"num"`                                 // 购买数量
	ScoreNum  int64  `json:"score_num"`                           // 总扣款积分 = num * 单个商品积分
	AddressID uint32 `gorm:"column:address_id" json:"address_id"` // 配送地址ID
}

type orderPayRsp struct {
	ProductOrder *models.ProductOrder `json:"product_order"`
	Score        *models.Score        `json:"score"`
}

// Pay 积分购买
// @Summary      积分购买
// @Description  积分购买
// @Accept       json
// @Produce      json
// @param        root  body  orderPayReq  true  "参数"
// @Tags         积分管理
// @Success      200  {object}  render.Response
// @Router       /api/score/pay [Post]
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

	var info models.ProductOrder
	info.ProductNum = param.Num
	info.ProductID = param.ProductID
	info.AddressID = param.AddressID
	info.CreatedUserID = p.MyId()

	// 下单- 积分订单直接购买,并完成订单

	info.OrderSn = order.Handle.CreatedOrderSN()
	// 该数据改为支付成功时填充
	//info.TransactionID = *paySuccess.TransactionId
	info.CreatedAt = time.Now().Unix()

	// fix: 没有给默认值
	info.PayStatus = models.PayStatusOK
	info.DeliveryStatus = models.DeliveryStatusDefault
	info.RefundStatus = models.RefundStatusDefault
	info.ProfitsharingStatus = models.ProfitsharingStatusFalse
	info.OrderType = models.OrderTypeScore

	// 创建积分扣款数据
	var Score models.Score

	// 判断积分是否够
	var num int64
	err = p.DB().Debug().Model(&models.Score{}).
		Where(models.ScoreColumns.UserID, p.MyId()).
		Select(fmt.Sprintf("SUM(%s)", models.ScoreColumns.ScoreNum)).Find(&num).Error
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}
	if num < param.ScoreNum {
		exce.ThrowSys(exce.CodeRequestError, "当前积分不足, 无法购买商品")
	}

	// 创建数据记录
	err = p.DB().Transaction(func(tx *gorm.DB) error {
		// 创建积分订单
		err = tx.Model(&info).Create(&info).Error
		if err != nil {
			return err
		}

		// 创建积分扣款记录
		Score.ScoreType = models.ScoreTypeShop
		Score.CreatedAt = time.Now().Unix()
		Score.UserID = p.MyId()
		Score.ScoreNum = product.Score * param.Num
		Score.ProductOrderID = info.ID
		if int64(Score.ScoreNum) != param.ScoreNum {
			exce.ThrowSys(exce.CodeRequestError, "商品积分计算错误, 请查证")
		}
		if Score.ScoreNum <= 0 {
			exce.ThrowSys(exce.CodeRequestError, "积分购物, ScoreNum不能为0")
		}
		return tx.Create(&Score).Error
	})

	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	res := orderPayRsp{}
	res.ProductOrder = &info
	res.Score = &Score

	p.SuccessWithData(res)

}
