package main

import (
	"fmt"
	"time"

	"github.com/smgqk/tbsdk"
)

//TestStruct 测试
type TestStruct struct {
	id   int
	name string
	date time.Time
}

const (
	appKey    = ""
	appSecret = ""
)

func main() {
	var client = tbsdk.NewClientWithAddr(appKey, appSecret, tbsdk.APIAddrTest)
	var session = "test"
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
