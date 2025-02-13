package service

import (
	"context"

	"github.com/zheyuanf/ecommerce-tiktok/app/product/biz/dal/mysql"
	"github.com/zheyuanf/ecommerce-tiktok/app/product/biz/model"
	product "github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/product"
)

type SearchProductsService struct {
	ctx context.Context
} // NewSearchProductsService new SearchProductsService
func NewSearchProductsService(ctx context.Context) *SearchProductsService {
	return &SearchProductsService{ctx: ctx}
}

// Run create note info
func (s *SearchProductsService) Run(req *product.SearchProductsReq) (resp *product.SearchProductsResp, err error) {
	// Finish your business logic.
	productQuery := model.NewProductQuery(s.ctx, mysql.DB)
	products, err := productQuery.SearchProduct(req.Query)
	if err != nil {
		return nil, err
	}
	var results []*product.Product

	for _, v := range products {
		results = append(results, &product.Product{
			Id:          uint32(v.ID),
			Name:        v.Name,
			Description: v.Description,
			Picture:     v.Picture,
			Price:       v.Price,
			Storage:     v.Storage,
		})
	}

	return &product.SearchProductsResp{
		Results: results,
	}, nil
}
