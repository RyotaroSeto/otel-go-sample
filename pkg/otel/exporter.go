package otel

import (
	"context"
	"io"

	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/trace"
)

// Exporter
// テレメトリをOpenTelemetry Collectorに送信
// OTLPエンドポイントに送信するには、エンドポイントに送信するエクスポータを設定する必要がある。
// バイナリprotobufペイロードを持つHTTPを使用するOTLPメトリクス・エクスポーターの実装が含まれている。

// トレースのエクスポーター(http)
func newTracesHttpExporter(ctx context.Context) (trace.SpanExporter, error) {
	// endpointOpt := otlptracehttp.WithEndpoint(otlpEndpoint) // エンドポイントがデフォルト(localhost:4318)の場合は不要
	return otlptracehttp.New(ctx, otlptracehttp.WithInsecure())
}

// トレースのエクスポーター(grpc)
func newTracesGrpcExporter(ctx context.Context) (trace.SpanExporter, error) {
	// endpointOpt := otlptracegrpc.WithEndpoint(otlpEndpoint) // エンドポイントがデフォルト(localhost:4317)の場合は不要
	return otlptracegrpc.New(ctx, otlptracegrpc.WithInsecure())
}

// トレースのエクスポーター(jaeger)
func newTracesJaegerExporter() (trace.SpanExporter, error) {
	return jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint("http://jaeger:14268/api/traces")))
}

// トレースのエクスポーター(jaeger)
func newTracesWriterExporter(w io.Writer) (trace.SpanExporter, error) {
	return stdouttrace.New(
		stdouttrace.WithWriter(w),
		stdouttrace.WithPrettyPrint(),
	)
}

// メトリクスのエクスポーター(http)
func newMetricHttpExporter(ctx context.Context) (metric.Exporter, error) {
	return otlpmetrichttp.New(ctx)
}

// メトリクスのエクスポーター(grpc)
func newMetricGrpcExporter(ctx context.Context) (metric.Exporter, error) {
	return otlpmetricgrpc.New(ctx)
}

// Prometheus メトリクスのエクスポーター
// プロメテウス・エクスポーターの使い方については以下を参照
// https://github.com/open-telemetry/opentelemetry-go/tree/main/example/prometheus
func newMetricPrometheusExporter(ctx context.Context) (metric.Reader, error) {
	return prometheus.New()
}
