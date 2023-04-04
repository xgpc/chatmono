package models

import (
	"gorm.io/datatypes"
)

/******sql******
CREATE TABLE `admin` (
  `user_id` int unsigned NOT NULL,
  `super` int NOT NULL DEFAULT '0' COMMENT '超级管理员:[0 普通, 1 超管]',
  `rule` json DEFAULT NULL COMMENT '权限规则',
  PRIMARY KEY (`user_id`),
  UNIQUE KEY `admin_user_id_uindex` (`user_id`),
  CONSTRAINT `admin_user_id_fk` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
******sql******/
// Admin [...]
type Admin struct {
	UserID uint32         `gorm:"primaryKey;column:user_id" json:"user_id"`
	User   User           `gorm:"joinForeignKey:user_id;foreignKey:id;references:UserID" json:"user_list"` // 用户表
	Super  int            `gorm:"column:super" json:"super"`                                               // 超级管理员:[0 普通, 1 超管]
	Rule   datatypes.JSON `gorm:"column:rule" json:"rule"`                                                 // 权限规则
}

// TableName get sql table name.获取数据库表名
func (m *Admin) TableName() string {
	return "admin"
}

// AdminColumns get sql column name.获取数据库列名
var AdminColumns = struct {
	UserID string
	Super  string
	Rule   string
}{
	UserID: "user_id",
	Super:  "super",
	Rule:   "rule",
}

/******sql******
CREATE TABLE `dict` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `label` varchar(200) DEFAULT NULL,
  `dict_type` varchar(200) DEFAULT NULL,
  `value_data` json DEFAULT NULL,
  `created_user_id` int unsigned DEFAULT NULL,
  `parent_id` int unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=28 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
******sql******/
// Dict [...]
type Dict struct {
	ID            uint32         `gorm:"primaryKey;column:id" json:"id"`
	Label         string         `gorm:"column:label" json:"label"`
	DictType      string         `gorm:"column:dict_type" json:"dict_type"`
	ValueData     datatypes.JSON `gorm:"column:value_data" json:"value_data"`
	CreatedUserID uint32         `gorm:"column:created_user_id" json:"created_user_id"`
	ParentID      uint32         `gorm:"column:parent_id" json:"parent_id"`
}

// TableName get sql table name.获取数据库表名
func (m *Dict) TableName() string {
	return "dict"
}

// DictColumns get sql column name.获取数据库列名
var DictColumns = struct {
	ID            string
	Label         string
	DictType      string
	ValueData     string
	CreatedUserID string
	ParentID      string
}{
	ID:            "id",
	Label:         "label",
	DictType:      "dict_type",
	ValueData:     "value_data",
	CreatedUserID: "created_user_id",
	ParentID:      "parent_id",
}

/******sql******
CREATE TABLE `product` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '商品ID',
  `product_name` varchar(200) NOT NULL COMMENT '商品名称',
  `price` decimal(10,2) NOT NULL COMMENT '商品售价',
  `product_type` int NOT NULL COMMENT '商品类型: [1销售 2积分]',
  `product_body` text COMMENT '商品数据',
  `product_status` int NOT NULL COMMENT '商品状态:[1上架, 2下架]',
  `created_at` bigint DEFAULT NULL COMMENT '创建时间',
  `updated_at` bigint DEFAULT NULL COMMENT '修改时间',
  `created_user_id` int unsigned NOT NULL COMMENT '创建人',
  `score` int DEFAULT '0' COMMENT '商品积分',
  `rebate_rate` double(10,2) DEFAULT NULL COMMENT '佣金(每份商品可得佣金)',
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='商品'
******sql******/
// Product 商品
type Product struct {
	ID            uint32  `gorm:"primaryKey;column:id" json:"id"`                // 商品ID
	ProductName   string  `gorm:"column:product_name" json:"product_name"`       // 商品名称
	Price         float64 `gorm:"column:price" json:"price"`                     // 商品售价
	ProductType   int     `gorm:"column:product_type" json:"product_type"`       // 商品类型: [1销售 2积分]
	ProductBody   string  `gorm:"column:product_body" json:"product_body"`       // 商品数据
	ProductStatus int     `gorm:"column:product_status" json:"product_status"`   // 商品状态:[1上架, 2下架]
	CreatedAt     int64   `gorm:"column:created_at" json:"created_at"`           // 创建时间
	UpdatedAt     int64   `gorm:"column:updated_at" json:"updated_at"`           // 修改时间
	CreatedUserID uint32  `gorm:"column:created_user_id" json:"created_user_id"` // 创建人
	Score         int     `gorm:"column:score" json:"score"`                     // 商品积分
	RebateRate    float64 `gorm:"column:rebate_rate" json:"rebate_rate"`         // 佣金(每份商品可得佣金)
}

