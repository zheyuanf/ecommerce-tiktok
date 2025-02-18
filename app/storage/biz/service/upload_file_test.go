package service

import (
	"context"
	"os"
	"testing"

	"github.com/zheyuanf/ecommerce-tiktok/app/storage/infra/store"
	storage "github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/storage"
)

func TestUploadFile_Run(t *testing.T) {
	os.Chdir("../..")
	store.Init()
	ctx := context.Background()
	s := NewUploadFileService(ctx)
	// init req and assert value
	req := &storage.UploadFileRequest{
		FileName: "test-123.txt",
		FileType: "txt",
		Checksum: "",
		FileData: []byte("hello world"),
	}

	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)
}
