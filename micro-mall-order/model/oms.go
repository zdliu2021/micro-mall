package model

import (
	"time"
)

type MqMessage struct {
	MessageId     string    `gorm:"column:message_id;type:char(32);primary_key" json:"message_id"`
	Content       string    `gorm:"column:content;type:text" json:"content"`
	ToExchane     string    `gorm:"column:to_exchane;type:varchar(255)" json:"to_exchane"`
	RoutingKey    string    `gorm:"column:routing_key;type:varchar(255)" json:"routing_key"`
	ClassType     string    `gorm:"column:class_type;type:varchar(255)" json:"class_type"`
	MessageStatus int       `gorm:"column:message_status;type:int(1);default:0" json:"message_status"` // 0-新建 1-已发送 2-错误抵达 3-已抵达
	CreateTime    time.Time `gorm:"column:create_time;type:datetime" json:"create_time"`
	UpdateTime    time.Time `gorm:"column:update_time;type:datetime" json:"update_time"`
}

func (m *MqMessage) TableName() string {
	return "mq_message"
}

// 订单
type OmsOrder struct {
	Id                    int64     `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"id"`                  // id
	MemberId              int64     `gorm:"column:member_id;type:bigint(20)" json:"member_id"`                               // member_id
	OrderSn               string    `gorm:"column:order_sn;type:char(64)" json:"order_sn"`                                   // 订单号
	CouponId              int64     `gorm:"column:coupon_id;type:bigint(20)" json:"coupon_id"`                               // 使用的优惠券
	CreateTime            time.Time `gorm:"column:create_time;type:datetime" json:"create_time"`                             // create_time
	MemberUsername        string    `gorm:"column:member_username;type:varchar(200)" json:"member_username"`                 // 用户名
	TotalAmount           float64   `gorm:"column:total_amount;type:decimal(18,4)" json:"total_amount"`                      // 订单总额
	PayAmount             float64   `gorm:"column:pay_amount;type:decimal(18,4)" json:"pay_amount"`                          // 应付总额
	FreightAmount         float64   `gorm:"column:freight_amount;type:decimal(18,4)" json:"freight_amount"`                  // 运费金额
	PromotionAmount       float64   `gorm:"column:promotion_amount;type:decimal(18,4)" json:"promotion_amount"`              // 促销优化金额（促销价、满减、阶梯价）
	IntegrationAmount     float64   `gorm:"column:integration_amount;type:decimal(18,4)" json:"integration_amount"`          // 积分抵扣金额
	CouponAmount          float64   `gorm:"column:coupon_amount;type:decimal(18,4)" json:"coupon_amount"`                    // 优惠券抵扣金额
	DiscountAmount        float64   `gorm:"column:discount_amount;type:decimal(18,4)" json:"discount_amount"`                // 后台调整订单使用的折扣金额
	PayType               int       `gorm:"column:pay_type;type:tinyint(4)" json:"pay_type"`                                 // 支付方式【1->支付宝；2->微信；3->银联； 4->货到付款；】
	SourceType            int       `gorm:"column:source_type;type:tinyint(4)" json:"source_type"`                           // 订单来源[0->PC订单；1->app订单]
	Status                int       `gorm:"column:status;type:tinyint(4)" json:"status"`                                     // 订单状态【0->待付款；1->待发货；2->已发货；3->已完成；4->已关闭；5->无效订单】
	DeliveryCompany       string    `gorm:"column:delivery_company;type:varchar(64)" json:"delivery_company"`                // 物流公司(配送方式)
	DeliverySn            string    `gorm:"column:delivery_sn;type:varchar(64)" json:"delivery_sn"`                          // 物流单号
	AutoConfirmDay        int       `gorm:"column:auto_confirm_day;type:int(11)" json:"auto_confirm_day"`                    // 自动确认时间（天）
	Integration           int       `gorm:"column:integration;type:int(11)" json:"integration"`                              // 可以获得的积分
	Growth                int       `gorm:"column:growth;type:int(11)" json:"growth"`                                        // 可以获得的成长值
	BillType              int       `gorm:"column:bill_type;type:tinyint(4)" json:"bill_type"`                               // 发票类型[0->不开发票；1->电子发票；2->纸质发票]
	BillHeader            string    `gorm:"column:bill_header;type:varchar(255)" json:"bill_header"`                         // 发票抬头
	BillContent           string    `gorm:"column:bill_content;type:varchar(255)" json:"bill_content"`                       // 发票内容
	BillReceiverPhone     string    `gorm:"column:bill_receiver_phone;type:varchar(32)" json:"bill_receiver_phone"`          // 收票人电话
	BillReceiverEmail     string    `gorm:"column:bill_receiver_email;type:varchar(64)" json:"bill_receiver_email"`          // 收票人邮箱
	ReceiverName          string    `gorm:"column:receiver_name;type:varchar(100)" json:"receiver_name"`                     // 收货人姓名
	ReceiverPhone         string    `gorm:"column:receiver_phone;type:varchar(32)" json:"receiver_phone"`                    // 收货人电话
	ReceiverPostCode      string    `gorm:"column:receiver_post_code;type:varchar(32)" json:"receiver_post_code"`            // 收货人邮编
	ReceiverProvince      string    `gorm:"column:receiver_province;type:varchar(32)" json:"receiver_province"`              // 省份/直辖市
	ReceiverCity          string    `gorm:"column:receiver_city;type:varchar(32)" json:"receiver_city"`                      // 城市
	ReceiverRegion        string    `gorm:"column:receiver_region;type:varchar(32)" json:"receiver_region"`                  // 区
	ReceiverDetailAddress string    `gorm:"column:receiver_detail_address;type:varchar(200)" json:"receiver_detail_address"` // 详细地址
	Note                  string    `gorm:"column:note;type:varchar(500)" json:"note"`                                       // 订单备注
	ConfirmStatus         int       `gorm:"column:confirm_status;type:tinyint(4)" json:"confirm_status"`                     // 确认收货状态[0->未确认；1->已确认]
	DeleteStatus          int       `gorm:"column:delete_status;type:tinyint(4)" json:"delete_status"`                       // 删除状态【0->未删除；1->已删除】
	UseIntegration        int       `gorm:"column:use_integration;type:int(11)" json:"use_integration"`                      // 下单时使用的积分
	PaymentTime           time.Time `gorm:"column:payment_time;type:datetime" json:"payment_time"`                           // 支付时间
	DeliveryTime          time.Time `gorm:"column:delivery_time;type:datetime" json:"delivery_time"`                         // 发货时间
	ReceiveTime           time.Time `gorm:"column:receive_time;type:datetime" json:"receive_time"`                           // 确认收货时间
	CommentTime           time.Time `gorm:"column:comment_time;type:datetime" json:"comment_time"`                           // 评价时间
	ModifyTime            time.Time `gorm:"column:modify_time;type:datetime" json:"modify_time"`                             // 修改时间
}

