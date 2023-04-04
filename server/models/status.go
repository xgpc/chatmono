/**
 * @Author: smono
 * @Description:
 * @File:  status
 * @Version: 1.0.0
 * @Date: 2022/10/11 23:00
 */

package models

const (
	PayStatusWait            = 1 // 待支付
	PayStatusTimeOutOrCancel = 2 // 超时或取消
	PayStatusOK              = 3 // 已支付
	PayStatusQuit            = 4 // 已退款
)

// 配送状态:[1未接单, 2已接单/已发货, 3已送达, 4已取消]
const (
	DeliveryStatusDefault = 1 //未接单
	DeliveryStatusSend    = 2 //已接单
	DeliveryStatusOk      = 3 //已送达
	DeliveryStatusCancel  = 4 //已取消
)

// 1 正常, 2申请退款, 3驳回退款, 4同意退款]
const (
	RefundStatusDefault     = 1 //正常
	RefundStatusQuitExamine = 2 //申请退款
	RefundStatusQuitRe      = 3 //驳回退款
	RefundStatusQuitOK      = 4 //同意退款
)

const (
	IsDefaultFalse = 1
	IsDefaultTrue  = 2
)

const (
	ScoreTypeSign  = 1 // 打卡
	ScoreTypeShare = 2 // 分享
	ScoreTypeShop  = 3 // 购物
)

// 是否开通分账
const (
	IsProfitsharingFalse = 1
	IsProfitsharingTrue  = 2
)

// 是否已经分账

const (
	ProfitsharingStatusFalse = 1
	ProfitsharingStatusTrue  = 2
)

// 订单类型
const (
	OrderTypeMoney = 1 // 现金
	OrderTypeScore = 2
)
