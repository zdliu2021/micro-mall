package model

import (
	"time"
)

// 优惠券信息
type SmsCoupon struct {
	Id              int64     `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"id"`  // id
	CouponType      int       `gorm:"column:coupon_type;type:tinyint(1)" json:"coupon_type"`           // 优惠卷类型[0->全场赠券；1->会员赠券；2->购物赠券；3->注册赠券]
	CouponImg       string    `gorm:"column:coupon_img;type:varchar(2000)" json:"coupon_img"`          // 优惠券图片
	CouponName      string    `gorm:"column:coupon_name;type:varchar(100)" json:"coupon_name"`         // 优惠卷名字
	Num             int       `gorm:"column:num;type:int(11)" json:"num"`                              // 数量
	Amount          float64   `gorm:"column:amount;type:decimal(18,4)" json:"amount"`                  // 金额
	PerLimit        int       `gorm:"column:per_limit;type:int(11)" json:"per_limit"`                  // 每人限领张数
	MinPoint        float64   `gorm:"column:min_point;type:decimal(18,4)" json:"min_point"`            // 使用门槛
	StartTime       time.Time `gorm:"column:start_time;type:datetime" json:"start_time"`               // 开始时间
	EndTime         time.Time `gorm:"column:end_time;type:datetime" json:"end_time"`                   // 结束时间
	UseType         int       `gorm:"column:use_type;type:tinyint(1)" json:"use_type"`                 // 使用类型[0->全场通用；1->指定分类；2->指定商品]
	Note            string    `gorm:"column:note;type:varchar(200)" json:"note"`                       // 备注
	PublishCount    int       `gorm:"column:publish_count;type:int(11)" json:"publish_count"`          // 发行数量
	UseCount        int       `gorm:"column:use_count;type:int(11)" json:"use_count"`                  // 已使用数量
	ReceiveCount    int       `gorm:"column:receive_count;type:int(11)" json:"receive_count"`          // 领取数量
	EnableStartTime time.Time `gorm:"column:enable_start_time;type:datetime" json:"enable_start_time"` // 可以领取的开始日期
	EnableEndTime   time.Time `gorm:"column:enable_end_time;type:datetime" json:"enable_end_time"`     // 可以领取的结束日期
	Code            string    `gorm:"column:code;type:varchar(64)" json:"code"`                        // 优惠码
	MemberLevel     int       `gorm:"column:member_level;type:tinyint(1)" json:"member_level"`         // 可以领取的会员等级[0->不限等级，其他-对应等级]
	Publish         int       `gorm:"column:publish;type:tinyint(1)" json:"publish"`                   // 发布状态[0-未发布，1-已发布]
}

func (m *SmsCoupon) TableName() string {
	return "sms_coupon"
}

// 优惠券领取历史记录
type SmsCouponHistory struct {
	Id             int64     `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"id"`   // id
	CouponId       int64     `gorm:"column:coupon_id;type:bigint(20)" json:"coupon_id"`                // 优惠券id
	MemberId       int64     `gorm:"column:member_id;type:bigint(20)" json:"member_id"`                // 会员id
	MemberNickName string    `gorm:"column:member_nick_name;type:varchar(64)" json:"member_nick_name"` // 会员名字
	GetType        int       `gorm:"column:get_type;type:tinyint(1)" json:"get_type"`                  // 获取方式[0->后台赠送；1->主动领取]
	CreateTime     time.Time `gorm:"column:create_time;type:datetime" json:"create_time"`              // 创建时间
	UseType        int       `gorm:"column:use_type;type:tinyint(1)" json:"use_type"`                  // 使用状态[0->未使用；1->已使用；2->已过期]
	UseTime        time.Time `gorm:"column:use_time;type:datetime" json:"use_time"`                    // 使用时间
	OrderId        int64     `gorm:"column:order_id;type:bigint(20)" json:"order_id"`                  // 订单id
	OrderSn        int64     `gorm:"column:order_sn;type:bigint(20)" json:"order_sn"`                  // 订单号
}

func (m *SmsCouponHistory) TableName() string {
	return "sms_coupon_history"
}

// 优惠券分类关联
type SmsCouponSpuCategoryRelation struct {
	Id           int64  `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"id"` // id
	CouponId     int64  `gorm:"column:coupon_id;type:bigint(20)" json:"coupon_id"`              // 优惠券id
	CategoryId   int64  `gorm:"column:category_id;type:bigint(20)" json:"category_id"`          // 产品分类id
	CategoryName string `gorm:"column:category_name;type:varchar(64)" json:"category_name"`     // 产品分类名称
}

func (m *SmsCouponSpuCategoryRelation) TableName() string {
	return "sms_coupon_spu_category_relation"
}

// 优惠券与产品关联
type SmsCouponSpuRelation struct {
	Id       int64  `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"id"` // id
	CouponId int64  `gorm:"column:coupon_id;type:bigint(20)" json:"coupon_id"`              // 优惠券id
	SpuId    int64  `gorm:"column:spu_id;type:bigint(20)" json:"spu_id"`                    // spu_id
	SpuName  string `gorm:"column:spu_name;type:varchar(255)" json:"spu_name"`              // spu_name
}

