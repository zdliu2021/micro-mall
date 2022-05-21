package model

import (
	"time"
)

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

// 采购信息
type WmsPurchase struct {
	Id           int64     `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"id"`
	AssigneeId   int64     `gorm:"column:assignee_id;type:bigint(20)" json:"assignee_id"`
	AssigneeName string    `gorm:"column:assignee_name;type:varchar(255)" json:"assignee_name"`
	Phone        string    `gorm:"column:phone;type:char(13)" json:"phone"`
	Priority     int       `gorm:"column:priority;type:int(4)" json:"priority"`
	Status       int       `gorm:"column:status;type:int(4)" json:"status"`
	WareId       int64     `gorm:"column:ware_id;type:bigint(20)" json:"ware_id"`
	Amount       float64   `gorm:"column:amount;type:decimal(18,4)" json:"amount"`
	CreateTime   time.Time `gorm:"column:create_time;type:datetime" json:"create_time"`
	UpdateTime   time.Time `gorm:"column:update_time;type:datetime" json:"update_time"`
}

func (m *WmsPurchase) TableName() string {
	return "wms_purchase"
}

type WmsPurchaseDetail struct {
	Id         int64   `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"id"`
	PurchaseId int64   `gorm:"column:purchase_id;type:bigint(20)" json:"purchase_id"` // 采购单id
	SkuId      int64   `gorm:"column:sku_id;type:bigint(20)" json:"sku_id"`           // 采购商品id
	SkuNum     int     `gorm:"column:sku_num;type:int(11)" json:"sku_num"`            // 采购数量
	SkuPrice   float64 `gorm:"column:sku_price;type:decimal(18,4)" json:"sku_price"`  // 采购金额
	WareId     int64   `gorm:"column:ware_id;type:bigint(20)" json:"ware_id"`         // 仓库id
	Status     int     `gorm:"column:status;type:int(11)" json:"status"`              // 状态[0新建，1已分配，2正在采购，3已完成，4采购失败]
}

func (m *WmsPurchaseDetail) TableName() string {
	return "wms_purchase_detail"
}

// 仓库信息

type WmsWareInfo struct {
	Id       int64  `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"id"` // id
	Name     string `gorm:"column:name;type:varchar(255)" json:"name"`                      // 仓库名
	Address  string `gorm:"column:address;type:varchar(255)" json:"address"`                // 仓库地址
	Areacode string `gorm:"column:areacode;type:varchar(20)" json:"areacode"`               // 区域编码
}

func (m *WmsWareInfo) TableName() string {
	return "wms_ware_info"
}

// 库存工作单
type WmsWareOrderTask struct {
	Id              int64     `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"id"`    // id
	OrderId         int64     `gorm:"column:order_id;type:bigint(20)" json:"order_id"`                   // order_id
	OrderSn         string    `gorm:"column:order_sn;type:varchar(255)" json:"order_sn"`                 // order_sn
	Consignee       string    `gorm:"column:consignee;type:varchar(100)" json:"consignee"`               // 收货人
	ConsigneeTel    string    `gorm:"column:consignee_tel;type:char(15)" json:"consignee_tel"`           // 收货人电话
	DeliveryAddress string    `gorm:"column:delivery_address;type:varchar(500)" json:"delivery_address"` // 配送地址
	OrderComment    string    `gorm:"column:order_comment;type:varchar(200)" json:"order_comment"`       // 订单备注
	PaymentWay      int       `gorm:"column:payment_way;type:tinyint(1)" json:"payment_way"`             // 付款方式【 1:在线付款 2:货到付款】
	TaskStatus      int       `gorm:"column:task_status;type:tinyint(2)" json:"task_status"`             // 任务状态
	OrderBody       string    `gorm:"column:order_body;type:varchar(255)" json:"order_body"`             // 订单描述
	TrackingNo      string    `gorm:"column:tracking_no;type:char(30)" json:"tracking_no"`               // 物流单号
	CreateTime      time.Time `gorm:"column:create_time;type:datetime" json:"create_time"`               // create_time
	WareId          int64     `gorm:"column:ware_id;type:bigint(20)" json:"ware_id"`                     // 仓库id
	TaskComment     string    `gorm:"column:task_comment;type:varchar(500)" json:"task_comment"`         // 工作单备注
}

func (m *WmsWareOrderTask) TableName() string {
	return "wms_ware_order_task"
}

// 库存工作单
type WmsWareOrderTaskDetail struct {
	Id         int64  `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"id"` // id
	SkuId      int64  `gorm:"column:sku_id;type:bigint(20)" json:"sku_id"`                    // sku_id
	SkuName    string `gorm:"column:sku_name;type:varchar(255)" json:"sku_name"`              // sku_name
	SkuNum     int    `gorm:"column:sku_num;type:int(11)" json:"sku_num"`                     // 购买个数
	TaskId     int64  `gorm:"column:task_id;type:bigint(20)" json:"task_id"`                  // 工作单id
	WareId     int64  `gorm:"column:ware_id;type:bigint(20)" json:"ware_id"`                  // 仓库id
	LockStatus int    `gorm:"column:lock_status;type:int(1)" json:"lock_status"`              // 1-已锁定  2-已解锁  3-扣减
}

func (m *WmsWareOrderTaskDetail) TableName() string {
	return "wms_ware_order_task_detail"
}

// 商品库存
type WmsWareSku struct {
	Id          int64  `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"id"` // id
	SkuId       int64  `gorm:"column:sku_id;type:bigint(20)" json:"sku_id"`                    // sku_id
	WareId      int64  `gorm:"column:ware_id;type:bigint(20)" json:"ware_id"`                  // 仓库id
	Stock       int    `gorm:"column:stock;type:int(11)" json:"stock"`                         // 库存数
	SkuName     string `gorm:"column:sku_name;type:varchar(200)" json:"sku_name"`              // sku_name
	StockLocked int    `gorm:"column:stock_locked;type:int(11);default:0" json:"stock_locked"` // 锁定库存
}

func (m *WmsWareSku) TableName() string {
	return "wms_ware_sku"
}
