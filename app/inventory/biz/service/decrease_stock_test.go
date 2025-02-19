package service

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/zheyuanf/ecommerce-tiktok/app/inventory/biz/dal/mysql"
	"github.com/zheyuanf/ecommerce-tiktok/app/inventory/biz/model"
	inventory "github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/inventory"
)

func MockDataDecrease(ctx context.Context) {
	model.CreateInventory(ctx, mysql.DB, &model.Inventory{
		ProductID:   1,
		Stock:       1000000,
		LockedStock: 0,
	})
}

func TestDecreaseStock_Run(t *testing.T) {
	ctx := context.Background()
	mysql.MockInit()
	MockDataDecrease(ctx)
	s := NewDecreaseStockService(ctx)
	// init req and assert value

	req := &inventory.DecreaseStockRequest{
		ProductId: 1,
		Quantity:  1,
	}
	wq := sync.WaitGroup{}
	var success int32 = 0
	var fail int32 = 0
	// 测试并发扣减库存
	for i := 0; i < 100; i++ {
		wq.Add(1)
		go func() {
			for j := 0; j < 100; j++ {
				_, err := s.Run(req)
				if err != nil {
					atomic.AddInt32(&fail, 1)
					continue
				}
				atomic.AddInt32(&success, 1)
			}
			wq.Done()
		}()
	}
	wq.Wait()
	ivt, err := model.QueryInventory(ctx, mysql.DB, 1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("success: %d, fail: %d\n", success, fail)
	fmt.Printf("stock: %d, lockedStock: %d\n", ivt.Stock, ivt.LockedStock)
}
