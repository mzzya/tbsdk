package main

import (
	"fmt"

	"github.com/smgqk/tbsdk"
)

//TestStruct 测试
type TestStruct struct {
	id   int
	name string
}

const (
	appKey    = ""
	appSecret = ""
)

func main() {
	var client = tbsdk.NewClientWithProxy(appKey, appSecret, "http://gw.api.tbsandbox.com/router/rest")
	var session = ""
	var req = new(tbsdk.TaobaoItemcatsGetRequest)
	var resp = new(tbsdk.TaobaoItemcatsGetResponse)
	_, err := client.DoPostObj(req, session, resp)
	ErrorHandler(err)
	fmt.Printf("%+v\n", resp)
}
func ErrorHandler(err error) {
	if err != nil {
		panic(err)
	}
}
