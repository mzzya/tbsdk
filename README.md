# 淘宝sdk
基于淘宝官方ApiMetadata.xml接口文件生成

#### 使用方法
淘宝接口众多，使用前请测试，建议将生成文件复制至现有项目中使用。
默认生成 ./tbapi.go 为接口request与response结构
./tbmodel.go 为接口共用数据结构

```
const (
	appKey    = ""
	appSecret = ""
)

func main() {
	var client = tbsdk.NewClientWithProxy(appKey, appSecret, tbsdk.APIAddrTest)
	var session = "test"
	var req = new(tbsdk.TaobaoItemcatsGetRequest)
	var resp = new(tbsdk.TaobaoItemcatsGetResponse)
	_, err := client.DoPostObj(req, session, resp)
	ErrorHandler(err)
	fmt.Printf("%+v\n", resp)
}
```