package model

type SkuEsModel_Attr struct {
	AttrId    int64
	AttrName  string
	AttrValue string
}

type SkuEsModel struct {
	SkuId       int64
	SpuId       int64
	SkuTitle    string
	SkuPrice    float64
	SkuImg      string
	SaleCount   int64
	HasStock    bool
	BrandId     int64
	CatalogId   int64
	BrandName   string
	BrandImg    string
	CatalogName string
	HotScore    int64
	Attrs       []SkuEsModel_Attr
}
