package tbsdk

// TaobaoTaeBillsGetRequest tae查询账单明细
type TaobaoTaeBillsGetRequest struct {
    
    /* current_page optional页数,建议不要超过100页,越大性能越低,有可能会超时 */
    current_page int64 `json:"current_page";xml:"current_page"`
    
    /* fields required传入需要返回的字段,参见Bill结构体 */
    fields string `json:"fields";xml:"fields"`
    
    /* item_id optional科目编号 */
    item_id int64 `json:"item_id";xml:"item_id"`
    
    /* p_trade_id optional交易编号 */
    p_trade_id int64 `json:"p_trade_id";xml:"p_trade_id"`
    
    /* page_size optional每页大小,默认40条,可选范围 ：40~100 */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* query_date_type optional查询条件中的时间类型:1-交易订单完成时间biz_time 2-支付宝扣款时间pay_time 如果不填默认为2即根据支付时间查询,查询的结果会根据该时间倒排序 */
    query_date_type int64 `json:"query_date_type";xml:"query_date_type"`
    
    /* query_end_date required结束时间,限制:结束时间-开始时间不能大于1天(根据order_id或者trade_id查询除外) */
    query_end_date Date `json:"query_end_date";xml:"query_end_date"`
    
    /* query_start_date required开始时间 */
    query_start_date Date `json:"query_start_date";xml:"query_start_date"`
    
    /* trade_id optional子订单编号 */
    trade_id int64 `json:"trade_id";xml:"trade_id"`
    
}

func (req *TaobaoTaeBillsGetRequest) GetAPIName() string {
	return "taobao.tae.bills.get"
}

// TaobaoTaeBillsGetResponse tae查询账单明细
type TaobaoTaeBillsGetResponse struct {
    
    /* bills Object Array账单列表 */
    bills BillDto `json:"bills";xml:"bills"`
    
    /* has_next Basic是否存在下一页 */
    has_next bool `json:"has_next";xml:"has_next"`
    
    /* total_results Basic当前页查询返回的结果数(0-100)。相同的查询时间段条件下，最大只能获取总共5000条记录。所以当大于等于5000时 ISV可以通过start_time及end_time来进行拆分，以保证可以查询到全部数据 */
    total_results int64 `json:"total_results";xml:"total_results"`
    
}
