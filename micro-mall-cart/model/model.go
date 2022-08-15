package cart

import "strconv"

// 购物车结构
// Map<k1,Map<k2,obj>>
// k1 是用户ID
// k2 是sku_id
//obj 是json序列化之后的cartInfo

type CartInfo struct {
	ID           int64  `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	UserID       string `gorm:"column:user_id;NOT NULL"` // 用户id或者userKey
	SkuID        int64  `gorm:"column:sku_id;NOT NULL"`  // skuId
	Check        int    `gorm:"column:check;NOT NULL"`   // 选中状态
	Title        string `gorm:"column:title;NOT NULL"`   // 标题
	DefaultImage string `gorm:"column:default_image"`    // 默认图片
	Price        string `gorm:"column:price;NOT NULL"`   // 加入购物车时价格
	Count        int    `gorm:"column:count;NOT NULL"`   // 数量
	Store        int    `gorm:"column:store;NOT NULL"`   // 是否有货
	SaleAttrs    string `gorm:"column:sale_attrs"`       // 销售属性（json格式）
	Sales        string `gorm:"column:sales"`            // 营销信息（json格式）
}

func (m *CartInfo) TableName() string {
	return "cart_info"
}

type Cart struct {
	Items       []CartInfo
	CountNum    int     //商品数量
	CountType   int     //商品类型数量
	totalAmount float32 //商品总价
	Reduce      float32 //减免价格
}

func (cart *Cart) GetCountType() int {
	if cart.Items == nil {
		return 0
	} else {
		return len(cart.Items)
	}
}

func (cart *Cart) GetCountNum() int {
	if cart.Items == nil {
		return 0
	} else {
		var ans int = 0
		for i := 0; i < len(cart.Items); i++ {
			ans += cart.Items[i].Count
		}
		return ans
	}
}

func (cart *Cart) GetTotalAmount() float32 {
	if cart.Items == nil {
		return 0
	} else {
		var ans float32 = 0
		for i := 0; i < len(cart.Items); i++ {
			price, _ := strconv.ParseFloat(cart.Items[i].Price, 32)
			ans += float32(cart.Items[i].Count) * float32(price)
		}
		return ans
	}
}

func (cart *Cart) GetReduce() float32 {
	return cart.Reduce
}

type SkuInfo struct {
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

type UserInfo struct {
	UserId   int64
	UserKey  string
	tempUser bool
}
