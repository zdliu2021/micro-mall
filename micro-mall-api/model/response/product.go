package response

type ListCategoryTreeResponse struct {
	CatId        int64                       ` json:"catId"`        // 分类id
	Name         string                      `json:"name"`          // 分类名称
	ParentCid    int64                       `json:"parentCid"`     // 父分类id
	CatLevel     int32                       ` json:"catLevel"`     // 层级
	ShowStatus   int32                       ` json:"showStatus"`   // 是否显示[0-不显示，1显示]
	Sort         int32                       `json:"sort"`          // 排序
	Icon         string                      ` json:"icon"`         // 图标地 g  P
	ProductUnit  string                      ` json:"productUnit"`  // 计量单位
	ProductCount int32                       ` json:"productCount"` // 商品数量
	Children     []*ListCategoryTreeResponse `json:"children,omitempty"`
}

type GetCategoryInfoResponse struct {
	CatId        int64  ` json:"catId"`        // 分类id
	Name         string `json:"name"`          // 分类名称
	ParentCid    int64  `json:"parentCid"`     // 父分类id
	CatLevel     int32  ` json:"catLevel"`     // 层级
	ShowStatus   int32  ` json:"showStatus"`   // 是否显示[0-不显示，1显示]
	Sort         int32  `json:"sort"`          // 排序
	Icon         string ` json:"icon"`         // 图标地 g  P
	ProductUnit  string ` json:"productUnit"`  // 计量单位
	ProductCount int32  ` json:"productCount"` // 商品数量
}

type BrandEntity struct {
	BrandId     int64  ` json:"brandId"`    // 品牌id
	Name        string ` json:"name"`       // 品牌名
	Logo        string ` json:"logo"`       // 品牌logo地址
	Descript    string ` json:"descript"`   // 介绍
	ShowStatus  int32  ` json:"showStatus"` // 显示状态[0-不显示；1-显示]
	FirstLetter string `json:"firstLetter"` // 检索首字母
	Sort        int32  ` json:"sort"`       // 排序
}

type GetBrandListResponse struct {
	List       []BrandEntity `json:"list"`
	TotalCount int32         `json:"totalCount"`
}

type GetBrandInfoRequest struct {
	Brand BrandEntity `json:"brand"`
}

type GetBrandRelationedCateLogResponse struct {
	Id          int64  `json:"id"`
	CateLogId   int64  `json:"catelogId"`
	CateLogName string `json:"catelogName"`
	BrandId     int64  `json:"brandId"`
	BrandName   string `json:"brandName"`
}

type ListCatelogAttrGroupResponse struct {
	AttrGroupId   int64   `json:"attrGroupId"`
	AttrGroupName string  `json:"attrGroupName"`
	CatelogId     int64   `json:"catelogId"`
	Descript      string  `json:"descript"`
	Icon          string  `json:"icon"`
	Sort          int32   ` json:"sort"`
	CatelogPath   []int64 `json:"catelogPath"`
}

type AttrEntity struct {
	AttrId      int64  `json:"attrId"`      //属性id
	AttrName    string `json:"attrName"`    //属性名
	AttrType    int32  `json:"attrType"`    //属性类型，0-销售属性，1-基本属性
	CatelogName string `json:"catelogName"` //所属分类名字
	GroupName   string `json:"groupName"`   //所属分组名字
	Enable      int32  `json:"enable"`      //是否启用
	Icon        string `json:"icon"`        //图标
	SearchType  int32  `json:"searchType"`  //是否需要检索[0-不需要，1-需要]
	ShowDesc    int32  `json:"showDesc"`    //是否展示在介绍上；0-否 1-是
	ValueSelect string `json:"valueSelect"` //可选值列表[用逗号分隔]
	ValueType   int32  `json:"valueType"`   //值类型[0-为单个值，1-可以选择多个值]
}

type GetAttrInfoResponse struct {
	AttrId      int64   `json:"attrId"`      //属性id
	AttrName    string  `json:"attrName"`    //属性名
	AttrType    int32   `json:"attrType"`    //属性类型，0-销售属性，1-基本属性
	Enable      int32   `json:"enable"`      //是否启用
	Icon        string  `json:"icon"`        //图标
	SearchType  int32   `json:"searchType"`  //是否需要检索[0-不需要，1-需要]
	ShowDesc    int32   `json:"showDesc"`    //是否展示在介绍上；0-否 1-是
	ValueSelect string  `json:"valueSelect"` //可选值列表[用逗号分隔]
	ValueType   int32   ` json:"valueType"`  //值类型[0-为单个值，1-可以选择多个值]
	CatelogId   int64   `json:"catelogId"`
	AttrGroupId int32   `json:"attrGroupId"`
	CatelogPath []int64 `json:"catelogPath"`
}

type GetAttrNotCorrelationResponse struct {
	AttrId      int64  `json:"attrId"`      //属性id
	AttrName    string `json:"attrName"`    //属性名
	AttrType    int32  `json:"attrType"`    //属性类型，0-销售属性，1-基本属性
	Enable      int32  `json:"enable"`      //是否启用
	Icon        string `json:"icon"`        //图标
	SearchType  int32  `json:"searchType"`  //是否需要检索[0-不需要，1-需要]
	ShowDesc    int32  `json:"showDesc"`    //是否展示在介绍上；0-否 1-是
	ValueSelect string `json:"valueSelect"` //可选值列表[用逗号分隔]
	ValueType   int32  ` json:"valueType"`  //值类型[0-为单个值，1-可以选择多个值]
	CatelogId   int64  `json:"catelogId"`
	AttrGroupId int32  `json:"attrGroupId"`
}

