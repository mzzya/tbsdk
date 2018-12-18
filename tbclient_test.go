package tbsdk_test

import (
	"testing"

	"github.com/smgqk/tbsdk"
)

func TestSignString(t *testing.T) {
	var param = "app_key12345678fieldsnum_iid,title,nick,price,numformatjsonmethodtaobao.item.seller.getnum_iid11223344sessiontestsign_methodmd5timestamp2016-01-01 12:00:00v2.0"
	var signResult = "66987CB115214E59E6EC978214934FB8"
	var signStr = tbsdk.SignString(param, "helloworld", "md5")
	if signStr != signResult {
		t.FailNow()
	}
}
