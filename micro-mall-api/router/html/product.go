package html

import (
	"context"
	"github.com/gin-gonic/gin"
	response2 "mall-demo/micro-mall-api/model/response"
	proto_product "mall-demo/micro-mall-api/proto/micro-mall-product-proto"
	rpc_client "mall-demo/micro-mall-api/rpc-client"
	"net/http"
)

func Login(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login/index.html", gin.H{})
}

func Index(ctx *gin.Context) {
	categories := Index_GetCategory()
	ctx.HTML(http.StatusOK, "index/index.html", gin.H{
		"categories": categories,
	})
}

func SearchList(ctx *gin.Context) {

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