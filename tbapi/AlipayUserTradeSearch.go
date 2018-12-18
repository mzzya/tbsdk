package tbsdk

// AlipayUserTradeSearchRequest 查询支付宝账户交易记录
type AlipayUserTradeSearchRequest struct {
    
    /* alipay_order_no optional支付宝订单号，为空查询所有记录 */
    alipay_order_no string `json:"alipay_order_no";xml:"alipay_order_no"`
    
    /* end_time required结束时间。与开始时间间隔在七天之内 */
    end_time string `json:"end_time";xml:"end_time"`
    
    /* merchant_order_no optional商户订单号，为空查询所有记录 */
    merchant_order_no string `json:"merchant_order_no";xml:"merchant_order_no"`
    
    /* order_from optional订单来源，为空查询所有来源。淘宝(TAOBAO)，支付宝(ALIPAY)，其它(OTHER) */
    order_from string `json:"order_from";xml:"order_from"`
    
    /* order_status optional订单状态，为空查询所有状态订单 */
    order_status string `json:"order_status";xml:"order_status"`
    
    /* order_type optional订单类型，为空查询所有类型订单。 */
    order_type string `json:"order_type";xml:"order_type"`
    
    /* page_no required页码。取值范围:大于零的整数; 默认值1 */
    page_no string `json:"page_no";xml:"page_no"`
    
    /* page_size required每页获取条数。最大值500。 */
    page_size string `json:"page_size";xml:"page_size"`
    
    /* start_time required开始时间，时间必须是今天范围之内。格式为yyyy-MM-dd HH:mm:ss，精确到秒 */
    start_time string `json:"start_time";xml:"start_time"`
    
}

func (req *AlipayUserTradeSearchRequest) GetAPIName() string {
	return "alipay.user.trade.search"
}

// AlipayUserTradeSearchResponse 查询支付宝账户交易记录
type AlipayUserTradeSearchResponse struct {
    
    /* total_pages Basic总页数 */
    total_pages string `json:"total_pages";xml:"total_pages"`
    
    /* total_results Basic总记录数 */
    total_results string `json:"total_results";xml:"total_results"`
    
    /* trade_records Object Array交易记录列表 */
    trade_records TradeRecord `json:"trade_records";xml:"trade_records"`
    
}
