# Douyin OpenAPI SDK for Go
本项目是抖音开放平台OpenAPI SDK 的Go语言实现。


## 要求
您需要确保本地安装的 go 环境版本大于等于 1.21.0.

## 安装
你可以使用 go mod 来管理你的依赖
```
go get github.com/bytedance/douyin-openapi-sdk-go
```

## 项目结构

调用入口位于 `client` 目录（与上游一致）：

```go
import openApiSdkClient "github.com/bytedance/douyin-openapi-sdk-go/client"

opt := new(credential.Config).SetClientKey("...").SetClientSecret("...")
sdkClient, err := openApiSdkClient.NewClient(opt)
```

```
client/client.go                Client、NewClient、invoke、CommonOpenAPI
client/api_<范围>.go             全部 639 个 API 方法，按方法名字母区间分成 2 个文件
client/models_<范围>.go          请求/响应结构体，按类型名字母区间分成 13 个文件
internal/transport/             共用的请求流水线：重试、错误判定、响应解码
```

所有方法都是 `*Client` 的方法，因此必须位于同一个包、同一个目录中，无法再按业务
域拆到子目录里。文件按名称的字母区间切分，每个文件约 1 万行；同一个类型的结构体
定义和它的 setter 始终在同一个文件内。

所有 API 方法都通过 `internal/transport` 发起请求。单个方法只声明与自己相关的
信息（HTTP 方法、路径、Host、Content-Type、Body 编码、access-token 头以及
query/body 参数），重试与错误处理逻辑只存在一份：

```go
func (client *Client) JsGetticket(request *JsGetticketRequest) (_result *JsGetticketResponse, _err error) {
	return invoke[JsGetticketResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/js/getticket/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}
```

## 发行说明
每个版本的详细更改记录在[发行说明](./ChangeLog.txt)中。


## 许可证
[Apache-2.0](http://www.apache.org/licenses/LICENSE-2.0)

Copyright 2024 ByteDance Ltd. and/or its affiliates.
