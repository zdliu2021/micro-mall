package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/olivere/elastic/v7"
	"mall-demo/micro-mall-search/global"
	"mall-demo/micro-mall-search/model"
	proto_search "mall-demo/micro-mall-search/proto/micro-mall-search-proto"
	"math"
	"strconv"
	"strings"
)

type BaseService struct {
}

func (bs *BaseService) ProductStatusUp(ctx context.Context, req *proto_search.ProductStatusUpRequest) (*proto_search.ProductStatusUpResponse, error) {
	fmt.Printf("%+v", req)
	var esModels []model.SkuEsModel
	copier.CopyWithOption(&esModels, &req.Entities, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	reqEs := global.GVA_ES.Bulk().Index(global.ProductIndex)
	for _, esModel := range esModels {
		doc := elastic.NewBulkIndexRequest().Id(strconv.Itoa(int(esModel.SkuId))).Doc(esModel)
		reqEs.Add(doc)
	}
	if _, err := reqEs.Do(context.TODO()); err != nil {
		global.GVA_LOG.Error(err.Error())
	}
	return &proto_search.ProductStatusUpResponse{}, nil
}

func (bs *BaseService) SearchProduct(ctx context.Context, req *proto_search.SearchProductRequest) (*proto_search.SearchProductResponse, error) {
	reqEs := global.GVA_ES.Search().Index(global.ProductIndex)

	//1.keyword
	boolQuery := elastic.NewBoolQuery()
	if req.Keyword != "" {
		boolQuery.Must(elastic.NewMatchQuery("skuTitle", req.Keyword))
	}
	// 2. catalogId
	if req.CatalogId != 0 {
		boolQuery.Filter(elastic.NewTermsQuery("catalogId", req.CatalogId))
	}
	// 3. brandId
	if len(req.BrandId) > 0 {
		brandIds := make([]interface{}, 0, 10)
		for i := 0; i < len(req.BrandId); i++ {
			brandIds = append(brandIds, req.BrandId[i])
		}
		boolQuery.Filter(elastic.NewTermsQuery("brandId", brandIds...))
	}
	// 4. attrs
	if len(req.Attrs) > 0 {
		for _, attr := range req.Attrs {
			nestedBoolQuery := elastic.NewBoolQuery()
			ss := strings.Split(attr, "_")
			attrId, _ := strconv.ParseInt(ss[0], 10, 64)
			attrValue := strings.Split(ss[1], ":")
			nestedBoolQuery.Must(elastic.NewTermQuery("attrs.attrId", attrId))
			nestedBoolQuery.Must(elastic.NewTermsQueryFromStrings("attrs.attrValue", attrValue...))
			nestedQuery := elastic.NewNestedQuery("attrs", nestedBoolQuery)
			boolQuery.Filter(nestedQuery)
		}
	}
	// 5. hasStock
	if req.HasStock == 1 {
		boolQuery.Filter(elastic.NewTermQuery("hasStock", req.HasStock))
	}

	// 6. skuPrice
	if req.SkuPrice != "" {
		rangeQuery := elastic.NewRangeQuery("skuPrice")
		s := strings.Split(req.SkuPrice, "_")
		if strings.HasPrefix(req.SkuPrice, "_") {
			maxs, _ := strconv.ParseFloat(s[1], 64)
			rangeQuery.Lte(maxs)
		} else if strings.HasSuffix(req.SkuPrice, "_") {
			mins, _ := strconv.ParseFloat(s[0], 64)
			rangeQuery.Gte(mins)
		} else {
			maxs, _ := strconv.ParseFloat(s[1], 64)
			mins, _ := strconv.ParseFloat(s[0], 64)
			rangeQuery.Lte(maxs).Gte(mins)
		}
		boolQuery.Filter(rangeQuery)
	}

	// merge all query
	reqEs.Query(boolQuery)

	// 7. sort
	if req.Sort != "" {
		s := strings.Split(req.Sort, "_")
		ascending := s[1] == "asc"
		reqEs.Sort(s[0], ascending)
	}
	// 8. 分页
	if req.PageNum == 0 {
		req.PageNum = 1
	}

	reqEs.From(int((req.PageNum - 1) * 10)).Size(10)

	// 9. highlight
	if req.Keyword != "" {
		highLightQuery := elastic.NewHighlight()
		highLightQuery.Field("skuTitle")
		highLightQuery.PreTags("<b style='color:red'>")
		highLightQuery.PostTags("</b>")
		reqEs.Highlight(highLightQuery)
	}

	//// 聚合分析
	brand_agg := elastic.NewTermsAggregation().Field("brandId").Size(50)
	brand_agg.SubAggregation("brand_name", elastic.NewTermsAggregation().Field("brandName").Size(1))
	brand_agg.SubAggregation("brand_img", elastic.NewTermsAggregation().Field("brandImg").Size(1))
	reqEs.Aggregation("brand_agg", brand_agg)

	catalog_agg := elastic.NewTermsAggregation().Field("catalogId").Size(50)
	catalog_agg.SubAggregation("catalog_name", elastic.NewTermsAggregation().Field("catalogName").Size(1))
	reqEs.Aggregation("catalog_agg", catalog_agg)

	nestedAgg := elastic.NewNestedAggregation().Path("attrs")
	attr_agg := elastic.NewTermsAggregation().Field("attrs.attrId")
	attr_agg.SubAggregation("attr_name", elastic.NewTermsAggregation().Field("attrs.attrName"))
	attr_agg.SubAggregation("attr_value", elastic.NewTermsAggregation().Field("attrs.attrValue"))
	nestedAgg.SubAggregation("attr_agg", attr_agg)
	reqEs.Aggregation("nestedAttr_agg", nestedAgg)

	result, err := reqEs.Do(context.TODO())
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return &proto_search.SearchProductResponse{}, err
	} else {
		return BuildSearchProductResult(result, req), nil
	}
}

