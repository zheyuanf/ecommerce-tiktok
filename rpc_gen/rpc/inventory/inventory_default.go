package inventory

import (
	"context"
	inventory "github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/inventory"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func GetStock(ctx context.Context, req *inventory.GetStockRequest, callOptions ...callopt.Option) (resp *inventory.GetStockResponse, err error) {
	resp, err = defaultClient.GetStock(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetStock call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func DecreaseStock(ctx context.Context, req *inventory.DecreaseStockRequest, callOptions ...callopt.Option) (resp *inventory.DecreaseStockResponse, err error) {
	resp, err = defaultClient.DecreaseStock(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "DecreaseStock call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func IncreaseStock(ctx context.Context, req *inventory.IncreaseStockRequest, callOptions ...callopt.Option) (resp *inventory.IncreaseStockResponse, err error) {
	resp, err = defaultClient.IncreaseStock(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "IncreaseStock call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func BatchGetStock(ctx context.Context, req *inventory.BatchGetStockRequest, callOptions ...callopt.Option) (resp *inventory.BatchGetStockResponse, err error) {
	resp, err = defaultClient.BatchGetStock(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "BatchGetStock call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
