package tbsdk

// AlipayUserAccountGetRequest 查询支付宝账户余额
type AlipayUserAccountGetRequest struct {
    
}

func (req *AlipayUserAccountGetRequest) GetAPIName() string {
	return "alipay.user.account.get"
}

// AlipayUserAccountGetResponse 查询支付宝账户余额
type AlipayUserAccountGetResponse struct {
    
    /* alipay_account Object支付宝用户账户信息 */
    alipay_account AlipayAccount `json:"alipay_account";xml:"alipay_account"`
    
}