type GetAttrGroupInfoResponse struct {
	AttrGroupId   int64   `json:"attrGroupId"`
	AttrGroupName string  `json:"attrGroupName"`
	CatelogId     int64   `json:"catelogId"`
	Descript      string  `json:"descript"`
	Icon          string  `json:"icon"`
	Sort          int32   `json:"sort"`
	CatelogPath   []int64 `json:"catelogPath"`
}

type GetAttrRelatedAttrGroupResponse struct {
	AttrId      int64  ` json:"attrId"`     //属性id
	AttrName    string ` json:"attrName"`   //属性名
	AttrType    int32  `json:"attrType"`    //属性类型，0-销售属性，1-基本属性
	CatelogId   int64  ` json:"catelogId"`  //所属分类名字
	Enable      int32  ` json:"enable"`     //是否启用
	Icon        string ` json:"icon"`       //图标
	SearchType  int32  ` json:"searchType"` //是否需要检索[0-不需要，1-需要]
	ShowDesc    int32  ` json:"showDesc"`   //是否展示在介绍上；0-否 1-是
	ValueSelect string `json:"valueSelect"` //可选值列表[用逗号分隔]
	ValueType   int32  ` json:"valueType"`  //值类型[0-为单个值，1-可以选择多个值]
}

type BrandCatelogRelation struct {
	Id          int64  `json:"id,omitempty"`
	BrandId     int64  `json:"brandId,omitempty"`
	BrandName   string `json:"brandName,omitempty"`
	CatelogId   int64  `json:"catelogId,omitempty"`
	CatelogName string `json:"catelogName,omitempty"`
}

type PmsAttrEntity struct {
	AttrId      int64  ` json:"attrId"`     //属性id
	AttrName    string ` json:"attrName"`   //属性名
	AttrType    int32  `json:"attrType"`    //属性类型，0-销售属性，1-基本属性
	CatelogId   int64  ` json:"catelogId"`  //所属分类名字
	Enable      int32  ` json:"enable"`     //是否启用
	Icon        string ` json:"icon"`       //图标
	SearchType  int32  ` json:"searchType"` //是否需要检索[0-不需要，1-需要]
	ShowDesc    int32  ` json:"showDesc"`   //是否展示在介绍上；0-否 1-是
	ValueSelect string `json:"valueSelect"` //可选值列表[用逗号分隔]
	ValueType   int32  ` json:"valueType"`  //值类型[0-为单个值，1-可以选择多个值]
}

type AttrGroupAndAttrsEntity struct {
	AttrGroupId   int64           `json:"attrGroupId"`
	AttrGroupName string          `json:"attrGroupName"`
	CatelogId     int64           `json:"catelogId"`
	Descript      string          `json:"descript"`
	Icon          string          `json:"icon"`
	Sort          int32           `json:"sort"`
	Attrs         []PmsAttrEntity `json:"attrs"`
}

type GetSpuInfoResponse_List struct {
	BrandId        int64   `json:"brandId"`
	BrandName      string  `json:"brandName"`
	CatalogId      int64   `json:"catalogId"`
	CatalogName    string  `json:"catalogName"`
	CreateTime     string  `json:"createTime"`
	Id             int64   `json:"id"`
	PublishStatus  int     `json:"publishStatus"`
	SpuDescription string  `json:"spuDescription"`
	SpuName        string  `json:"spuName"`
	UpdateTime     string  `json:"updateTime"`
	Weight         float64 `json:"weight"`
}
type GetSpuInfoResponse struct {
	TotalCount int                       `json:"totalCount"`
	PageSize   int                       `json:"pageSize"`
	TotalPage  int                       `json:"totalPage"`
	CurrPage   int                       `json:"currPage"`
	List       []GetSpuInfoResponse_List `json:"list"`
}

type SpuAttr struct {
	Id        int         `json:"id"`
	SpuId     int         `json:"spuId"`
	AttrId    int         `json:"attrId"`
	AttrName  string      `json:"attrName"`
	AttrValue string      `json:"attrValue"`
	AttrSort  interface{} `json:"attrSort"`
	QuickShow int         `json:"quickShow"`
}

type ListSkuResponse struct {
	TotalCount int `json:"totalCount"`
	PageSize   int `json:"pageSize"`
	TotalPage  int `json:"totalPage"`
	CurrPage   int `json:"currPage"`
	List       []struct {
		SkuId         int         `json:"skuId"`
		SpuId         int         `json:"spuId"`
		SkuName       string      `json:"skuName"`
		SkuDesc       interface{} `json:"skuDesc"`
		CatalogId     int         `json:"catalogId"`
		BrandId       int         `json:"brandId"`
		SkuDefaultImg string      `json:"skuDefaultImg"`
		SkuTitle      string      `json:"skuTitle"`
		SkuSubtitle   string      `json:"skuSubtitle"`
		Price         float64     `json:"price"`
		SaleCount     int         `json:"saleCount"`
	} `json:"list"`
}

type Category3D struct {
	Catalog2Id string `json:"catalog2Id"`
	Id         string `json:"id"`
	Name       string `json:"name"`
}

type Category2D struct {
	Catalog1Id   string       `json:"Catalog1Id"`
	Catalog3List []Category3D `json:"catalog3List"`
	Id           string       `json:"id"`
	Name         string       `json:"name"`
}
