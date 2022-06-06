package model

type SkuEsModel_Attr struct {
	AttrId    int64  `json:"attrId"`
	AttrName  string `json:"attrName"`
	AttrValue string `json:"attrValue"`
}

type SkuEsModel struct {
	SkuId       int64             `json:"skuId"`
	SpuId       int64             `json:"spuId"`
	SkuTitle    string            `json:"skuTitle"`
	SkuPrice    float64           `json:"skuPrice"`
	SkuImg      string            `json:"skuImg"`
	SaleCount   int64             `json:"saleCount"`
	HasStock    bool              `json:"hasStock"`
	BrandId     int64             `json:"brandId"`
	CatalogId   int64             `json:"catalogId"`
	BrandName   string            `json:"brandName"`
	BrandImg    string            `json:"brandImg"`
	CatalogName string            `json:"catalogName"`
	HotScore    int64             `json:"hotScore"`
	Attrs       []SkuEsModel_Attr `json:"attrs"`
}