// TableName get sql table name.获取数据库表名
func (m *Product) TableName() string {
	return "product"
}

// ProductColumns get sql column name.获取数据库列名
var ProductColumns = struct {
	ID            string
	ProductName   string
	Price         string
	ProductType   string
	ProductBody   string
	ProductStatus string
	CreatedAt     string
	UpdatedAt     string
	CreatedUserID string
	Score         string
	RebateRate    string
}{
	ID:            "id",
	ProductName:   "product_name",
	Price:         "price",
	ProductType:   "product_type",
	ProductBody:   "product_body",
	ProductStatus: "product_status",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
	CreatedUserID: "created_user_id",
	Score:         "score",
	RebateRate:    "rebate_rate",
}

/******sql******
CREATE TABLE `product_order` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `order_sn` varchar(200) NOT NULL COMMENT '订单编号',
  `price` decimal(10,2) NOT NULL COMMENT '该单商品总金额 (商品单价 * 购买数量)',
  `product_id` int unsigned NOT NULL COMMENT '商品ID',
  `product_num` int DEFAULT NULL COMMENT '商品购买数量',
  `created_user_id` int unsigned NOT NULL COMMENT '下单人ID',
  `pay_status` int NOT NULL DEFAULT '1' COMMENT '支付状态:[1未支付, 2已取消或超时, 3支付完成, 4已退款]',
  `transaction_id` varchar(500) DEFAULT NULL COMMENT '微信支付单号',
  `created_at` bigint NOT NULL COMMENT '下单时间',
  `delivery_status` int NOT NULL DEFAULT '1' COMMENT '配送状态:[1未接单, 2已接单/已发货, 3已送达, 4已取消]',
  `address_id` int unsigned DEFAULT NULL COMMENT '配送地址ID',
  `delivery_user_id` int unsigned NOT NULL COMMENT '配送人ID/发货人ID',
  `delivery_stat_at` bigint DEFAULT NULL COMMENT '接单时间',
  `delivery_end_at` bigint DEFAULT NULL COMMENT '送达时间',
  `open_id` varchar(500) NOT NULL,
  `express_sn` varchar(500) DEFAULT NULL COMMENT '快递单号',
  `delivery_type` int NOT NULL DEFAULT '1' COMMENT '配送类型:[1同城配送, 2快递配送]',
  `refund_status` int NOT NULL DEFAULT '1' COMMENT '退款状态:[1 正常, 2申请退款, 3驳回退款, 4同意退款]',
  `refund_reason` text COMMENT '退款原因',
  `refund_examine_user` int unsigned DEFAULT NULL COMMENT '退款操作人ID',
  `refund_detail` text COMMENT '退款详情',
  `referrer_id` int unsigned DEFAULT NULL COMMENT '推荐人ID',
  `profitsharing_status` int DEFAULT '1' COMMENT '分账状态:[1 未分账, 2 已分帐]',
  `profitsharing_money` decimal(10,2) DEFAULT NULL COMMENT '分账金额(商品分账金 * 购买数量)',
  `order_type` int NOT NULL COMMENT '订单类型:[1在线支付, 2积分支付]',
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`),
  UNIQUE KEY `order_sn` (`order_sn`),
  KEY `product_order_product_id_fk` (`product_id`),
  KEY `product_order_address_id_fk` (`address_id`)
) ENGINE=InnoDB AUTO_INCREMENT=64 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='商品订单'
******sql******/
// ProductOrder 商品订单
type ProductOrder struct {
	ID                  uint32  `gorm:"primaryKey;column:id" json:"id"`
	OrderSn             string  `gorm:"column:order_sn" json:"order_sn"`                 // 订单编号
	Price               float64 `gorm:"column:price" json:"price"`                       // 该单商品总金额 (商品单价 * 购买数量)
	ProductID           uint32  `gorm:"column:product_id" json:"product_id"`             // 商品ID
	ProductNum          int     `gorm:"column:product_num" json:"product_num"`           // 商品购买数量
	CreatedUserID       uint32  `gorm:"column:created_user_id" json:"created_user_id"`   // 下单人ID
	PayStatus           int     `gorm:"column:pay_status" json:"pay_status"`             // 支付状态:[1未支付, 2已取消或超时, 3支付完成, 4已退款]
	TransactionID       string  `gorm:"column:transaction_id" json:"transaction_id"`     // 微信支付单号
	CreatedAt           int64   `gorm:"column:created_at" json:"created_at"`             // 下单时间
	DeliveryStatus      int     `gorm:"column:delivery_status" json:"delivery_status"`   // 配送状态:[1未接单, 2已接单/已发货, 3已送达, 4已取消]
	AddressID           uint32  `gorm:"column:address_id" json:"address_id"`             // 配送地址ID
	DeliveryUserID      uint32  `gorm:"column:delivery_user_id" json:"delivery_user_id"` // 配送人ID/发货人ID
	DeliveryStatAt      int64   `gorm:"column:delivery_stat_at" json:"delivery_stat_at"` // 接单时间
	DeliveryEndAt       int64   `gorm:"column:delivery_end_at" json:"delivery_end_at"`   // 送达时间
	OpenID              string  `gorm:"column:open_id" json:"open_id"`
	ExpressSn           string  `gorm:"column:express_sn" json:"express_sn"`                     // 快递单号
	DeliveryType        int     `gorm:"column:delivery_type" json:"delivery_type"`               // 配送类型:[1同城配送, 2快递配送]
	RefundStatus        int     `gorm:"column:refund_status" json:"refund_status"`               // 退款状态:[1 正常, 2申请退款, 3驳回退款, 4同意退款]
	RefundReason        string  `gorm:"column:refund_reason" json:"refund_reason"`               // 退款原因
	RefundExamineUser   uint32  `gorm:"column:refund_examine_user" json:"refund_examine_user"`   // 退款操作人ID
	RefundDetail        string  `gorm:"column:refund_detail" json:"refund_detail"`               // 退款详情
	ReferrerID          uint32  `gorm:"column:referrer_id" json:"referrer_id"`                   // 推荐人ID
	ProfitsharingStatus int     `gorm:"column:profitsharing_status" json:"profitsharing_status"` // 分账状态:[1 未分账, 2 已分帐]
	ProfitsharingMoney  float64 `gorm:"column:profitsharing_money" json:"profitsharing_money"`   // 分账金额(商品分账金 * 购买数量)
	OrderType           int     `gorm:"column:order_type" json:"order_type"`                     // 订单类型:[1在线支付, 2积分支付]
}