func (m *OmsOrder) TableName() string {
	return "oms_order"
}

// 订单项信息
type OmsOrderItem struct {
	Id                int64   `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"id"`         // id
	OrderId           int64   `gorm:"column:order_id;type:bigint(20)" json:"order_id"`                        // order_id
	OrderSn           string  `gorm:"column:order_sn;type:char(64)" json:"order_sn"`                          // order_sn
	SpuId             int64   `gorm:"column:spu_id;type:bigint(20)" json:"spu_id"`                            // spu_id
	SpuName           string  `gorm:"column:spu_name;type:varchar(255)" json:"spu_name"`                      // spu_name
	SpuPic            string  `gorm:"column:spu_pic;type:varchar(500)" json:"spu_pic"`                        // spu_pic
	SpuBrand          string  `gorm:"column:spu_brand;type:varchar(200)" json:"spu_brand"`                    // 品牌
	CategoryId        int64   `gorm:"column:category_id;type:bigint(20)" json:"category_id"`                  // 商品分类id
	SkuId             int64   `gorm:"column:sku_id;type:bigint(20)" json:"sku_id"`                            // 商品sku编号
	SkuName           string  `gorm:"column:sku_name;type:varchar(255)" json:"sku_name"`                      // 商品sku名字
	SkuPic            string  `gorm:"column:sku_pic;type:varchar(500)" json:"sku_pic"`                        // 商品sku图片
	SkuPrice          float64 `gorm:"column:sku_price;type:decimal(18,4)" json:"sku_price"`                   // 商品sku价格
	SkuQuantity       int     `gorm:"column:sku_quantity;type:int(11)" json:"sku_quantity"`                   // 商品购买的数量
	SkuAttrsVals      string  `gorm:"column:sku_attrs_vals;type:varchar(500)" json:"sku_attrs_vals"`          // 商品销售属性组合（JSON）
	PromotionAmount   float64 `gorm:"column:promotion_amount;type:decimal(18,4)" json:"promotion_amount"`     // 商品促销分解金额
	CouponAmount      float64 `gorm:"column:coupon_amount;type:decimal(18,4)" json:"coupon_amount"`           // 优惠券优惠分解金额
	IntegrationAmount float64 `gorm:"column:integration_amount;type:decimal(18,4)" json:"integration_amount"` // 积分优惠分解金额
	RealAmount        float64 `gorm:"column:real_amount;type:decimal(18,4)" json:"real_amount"`               // 该商品经过优惠后的分解金额
	GiftIntegration   int     `gorm:"column:gift_integration;type:int(11)" json:"gift_integration"`           // 赠送积分
	GiftGrowth        int     `gorm:"column:gift_growth;type:int(11)" json:"gift_growth"`                     // 赠送成长值
}

