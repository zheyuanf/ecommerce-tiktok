package service

import (
	"bytes"
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/minio/minio-go/v7"

	"github.com/zheyuanf/ecommerce-tiktok/app/storage/infra/store"
	storage "github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/storage"
)

func TestNewMultiUpload_Run(t *testing.T) {
	os.Chdir("../..")
	store.Init()

	fileName := "multipart_upload.txt"
	fileSize := int32(1024 * 1024 * 1024) // 1GB
	chunkSize := int32(10 * 1024 * 1024)  // 10MB
	// 模拟大文件 1GB
	data := GetBigFile(fileName, fileSize)
	ctx := context.Background()
	newS := NewNewMultiUploadService(ctx)

	var chunks int32 = (fileSize + chunkSize - 1) / chunkSize
	newReq := &storage.NewMultiUploadRequest{
		FileName:    fileName,
		FileType:    "text/plain",
		TotalChunks: chunks,
	}
	newResp, err := newS.Run(newReq)
	if err != nil {
		t.Errorf("Error creating multipart upload: %v", err)
		return
	}
	uploadId := newResp.UploadId
	t.Logf("upload id: %v", uploadId)
	parts := []minio.CompletePart{}
	t.Logf("len of presigned urls: %v", len(newResp.PresignedUrls))
	for i, url := range newResp.PresignedUrls {
		chunk := data[i*int(chunkSize) : min((i+1)*int(chunkSize), int(fileSize))]
		httpReq, err := http.NewRequest("PUT", url.Url, bytes.NewReader(chunk))
		if err != nil {
			t.Errorf("Error creating HTTP request: %v", err)
			return
		}
		// 设置请求头（例如 Content-Type）
		httpReq.Header.Set("Content-Type", "application/octet-stream")

		// 发起上传请求
		client := &http.Client{}
		httpResp, err := client.Do(httpReq)
		if err != nil {
			t.Errorf("Error sending HTTP request: %v", err)
			return
		}
		defer httpResp.Body.Close()

		if httpResp.StatusCode != 200 {
			t.Errorf("Unexpected HTTP status code: %d", httpResp.StatusCode)
			return
		}

		partInfo := minio.CompletePart{
			PartNumber: int(url.ChunkNumber), // 当前分片的编号
			ETag:       httpResp.Header.Get("ETag"),
		}
		parts = append(parts, partInfo)
	}
	getS := NewGetMultiUploadProgressService(ctx)
	getResp, err := getS.Run(&storage.GetMultiUploadProgressRequest{
		UploadId: uploadId,
		FileName: fileName,
	})
	t.Logf("get progress resp: %v", getResp)

	mergeS := NewMergeFileChunksService(ctx)
	mergeReq := &storage.MergeFileChunksRequest{
		UploadId:    uploadId,
		FileName:    fileName,
		FileType:    "text/plain",
		TotalChunks: chunks,
	}
	mergeResp, err := mergeS.Run(mergeReq)
	if err != nil {
		t.Errorf("Error merging file chunks: %v", err)
		return
	}
	t.Logf("info: %v", mergeResp)
}

func GetBigFile(filename string, fileSize int32) []byte {
	// Mock fileSize bytes
	bs := make([]byte, fileSize)
	for i := range bs {
		bs[i] = byte(i%26 + 'A')
	}
	return bs
}
