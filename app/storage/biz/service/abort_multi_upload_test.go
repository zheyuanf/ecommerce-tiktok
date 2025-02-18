package service

import (
	"context"
	"os"
	"testing"

	"github.com/zheyuanf/ecommerce-tiktok/app/storage/infra/store"
	storage "github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/storage"
)

func TestAbortMultiUpload_Run(t *testing.T) {
	os.Chdir("../..")
	store.Init()
	ctx := context.Background()
	sAbort := NewAbortMultiUploadService(ctx)
	sNew := NewNewMultiUploadService(ctx)
	// init req and assert value
	fileName := "multi_upload.txt"
	newReq := &storage.NewMultiUploadRequest{
		FileName: fileName,
		FileType: "text/plain",
	}
	newResp, err := sNew.Run(newReq)
	uploadId := newResp.UploadId
	t.Logf("uploadId: %v", uploadId)

	req := &storage.AbortMultiUploadRequest{
		FileName: fileName,
		UploadId: uploadId,
	}
	resp, err := sAbort.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