func (m *OmsOrderItem) TableName() string {
	return "oms_order_item"
}

// 订单操作历史记录
type OmsOrderOperateHistory struct {
	Id          int64     `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"id"` // id
	OrderId     int64     `gorm:"column:order_id;type:bigint(20)" json:"order_id"`                // 订单id
	OperateMan  string    `gorm:"column:operate_man;type:varchar(100)" json:"operate_man"`        // 操作人[用户；系统；后台管理员]
	CreateTime  time.Time `gorm:"column:create_time;type:datetime" json:"create_time"`            // 操作时间
	OrderStatus int       `gorm:"column:order_status;type:tinyint(4)" json:"order_status"`        // 订单状态【0->待付款；1->待发货；2->已发货；3->已完成；4->已关闭；5->无效订单】
	Note        string    `gorm:"column:note;type:varchar(500)" json:"note"`                      // 备注
}

func (m *OmsOrderOperateHistory) TableName() string {
	return "oms_order_operate_history"
}

// 订单退货申请
type OmsOrderReturnApply struct {
	Id             int64     `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"id"`  // id
	OrderId        int64     `gorm:"column:order_id;type:bigint(20)" json:"order_id"`                 // order_id
	SkuId          int64     `gorm:"column:sku_id;type:bigint(20)" json:"sku_id"`                     // 退货商品id
	OrderSn        string    `gorm:"column:order_sn;type:char(32)" json:"order_sn"`                   // 订单编号
	CreateTime     time.Time `gorm:"column:create_time;type:datetime" json:"create_time"`             // 申请时间
	MemberUsername string    `gorm:"column:member_username;type:varchar(64)" json:"member_username"`  // 会员用户名
	ReturnAmount   float64   `gorm:"column:return_amount;type:decimal(18,4)" json:"return_amount"`    // 退款金额
	ReturnName     string    `gorm:"column:return_name;type:varchar(100)" json:"return_name"`         // 退货人姓名
	ReturnPhone    string    `gorm:"column:return_phone;type:varchar(20)" json:"return_phone"`        // 退货人电话
	Status         int       `gorm:"column:status;type:tinyint(1)" json:"status"`                     // 申请状态[0->待处理；1->退货中；2->已完成；3->已拒绝]
	HandleTime     time.Time `gorm:"column:handle_time;type:datetime" json:"handle_time"`             // 处理时间
	SkuImg         string    `gorm:"column:sku_img;type:varchar(500)" json:"sku_img"`                 // 商品图片
	SkuName        string    `gorm:"column:sku_name;type:varchar(200)" json:"sku_name"`               // 商品名称
	SkuBrand       string    `gorm:"column:sku_brand;type:varchar(200)" json:"sku_brand"`             // 商品品牌
	SkuAttrsVals   string    `gorm:"column:sku_attrs_vals;type:varchar(500)" json:"sku_attrs_vals"`   // 商品销售属性(JSON)
	SkuCount       int       `gorm:"column:sku_count;type:int(11)" json:"sku_count"`                  // 退货数量
	SkuPrice       float64   `gorm:"column:sku_price;type:decimal(18,4)" json:"sku_price"`            // 商品单价
	SkuRealPrice   float64   `gorm:"column:sku_real_price;type:decimal(18,4)" json:"sku_real_price"`  // 商品实际支付单价
	Reason         string    `gorm:"column:reason;type:varchar(200)" json:"reason"`                   // 原因
	Description    string    `gorm:"column:description述;type:varchar(500)" json:"description述"`       // 描述
	DescPics       string    `gorm:"column:desc_pics;type:varchar(2000)" json:"desc_pics"`            // 凭证图片，以逗号隔开
	HandleNote     string    `gorm:"column:handle_note;type:varchar(500)" json:"handle_note"`         // 处理备注
	HandleMan      string    `gorm:"column:handle_man;type:varchar(200)" json:"handle_man"`           // 处理人员
	ReceiveMan     string    `gorm:"column:receive_man;type:varchar(100)" json:"receive_man"`         // 收货人
	ReceiveTime    time.Time `gorm:"column:receive_time;type:datetime" json:"receive_time"`           // 收货时间
	ReceiveNote    string    `gorm:"column:receive_note;type:varchar(500)" json:"receive_note"`       // 收货备注
	ReceivePhone   string    `gorm:"column:receive_phone;type:varchar(20)" json:"receive_phone"`      // 收货电话
	CompanyAddress string    `gorm:"column:company_address;type:varchar(500)" json:"company_address"` // 公司收货地址
}

