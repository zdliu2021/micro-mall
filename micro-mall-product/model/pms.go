package model

import (
	"gorm.io/gorm"
	"time"
)

// 商品属性
type PmsAttr struct {
	AttrId      int64  `gorm:"column:attr_id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"attr_id"` // 属性id
	AttrName    string `gorm:"column:attr_name;type:char(30)" json:"attr_name"`                          // 属性名
	SearchType  int    `gorm:"column:search_type;type:tinyint(4)" json:"search_type"`                    // 是否需要检索[0-不需要，1-需要]
	ValueType   int    `gorm:"column:value_type;type:tinyint(4)" json:"value_type"`                      // 值类型[0-为单个值，1-可以选择多个值]
	Icon        string `gorm:"column:icon;type:varchar(255)" json:"icon"`                                // 属性图标
	ValueSelect string `gorm:"column:value_select;type:char(255)" json:"value_select"`                   // 可选值列表[用逗号分隔]
	AttrType    int    `gorm:"column:attr_type;type:tinyint(4)" json:"attr_type"`                        // 属性类型[0-销售属性，1-基本属性，2-既是销售属性又是基本属性]
	Enable      int64  `gorm:"column:enable;type:bigint(20)" json:"enable"`                              // 启用状态[0 - 禁用，1 - 启用]
	CatelogId   int64  `gorm:"column:catelog_id;type:bigint(20)" json:"catelog_id"`                      // 所属分类
	ShowDesc    int    `gorm:"column:show_desc;type:tinyint(4)" json:"show_desc"`                        // 快速展示【是否展示在介绍上；0-否 1-是】，在sku中仍然可以调整
}

func (m *PmsAttr) TableName() string {
	return "pms_attr"
}

// 属性&属性分组关联
type PmsAttrAttrgroupRelation struct {
	Id          int64 `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"id"` // id
	AttrId      int64 `gorm:"column:attr_id;type:bigint(20)" json:"attr_id"`                  // 属性id
	AttrGroupId int64 `gorm:"column:attr_group_id;type:bigint(20)" json:"attr_group_id"`      // 属性分组id
	AttrSort    int   `gorm:"column:attr_sort;type:int(11)" json:"attr_sort"`                 // 属性组内排序
}

func (m *PmsAttrAttrgroupRelation) TableName() string {
	return "pms_attr_attrgroup_relation"
}

