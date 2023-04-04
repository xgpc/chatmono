/**
 * @Author: smono
 * @Description:
 * @File:  taskChechUp
 * @Version: 1.0.0
 * @Date: 2022/10/11 11:37
 */

package productOrder

import (
	"chatmono/models"
	"chatmono/services/order"
	"github.com/xgpc/dsg/exce"
	"github.com/xgpc/dsg/frame"
	"github.com/xgpc/dsg/log"
	"time"
)

// 定时查询已支付订单状态
func taskCheckUP() {
	defer func() {
		if e := recover(); e != nil {
			log.Errorf("task Error:%s", e)
		}
	}()

	var list []models.ProductOrder
	now := time.Now()
	// 注意：
	//• 关单没有时间限制，建议在订单生成后间隔几分钟（最短5分钟）再调用关单接口，避免出现订单状态同步不及时导致关单失败。
	// 设置为6分钟
	duration, err := time.ParseDuration("-6m")
	if err != nil {
		panic(err)
	}

	curTime := now.Add(duration).Unix()

	// 定期清理 任务
	frame.MySqlDefault().Model(&models.ProductOrder{}).
		Where(models.ProductOrderColumns.PayStatus, models.PayStatusWait).
		Where(models.ProductOrderColumns.OrderType, models.OrderTypeMoney).
		Where(models.ProductOrderColumns.CreatedAt+" < (?)", curTime).
		Find(&list)
	if len(list) == 0 {
		return
	}

	var num int
	for _, v := range list {
		num++
		// 先查询是否已经付款
		DealCheckUP(&v)

		if v.PayStatus == models.PayStatusWait {
			res := CancelCheckUP(&v)
			if !res {
				log.Error("取消订单失败:%v", v)
			}
		}

	}

	log.Info("Task End Total:", num)
}

// DealCheckUP 处理订单
func DealCheckUP(info *models.ProductOrder) {
	res, err := order.Handle.Query(info.OrderSn)
	if err != nil {
		exce.ThrowSys(exce.CodeSysBusy, err.Error())
	}

	switch *res.Transaction.TradeState {
	case order.SUCCESS:
		if info.PayStatus != models.PayStatusOK {
			info.PayStatus = models.PayStatusOK
			info.TransactionID = *res.TransactionId
			err := db().Model(&info).Where(models.ProductOrderColumns.ID, info.ID).Save(info).Error
			if err != nil {
				exce.ThrowSys(exce.CodeSysBusy, err.Error())
			}
		}
	case order.CLOSED:
		if info.PayStatus == models.PayStatusWait {
			info.PayStatus = models.PayStatusTimeOutOrCancel
			err := db().Model(&info).Where(models.ProductOrderColumns.ID, info.ID).Save(info).Error
			if err != nil {
				exce.ThrowSys(exce.CodeSysBusy, err.Error())
			}
		}
	case order.NOTPAY, order.USERPAYING:
		// 未支付状态的订单暂不处理
	default:
		exce.ThrowSys(exce.CodeSysBusy, "未知状态: "+*res.Transaction.TradeState+",请联系管理员")

	}
}

// CancelCheckUP 取消订单
func CancelCheckUP(info *models.ProductOrder) bool {
	ok, err := order.Handle.CancelOrder(info.OrderSn)
	if err != nil {
		exce.ThrowSys(exce.CodeSysBusy, err.Error())
	}
	if ok {
		info.PayStatus = models.PayStatusTimeOutOrCancel

		err = db().Model(info).Where(models.ProductOrderColumns.ID, info.ID).Save(info).Error
		if err != nil {
			exce.ThrowSys(exce.CodeSysBusy, err.Error())
		}
		return true
	}

	return false
}
