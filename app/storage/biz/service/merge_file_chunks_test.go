package service

import (
	"context"
	"testing"
	storage "github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/storage"
)

func TestMergeFileChunks_Run(t *testing.T) {
	ctx := context.Background()
	s := NewMergeFileChunksService(ctx)
	// init req and assert value

	req := &storage.MergeFileChunksRequest{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
