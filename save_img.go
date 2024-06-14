package wechatopen

import (
	"context"
	"go.dtapp.net/gorequest"
	"log"
	"os"
)

type SaveImgResponse struct {
	Path string
	Name string
}

func (c *Client) SaveImg(ctx context.Context, resp gorequest.Response, dir, saveName string) SaveImgResponse {

	// OpenTelemetry链路追踪
	ctx = c.TraceStartSpan(ctx, "SaveImg")
	defer c.TraceEndSpan()

	// 返回是二进制图片，或者json错误
	if resp.ResponseHeader.Get("Content-Type") == "image/jpeg" || resp.ResponseHeader.Get("Content-Type") == "image/png" {
		// 保存在output目录
		outputFileName := saveName

		if resp.ResponseHeader.Get("Content-Type") == "image/jpeg" {
			outputFileName = outputFileName + ".jpg"
		} else {
			outputFileName = outputFileName + ".png"
		}
	here:
		log.Println(dir + outputFileName)
		f, err := os.OpenFile(dir+outputFileName, os.O_CREATE|os.O_RDWR, 0666)
		log.Println(err)
		if err != nil {
			os.Mkdir(dir, 0755)
			goto here
		}
		f.Write(resp.ResponseBody)
		f.Close()
		return SaveImgResponse{
			Path: dir + outputFileName,
			Name: outputFileName,
		}
	}
	return SaveImgResponse{}
}
