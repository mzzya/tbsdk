package tbsdk_test

import (
	"log"
	"testing"

	"github.com/smgqk/tbsdk"
)

func TestSignStringMap(t *testing.T) {
	param := make(map[string]string, 10)
	param["method"] = "taobao.item.seller.get"
	param["app_key"] = "12345678"
	param["session"] = "test"
	param["timestamp"] = "2016-01-01 12:00:00"
	param["format"] = "json"
	param["v"] = "2.0"
	param["sign_method"] = "md5"
	param["fields"] = "num_iid,title,nick,price,num"
	param["num_iid"] = "11223344"
	var signResult = "66987CB115214E59E6EC978214934FB8"
	var signStr = tbsdk.SignStringMap(param, "helloworld", "md5")
	log.Printf("%s---true\n", signResult)
	log.Printf("%s---false\n", signStr)
	if signStr != signResult {
		t.FailNow()
	}
}
func TestSignString(t *testing.T) {
	var param = "app_key12345678fieldsnum_iid,title,nick,price,numformatjsonmethodtaobao.item.seller.getnum_iid11223344sessiontestsign_methodmd5timestamp2016-01-01 12:00:00v2.0"
	var signResult = "66987CB115214E59E6EC978214934FB8"
	var signStr = tbsdk.SignString(param, "helloworld", "md5")
	log.Printf("%s---true\n", signResult)
	log.Printf("%s---false\n", signStr)
	if signStr != signResult {
		t.FailNow()
	}
}
