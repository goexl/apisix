package apisix

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"strings"

	"github.com/goexl/exc"
	"github.com/goexl/gox/field"
)

func (c *Client) UploadDescriptor(
	ctx context.Context,
	filename string, id string, description string,
) (rsp *Response, err error) {
	req := c.http.R()
	if content, re := os.ReadFile(filename); nil != re {
		err = re
	} else {
		pr := new(descriptorReq)
		pr.Content = base64.StdEncoding.EncodeToString(content)
		pr.Description = description
		req.SetBody(pr)
	}
	if nil != err {
		return
	}

	rsp = new(Response)
	url := fmt.Sprintf("%s/apisix/admin/protos/%s", strings.TrimSuffix(c.endpoint, "/"), id)
	req.SetHeader("X-API-KEY", c.apiKey)
	if hr, he := req.SetContext(ctx).SetResult(rsp).Put(url); nil != he {
		err = he
	} else if hr.IsError() {
		err = exc.NewFields(
			"Apisix返回错误",
			field.New("url", hr.Request.URL),
			field.New("code", hr.StatusCode()),
			field.New("body", hr.Body()),
		)
	} else {
		c.logger.Debug("上传Protobuf文件到Apisix成功", field.New("rsp", rsp))
	}

	return
}
