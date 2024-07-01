package wechatopen

import (
	"context"
	"go.dtapp.net/gorequest"
	"go.opentelemetry.io/otel/trace"
)

// TraceStartSpan 开始OpenTelemetry链路追踪状态
func TraceStartSpan(ctx context.Context, spanName string) (context.Context, trace.Span) {
	return gorequest.TraceNewSpan(ctx, "go.dtapp.net/wechatopen", "wechatopen.", spanName, Version, trace.SpanKindClient)
}
