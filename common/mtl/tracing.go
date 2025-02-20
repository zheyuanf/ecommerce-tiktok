package mtl

import (
	"github.com/kitex-contrib/obs-opentelemetry/provider"
)

func InitTracing(serviceName string, endPoint string) provider.OtelProvider {
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(serviceName),
		provider.WithInsecure(),
		provider.WithEnableMetrics(false),
		provider.WithExportEndpoint(endPoint),
	)
	return p
}