func (m *OmsOrderReturnApply) TableName() string {
	return "oms_order_return_apply"
}

// 退货原因
type OmsOrderReturnReason struct {
	Id         int64     `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"id"` // id
	Name       string    `gorm:"column:name;type:varchar(200)" json:"name"`                      // 退货原因名
	Sort       int       `gorm:"column:sort;type:int(11)" json:"sort"`                           // 排序
	Status     int       `gorm:"column:status;type:tinyint(1)" json:"status"`                    // 启用状态
	CreateTime time.Time `gorm:"column:create_time;type:datetime" json:"create_time"`            // create_time
}

func (m *OmsOrderReturnReason) TableName() string {
	return "oms_order_return_reason"
}

// 订单配置信息
type OmsOrderSetting struct {
	Id                  int64 `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"id"`         // id
	FlashOrderOvertime  int   `gorm:"column:flash_order_overtime;type:int(11)" json:"flash_order_overtime"`   // 秒杀订单超时关闭时间(分)
	NormalOrderOvertime int   `gorm:"column:normal_order_overtime;type:int(11)" json:"normal_order_overtime"` // 正常订单超时时间(分)
	ConfirmOvertime     int   `gorm:"column:confirm_overtime;type:int(11)" json:"confirm_overtime"`           // 发货后自动确认收货时间（天）
	FinishOvertime      int   `gorm:"column:finish_overtime;type:int(11)" json:"finish_overtime"`             // 自动完成交易时间，不能申请退货（天）
	CommentOvertime     int   `gorm:"column:comment_overtime;type:int(11)" json:"comment_overtime"`           // 订单完成后自动好评时间（天）
	MemberLevel         int   `gorm:"column:member_level;type:tinyint(2)" json:"member_level"`                // 会员等级【0-不限会员等级，全部通用；其他-对应的其他会员等级】
}

