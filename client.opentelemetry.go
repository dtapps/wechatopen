package wechatopen

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

// SetTrace 设置OpenTelemetry链路追踪
func (c *Client) SetTrace(trace bool) {

}

// TraceStartSpan 开始OpenTelemetry链路追踪状态
func TraceStartSpan(ctx context.Context, spanName string) (context.Context, trace.Span) {
	tr := otel.Tracer("go.dtapp.net/wechatopen", trace.WithInstrumentationVersion(Version))
	ctx, span := tr.Start(ctx, "wechatopen."+spanName, trace.WithSpanKind(trace.SpanKindClient))
	return ctx, span
}
