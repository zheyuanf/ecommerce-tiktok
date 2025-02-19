package service

import (
	"context"
	"testing"

	storage "github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/storage"
)

func TestGetMultiUploadProgress_Run(t *testing.T) {
	ctx := context.Background()
	s := NewGetMultiUploadProgressService(ctx)
	// init req and assert value

	req := &storage.GetMultiUploadProgressRequest{
		FileName: "test.jpg",
		UploadId: "1234567890",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
