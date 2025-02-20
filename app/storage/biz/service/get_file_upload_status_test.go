package service

import (
	"context"
	"testing"

	storage "github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/storage"
)

func TestGetFileUploadStatus_Run(t *testing.T) {
	ctx := context.Background()
	s := NewGetFileUploadStatusService(ctx)
	// init req and assert value

	req := &storage.GetFileUploadStatusRequest{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
