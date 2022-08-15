package router

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"mall-demo/micro-mall-api/middleware"
	"mall-demo/micro-mall-api/router/api/v1"
	"mall-demo/micro-mall-api/router/html"
	rpc_client "mall-demo/micro-mall-api/rpc-client"
	"time"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	// 开启跨域
	router.Use(middleware.Cors())
	router.Use(middleware.RateLimitMiddleware(time.Second, 100, 100))

	// 设置session
	store, err := redis.NewStore(10, "tcp", "localhost:6399", "", []byte("secret"))
	if err != nil {
		fmt.Println("#####################")
		fmt.Println(err.Error())
	}
	router.Use(sessions.Sessions("sessionid", store))

	// HTML
	router.LoadHTMLGlob("templates/**/*.html")
	router.Static("/assets", "templates/")
	htmlRouter := router.Group("")
	{
		htmlRouter.GET("login.html", html.Login)
		htmlRouter.GET("index.html", html.Index)
		htmlRouter.GET("reg.html", html.Reg)
		htmlRouter.GET("search/list.html", html.SearchList)
		htmlRouter.GET("index/json/catalog.json", html.IndexCatalog)
		htmlRouter.GET("item/:skuId", html.SkuItem)
	}

	frontend_route := router.Group("")
	{
		frontend_route.GET("sms/sendCode", v1.SendCode)
		frontend_route.POST("register", v1.Register)
		frontend_route.POST("login", v1.Login)
		frontend_route.GET("oauth2/gitte/success", v1.OAuthGitteSuccess)

		//cart
		frontend_route.GET("currentUserCartItems", v1.CurrentUserCartItems)
		frontend_route.GET("cartList", v1.CartList)
		frontend_route.GET("addCartItem", v1.AddCartItem)
		frontend_route.GET("checkItem", v1.CheckItem)
		frontend_route.GET("deleteItem", v1.DeleteItem)
		frontend_route.GET("getCartItem", v1.GetCartItem)
	}

	// api
	route := router.Group("renren-fast")
	ProductGroup := route.Group("product")
	{
		ProductCategoryGroup := ProductGroup.Group("category")
		{
			ProductCategoryGroup.GET("list/tree", v1.ListCategoryTree)   //1
			ProductCategoryGroup.POST("delete", v1.DeleteCategory)       //1
			ProductCategoryGroup.POST("save", v1.SaveCategory)           //1
			ProductCategoryGroup.POST("update", v1.UpdateCategory)       //1
			ProductCategoryGroup.GET("info/:cat_id", v1.GetCategoryInfo) //1
		}
		ProductAttrGroupGroup := ProductGroup.Group("attrgroup")
		{
			ProductAttrGroupGroup.GET("list/:catelogId", v1.ListCatelogAttrGroup)               //1
			ProductAttrGroupGroup.GET("info/:attrGroupID", v1.GetAttrGroupInfo)                 // 1
			ProductAttrGroupGroup.POST("delete", v1.DeleteAttrGroup)                            // 1
			ProductAttrGroupGroup.POST("update", v1.UpdateAttrGroup)                            //1
			ProductAttrGroupGroup.POST("save", v1.SaveAttrGroup)                                //1
			ProductAttrGroupGroup.GET("noattr/relation/:attrgroupId", v1.GetAttrNotCorrelation) //1
			ProductAttrGroupGroup.POST("attr/relation/delete", v1.DeleteAttrAttrGroupRelation)  //1
			ProductAttrGroupGroup.POST("attr/relation", v1.SaveAttrAttrGroupRelation)           //1
			ProductAttrGroupGroup.GET(":attrGroupId/attr/relation", v1.GetAttrRelatedAttrGroup) //1
			ProductAttrGroupGroup.GET("withattr/:catelogId", v1.GetAllGroupAndAttrRelatedWithCatelog)

		}

		ProductBrandGroup := ProductGroup.Group("brand")
		{
			ProductBrandGroup.GET("list", v1.GetBrandList)           //1
			ProductBrandGroup.GET("info/:brand_id", v1.GetBrandInfo) //1
			ProductBrandGroup.POST("update", v1.UpdateBrand)         //1
			ProductBrandGroup.POST("save", v1.SaveBrand)             //1
			ProductBrandGroup.POST("delete", v1.DeleteBrand)         //1
		}
		ProductCategoryBrandRelationGroup := ProductGroup.Group("categorybrandrelation")
		{
			ProductCategoryBrandRelationGroup.GET("brands/list", v1.GetBrandRelatedWithCatelog)
			ProductCategoryBrandRelationGroup.GET("catelog/list", v1.GetBrandRelationedCateLog) //1
			ProductCategoryBrandRelationGroup.POST("save", v1.SaveBrandCatelogRelation)         //1
			ProductCategoryBrandRelationGroup.POST("delete", v1.DeleteBrandCatelogRelation)     //1
		}
		ProductAttrGroup := ProductGroup.Group("attr")
		{
			ProductAttrGroup.GET("base/list/:catelogId", v1.GetAttrListRelatedWithCatelog)     //1
			ProductAttrGroup.GET("sale/list/:catelogId", v1.GetSaleAttrListRelatedWithCatelog) //1
			ProductAttrGroup.POST("save", v1.SaveProductAttr)
			ProductAttrGroup.GET("info/:attrId", v1.GetAttrInfo) //1
			ProductAttrGroup.POST("update", v1.UpdateAttr)       //1
			ProductAttrGroup.POST("delete", v1.DeleteAttr)       //1
			ProductAttrGroup.GET("base/listforspu/:spuId", v1.GetSpuAttrsInfo)
			ProductAttrGroup.POST("update/:spuId", v1.UpdateSpuAttrs)
		}
		ProductSpuInfoGroup := ProductGroup.Group("spuinfo")
		{
			///micro-mall-micro-mall-product-proto/spuinfo/list
			ProductSpuInfoGroup.GET("list", v1.ListSpu)
			// /micro-mall-micro-mall-product-proto/spuinfo/save
			ProductSpuInfoGroup.POST("save", v1.SaveSpu)
			ProductSpuInfoGroup.POST(":spuId/up", v1.UpSpu)
		}
		ProductSkuInfoGroup := ProductGroup.Group("skuinfo")
		{
			///micro-mall-micro-mall-product-proto/skuinfo/list
			ProductSkuInfoGroup.GET("list", v1.ListSku)
		}
	}
	MemberGroup := route.Group("member")
	{
		MemberGroup.GET("memberlevel/list", v1.GetMemberLevel)
		MemberGroup.POST("memberlevel/save", v1.SaveMemberLevel)
		MemberGroup.POST("memberlevel/delete", v1.DeleteMemberLevel)
		MemberGroup.POST("memberlevel/update", v1.UpdateMemberLevel)

		MemberGroup.GET("member/list", v1.GetMember)
		MemberGroup.POST("member/save", v1.SaveMember)
		MemberGroup.POST("member/delete", v1.DeleteMember)
		MemberGroup.POST("member/update", v1.UpdateMember)
	}

	WareGroup := route.Group("ware")
	{
		WareGroup.GET("wareinfo/list", v1.GetWareInfoList)
		WareGroup.POST("wareinfo/save", v1.SaveWare)
		WareGroup.GET("waresku/list", v1.GetSkuWareInfo)
		WareGroup.GET("purchasedetail/list", v1.GetPurchaseDetailedInfo)
		WareGroup.POST("purchasedetail/save", v1.SavePurchaseDetail)

		WareGroup.GET("purchase/list", v1.GetPurchaseList)
		WareGroup.POST("purchase/save", v1.SavePurchase)
		WareGroup.POST("purchase/merge", v1.MergePurchase)
		WareGroup.GET("purchase/unreceive/list", v1.GetUnReceivedPurchaseInfo)
		WareGroup.POST("purchase/received", v1.ReceivePurchase)
		WareGroup.POST("purchase/update", v1.UpdatePurchase) //分配采购单
	}

	ThirdPartyGroup := route.Group("thirdparty")
	{
		ThirdPartyOSSGroup := ThirdPartyGroup.Group("oss")
		{
			ThirdPartyOSSGroup.GET("policy", v1.GetOSSPolicy)
		}
	}

	return router
}

func InitRpcClients() {
	rpc_client.InitMemberRpcClient()
	rpc_client.InitCouponRpcClient()
	rpc_client.InitWareRpcClient()
	//rpc_client.InitSearchRpcClient()
	rpc_client.InitProductRpcClient()
	rpc_client.InitThirdPartyRpcClient()
	rpc_client.InitAuthServerRpcClient()
	rpc_client.InitCartRpcClient()
}
