package service

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/olivere/elastic/v7"
	"mall-demo/micro-mall-search/global"
	"mall-demo/micro-mall-search/model"
	proto_search "mall-demo/micro-mall-search/proto/micro-mall-search-proto"
	"strconv"
)

type BaseService struct {
}

func (bs *BaseService) ProductStatusUp(ctx context.Context, req *proto_search.ProductStatusUpRequest) (*proto_search.ProductStatusUpResponse, error) {
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