// 属性分组
type PmsAttrGroup struct {
	AttrGroupId   int64  `gorm:"column:attr_group_id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"attr_group_id"` // 分组id
	AttrGroupName string `gorm:"column:attr_group_name;type:char(20)" json:"attr_group_name"`                          // 组名
	Sort          int    `gorm:"column:sort;type:int(11)" json:"sort"`                                                 // 排序
	Descript      string `gorm:"column:descript;type:varchar(255)" json:"descript"`                                    // 描述
	Icon          string `gorm:"column:icon;type:varchar(255)" json:"icon"`                                            // 组图标
	CatelogId     int64  `gorm:"column:catelog_id;type:bigint(20)" json:"catelog_id"`                                  // 所属分类id

	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (m *PmsAttrGroup) TableName() string {
	return "pms_attr_group"
}

// 品牌
type PmsBrand struct {
	BrandId     int64  `gorm:"column:brand_id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"brand_id"` // 品牌id
	Name        string `gorm:"column:name;type:char(50)" json:"name"`                                      // 品牌名
	Logo        string `gorm:"column:logo;type:varchar(2000)" json:"logo"`                                 // 品牌logo地址
	Descript    string `gorm:"column:descript;type:longtext" json:"descript"`                              // 介绍
	ShowStatus  int    `gorm:"column:show_status;type:tinyint(4)" json:"show_status"`                      // 显示状态[0-不显示；1-显示]
	FirstLetter string `gorm:"column:first_letter;type:char(1)" json:"first_letter"`                       // 检索首字母
	Sort        int    `gorm:"column:sort;type:int(11)" json:"sort"`                                       // 排序

	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (m *PmsBrand) TableName() string {
	return "pms_brand"
}

// 商品三级分类
type PmsCategory struct {
	CatId        int64  `gorm:"column:cat_id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"cat_id"` // 分类id
	Name         string `gorm:"column:name;type:char(50)" json:"name"`                                  // 分类名称
	ParentCid    int64  `gorm:"column:parent_cid;type:bigint(20)" json:"parent_cid"`                    // 父分类id
	CatLevel     int    `gorm:"column:cat_level;type:int(11)" json:"cat_level"`                         // 层级
	ShowStatus   int    `gorm:"column:show_status;type:tinyint(4)" json:"show_status"`                  // 是否显示[0-不显示，1显示]
	Sort         int    `gorm:"column:sort;type:int(11)" json:"sort"`                                   // 排序
	Icon         string `gorm:"column:icon;type:char(255)" json:"icon"`                                 // 图标地址
	ProductUnit  string `gorm:"column:product_unit;type:char(50)" json:"product_unit"`                  // 计量单位
	ProductCount int    `gorm:"column:product_count;type:int(11)" json:"product_count"`                 // 商品数量

	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (m *PmsCategory) TableName() string {
	return "pms_category"
}

// 品牌分类关联
type PmsCategoryBrandRelation struct {
	Id          int64  `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"id"`
	BrandId     int64  `gorm:"column:brand_id;type:bigint(20)" json:"brand_id"`     // 品牌id
	CatelogId   int64  `gorm:"column:catelog_id;type:bigint(20)" json:"catelog_id"` // 分类id
	BrandName   string `gorm:"column:brand_name;type:varchar(255)" json:"brand_name"`
	CatelogName string `gorm:"column:catelog_name;type:varchar(255)" json:"catelog_name"`
}

func (m *PmsCategoryBrandRelation) TableName() string {
	return "pms_category_brand_relation"
}

// 商品评价回复关系
type PmsCommentReplay struct {
	Id        int64 `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"id"` // id
	CommentId int64 `gorm:"column:comment_id;type:bigint(20)" json:"comment_id"`            // 评论id
	ReplyId   int64 `gorm:"column:reply_id;type:bigint(20)" json:"reply_id"`                // 回复id
}

func (m *PmsCommentReplay) TableName() string {
	return "pms_comment_replay"
}

// spu属性值
type PmsProductAttrValue struct {
	Id        int64  `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"id"` // id
	SpuId     int64  `gorm:"column:spu_id;type:bigint(20)" json:"spu_id"`                    // 商品id
	AttrId    int64  `gorm:"column:attr_id;type:bigint(20)" json:"attr_id"`                  // 属性id
	AttrName  string `gorm:"column:attr_name;type:varchar(200)" json:"attr_name"`            // 属性名
	AttrValue string `gorm:"column:attr_value;type:varchar(200)" json:"attr_value"`          // 属性值
	AttrSort  int    `gorm:"column:attr_sort;type:int(11)" json:"attr_sort"`                 // 顺序
	QuickShow int    `gorm:"column:quick_show;type:tinyint(4)" json:"quick_show"`            // 快速展示【是否展示在介绍上；0-否 1-是】
}

func (m *PmsProductAttrValue) TableName() string {
	return "pms_product_attr_value"
}

// sku图片
type PmsSkuImages struct {
	Id         int64  `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"id"` // id
	SkuId      int64  `gorm:"column:sku_id;type:bigint(20)" json:"sku_id"`                    // sku_id
	ImgUrl     string `gorm:"column:img_url;type:varchar(255)" json:"img_url"`                // 图片地址
	ImgSort    int    `gorm:"column:img_sort;type:int(11)" json:"img_sort"`                   // 排序
	DefaultImg int    `gorm:"column:default_img;type:int(11)" json:"default_img"`             // 默认图[0 - 不是默认图，1 - 是默认图]
}

// sku信息
type PmsSkuInfo struct {
	SkuId         int64   `gorm:"column:sku_id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"sku_id"` // skuId
	SpuId         int64   `gorm:"column:spu_id;type:bigint(20)" json:"spu_id"`                            // spuId
	SkuName       string  `gorm:"column:sku_name;type:varchar(255)" json:"sku_name"`                      // sku名称
	SkuDesc       string  `gorm:"column:sku_desc;type:varchar(2000)" json:"sku_desc"`                     // sku介绍描述
	CatalogId     int64   `gorm:"column:catalog_id;type:bigint(20)" json:"catalog_id"`                    // 所属分类id
	BrandId       int64   `gorm:"column:brand_id;type:bigint(20)" json:"brand_id"`                        // 品牌id
	SkuDefaultImg string  `gorm:"column:sku_default_img;type:varchar(255)" json:"sku_default_img"`        // 默认图片
	SkuTitle      string  `gorm:"column:sku_title;type:varchar(255)" json:"sku_title"`                    // 标题
	SkuSubtitle   string  `gorm:"column:sku_subtitle;type:varchar(2000)" json:"sku_subtitle"`             // 副标题
	Price         float64 `gorm:"column:price;type:decimal(18,4)" json:"price"`                           // 价格
	SaleCount     int64   `gorm:"column:sale_count;type:bigint(20)" json:"sale_count"`                    // 销量
}

func (m *PmsSkuInfo) TableName() string {
	return "pms_sku_info"
}

// sku销售属性&值
type PmsSkuSaleAttrValue struct {
	Id        int64  `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"id"` // id
	SkuId     int64  `gorm:"column:sku_id;type:bigint(20)" json:"sku_id"`                    // sku_id
	AttrId    int64  `gorm:"column:attr_id;type:bigint(20)" json:"attr_id"`                  // attr_id
	AttrName  string `gorm:"column:attr_name;type:varchar(200)" json:"attr_name"`            // 销售属性名
	AttrValue string `gorm:"column:attr_value;type:varchar(200)" json:"attr_value"`          // 销售属性值
	AttrSort  int    `gorm:"column:attr_sort;type:int(11)" json:"attr_sort"`                 // 顺序
}

func (m *PmsSkuSaleAttrValue) TableName() string {
	return "pms_sku_sale_attr_value"
}

// 商品评价
type PmsSpuComment struct {
	Id             int64     `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"id"`    // id
	SkuId          int64     `gorm:"column:sku_id;type:bigint(20)" json:"sku_id"`                       // sku_id
	SpuId          int64     `gorm:"column:spu_id;type:bigint(20)" json:"spu_id"`                       // spu_id
	SpuName        string    `gorm:"column:spu_name;type:varchar(255)" json:"spu_name"`                 // 商品名字
	MemberNickName string    `gorm:"column:member_nick_name;type:varchar(255)" json:"member_nick_name"` // 会员昵称
	Star           int       `gorm:"column:star;type:tinyint(1)" json:"star"`                           // 星级
	MemberIp       string    `gorm:"column:member_ip;type:varchar(64)" json:"member_ip"`                // 会员ip
	CreateTime     time.Time `gorm:"column:create_time;type:datetime" json:"create_time"`               // 创建时间
	ShowStatus     int       `gorm:"column:show_status;type:tinyint(1)" json:"show_status"`             // 显示状态[0-不显示，1-显示]
	SpuAttributes  string    `gorm:"column:spu_attributes;type:varchar(255)" json:"spu_attributes"`     // 购买时属性组合
	LikesCount     int       `gorm:"column:likes_count;type:int(11)" json:"likes_count"`                // 点赞数
	ReplyCount     int       `gorm:"column:reply_count;type:int(11)" json:"reply_count"`                // 回复数
	Resources      string    `gorm:"column:resources;type:varchar(1000)" json:"resources"`              // 评论图片/视频[json数据；[{type:文件类型,url:资源路径}]]
	Content        string    `gorm:"column:content;type:text" json:"content"`                           // 内容
	MemberIcon     string    `gorm:"column:member_icon;type:varchar(255)" json:"member_icon"`           // 用户头像
	CommentType    int       `gorm:"column:comment_type;type:tinyint(4)" json:"comment_type"`           // 评论类型[0 - 对商品的直接评论，1 - 对评论的回复]
}

func (m *PmsSpuComment) TableName() string {
	return "pms_spu_comment"
}

// spu图片
type PmsSpuImages struct {
	Id         int64  `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"id"` // id
	SpuId      int64  `gorm:"column:spu_id;type:bigint(20)" json:"spu_id"`                    // spu_id
	ImgName    string `gorm:"column:img_name;type:varchar(200)" json:"img_name"`              // 图片名
	ImgUrl     string `gorm:"column:img_url;type:varchar(255)" json:"img_url"`                // 图片地址
	ImgSort    int    `gorm:"column:img_sort;type:int(11)" json:"img_sort"`                   // 顺序
	DefaultImg int    `gorm:"column:default_img;type:tinyint(4)" json:"default_img"`          // 是否默认图
}

// spu信息
type PmsSpuInfo struct {
	Id             int64     `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"id"`   // 商品id
	SpuName        string    `gorm:"column:spu_name;type:varchar(200)" json:"spu_name"`                // 商品名称
	SpuDescription string    `gorm:"column:spu_description;type:varchar(1000)" json:"spu_description"` // 商品描述
	CatalogId      int64     `gorm:"column:catalog_id;type:bigint(20)" json:"catalog_id"`              // 所属分类id
	BrandId        int64     `gorm:"column:brand_id;type:bigint(20)" json:"brand_id"`                  // 品牌id
	Weight         float64   `gorm:"column:weight;type:decimal(18,4)" json:"weight"`
	PublishStatus  int       `gorm:"column:publish_status;type:tinyint(4)" json:"publish_status"` // 上架状态[0--新建，1 - 上架，2 — 下架]
	CreateTime     time.Time `gorm:"column:create_time;type:datetime" json:"create_time"`
	UpdateTime     time.Time `gorm:"column:update_time;type:datetime" json:"update_time"`
}

func (m *PmsSpuInfo) TableName() string {
	return "pms_spu_info"
}

// spu信息介绍
type PmsSpuInfoDesc struct {
	SpuId   int64  `gorm:"column:spu_id;type:bigint(20);primary_key" json:"spu_id"` // 商品id
	Decript string `gorm:"column:decript;type:longtext" json:"decript"`             // 商品介绍
}

func (m *PmsSpuInfoDesc) TableName() string {
	return "pms_spu_info_desc"
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
