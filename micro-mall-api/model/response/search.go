package response

type SearchProductResponse struct {
	Products  []SearchProductResponse_Entity  `json:"products,omitempty"`
	PageNum   int32                           `json:"page_num,omitempty"`
	Total     int32                           `json:"total,omitempty"` //总记录数
	TotalPage int32                           `json:"total_page,omitempty"`
	Brands    []SearchProductResponse_Brand   `json:"brands,omitempty"`
	Attrs     []SearchProductResponse_Attr    `json:"attrs,omitempty"`
	Catalogs  []SearchProductResponse_Catalog `json:"catalogs,omitempty"`
}

type SearchProductResponse_Entity struct {
	SkuId       int64   `json:"sku_id,omitempty"`
	SpuId       int64   `json:"spu_id,omitempty"`
	SkuTitle    string  ` json:"sku_title,omitempty"`
	SkuPrice    float64 ` json:"sku_price,omitempty"`
	SkuImg      string
	SaleCount   int64                               `json:"sale_count,omitempty"`
	HasStock    bool                                `json:"has_stock,omitempty"`
	BrandId     int64                               `json:"brand_id,omitempty"`
	CatalogId   int64                               `json:"catalog_id,omitempty"`
	BrandName   string                              `json:"brand_name,omitempty"`
	BrandImg    string                              `json:"brand_img,omitempty"`
	CatalogName string                              `json:"catalog_name,omitempty"`
	HotScore    int64                               `json:"hot_score,omitempty"`
	Attrs       []SearchProductResponse_Entity_Attr `json:"attrs,omitempty"`
}

type SearchProductResponse_Entity_Attr struct {
	AttrId    int64  `json:"attr_id,omitempty"`
	AttrName  string `json:"attr_name,omitempty"`
	AttrValue string `json:"attr_value,omitempty"`
}

type SearchProductResponse_Brand struct {
	BrandId   int64  `json:"brand_id,omitempty"`
	BrandName string `json:"brand_name,omitempty"`
	BrandImg  string `json:"brand_img,omitempty"`
}

type SearchProductResponse_Attr struct {
	AttrId    int64    `json:"attr_id,omitempty"`
	AttrName  string   `json:"attr_name,omitempty"`
	AttrValue []string `json:"att_value,omitempty"`
}

type SearchProductResponse_Catalog struct {
	CatalogId   int64  `json:"catalog_id,omitempty"`
	CatalogName string `json:"catalog_name,omitempty"`
}
