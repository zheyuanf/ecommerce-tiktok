package inventory

import (
	"context"
	inventory "github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/inventory"

	"github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/inventory/inventoryservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() inventoryservice.Client
	Service() string
	GetStock(ctx context.Context, Req *inventory.GetStockRequest, callOptions ...callopt.Option) (r *inventory.GetStockResponse, err error)
	DecreaseStock(ctx context.Context, Req *inventory.DecreaseStockRequest, callOptions ...callopt.Option) (r *inventory.DecreaseStockResponse, err error)
	IncreaseStock(ctx context.Context, Req *inventory.IncreaseStockRequest, callOptions ...callopt.Option) (r *inventory.IncreaseStockResponse, err error)
	BatchGetStock(ctx context.Context, Req *inventory.BatchGetStockRequest, callOptions ...callopt.Option) (r *inventory.BatchGetStockResponse, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := inventoryservice.NewClient(dstService, opts...)
	if err != nil {
		return nil, err
	}
	cli := &clientImpl{
		service:     dstService,
		kitexClient: kitexClient,
	}

	return cli, nil
}

type clientImpl struct {
	service     string
	kitexClient inventoryservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() inventoryservice.Client {
	return c.kitexClient
}

func (c *clientImpl) GetStock(ctx context.Context, Req *inventory.GetStockRequest, callOptions ...callopt.Option) (r *inventory.GetStockResponse, err error) {
	return c.kitexClient.GetStock(ctx, Req, callOptions...)
}

func (c *clientImpl) DecreaseStock(ctx context.Context, Req *inventory.DecreaseStockRequest, callOptions ...callopt.Option) (r *inventory.DecreaseStockResponse, err error) {
	return c.kitexClient.DecreaseStock(ctx, Req, callOptions...)
}

func (c *clientImpl) IncreaseStock(ctx context.Context, Req *inventory.IncreaseStockRequest, callOptions ...callopt.Option) (r *inventory.IncreaseStockResponse, err error) {
	return c.kitexClient.IncreaseStock(ctx, Req, callOptions...)
}

func (c *clientImpl) BatchGetStock(ctx context.Context, Req *inventory.BatchGetStockRequest, callOptions ...callopt.Option) (r *inventory.BatchGetStockResponse, err error) {
	return c.kitexClient.BatchGetStock(ctx, Req, callOptions...)
}