func (m *SmsCouponSpuRelation) TableName() string {
	return "sms_coupon_spu_relation"
}

// 首页轮播广告
type SmsHomeAdv struct {
	Id          int64     `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"id"` // id
	Name        string    `gorm:"column:name;type:varchar(100)" json:"name"`                      // 名字
	Pic         string    `gorm:"column:pic;type:varchar(500)" json:"pic"`                        // 图片地址
	StartTime   time.Time `gorm:"column:start_time;type:datetime" json:"start_time"`              // 开始时间
	EndTime     time.Time `gorm:"column:end_time;type:datetime" json:"end_time"`                  // 结束时间
	Status      int       `gorm:"column:status;type:tinyint(1)" json:"status"`                    // 状态
	ClickCount  int       `gorm:"column:click_count;type:int(11)" json:"click_count"`             // 点击数
	Url         string    `gorm:"column:url;type:varchar(500)" json:"url"`                        // 广告详情连接地址
	Note        string    `gorm:"column:note;type:varchar(500)" json:"note"`                      // 备注
	Sort        int       `gorm:"column:sort;type:int(11)" json:"sort"`                           // 排序
	PublisherId int64     `gorm:"column:publisher_id;type:bigint(20)" json:"publisher_id"`        // 发布者
	AuthId      int64     `gorm:"column:auth_id;type:bigint(20)" json:"auth_id"`                  // 审核者
}

func (m *SmsHomeAdv) TableName() string {
	return "sms_home_adv"
}

// 首页专题表【jd首页下面很多专题，每个专题链接新的页面，展示专题商品信息】
type SmsHomeSubject struct {
	Id       int64  `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"id"` // id
	Name     string `gorm:"column:name;type:varchar(200)" json:"name"`                      // 专题名字
	Title    string `gorm:"column:title;type:varchar(255)" json:"title"`                    // 专题标题
	SubTitle string `gorm:"column:sub_title;type:varchar(255)" json:"sub_title"`            // 专题副标题
	Status   int    `gorm:"column:status;type:tinyint(1)" json:"status"`                    // 显示状态
	Url      string `gorm:"column:url;type:varchar(500)" json:"url"`                        // 详情连接
	Sort     int    `gorm:"column:sort;type:int(11)" json:"sort"`                           // 排序
	Img      string `gorm:"column:img;type:varchar(500)" json:"img"`                        // 专题图片地址
}

func (m *SmsHomeSubject) TableName() string {
	return "sms_home_subject"
}

// 专题商品
type SmsHomeSubjectSpu struct {
	Id        int64  `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"id"` // id
	Name      string `gorm:"column:name;type:varchar(200)" json:"name"`                      // 专题名字
	SubjectId int64  `gorm:"column:subject_id;type:bigint(20)" json:"subject_id"`            // 专题id
	SpuId     int64  `gorm:"column:spu_id;type:bigint(20)" json:"spu_id"`                    // spu_id
	Sort      int    `gorm:"column:sort;type:int(11)" json:"sort"`                           // 排序
}

func (m *SmsHomeSubjectSpu) TableName() string {
	return "sms_home_subject_spu"
}

// 商品会员价格
type SmsMemberPrice struct {
	Id              int64   `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"id"`      // id
	SkuId           int64   `gorm:"column:sku_id;type:bigint(20)" json:"sku_id"`                         // sku_id
	MemberLevelId   int64   `gorm:"column:member_level_id;type:bigint(20)" json:"member_level_id"`       // 会员等级id
	MemberLevelName string  `gorm:"column:member_level_name;type:varchar(100)" json:"member_level_name"` // 会员等级名
	MemberPrice     float64 `gorm:"column:member_price;type:decimal(18,4)" json:"member_price"`          // 会员对应价格
	AddOther        int     `gorm:"column:add_other;type:tinyint(1)" json:"add_other"`                   // 可否叠加其他优惠[0-不可叠加优惠，1-可叠加]
}

func (m *SmsMemberPrice) TableName() string {
	return "sms_member_price"
}

// 秒杀活动
type SmsSeckillPromotion struct {
	Id         int64     `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"id"` // id
	Title      string    `gorm:"column:title;type:varchar(255)" json:"title"`                    // 活动标题
	StartTime  time.Time `gorm:"column:start_time;type:datetime" json:"start_time"`              // 开始日期
	EndTime    time.Time `gorm:"column:end_time;type:datetime" json:"end_time"`                  // 结束日期
	Status     int       `gorm:"column:status;type:tinyint(4)" json:"status"`                    // 上下线状态
	CreateTime time.Time `gorm:"column:create_time;type:datetime" json:"create_time"`            // 创建时间
	UserId     int64     `gorm:"column:user_id;type:bigint(20)" json:"user_id"`                  // 创建人
}

func (m *SmsSeckillPromotion) TableName() string {
	return "sms_seckill_promotion"
}

