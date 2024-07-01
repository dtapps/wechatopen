package wechatopen

import (
	"context"
	"go.dtapp.net/gojson"
	"go.dtapp.net/gorequest"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

func (c *Client) request(ctx context.Context, span trace.Span, url string, param gorequest.Params, method string, response any) (resp gorequest.Response, err error) {

	// 请求地址
	uri := apiUrl + url

	// 设置请求地址
	c.httpClient.SetUri(uri)

	// 设置方式
	c.httpClient.SetMethod(method)

	// 设置格式
	c.httpClient.SetContentTypeJson()

	// 设置参数
	c.httpClient.SetParams(param)

	// OpenTelemetry链路追踪
	span.SetAttributes(attribute.String("http.request.url", uri))
	span.SetAttributes(attribute.String("http.request.method", method))
	span.SetAttributes(attribute.String("http.request.body", gojson.JsonEncodeNoError(param)))

	// 发起请求
	request, err := c.httpClient.Request(ctx)
	if err != nil {
		span.RecordError(err, trace.WithStackTrace(true))
		span.SetStatus(codes.Error, err.Error())
		return gorequest.Response{}, err
	}

	// 解析响应
	if request.HeaderIsImg() == false {
		err = gojson.Unmarshal(request.ResponseBody, &response)
		if err != nil {
			span.RecordError(err, trace.WithStackTrace(true))
			span.SetStatus(codes.Error, err.Error())
		}
	}

	// OpenTelemetry链路追踪
	span.SetAttributes(attribute.String("http.response.body", gojson.JsonEncodeNoError(response)))

	return request, err
}
