package store

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/zheyuanf/ecommerce-tiktok/app/storage/conf"
)

var (
	MinIOClient *minio.Client
	MinIOCore   *minio.Core
	BucketName  string
)

type MinIOClt struct {
	Client     *minio.Client
	BucketName string
}

func Init() {
	BucketName = conf.GetConf().MinIO.BucketName
	var err error
	MinIOCore, err = minio.NewCore(conf.GetConf().MinIO.Endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(
			conf.GetConf().MinIO.AccessKeyID,
			conf.GetConf().MinIO.SecretAccessKey,
			"",
		),
		Secure: conf.GetConf().MinIO.UseSSL,
	})
	if err != nil {
		panic(err)
	}
	MinIOClient, err = minio.New(conf.GetConf().MinIO.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(conf.GetConf().MinIO.AccessKeyID, conf.GetConf().MinIO.SecretAccessKey, ""),
		Secure: conf.GetConf().MinIO.UseSSL,
	})
	if err != nil {
		panic(err)
	}
}

func GetFileChecksum(data []byte) string {
	return ""
}