// TableName get sql table name.获取数据库表名
func (m *ProductOrder) TableName() string {
	return "product_order"
}

// ProductOrderColumns get sql column name.获取数据库列名
var ProductOrderColumns = struct {
	ID                  string
	OrderSn             string
	Price               string
	ProductID           string
	ProductNum          string
	CreatedUserID       string
	PayStatus           string
	TransactionID       string
	CreatedAt           string
	DeliveryStatus      string
	AddressID           string
	DeliveryUserID      string
	DeliveryStatAt      string
	DeliveryEndAt       string
	OpenID              string
	ExpressSn           string
	DeliveryType        string
	RefundStatus        string
	RefundReason        string
	RefundExamineUser   string
	RefundDetail        string
	ReferrerID          string
	ProfitsharingStatus string
	ProfitsharingMoney  string
	OrderType           string
}{
	ID:                  "id",
	OrderSn:             "order_sn",
	Price:               "price",
	ProductID:           "product_id",
	ProductNum:          "product_num",
	CreatedUserID:       "created_user_id",
	PayStatus:           "pay_status",
	TransactionID:       "transaction_id",
	CreatedAt:           "created_at",
	DeliveryStatus:      "delivery_status",
	AddressID:           "address_id",
	DeliveryUserID:      "delivery_user_id",
	DeliveryStatAt:      "delivery_stat_at",
	DeliveryEndAt:       "delivery_end_at",
	OpenID:              "open_id",
	ExpressSn:           "express_sn",
	DeliveryType:        "delivery_type",
	RefundStatus:        "refund_status",
	RefundReason:        "refund_reason",
	RefundExamineUser:   "refund_examine_user",
	RefundDetail:        "refund_detail",
	ReferrerID:          "referrer_id",
	ProfitsharingStatus: "profitsharing_status",
	ProfitsharingMoney:  "profitsharing_money",
	OrderType:           "order_type",
}

