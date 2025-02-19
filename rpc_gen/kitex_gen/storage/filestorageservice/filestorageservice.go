// Code generated by Kitex v0.9.1. DO NOT EDIT.

package filestorageservice

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	storage "github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/storage"
	proto "google.golang.org/protobuf/proto"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"UploadFile": kitex.NewMethodInfo(
		uploadFileHandler,
		newUploadFileArgs,
		newUploadFileResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"NewMultiUpload": kitex.NewMethodInfo(
		newMultiUploadHandler,
		newNewMultiUploadArgs,
		newNewMultiUploadResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"AbortMultiUpload": kitex.NewMethodInfo(
		abortMultiUploadHandler,
		newAbortMultiUploadArgs,
		newAbortMultiUploadResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"GetMultiUploadProgress": kitex.NewMethodInfo(
		getMultiUploadProgressHandler,
		newGetMultiUploadProgressArgs,
		newGetMultiUploadProgressResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"MergeFileChunks": kitex.NewMethodInfo(
		mergeFileChunksHandler,
		newMergeFileChunksArgs,
		newMergeFileChunksResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"GetDownloadUrl": kitex.NewMethodInfo(
		getDownloadUrlHandler,
		newGetDownloadUrlArgs,
		newGetDownloadUrlResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
}

var (
	fileStorageServiceServiceInfo                = NewServiceInfo()
	fileStorageServiceServiceInfoForClient       = NewServiceInfoForClient()
	fileStorageServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return fileStorageServiceServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return fileStorageServiceServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return fileStorageServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "FileStorageService"
	handlerType := (*storage.FileStorageService)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "storage",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.9.1",
		Extra:           extra,
	}
	return svcInfo
}

func uploadFileHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(storage.UploadFileRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(storage.FileStorageService).UploadFile(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *UploadFileArgs:
		success, err := handler.(storage.FileStorageService).UploadFile(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*UploadFileResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newUploadFileArgs() interface{} {
	return &UploadFileArgs{}
}

func newUploadFileResult() interface{} {
	return &UploadFileResult{}
}

type UploadFileArgs struct {
	Req *storage.UploadFileRequest
}

func (p *UploadFileArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(storage.UploadFileRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *UploadFileArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *UploadFileArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *UploadFileArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *UploadFileArgs) Unmarshal(in []byte) error {
	msg := new(storage.UploadFileRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var UploadFileArgs_Req_DEFAULT *storage.UploadFileRequest

func (p *UploadFileArgs) GetReq() *storage.UploadFileRequest {
	if !p.IsSetReq() {
		return UploadFileArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UploadFileArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *UploadFileArgs) GetFirstArgument() interface{} {
	return p.Req
}

type UploadFileResult struct {
	Success *storage.UploadFileResponse
}

var UploadFileResult_Success_DEFAULT *storage.UploadFileResponse

func (p *UploadFileResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(storage.UploadFileResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *UploadFileResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *UploadFileResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *UploadFileResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *UploadFileResult) Unmarshal(in []byte) error {
	msg := new(storage.UploadFileResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UploadFileResult) GetSuccess() *storage.UploadFileResponse {
	if !p.IsSetSuccess() {
		return UploadFileResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UploadFileResult) SetSuccess(x interface{}) {
	p.Success = x.(*storage.UploadFileResponse)
}

func (p *UploadFileResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UploadFileResult) GetResult() interface{} {
	return p.Success
}

func newMultiUploadHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(storage.NewMultiUploadRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(storage.FileStorageService).NewMultiUpload(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *NewMultiUploadArgs:
		success, err := handler.(storage.FileStorageService).NewMultiUpload(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*NewMultiUploadResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newNewMultiUploadArgs() interface{} {
	return &NewMultiUploadArgs{}
}

func newNewMultiUploadResult() interface{} {
	return &NewMultiUploadResult{}
}

type NewMultiUploadArgs struct {
	Req *storage.NewMultiUploadRequest
}

func (p *NewMultiUploadArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(storage.NewMultiUploadRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *NewMultiUploadArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *NewMultiUploadArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *NewMultiUploadArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *NewMultiUploadArgs) Unmarshal(in []byte) error {
	msg := new(storage.NewMultiUploadRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var NewMultiUploadArgs_Req_DEFAULT *storage.NewMultiUploadRequest

func (p *NewMultiUploadArgs) GetReq() *storage.NewMultiUploadRequest {
	if !p.IsSetReq() {
		return NewMultiUploadArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *NewMultiUploadArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *NewMultiUploadArgs) GetFirstArgument() interface{} {
	return p.Req
}

type NewMultiUploadResult struct {
	Success *storage.NewMultiUploadResponse
}

var NewMultiUploadResult_Success_DEFAULT *storage.NewMultiUploadResponse

func (p *NewMultiUploadResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(storage.NewMultiUploadResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *NewMultiUploadResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *NewMultiUploadResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *NewMultiUploadResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *NewMultiUploadResult) Unmarshal(in []byte) error {
	msg := new(storage.NewMultiUploadResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *NewMultiUploadResult) GetSuccess() *storage.NewMultiUploadResponse {
	if !p.IsSetSuccess() {
		return NewMultiUploadResult_Success_DEFAULT
	}
	return p.Success
}

func (p *NewMultiUploadResult) SetSuccess(x interface{}) {
	p.Success = x.(*storage.NewMultiUploadResponse)
}

func (p *NewMultiUploadResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *NewMultiUploadResult) GetResult() interface{} {
	return p.Success
}

func abortMultiUploadHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(storage.AbortMultiUploadRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(storage.FileStorageService).AbortMultiUpload(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *AbortMultiUploadArgs:
		success, err := handler.(storage.FileStorageService).AbortMultiUpload(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*AbortMultiUploadResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newAbortMultiUploadArgs() interface{} {
	return &AbortMultiUploadArgs{}
}

func newAbortMultiUploadResult() interface{} {
	return &AbortMultiUploadResult{}
}

type AbortMultiUploadArgs struct {
	Req *storage.AbortMultiUploadRequest
}

func (p *AbortMultiUploadArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(storage.AbortMultiUploadRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *AbortMultiUploadArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *AbortMultiUploadArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *AbortMultiUploadArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *AbortMultiUploadArgs) Unmarshal(in []byte) error {
	msg := new(storage.AbortMultiUploadRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var AbortMultiUploadArgs_Req_DEFAULT *storage.AbortMultiUploadRequest

func (p *AbortMultiUploadArgs) GetReq() *storage.AbortMultiUploadRequest {
	if !p.IsSetReq() {
		return AbortMultiUploadArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AbortMultiUploadArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *AbortMultiUploadArgs) GetFirstArgument() interface{} {
	return p.Req
}

type AbortMultiUploadResult struct {
	Success *storage.AbortMultiUploadResponse
}

var AbortMultiUploadResult_Success_DEFAULT *storage.AbortMultiUploadResponse

func (p *AbortMultiUploadResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(storage.AbortMultiUploadResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *AbortMultiUploadResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *AbortMultiUploadResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *AbortMultiUploadResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *AbortMultiUploadResult) Unmarshal(in []byte) error {
	msg := new(storage.AbortMultiUploadResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AbortMultiUploadResult) GetSuccess() *storage.AbortMultiUploadResponse {
	if !p.IsSetSuccess() {
		return AbortMultiUploadResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AbortMultiUploadResult) SetSuccess(x interface{}) {
	p.Success = x.(*storage.AbortMultiUploadResponse)
}

func (p *AbortMultiUploadResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AbortMultiUploadResult) GetResult() interface{} {
	return p.Success
}

func getMultiUploadProgressHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(storage.GetMultiUploadProgressRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(storage.FileStorageService).GetMultiUploadProgress(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *GetMultiUploadProgressArgs:
		success, err := handler.(storage.FileStorageService).GetMultiUploadProgress(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetMultiUploadProgressResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newGetMultiUploadProgressArgs() interface{} {
	return &GetMultiUploadProgressArgs{}
}

func newGetMultiUploadProgressResult() interface{} {
	return &GetMultiUploadProgressResult{}
}

type GetMultiUploadProgressArgs struct {
	Req *storage.GetMultiUploadProgressRequest
}

func (p *GetMultiUploadProgressArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(storage.GetMultiUploadProgressRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetMultiUploadProgressArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetMultiUploadProgressArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetMultiUploadProgressArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *GetMultiUploadProgressArgs) Unmarshal(in []byte) error {
	msg := new(storage.GetMultiUploadProgressRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetMultiUploadProgressArgs_Req_DEFAULT *storage.GetMultiUploadProgressRequest

func (p *GetMultiUploadProgressArgs) GetReq() *storage.GetMultiUploadProgressRequest {
	if !p.IsSetReq() {
		return GetMultiUploadProgressArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetMultiUploadProgressArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetMultiUploadProgressArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetMultiUploadProgressResult struct {
	Success *storage.GetMultiUploadProgressResponse
}

var GetMultiUploadProgressResult_Success_DEFAULT *storage.GetMultiUploadProgressResponse

func (p *GetMultiUploadProgressResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(storage.GetMultiUploadProgressResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetMultiUploadProgressResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetMultiUploadProgressResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetMultiUploadProgressResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *GetMultiUploadProgressResult) Unmarshal(in []byte) error {
	msg := new(storage.GetMultiUploadProgressResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetMultiUploadProgressResult) GetSuccess() *storage.GetMultiUploadProgressResponse {
	if !p.IsSetSuccess() {
		return GetMultiUploadProgressResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetMultiUploadProgressResult) SetSuccess(x interface{}) {
	p.Success = x.(*storage.GetMultiUploadProgressResponse)
}

func (p *GetMultiUploadProgressResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetMultiUploadProgressResult) GetResult() interface{} {
	return p.Success
}

func mergeFileChunksHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(storage.MergeFileChunksRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(storage.FileStorageService).MergeFileChunks(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *MergeFileChunksArgs:
		success, err := handler.(storage.FileStorageService).MergeFileChunks(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*MergeFileChunksResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newMergeFileChunksArgs() interface{} {
	return &MergeFileChunksArgs{}
}

func newMergeFileChunksResult() interface{} {
	return &MergeFileChunksResult{}
}

type MergeFileChunksArgs struct {
	Req *storage.MergeFileChunksRequest
}

func (p *MergeFileChunksArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(storage.MergeFileChunksRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *MergeFileChunksArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *MergeFileChunksArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *MergeFileChunksArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *MergeFileChunksArgs) Unmarshal(in []byte) error {
	msg := new(storage.MergeFileChunksRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var MergeFileChunksArgs_Req_DEFAULT *storage.MergeFileChunksRequest

func (p *MergeFileChunksArgs) GetReq() *storage.MergeFileChunksRequest {
	if !p.IsSetReq() {
		return MergeFileChunksArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MergeFileChunksArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *MergeFileChunksArgs) GetFirstArgument() interface{} {
	return p.Req
}

type MergeFileChunksResult struct {
	Success *storage.MergeFileChunksResponse
}

var MergeFileChunksResult_Success_DEFAULT *storage.MergeFileChunksResponse

func (p *MergeFileChunksResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(storage.MergeFileChunksResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *MergeFileChunksResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *MergeFileChunksResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *MergeFileChunksResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *MergeFileChunksResult) Unmarshal(in []byte) error {
	msg := new(storage.MergeFileChunksResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MergeFileChunksResult) GetSuccess() *storage.MergeFileChunksResponse {
	if !p.IsSetSuccess() {
		return MergeFileChunksResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MergeFileChunksResult) SetSuccess(x interface{}) {
	p.Success = x.(*storage.MergeFileChunksResponse)
}

func (p *MergeFileChunksResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MergeFileChunksResult) GetResult() interface{} {
	return p.Success
}

func getDownloadUrlHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(storage.GetDownloadUrlRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(storage.FileStorageService).GetDownloadUrl(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *GetDownloadUrlArgs:
		success, err := handler.(storage.FileStorageService).GetDownloadUrl(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetDownloadUrlResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newGetDownloadUrlArgs() interface{} {
	return &GetDownloadUrlArgs{}
}

func newGetDownloadUrlResult() interface{} {
	return &GetDownloadUrlResult{}
}

type GetDownloadUrlArgs struct {
	Req *storage.GetDownloadUrlRequest
}

func (p *GetDownloadUrlArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(storage.GetDownloadUrlRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetDownloadUrlArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetDownloadUrlArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetDownloadUrlArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *GetDownloadUrlArgs) Unmarshal(in []byte) error {
	msg := new(storage.GetDownloadUrlRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetDownloadUrlArgs_Req_DEFAULT *storage.GetDownloadUrlRequest

func (p *GetDownloadUrlArgs) GetReq() *storage.GetDownloadUrlRequest {
	if !p.IsSetReq() {
		return GetDownloadUrlArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetDownloadUrlArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetDownloadUrlArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetDownloadUrlResult struct {
	Success *storage.GetDownloadUrlResponse
}

var GetDownloadUrlResult_Success_DEFAULT *storage.GetDownloadUrlResponse

func (p *GetDownloadUrlResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(storage.GetDownloadUrlResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetDownloadUrlResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetDownloadUrlResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetDownloadUrlResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *GetDownloadUrlResult) Unmarshal(in []byte) error {
	msg := new(storage.GetDownloadUrlResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetDownloadUrlResult) GetSuccess() *storage.GetDownloadUrlResponse {
	if !p.IsSetSuccess() {
		return GetDownloadUrlResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetDownloadUrlResult) SetSuccess(x interface{}) {
	p.Success = x.(*storage.GetDownloadUrlResponse)
}

func (p *GetDownloadUrlResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetDownloadUrlResult) GetResult() interface{} {
	return p.Success
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) UploadFile(ctx context.Context, Req *storage.UploadFileRequest) (r *storage.UploadFileResponse, err error) {
	var _args UploadFileArgs
	_args.Req = Req
	var _result UploadFileResult
	if err = p.c.Call(ctx, "UploadFile", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) NewMultiUpload(ctx context.Context, Req *storage.NewMultiUploadRequest) (r *storage.NewMultiUploadResponse, err error) {
	var _args NewMultiUploadArgs
	_args.Req = Req
	var _result NewMultiUploadResult
	if err = p.c.Call(ctx, "NewMultiUpload", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AbortMultiUpload(ctx context.Context, Req *storage.AbortMultiUploadRequest) (r *storage.AbortMultiUploadResponse, err error) {
	var _args AbortMultiUploadArgs
	_args.Req = Req
	var _result AbortMultiUploadResult
	if err = p.c.Call(ctx, "AbortMultiUpload", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetMultiUploadProgress(ctx context.Context, Req *storage.GetMultiUploadProgressRequest) (r *storage.GetMultiUploadProgressResponse, err error) {
	var _args GetMultiUploadProgressArgs
	_args.Req = Req
	var _result GetMultiUploadProgressResult
	if err = p.c.Call(ctx, "GetMultiUploadProgress", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MergeFileChunks(ctx context.Context, Req *storage.MergeFileChunksRequest) (r *storage.MergeFileChunksResponse, err error) {
	var _args MergeFileChunksArgs
	_args.Req = Req
	var _result MergeFileChunksResult
	if err = p.c.Call(ctx, "MergeFileChunks", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetDownloadUrl(ctx context.Context, Req *storage.GetDownloadUrlRequest) (r *storage.GetDownloadUrlResponse, err error) {
	var _args GetDownloadUrlArgs
	_args.Req = Req
	var _result GetDownloadUrlResult
	if err = p.c.Call(ctx, "GetDownloadUrl", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
