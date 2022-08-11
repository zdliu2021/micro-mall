package server

import (
	"context"
	proto_product "mall-demo/micro-mall-product/proto/micro-mall-product-proto"
	"mall-demo/micro-mall-product/service"
)

// server is used to implement helloworld.GreeterServer.
type Server struct {
	proto_product.UnimplementedProductRpcServer
	categoryService  service.CategoryService
	brandService     service.BrandService
	attrGroupService service.AttrGroupService
	attrService      service.AttrService
	spuService       service.SpuService
}

// category

func (s *Server) ListCategoryTree(ctx context.Context, req *proto_product.ListCategoryTreeRequest) (*proto_product.ListCategoryTreeResponse, error) {
	return s.categoryService.ListCategoryTree(ctx, req)
}

func (s *Server) DeleteCategory(ctx context.Context, req *proto_product.DeleteCategoryRequest) (*proto_product.DeleteCategoryResponse, error) {
	return s.categoryService.DeleteCategory(ctx, req)
}

func (s *Server) SaveCategory(ctx context.Context, req *proto_product.SaveCategoryRequest) (*proto_product.SaveCategoryResponse, error) {
	return s.categoryService.SaveCategory(ctx, req)
}

func (s *Server) UpdateCategory(ctx context.Context, req *proto_product.UpdateCategoryRequest) (*proto_product.UpdateCategoryResponse, error) {
	return s.categoryService.UpdateCategory(ctx, req)
}

func (s *Server) GetCategoryInfo(ctx context.Context, req *proto_product.GetCategoryInfoRequest) (*proto_product.GetCategoryInfoResponse, error) {
	return s.categoryService.GetCategoryInfo(ctx, req)
}

// brand

func (s *Server) GetBrandList(ctx context.Context, req *proto_product.GetBrandListRequest) (*proto_product.GetBrandListResponse, error) {
	return s.brandService.GetBrandList(ctx, req)
}

func (s *Server) GetBrandInfo(ctx context.Context, req *proto_product.GetBrandInfoRequest) (*proto_product.GetBrandInfoResponse, error) {
	return s.brandService.GetBrandInfo(ctx, req)
}

func (s *Server) UpdateBrand(ctx context.Context, req *proto_product.UpdateBrandRequest) (*proto_product.UpdateBrandResponse, error) {
	return s.brandService.UpdateBrand(ctx, req)
}

func (s *Server) SaveBrand(ctx context.Context, req *proto_product.SaveBrandRequest) (*proto_product.SaveBrandResponse, error) {
	return s.brandService.SaveBrand(ctx, req)
}
func (s *Server) DeleteBrand(ctx context.Context, req *proto_product.DeleteBrandRequest) (*proto_product.DeleteBrandResponse, error) {
	return s.brandService.DeleteBrand(ctx, req)
}
func (s *Server) UpdateBrandStatus(ctx context.Context, req *proto_product.UpdateBrandStatusRequest) (*proto_product.UpdateBrandStatusResponse, error) {
	return s.brandService.UpdateBrandStatus(ctx, req)
}

func (s *Server) GetBrandRelatedCateLog(ctx context.Context, req *proto_product.GetBrandRelatedCateLogRequest) (*proto_product.GetBrandRelatedCateLogResponse, error) {
	return s.brandService.GetBrandRelatedCateLog(ctx, req)
}

func (s *Server) SaveBrandCatelogRelation(ctx context.Context, req *proto_product.SaveBrandCatelogRelationRequest) (*proto_product.SaveBrandCatelogRelationResponse, error) {
	return s.brandService.SaveBrandCatelogRelation(ctx, req)
}

func (s *Server) DeleteBrandCatelogRelation(ctx context.Context, req *proto_product.DeleteBrandCatelogRelationRequest) (*proto_product.DeleteBrandCatelogRelationResponse, error) {
	return s.brandService.DeleteBrandCatelogRelation(ctx, req)
}
func (s *Server) GetBrandRelatedWithCatelog(ctx context.Context, req *proto_product.GetBrandRelatedWithCatelogRequest) (*proto_product.GetBrandRelatedWithCatelogResponse, error) {
	return s.brandService.GetBrandRelatedWithCatelog(ctx, req)
}

// attrgorup

func (s *Server) ListCatelogAttrGroup(ctx context.Context, req *proto_product.ListCatelogAttrGroupRequest) (*proto_product.ListCatelogAttrGroupResponse, error) {
	return s.attrGroupService.ListCatelogAttrGroup(ctx, req)
}

func (s *Server) GetAttrGroupInfo(ctx context.Context, req *proto_product.GetAttrGroupInfoRequest) (*proto_product.GetAttrGroupInfoResponse, error) {
	return s.attrGroupService.GetAttrGroupInfo(ctx, req)
}