/******sql******
CREATE TABLE `session_templates` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(200) DEFAULT NULL,
  `body` text,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COMMENT='会话模板'
******sql******/
// SessionTemplates 会话模板
type SessionTemplates struct {
	ID    uint32 `gorm:"primaryKey;column:id" json:"id"`
	Title string `gorm:"column:title" json:"title"`
	Body  string `gorm:"column:body" json:"body"`
}

// TableName get sql table name.获取数据库表名
func (m *SessionTemplates) TableName() string {
	return "session_templates"
}

// SessionTemplatesColumns get sql column name.获取数据库列名
var SessionTemplatesColumns = struct {
	ID    string
	Title string
	Body  string
}{
	ID:    "id",
	Title: "title",
	Body:  "body",
}

/******sql******
CREATE TABLE `user` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `user_name` varchar(20) NOT NULL,
  `mobile_data` blob,
  `user_img` varchar(500) DEFAULT NULL COMMENT '用户头像',
  `area_code` varchar(12) DEFAULT NULL,
  `city` varchar(200) DEFAULT NULL,
  `mobile` varchar(11) DEFAULT NULL,
  `password` blob COMMENT '密码',
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_id_uindex` (`id`),
  UNIQUE KEY `user_name` (`user_name`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户表'
******sql******/
// User 用户表
type User struct {
	ID         uint32 `gorm:"primaryKey;column:id" json:"id"`
	UserName   string `gorm:"column:user_name" json:"user_name"`
	MobileData []byte `gorm:"column:mobile_data" json:"mobile_data"`
	UserImg    string `gorm:"column:user_img" json:"user_img"` // 用户头像
	AreaCode   string `gorm:"column:area_code" json:"area_code"`
	City       string `gorm:"column:city" json:"city"`
	Mobile     string `gorm:"column:mobile" json:"mobile"`
	Password   []byte `gorm:"column:password" json:"password"` // 密码
}

// TableName get sql table name.获取数据库表名
func (m *User) TableName() string {
	return "user"
}

// UserColumns get sql column name.获取数据库列名
var UserColumns = struct {
	ID         string
	UserName   string
	MobileData string
	UserImg    string
	AreaCode   string
	City       string
	Mobile     string
	Password   string
}{
	ID:         "id",
	UserName:   "user_name",
	MobileData: "mobile_data",
	UserImg:    "user_img",
	AreaCode:   "area_code",
	City:       "city",
	Mobile:     "mobile",
	Password:   "password",
}

/******sql******
CREATE TABLE `wx_user` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int unsigned NOT NULL,
  `wx_appid` varchar(200) DEFAULT NULL,
  `open_id` varchar(200) NOT NULL,
  `is_profitsharing` int NOT NULL DEFAULT '1' COMMENT '是否开通分账:[1否, 2是]',
  PRIMARY KEY (`id`),
  UNIQUE KEY `wx_user_id_uindex` (`id`),
  UNIQUE KEY `user_id` (`user_id`),
  CONSTRAINT `wx_user_user_id_fk` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
******sql******/
// WxUser [...]
type WxUser struct {
	ID              uint32 `gorm:"primaryKey;column:id" json:"id"`
	UserID          uint32 `gorm:"column:user_id" json:"user_id"`
	User            User   `gorm:"joinForeignKey:user_id;foreignKey:id;references:UserID" json:"user_list"` // 用户表
	WxAppid         string `gorm:"column:wx_appid" json:"wx_appid"`
	OpenID          string `gorm:"column:open_id" json:"open_id"`
	IsProfitsharing int    `gorm:"column:is_profitsharing" json:"is_profitsharing"` // 是否开通分账:[1否, 2是]
}

// TableName get sql table name.获取数据库表名
func (m *WxUser) TableName() string {
	return "wx_user"
}

// WxUserColumns get sql column name.获取数据库列名
var WxUserColumns = struct {
	ID              string
	UserID          string
	WxAppid         string
	OpenID          string
	IsProfitsharing string
}{
	ID:              "id",
	UserID:          "user_id",
	WxAppid:         "wx_appid",
	OpenID:          "open_id",
	IsProfitsharing: "is_profitsharing",
}
