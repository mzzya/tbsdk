# 淘宝sdk
=======
基于淘宝官方ApiMetadata.xml接口文件生成

#### 使用方法
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