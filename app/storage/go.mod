module github.com/zheyuanf/ecommerce-tiktok/app/storage

go 1.23.3

replace github.com/apache/thrift => github.com/apache/thrift v0.13.0

require (
	github.com/cloudwego/kitex v0.12.2
	github.com/kr/pretty v0.3.1
	github.com/minio/minio-go/v7 v7.0.86
	github.com/redis/go-redis/v9 v9.7.0
	github.com/zheyuanf/ecommerce-tiktok/rpc_gen v0.0.0-20250219070401-19a48eb11a71
	gopkg.in/validator.v2 v2.0.1
	gopkg.in/yaml.v2 v2.4.0
	gorm.io/driver/mysql v1.5.7
	gorm.io/gorm v1.25.12
)

require (
	github.com/bytedance/gopkg v0.1.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/cloudwego/fastpb v0.0.5 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/go-ini/ini v1.67.0 // indirect
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/goccy/go-json v0.10.5 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/klauspost/compress v1.17.11 // indirect
	github.com/klauspost/cpuid/v2 v2.2.9 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/minio/crc64nvme v1.0.0 // indirect
	github.com/minio/md5-simd v1.1.2 // indirect
	github.com/rogpeppe/go-internal v1.9.0 // indirect
	github.com/rs/xid v1.6.0 // indirect
	golang.org/x/crypto v0.33.0 // indirect
	golang.org/x/net v0.35.0 // indirect
	golang.org/x/sys v0.30.0 // indirect
	golang.org/x/text v0.22.0 // indirect
	google.golang.org/protobuf v1.33.0 // indirect
)
