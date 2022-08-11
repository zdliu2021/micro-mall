package model

import (
	"time"
)

// 成长值变化历史记录

type UmsGrowthChangeHistory struct {
	Id          int64     `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"id"` // id
	MemberId    int64     `gorm:"column:member_id;type:bigint(20)" json:"member_id"`              // member_id
	CreateTime  time.Time `gorm:"column:create_time;type:datetime" json:"create_time"`            // create_time
	ChangeCount int       `gorm:"column:change_count;type:int(11)" json:"change_count"`           // 改变的值（正负计数）
	Note        string    `gorm:"column:note;type:varchar(5)" json:"note"`                        // 备注
	SourceType  int       `gorm:"column:source_type;type:tinyint(4)" json:"source_type"`          // 积分来源[0-购物，1-管理员修改]
}

func (m *UmsGrowthChangeHistory) TableName() string {
	return "ums_growth_change_history"
}

// 积分变化历史记录

type UmsIntegrationChangeHistory struct {
	Id          int64     `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"id"` // id
	MemberId    int64     `gorm:"column:member_id;type:bigint(20)" json:"member_id"`              // member_id
	CreateTime  time.Time `gorm:"column:create_time;type:datetime" json:"create_time"`            // create_time
	ChangeCount int       `gorm:"column:change_count;type:int(11)" json:"change_count"`           // 改变的值（正负计数）
	Note        string    `gorm:"column:note;type:varchar(255)" json:"note"`                      // 备注
	SourceTyoe  int       `gorm:"column:source_tyoe;type:tinyint(4)" json:"source_tyoe"`          // 积分来源[0-购物，1-管理员修改]
}

func (m *UmsIntegrationChangeHistory) TableName() string {
	return "ums_integration_change_history"
}

// 会员

type UmsMember struct {
	Id          int64     `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"id"` // id
	LevelId     int64     `gorm:"column:level_id;type:bigint(20)" json:"level_id"`                // 会员等级ID
	Username    string    `gorm:"column:username;type:char(64)" json:"username"`                  // 用户名
	Password    string    `gorm:"column:password;type:varchar(64)" json:"password"`               // 密码
	Nickname    string    `gorm:"column:nickname;type:varchar(64)" json:"nickname"`               // 昵称
	Mobile      string    `gorm:"column:mobile;type:varchar(20)" json:"mobile"`                   // 手机号码
	Email       string    `gorm:"column:email;type:varchar(64)" json:"email"`                     // 邮箱
	Header      string    `gorm:"column:header;type:varchar(500)" json:"header"`                  // 头像
	Gender      int       `gorm:"column:gender;type:tinyint(4)" json:"gender"`                    // 性别
	Birth       time.Time `gorm:"column:birth;type:date" json:"birth"`                            // 生日
	City        string    `gorm:"column:city;type:varchar(500)" json:"city"`                      // 所在城市
	Job         string    `gorm:"column:job;type:varchar(255)" json:"job"`                        // 职业
	Sign        string    `gorm:"column:sign;type:varchar(255)" json:"sign"`                      // 个性签名
	SourceType  int       `gorm:"column:source_type;type:tinyint(4)" json:"source_type"`          // 用户来源
	Integration int       `gorm:"column:integration;type:int(11)" json:"integration"`             // 积分
	Growth      int       `gorm:"column:growth;type:int(11)" json:"growth"`                       // 成长值
	Status      int       `gorm:"column:status;type:tinyint(4)" json:"status"`                    // 启用状态
	CreateTime  time.Time `gorm:"column:create_time;type:datetime" json:"create_time"`            // 注册时间
	SocialID    int64     `gorm:"column:social_id;type:bigint(20);AUTO_INCREMENT" json:"social_id"`
}

func (m *UmsMember) TableName() string {
	return "ums_member"
}

// 会员收藏的商品

type UmsMemberCollectSpu struct {
	Id         int64     `gorm:"column:id;type:bigint(20);primary_key" json:"id"`     // id
	MemberId   int64     `gorm:"column:member_id;type:bigint(20)" json:"member_id"`   // 会员ID
	SpuId      int64     `gorm:"column:spu_id;type:bigint(20)" json:"spu_id"`         // spu_id
	SpuName    string    `gorm:"column:spu_name;type:varchar(500)" json:"spu_name"`   // spu_name
	SpuImg     string    `gorm:"column:spu_img;type:varchar(500)" json:"spu_img"`     // spu_img
	CreateTime time.Time `gorm:"column:create_time;type:datetime" json:"create_time"` // create_time
}

func (m *UmsMemberCollectSpu) TableName() string {
	return "ums_member_collect_spu"
}

//会员收藏的专题活动

type UmsMemberCollectSubject struct {
	Id          int64  `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"id"` // id
	SubjectId   int64  `gorm:"column:subject_id;type:bigint(20)" json:"subject_id"`            // subject_id
	SubjectName string `gorm:"column:subject_name;type:varchar(255)" json:"subject_name"`      // subject_name
	SubjectImg  string `gorm:"column:subject_img;type:varchar(500)" json:"subject_img"`        // subject_img
	SubjectUrll string `gorm:"column:subject_urll;type:varchar(500)" json:"subject_urll"`      // 活动url
}

func (m *UmsMemberCollectSubject) TableName() string {
	return "ums_member_collect_subject"
}

// 会员等级

