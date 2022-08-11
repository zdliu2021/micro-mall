package html

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	response2 "mall-demo/micro-mall-api/model/response"
	proto_product "mall-demo/micro-mall-api/proto/micro-mall-product-proto"
	proto_search "mall-demo/micro-mall-api/proto/micro-mall-search-proto"
	rpc_client "mall-demo/micro-mall-api/rpc-client"
	"net/http"
	"strconv"
	"strings"
)

func Login(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login/login.html", gin.H{})
}

func Index(ctx *gin.Context) {
	categories := Index_GetCategory()
	ctx.HTML(http.StatusOK, "index/index.html", gin.H{
		"categories": categories,
	})
}

func Reg(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "reg/reg.html", gin.H{})
}

func IndexCatalog(ctx *gin.Context) {
	categories := Index_GetCategory()
	res := make(map[string][]response2.Category2D)
	for i := 0; i < len(categories); i++ {
		tmp := make([]response2.Category2D, 0, 10)
		for j := 0; j < len(categories[i].Children); j++ {
			var tmp2 response2.Category2D
			tmp2.Catalog3List = make([]response2.Category3D, 0, 10)
			tmp2.Id = strconv.Itoa(int(categories[i].Children[j].CatId))
			tmp2.Name = categories[i].Children[j].Name
			tmp2.Catalog1Id = strconv.Itoa(int(categories[i].CatId))
			for k := 0; k < len(categories[i].Children[j].Children); k++ {
				var tmp3 response2.Category3D
				tmp3.Catalog2Id = strconv.Itoa(int(categories[i].Children[j].CatId))
				tmp3.Id = strconv.Itoa(int(categories[i].Children[j].Children[k].CatId))
				tmp3.Name = categories[i].Children[j].Children[k].Name
				tmp2.Catalog3List = append(tmp2.Catalog3List, tmp3)
			}
			tmp = append(tmp, tmp2)
		}
		res[strconv.Itoa(int(categories[i].CatId))] = tmp
	}
	ctx.JSON(http.StatusOK, res)
}

func SearchList(ctx *gin.Context) {
	rpcClient := rpc_client.GetSearchClient()
	var req proto_search.SearchProductRequest

	req.Keyword = ctx.Query("keyword")
	req.CatalogId, _ = strconv.ParseInt(ctx.Query("catalog3Id"), 10, 64)
	req.Sort = ctx.Query("sort")
	hasStock, _ := strconv.Atoi(ctx.Query("hasStock"))
	req.HasStock = int32(hasStock)
	req.SkuPrice = ctx.Query("skuPrice")
	pageNum, _ := strconv.Atoi(ctx.Query("pageNum"))
	req.PageNum = int32(pageNum)
	req.Attrs = ctx.QueryArray("attrs")
	brandIds := ctx.QueryArray("brandId")
	req.BrandId = make([]int64, 0, 10)
	for i := 0; i < len(brandIds); i++ {
		brandId, _ := strconv.ParseInt(brandIds[i], 10, 64)
		req.BrandId = append(req.BrandId, brandId)
	}

	resp, err := rpcClient.SearchProduct(context.TODO(), &req)

	if err != nil {
		response2.FailWithMessage(err.Error(), ctx)
		return
	}
	//response2.OkWithData(resp, ctx)
	res := response2.SearchProductResponse{}
	copier.CopyWithOption(&res, &resp, copier.Option{IgnoreEmpty: true, DeepCopy: true})

	ctx.HTML(http.StatusOK, "search/list.html", gin.H{"result": res})
}

func Index_GetCategory() (res []response2.ListCategoryTreeResponse) {
	var solve func(*response2.ListCategoryTreeResponse, *proto_product.CategoryEntity)
	solve = func(r1 *response2.ListCategoryTreeResponse, r2 *proto_product.CategoryEntity) {
		r1.CatId = r2.CatId
		r1.Name = r2.Name
		r1.ShowStatus = r2.ShowStatus
		r1.CatLevel = r2.CatLevel
		r1.Icon = r2.Icon
		r1.ProductCount = r2.ProductCount
		r1.Sort = r2.Sort
		r1.ParentCid = r2.ParentCid
		r1.Children = make([]*response2.ListCategoryTreeResponse, 0, 10)
		for i := 0; i < len(r2.Children); i++ {
			var tmp response2.ListCategoryTreeResponse
			solve(&tmp, r2.Children[i])
			r1.Children = append(r1.Children, &tmp)
		}
	}

	rpcClient := rpc_client.GetProductClient()
	req := &proto_product.ListCategoryTreeRequest{}
	resp, err := rpcClient.ListCategoryTree(context.TODO(), req)
	if err != nil {
		return
	}
	res = make([]response2.ListCategoryTreeResponse, 0, 10)
	if len(resp.CategoryEntities) != 0 {
		for i := 0; i < len(resp.CategoryEntities); i++ {
			var tmp response2.ListCategoryTreeResponse
			solve(&tmp, resp.CategoryEntities[i])
			res = append(res, tmp)
		}
	}
	return
}

func SkuItem(ctx *gin.Context) {
	skuId, _ := strconv.Atoi(strings.Split(ctx.Query("skuId"), ".")[0])
	rpcClient := rpc_client.GetProductClient()
	resp, err := rpcClient.GetSkuItem(context.TODO(), &proto_product.GetSkuItemRequest{
		SkuId: int64(skuId),
	})
	if err != nil {
		response2.FailWithMessage(err.Error(), ctx)
		return
	}

	ctx.HTML(http.StatusOK, "item/item.html", gin.H{
		"item": resp,
	})
}
