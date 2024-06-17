package wechatopen

import (
	"context"
	"go.dtapp.net/gorequest"
	"net/http"
)

type DataCubeGetWeAnAlySisAppidMonthlyRetainInfoResponse struct {
	RefDate    string `json:"ref_date"` // 日期
	VisitUvNew []struct {
		Key   int64 `json:"key"`
		Value int64 `json:"value"`
	} `json:"visit_uv_new"` // 新增用户留存
	VisitUv []struct {
		Key   int64 `json:"key"`
		Value int64 `json:"value"`
	} `json:"visit_uv"` // 活跃用户留存
}

type DataCubeGetWeAnAlySisAppidMonthlyRetainInfoResult struct {
	Result DataCubeGetWeAnAlySisAppidMonthlyRetainInfoResponse // 结果
	Body   []byte                                              // 内容
	Http   gorequest.Response                                  // 请求
}

func newDataCubeGetWeAnAlySisAppidMonthlyRetainInfoResult(result DataCubeGetWeAnAlySisAppidMonthlyRetainInfoResponse, body []byte, http gorequest.Response) *DataCubeGetWeAnAlySisAppidMonthlyRetainInfoResult {
	return &DataCubeGetWeAnAlySisAppidMonthlyRetainInfoResult{Result: result, Body: body, Http: http}
}

// DataCubeGetWeAnAlySisAppidMonthlyRetainInfo 获取用户访问小程序月留存
// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/data-analysis/visit-retain/getMonthlyRetain.html
func (c *Client) DataCubeGetWeAnAlySisAppidMonthlyRetainInfo(ctx context.Context, authorizerAccessToken, beginDate, endDate string, notMustParams ...gorequest.Params) (*DataCubeGetWeAnAlySisAppidMonthlyRetainInfoResult, error) {

	// OpenTelemetry链路追踪
	ctx, span := TraceStartSpan(ctx, "datacube/getweanalysisappidmonthlyretaininfo")
	defer span.End()

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("begin_date", beginDate)
	params.Set("end_date", endDate)

	// 请求
	var response DataCubeGetWeAnAlySisAppidMonthlyRetainInfoResponse
	request, err := c.request(ctx, span, "datacube/getweanalysisappidmonthlyretaininfo?access_token="+authorizerAccessToken, params, http.MethodPost, &response)
	return newDataCubeGetWeAnAlySisAppidMonthlyRetainInfoResult(response, request.ResponseBody, request), err
}
