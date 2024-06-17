package wechatopen

import (
	"context"
	"go.dtapp.net/gorequest"
	"net/http"
)

type DataCubeGetWeAnAlySisAppidWeeklyVisitTrendResponse struct {
	List []struct {
		RefDate         string  `json:"ref_date"`          // 时间
		SessionCnt      int64   `json:"session_cnt"`       // 打开次数（自然周内汇总）
		VisitPv         int64   `json:"visit_pv"`          // 访问次数（自然周内汇总）
		VisitUv         int64   `json:"visit_uv"`          // 访问人数（自然周内去重）
		VisitUvNew      int64   `json:"visit_uv_new"`      // 新用户数（自然周内去重）
		StayTimeUv      float64 `json:"stay_time_uv"`      // 人均停留时长 (浮点型，单位：秒)
		StayTimeSession float64 `json:"stay_time_session"` // 次均停留时长 (浮点型，单位：秒)
		VisitDepth      float64 `json:"visit_depth"`       // 平均访问深度 (浮点型)
	} `json:"list"` // 数据列表
}

type DataCubeGetWeAnAlySisAppidWeeklyVisitTrendResult struct {
	Result DataCubeGetWeAnAlySisAppidWeeklyVisitTrendResponse // 结果
	Body   []byte                                             // 内容
	Http   gorequest.Response                                 // 请求
}

func newDataCubeGetWeAnAlySisAppidWeeklyVisitTrendResult(result DataCubeGetWeAnAlySisAppidWeeklyVisitTrendResponse, body []byte, http gorequest.Response) *DataCubeGetWeAnAlySisAppidWeeklyVisitTrendResult {
	return &DataCubeGetWeAnAlySisAppidWeeklyVisitTrendResult{Result: result, Body: body, Http: http}
}

// DataCubeGetWeAnAlySisAppidWeeklyVisitTrend 获取用户访问小程序数据周趋势
// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/data-analysis/visit-trend/getWeeklyVisitTrend.html
func (c *Client) DataCubeGetWeAnAlySisAppidWeeklyVisitTrend(ctx context.Context, authorizerAccessToken, beginDate, endDate string, notMustParams ...gorequest.Params) (*DataCubeGetWeAnAlySisAppidWeeklyVisitTrendResult, error) {

	// OpenTelemetry链路追踪
	ctx, span := TraceStartSpan(ctx, "datacube/getweanalysisappidweeklyvisittrend")
	defer span.End()

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("begin_date", beginDate)
	params.Set("end_date", endDate)

	// 请求
	var response DataCubeGetWeAnAlySisAppidWeeklyVisitTrendResponse
	request, err := c.request(ctx, span, "datacube/getweanalysisappidweeklyvisittrend?access_token="+authorizerAccessToken, params, http.MethodPost, &response)
	return newDataCubeGetWeAnAlySisAppidWeeklyVisitTrendResult(response, request.ResponseBody, request), err
}