func (m *OmsOrderSetting) TableName() string {
	return "oms_order_setting"
}

// 支付信息表
type OmsPaymentInfo struct {
	Id              int64     `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"id"`     // id
	OrderSn         string    `gorm:"column:order_sn;type:char(64)" json:"order_sn"`                      // 订单号（对外业务号）
	OrderId         int64     `gorm:"column:order_id;type:bigint(20)" json:"order_id"`                    // 订单id
	AlipayTradeNo   string    `gorm:"column:alipay_trade_no;type:varchar(50)" json:"alipay_trade_no"`     // 支付宝交易流水号
	TotalAmount     float64   `gorm:"column:total_amount;type:decimal(18,4)" json:"total_amount"`         // 支付总金额
	Subject         string    `gorm:"column:subject;type:varchar(200)" json:"subject"`                    // 交易内容
	PaymentStatus   string    `gorm:"column:payment_status;type:varchar(20)" json:"payment_status"`       // 支付状态
	CreateTime      time.Time `gorm:"column:create_time;type:datetime" json:"create_time"`                // 创建时间
	ConfirmTime     time.Time `gorm:"column:confirm_time;type:datetime" json:"confirm_time"`              // 确认时间
	CallbackContent string    `gorm:"column:callback_content;type:varchar(4000)" json:"callback_content"` // 回调内容
	CallbackTime    time.Time `gorm:"column:callback_time;type:datetime" json:"callback_time"`            // 回调时间
}

func (m *OmsPaymentInfo) TableName() string {
	return "oms_payment_info"
}

// 退款信息
type OmsRefundInfo struct {
	Id            int64   `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"id"` // id
	OrderReturnId int64   `gorm:"column:order_return_id;type:bigint(20)" json:"order_return_id"`  // 退款的订单
	Refund        float64 `gorm:"column:refund;type:decimal(18,4)" json:"refund"`                 // 退款金额
	RefundSn      string  `gorm:"column:refund_sn;type:varchar(64)" json:"refund_sn"`             // 退款交易流水号
	RefundStatus  int     `gorm:"column:refund_status;type:tinyint(1)" json:"refund_status"`      // 退款状态
	RefundChannel int     `gorm:"column:refund_channel;type:tinyint(4)" json:"refund_channel"`    // 退款渠道[1-支付宝，2-微信，3-银联，4-汇款]
	RefundContent string  `gorm:"column:refund_content;type:varchar(5000)" json:"refund_content"`
}

func (m *OmsRefundInfo) TableName() string {
	return "oms_refund_info"
}

type UndoLog struct {
	Id           int64     `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"id"`
	BranchId     int64     `gorm:"column:branch_id;type:bigint(20);NOT NULL" json:"branch_id"`
	Xid          string    `gorm:"column:xid;type:varchar(100);NOT NULL" json:"xid"`
	Context      string    `gorm:"column:context;type:varchar(128);NOT NULL" json:"context"`
	RollbackInfo string    `gorm:"column:rollback_info;type:longblob;NOT NULL" json:"rollback_info"`
	LogStatus    int       `gorm:"column:log_status;type:int(11);NOT NULL" json:"log_status"`
	LogCreated   time.Time `gorm:"column:log_created;type:datetime;NOT NULL" json:"log_created"`
	LogModified  time.Time `gorm:"column:log_modified;type:datetime;NOT NULL" json:"log_modified"`
	Ext          string    `gorm:"column:ext;type:varchar(100)" json:"ext"`
}

func (m *UndoLog) TableName() string {
	return "undo_log"
}