func BuildSearchProductResult(result *elastic.SearchResult, req *proto_search.SearchProductRequest) (resp *proto_search.SearchProductResponse) {
	resp = &proto_search.SearchProductResponse{}
	// 1. 设置得到的商品
	hits := result.Hits
	var esModels []model.SkuEsModel
	esModels = make([]model.SkuEsModel, 0, 10)
	for _, hit := range hits.Hits {
		esModel := model.SkuEsModel{}
		json.Unmarshal(hit.Source, &esModel)
		if len(req.Keyword) != 0 {
			esModel.SkuTitle = hit.Highlight["skuTitle"][0]
		}
		esModels = append(esModels, esModel)
	}
	fmt.Printf(" esModels : %+v \n\n", esModels)

	copier.CopyWithOption(&resp.Products, &esModels, copier.Option{IgnoreEmpty: true, DeepCopy: true})

	// 2. 当前商品涉及到的所有属性信息
	resp.Attrs = make([]*proto_search.SearchProductResponse_Attr, 0, 10)
	attr_Agg, _ := result.Aggregations.Nested("nestedAttr_agg")
	attr_id_agg, _ := attr_Agg.Aggregations.Terms("attr_agg")
	for _, bucket := range attr_id_agg.Buckets {
		attr := proto_search.SearchProductResponse_Attr{}
		// 得到属性的id
		attrId, _ := bucket.KeyNumber.Int64()
		attr.AttrId = attrId

		// 得到属性的名字
		attr_name_agg, _ := bucket.Aggregations.Terms("attr_name")
		attr_name := attr_name_agg.Buckets[0].Key.(string)
		attr.AttrName = attr_name

		// 得到属性的所有值
		attr.AttrValue = make([]string, 0, 10)
		attr_value_agg, _ := bucket.Aggregations.Terms("attr_value")
		for _, _bucket := range attr_value_agg.Buckets {
			attr.AttrValue = append(attr.AttrValue, _bucket.Key.(string))
		}
		resp.Attrs = append(resp.Attrs, &attr)
	}

	// 3. 当前商品涉及到的品牌信息
	resp.Brands = make([]*proto_search.SearchProductResponse_Brand, 0, 10)
	brand_aggs, _ := result.Aggregations.Terms("brand_agg")
	for _, brand_agg := range brand_aggs.Buckets {
		brand := proto_search.SearchProductResponse_Brand{}

		// 得到品牌的id
		brandId, _ := brand_agg.KeyNumber.Int64()
		brand.BrandId = brandId

		// 得到品牌的名字
		brand_name_aggs, _ := brand_agg.Aggregations.Terms("brand_name")
		brand.BrandName = brand_name_aggs.Buckets[0].Key.(string)

		// 得到品牌的图片
		brand_img_aggs, _ := brand_agg.Aggregations.Terms("brand_img")
		brand.BrandImg = brand_img_aggs.Buckets[0].Key.(string)
		resp.Brands = append(resp.Brands, &brand)
	}

	// 4. 当前商品涉及的分类信息
	resp.Catalogs = make([]*proto_search.SearchProductResponse_Catalog, 0, 10)
	catalog_aggs, _ := result.Aggregations.Terms("catalog_agg")
	for _, catalog_agg := range catalog_aggs.Buckets {
		catalog := proto_search.SearchProductResponse_Catalog{}

		// 得到分类的id
		catalogId, _ := catalog_agg.KeyNumber.Int64()
		catalog.CatalogId = catalogId

		// 得到分类的名字
		catalog_name_aggs, _ := catalog_agg.Aggregations.Terms("catalog_name")
		catalog.CatalogName = catalog_name_aggs.Buckets[0].Key.(string)
		resp.Catalogs = append(resp.Catalogs, &catalog)
	}
	resp.PageNum = req.PageNum
	resp.Total = int32(result.Hits.TotalHits.Value)
	resp.TotalPage = int32(math.Ceil(float64(resp.Total) / 10.0))
	return resp
}