type UmsMemberLevel struct {
	Id                    int64   `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"id"`                // id
	Name                  string  `gorm:"column:name;type:varchar(100)" json:"name"`                                     // 等级名称
	GrowthPoint           int     `gorm:"column:growth_point;type:int(11)" json:"growth_point"`                          // 等级需要的成长值
	DefaultStatus         int     `gorm:"column:default_status;type:tinyint(4)" json:"default_status"`                   // 是否为默认等级【0：不是，1：是】
	FreeFreightPoint      float64 `gorm:"column:free_freight_point;type:decimal(18,4)" json:"free_freight_point"`        // 免运费标准
	CommentGrowthPoint    int     `gorm:"column:comment_growth_point;type:int(11)" json:"comment_growth_point"`          // 每次评价获取的成长值
	PriviledgeFreeFreight int     `gorm:"column:priviledge_free_freight;type:tinyint(4)" json:"priviledge_free_freight"` // 是否有免邮特权
	PriviledgeMemberPrice int     `gorm:"column:priviledge_member_price;type:tinyint(4)" json:"priviledge_member_price"` // 是否有会员价格特权
	PriviledgeBirthday    int     `gorm:"column:priviledge_birthday;type:tinyint(4)" json:"priviledge_birthday"`         // 是否有生日特权
	Note                  string  `gorm:"column:note;type:varchar(255)" json:"note"`                                     // 备注
}

func (m *UmsMemberLevel) TableName() string {
	return "ums_member_level"
}

// 会员登录记录

type UmsMemberLoginLog struct {
	Id         int64     `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"id"` // id
	MemberId   int64     `gorm:"column:member_id;type:bigint(20)" json:"member_id"`              // member_id
	CreateTime time.Time `gorm:"column:create_time;type:datetime" json:"create_time"`            // 创建时间
	Ip         string    `gorm:"column:ip;type:varchar(64)" json:"ip"`                           // ip
	City       string    `gorm:"column:city;type:varchar(64)" json:"city"`                       // city
	LoginType  int       `gorm:"column:login_type;type:tinyint(1)" json:"login_type"`            // 登录类型【1-web,2-app】
}

func (m *UmsMemberLoginLog) TableName() string {
	return "ums_member_login_log"
}

//会员收货地址

type UmsMemberReceiveAddress struct {
	Id            int64  `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"id"` // id
	MemberId      int64  `gorm:"column:member_id;type:bigint(20)" json:"member_id"`              // member_id
	Name          string `gorm:"column:name;type:varchar(255)" json:"name"`                      // 收货人姓名
	Phone         string `gorm:"column:phone;type:varchar(64)" json:"phone"`                     // 电话
	PostCode      string `gorm:"column:post_code;type:varchar(64)" json:"post_code"`             // 邮政编码
	Province      string `gorm:"column:province;type:varchar(100)" json:"province"`              // 省份/直辖市
	City          string `gorm:"column:city;type:varchar(100)" json:"city"`                      // 城市
	Region        string `gorm:"column:region;type:varchar(100)" json:"region"`                  // 区
	DetailAddress string `gorm:"column:detail_address;type:varchar(255)" json:"detail_address"`  // 详细地址（街道）
	Areacode      string `gorm:"column:areacode;type:varchar(15)" json:"areacode"`               // 省市区代码
	DefaultStatus int    `gorm:"column:default_status;type:tinyint(1)" json:"default_status"`    // 是否默认
}

func (m *UmsMemberReceiveAddress) TableName() string {
	return "ums_member_receive_address"
}

// 会员统计信息

type UmsMemberStatisticsInfo struct {
	Id                  int64   `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"id"`         // id
	MemberId            int64   `gorm:"column:member_id;type:bigint(20)" json:"member_id"`                      // 会员id
	ConsumeAmount       float64 `gorm:"column:consume_amount;type:decimal(18,4)" json:"consume_amount"`         // 累计消费金额
	CouponAmount        float64 `gorm:"column:coupon_amount;type:decimal(18,4)" json:"coupon_amount"`           // 累计优惠金额
	OrderCount          int     `gorm:"column:order_count;type:int(11)" json:"order_count"`                     // 订单数量
	CouponCount         int     `gorm:"column:coupon_count;type:int(11)" json:"coupon_count"`                   // 优惠券数量
	CommentCount        int     `gorm:"column:comment_count;type:int(11)" json:"comment_count"`                 // 评价数
	ReturnOrderCount    int     `gorm:"column:return_order_count;type:int(11)" json:"return_order_count"`       // 退货数量
	LoginCount          int     `gorm:"column:login_count;type:int(11)" json:"login_count"`                     // 登录次数
	AttendCount         int     `gorm:"column:attend_count;type:int(11)" json:"attend_count"`                   // 关注数量
	FansCount           int     `gorm:"column:fans_count;type:int(11)" json:"fans_count"`                       // 粉丝数量
	CollectProductCount int     `gorm:"column:collect_product_count;type:int(11)" json:"collect_product_count"` // 收藏的商品
	CollectSubjectCount int     `gorm:"column:collect_subject_count;type:int(11)" json:"collect_subject_count"` // 收藏的专题活动数量
	CollectCommentCount int     `gorm:"column:collect_comment_count;type:int(11)" json:"collect_comment_count"` // 收藏的评论数量
	InviteFriendCount   int     `gorm:"column:invite_friend_count;type:int(11)" json:"invite_friend_count"`     // 邀请的朋友数量
}

func (m *UmsMemberStatisticsInfo) TableName() string {
	return "ums_member_statistics_info"
}
