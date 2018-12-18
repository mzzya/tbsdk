package tbsdk

// AlipayUserAccountFreezeGetRequest 查询支付宝账户冻结类型的冻结金额。
type AlipayUserAccountFreezeGetRequest struct {
    
    /* freeze_type optional冻结类型，多个用,分隔。不传返回所有类型的冻结金额。
DEPOSIT_FREEZE,充值冻结
WITHDRAW_FREEZE,提现冻结
PAYMENT_FREEZE,支付冻结
BAIL_FREEZE,保证金冻结
CHARGE_FREEZE,收费冻结
PRE_DEPOSIT_FREEZE,预存款冻结
LOAN_FREEZE,贷款冻结
OTHER_FREEZE,其它 */
    freeze_type string `json:"freeze_type";xml:"freeze_type"`
    
}

func (req *AlipayUserAccountFreezeGetRequest) GetAPIName() string {
	return "alipay.user.account.freeze.get"
}

// AlipayUserAccountFreezeGetResponse 查询支付宝账户冻结类型的冻结金额。
type AlipayUserAccountFreezeGetResponse struct {
    
    /* freeze_items Object Array冻结金额列表 */
    freeze_items AccountFreeze `json:"freeze_items";xml:"freeze_items"`
    
    /* total_results Basic冻结总条数 */
    total_results string `json:"total_results";xml:"total_results"`
    
}