// 秒杀活动场次
type SmsSeckillSession struct {
	Id         int64     `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"id"` // id
	Name       string    `gorm:"column:name;type:varchar(200)" json:"name"`                      // 场次名称
	StartTime  time.Time `gorm:"column:start_time;type:datetime" json:"start_time"`              // 每日开始时间
	EndTime    time.Time `gorm:"column:end_time;type:datetime" json:"end_time"`                  // 每日结束时间
	Status     int       `gorm:"column:status;type:tinyint(1)" json:"status"`                    // 启用状态
	CreateTime time.Time `gorm:"column:create_time;type:datetime" json:"create_time"`            // 创建时间
}

func (m *SmsSeckillSession) TableName() string {
	return "sms_seckill_session"
}

// 秒杀商品通知订阅
type SmsSeckillSkuNotice struct {
	Id           int64     `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"id"` // id
	MemberId     int64     `gorm:"column:member_id;type:bigint(20)" json:"member_id"`              // member_id
	SkuId        int64     `gorm:"column:sku_id;type:bigint(20)" json:"sku_id"`                    // sku_id
	SessionId    int64     `gorm:"column:session_id;type:bigint(20)" json:"session_id"`            // 活动场次id
	SubcribeTime time.Time `gorm:"column:subcribe_time;type:datetime" json:"subcribe_time"`        // 订阅时间
	SendTime     time.Time `gorm:"column:send_time;type:datetime" json:"send_time"`                // 发送时间
	NoticeType   int       `gorm:"column:notice_type;type:tinyint(1)" json:"notice_type"`          // 通知方式[0-短信，1-邮件]
}

func (m *SmsSeckillSkuNotice) TableName() string {
	return "sms_seckill_sku_notice"
}

// 秒杀活动商品关联
type SmsSeckillSkuRelation struct {
	Id                 int64   `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"id"`          // id
	PromotionId        int64   `gorm:"column:promotion_id;type:bigint(20)" json:"promotion_id"`                 // 活动id
	PromotionSessionId int64   `gorm:"column:promotion_session_id;type:bigint(20)" json:"promotion_session_id"` // 活动场次id
	SkuId              int64   `gorm:"column:sku_id;type:bigint(20)" json:"sku_id"`                             // 商品id
	SeckillPrice       float64 `gorm:"column:seckill_price;type:decimal(10,4)" json:"seckill_price"`            // 秒杀价格
	SeckillCount       int     `gorm:"column:seckill_count;type:int(11)" json:"seckill_count"`                  // 秒杀总量
	SeckillLimit       int     `gorm:"column:seckill_limit;type:int(11)" json:"seckill_limit"`                  // 每人限购数量
	SeckillSort        int     `gorm:"column:seckill_sort;type:int(11)" json:"seckill_sort"`                    // 排序
}

func (m *SmsSeckillSkuRelation) TableName() string {
	return "sms_seckill_sku_relation"
}

// 商品满减信息
type SmsSkuFullReduction struct {
	Id          int64   `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"id"` // id
	SkuId       int64   `gorm:"column:sku_id;type:bigint(20)" json:"sku_id"`                    // spu_id
	FullPrice   float64 `gorm:"column:full_price;type:decimal(18,4)" json:"full_price"`         // 满多少
	ReducePrice float64 `gorm:"column:reduce_price;type:decimal(18,4)" json:"reduce_price"`     // 减多少
	AddOther    int     `gorm:"column:add_other;type:tinyint(1)" json:"add_other"`              // 是否参与其他优惠
}

func (m *SmsSkuFullReduction) TableName() string {
	return "sms_sku_full_reduction"
}

// 商品阶梯价格
type SmsSkuLadder struct {
	Id        int64   `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"id"` // id
	SkuId     int64   `gorm:"column:sku_id;type:bigint(20)" json:"sku_id"`                    // spu_id
	FullCount int     `gorm:"column:full_count;type:int(11)" json:"full_count"`               // 满几件
	Discount  float64 `gorm:"column:discount;type:decimal(4,2)" json:"discount"`              // 打几折
	Price     float64 `gorm:"column:price;type:decimal(18,4)" json:"price"`                   // 折后价
	AddOther  int     `gorm:"column:add_other;type:tinyint(1)" json:"add_other"`              // 是否叠加其他优惠[0-不可叠加，1-可叠加]
}

func (m *SmsSkuLadder) TableName() string {
	return "sms_sku_ladder"
}

// 商品spu积分设置
type SmsSpuBounds struct {
	Id         int64   `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"id"` // id
	SpuId      int64   `gorm:"column:spu_id;type:bigint(20)" json:"spu_id"`
	GrowBounds float64 `gorm:"column:grow_bounds;type:decimal(18,4)" json:"grow_bounds"` // 成长积分
	BuyBounds  float64 `gorm:"column:buy_bounds;type:decimal(18,4)" json:"buy_bounds"`   // 购物积分
	Work       int     `gorm:"column:work;type:tinyint(1)" json:"work"`                  // 优惠生效情况[1111（四个状态位，从右到左）;0 - 无优惠，成长积分是否赠送;1 - 无优惠，购物积分是否赠送;2 - 有优惠，成长积分是否赠送;3 - 有优惠，购物积分是否赠送【状态位0：不赠送，1：赠送】]
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