func (s *Server) UpdateAttrGroup(ctx context.Context, req *proto_product.UpdateAttrGroupRequest) (*proto_product.UpdateAttrGroupResponse, error) {
	return s.attrGroupService.UpdateAttrGroup(ctx, req)
}
func (s *Server) SaveAttrGroup(ctx context.Context, req *proto_product.SaveAttrGroupRequest) (*proto_product.SaveAttrGroupResponse, error) {
	return s.attrGroupService.SaveAttrGroup(ctx, req)
}
func (s *Server) DeleteAttrGroup(ctx context.Context, req *proto_product.DeleteAttrGroupRequest) (*proto_product.DeleteAttrGroupResponse, error) {
	return s.attrGroupService.DeleteAttrGroup(ctx, req)
}
func (s *Server) GetAllGroupAndAttrRelatedWithCatelog(ctx context.Context, req *proto_product.GetAllGroupAndAttrRelatedWithCatelogRequest) (*proto_product.GetAllGroupAndAttrRelatedWithCatelogResponse, error) {
	return s.attrGroupService.GetAllGroupAndAttrRelatedWithCatelog(ctx, req)
}

//attr
func (s *Server) GetAttrListRelatedWithCatelog(ctx context.Context, req *proto_product.GetAttrListRelatedWithCatelogRequest) (*proto_product.GetAttrListRelatedWithCatelogResponse, error) {
	return s.attrService.GetAttrListRelatedWithCatelog(ctx, req)
}

func (s *Server) SaveProductAttr(ctx context.Context, req *proto_product.SaveProductAttrRequest) (*proto_product.SaveProductAttrResponse, error) {
	return s.attrService.SaveProductAttr(ctx, req)
}
func (s *Server) GetAttrInfo(ctx context.Context, req *proto_product.GetAttrInfoRequest) (*proto_product.GetAttrInfoResponse, error) {
	return s.attrService.GetAttrInfo(ctx, req)
}
func (s *Server) GetAttrNotCorrelation(ctx context.Context, req *proto_product.GetAttrNotCorrelationRequest) (*proto_product.GetAttrNotCorrelationResponse, error) {
	return s.attrService.GetAttrNotCorrelation(ctx, req)
}
func (s *Server) DeleteAttrAttrGroupRelation(ctx context.Context, req *proto_product.DeleteAttrAttrGroupRelationRequest) (*proto_product.DeleteAttrAttrGroupRelationResponse, error) {
	return s.attrService.DeleteAttrAttrGroupRelation(ctx, req)
}
func (s *Server) SaveAttrAttrGroupRelation(ctx context.Context, req *proto_product.SaveAttrAttrGroupRelationRequest) (*proto_product.SaveAttrAttrGroupRelationResponse, error) {
	return s.attrService.SaveAttrAttrGroupRelation(ctx, req)
}
func (s *Server) GetAttrRelatedAttrGroup(ctx context.Context, req *proto_product.GetAttrRelatedAttrGroupRequest) (*proto_product.GetAttrRelatedAttrGroupResponse, error) {
	return s.attrService.GetAttrRelatedAttrGroup(ctx, req)
}
func (s *Server) UpdateAttr(ctx context.Context, req *proto_product.UpdateAttrRequest) (*proto_product.UpdateAttrResponse, error) {
	return s.attrService.UpdateAttr(ctx, req)
}

func (s *Server) DeleteAttr(ctx context.Context, req *proto_product.DeleteAttrRequest) (*proto_product.DeleteAttrResponse, error) {
	return s.attrService.DeleteAttr(ctx, req)
}

//spu

func (s *Server) SaveSpu(ctx context.Context, req *proto_product.SaveSpuRequest) (*proto_product.SaveSpuResponse, error) {
	return s.spuService.SaveSpu(ctx, req)
}
func (s *Server) SearchSpuInfo(ctx context.Context, req *proto_product.SearchSpuInfoRequest) (*proto_product.SearchSpuInfoResponse, error) {
	return s.spuService.SearchSpuInfo(ctx, req)
}
func (s *Server) UpSpu(ctx context.Context, req *proto_product.UpSpuRequest) (*proto_product.UpSpuResponse, error) {
	return s.spuService.UpSpu(ctx, req)
}
func (s *Server) SearchSkuInfo(ctx context.Context, req *proto_product.SearchSkuInfoRequest) (*proto_product.SearchSkuInfoResponse, error) {
	return s.spuService.SearchSkuInfo(ctx, req)
}
func (s *Server) GetSpuInfo(ctx context.Context, req *proto_product.GetSpuInfoRequest) (*proto_product.GetSpuInfoResponse, error) {
	return s.spuService.GetSpuInfo(ctx, req)
}
func (s *Server) UpdateSkuAttrs(ctx context.Context, req *proto_product.UpdateSpuAttrsRequest) (*proto_product.UpdateSpuAttrsResponse, error) {
	return s.spuService.UpdateSpuAttrs(ctx, req)
}

func (s *Server) GetSkuItem(ctx context.Context, req *proto_product.GetSkuItemRequest) (*proto_product.GetSkuItemResponse, error) {
	return s.spuService.GetSkuItem(ctx, req)
}
