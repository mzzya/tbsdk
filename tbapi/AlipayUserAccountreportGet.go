package tbsdk

// AlipayUserAccountreportGetRequest 获取支付宝余额明细记录，不包含用户通过银行卡快捷支付进行交易的交易明细
type AlipayUserAccountreportGetRequest struct {
    
    /* end_time required对账单结束时间，其中end_time - start_time <= 1天，对于对账记录比较多的情况，强烈建议按天查询，否则会出现超时的情况。 */
    end_time Date `json:"end_time";xml:"end_time"`
    
    /* fields required需要返回的字段列表。create_time:创建时间,type：账务类型,business_type:子业务类型,balance:当时支付宝账户余额,in_amount:收入金额,out_amount:支出金额,alipay_order_no:支付宝订单号,merchant_order_no:商户订单号,self_user_id:自己的支付宝ID,opt_user_id:对方的支付宝ID,memo:账号备注 */
    fields string `json:"fields";xml:"fields"`
    
    /* page_no optional要获取的对账单页码 */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size optional每次查询获取对账记录数量 */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* start_time required对账单开始时间。最近一个月内的日期。 */
    start_time Date `json:"start_time";xml:"start_time"`
    
    /* _type optional账务类型。多个类型是，用逗号分隔，不传则查询所有类型的。PAYMENT:在线支付，TRANSFER:转账，DEPOSIT:充值，WITHDRAW:提现，CHARGE:收费，PREAUTH:预授权，OTHER：其它。 */
    _type string `json:"type";xml:"type"`
    
}

func (req *AlipayUserAccountreportGetRequest) GetAPIName() string {
	return "alipay.user.accountreport.get"
}

// AlipayUserAccountreportGetResponse 获取支付宝余额明细记录，不包含用户通过银行卡快捷支付进行交易的交易明细
type AlipayUserAccountreportGetResponse struct {
    
    /* alipay_records Object Array对账记录列表 */
    alipay_records AlipayRecord `json:"alipay_records";xml:"alipay_records"`
    
    /* total_pages Basic总页数 */
    total_pages int64 `json:"total_pages";xml:"total_pages"`
    
    /* total_results Basic总记录数 */
    total_results int64 `json:"total_results";xml:"total_results"`
    
}
