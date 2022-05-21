package request

type SaveCategoryRequest struct {
	Name        string `json:"name"`                   // 分类名称
	ParentCid   int64  `json:"parentCid"`              // 父分类id
	CatLevel    int    `json:"catLevel"`               // 层级
	ShowStatus  int    `json:"showStatus"`             // 是否显示[0-不显示，1显示]
	Sort        int    `json:"sort"`                   // 排序
	Icon        string ` json:"icon"`                  // 图标地址
	ProductUnit string ` json:"productUnit,omitempty"` // 计量单位
}

type UpdateCategoryRequest struct {
	CatID       int64  `json:"catId"`
	Name        string `json:"name"`                   // 分类名称
	ParentCid   int64  `json:"parentCid"`              // 父分类id
	CatLevel    int32  `json:"catLevel"`               // 层级
	ShowStatus  int32  `json:"showStatus"`             // 是否显示[0-不显示，1显示]
	Sort        int32  `json:"sort"`                   // 排序
	Icon        string ` json:"icon"`                  // 图标地址
	ProductUnit string ` json:"productUnit,omitempty"` // 计量单位
}
type UpdateBrandRequest struct {
	BrandId     int64  ` json:"brandId,omitempty"`    // 品牌id
	Name        string ` json:"name,omitempty"`       // 品牌名
	Logo        string ` json:"logo,omitempty"`       // 品牌logo地址
	Descript    string ` json:"descript,omitempty"`   // 介绍
	ShowStatus  int32  ` json:"showStatus,omitempty"` // 显示状态[0-不显示；1-显示]
	FirstLetter string `json:"firstLetter,omitempty"` // 检索首字母
	Sort        int32  ` json:"sort,omitempty"`       // 排序
}
type SaveBrandRequest struct {
	BrandId     int64  ` json:"brandId,omitempty"`    // 品牌id
	Name        string ` json:"name,omitempty"`       // 品牌名
	Logo        string ` json:"logo,omitempty"`       // 品牌logo地址
	Descript    string ` json:"descript,omitempty"`   // 介绍
	ShowStatus  int32  ` json:"showStatus,omitempty"` // 显示状态[0-不显示；1-显示]
	FirstLetter string `json:"firstLetter,omitempty"` // 检索首字母
	Sort        int32  ` json:"sort,omitempty"`       // 排序
}

type DeleteBrandRequest struct {
	BrandIds []int64 `json:"brandIds"`
}

type SaveBrandCatelogRelationRequest struct {
	BrandId   int64 `json:"brandId"`
	CateLogId int64 `json:"catelogId"`
}

type DeleteAttrAttrGroupRelationRequest struct {
	AttrId      int64 `json:"attrId"`
	AttrGroupID int64 `json:"attrGroupId"`
}

type SaveAttrAttrGroupRelationRequest struct {
	AttrId      int64 `json:"attrId"`
	AttrGroupID int64 `json:"attrGroupId"`
}

type AttrEntity struct {
	AttrId      int64  ` json:"attrId,omitempty"`     //属性id
	AttrName    string ` json:"attrName,omitempty"`   //属性名
	AttrType    int32  `json:"attrType,omitempty"`    //属性类型，0-销售属性，1-基本属性
	CatelogId   int64  ` json:"catelogId,omitempty"`  //所属分类名字
	Enable      int32  ` json:"enable,omitempty"`     //是否启用
	Icon        string ` json:"icon,omitempty"`       //图标
	SearchType  int32  ` json:"searchType,omitempty"` //是否需要检索[0-不需要，1-需要]
	ShowDesc    int32  ` json:"showDesc,omitempty"`   //是否展示在介绍上；0-否 1-是
	ValueSelect string `json:"valueSelect,omitempty"` //可选值列表[用逗号分隔]
	ValueType   int32  ` json:"valueType,omitempty"`  //值类型[0-为单个值，1-可以选择多个值]
}

type AttrGroupEntity struct {
	AttrGroupId   int64  `json:"attrGroupId,omitempty"`
	AttrGroupName string `json:"attrGroupName,omitempty"`
	CatelogId     int64  `json:"catelogId,omitempty"`
	Descript      string `json:"descript"`
	Icon          string `json:"icon"`
	Sort          int32  `json:"sort"`
}

type DeleteAttrGroupRequest struct {
	AttrGroupIds []int64 `json:"attrGroupIds"`
}

type SaveSPU struct {
	SpuName        string   `json:"spuName"`
	SpuDescription string   `json:"spuDescription"`
	CatalogID      int64    `json:"catalogId"`
	BrandID        int64    `json:"brandId"`
	Weight         float64  `json:"weight"`
	PublishStatus  int32    `json:"publishStatus"`
	Decript        []string `json:"decript"`
	Images         []string `json:"images"`
	Bounds         struct {
		BuyBounds  int32 `json:"buyBounds"`
		GrowBounds int32 `json:"growBounds"`
	} `json:"bounds"`
	BaseAttrs []struct {
		AttrID     int32  `json:"attrId"`
		AttrValues string `json:"attrValues"`
		ShowDesc   int32  `json:"showDesc"`
	} `json:"baseAttrs"`
	Skus []struct {
		Attr []struct {
			AttrID    int32  `json:"attrId"`
			AttrName  string `json:"attrName"`
			AttrValue string `json:"attrValue"`
		} `json:"attr"`
		SkuName     string  `json:"skuName"`
		Price       float64 `json:"price"`
		SkuTitle    string  `json:"skuTitle"`
		SkuSubtitle string  `json:"skuSubtitle"`
		Images      []struct {
			ImgURL     string `json:"imgUrl"`
			DefaultImg int32  `json:"defaultImg"`
		} `json:"images"`
		Descar      []string `json:"descar"`
		FullCount   int32    `json:"fullCount"`
		Discount    float64  `json:"discount"`
		CountStatus int32    `json:"countStatus"`
		FullPrice   int32    `json:"fullPrice"`
		ReducePrice int32    `json:"reducePrice"`
		PriceStatus int32    `json:"priceStatus"`
		MemberPrice []struct {
			ID    int32   `json:"id"`
			Name  string  `json:"name"`
			Price float64 `json:"price"`
		} `json:"memberPrice"`
	} `json:"skus"`
}

type UpdateSpuAttrRequest []struct {
	AttrID    int    `json:"attrId"`
	AttrName  string `json:"attrName"`
	AttrValue string `json:"attrValue"`
	QuickShow int    `json:"quickShow"`
}
