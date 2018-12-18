package tbsdk

// AlibabaEinvoiceCreateResultsIncrementGetRequest 增量开票结果获取
type AlibabaEinvoiceCreateResultsIncrementGetRequest struct {
    
    /* end_modified required终止查询时间 */
    end_modified Date `json:"end_modified";xml:"end_modified"`
    
    /* page_no optional显示的页码 */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size optional页面大小(不能超过200) */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* payee_register_no required收款方税务登记证号 */
    payee_register_no string `json:"payee_register_no";xml:"payee_register_no"`
    
    /* start_modified required起始查询时间 */
    start_modified Date `json:"start_modified";xml:"start_modified"`
    
    /* status optional开票状态 (waiting = 开票中) 、(create_success = 开票成功)、(create_failed = 开票失败) */
    status string `json:"status";xml:"status"`
    
}

func (req *AlibabaEinvoiceCreateResultsIncrementGetRequest) GetAPIName() string {
	return "alibaba.einvoice.create.results.increment.get"
}

// AlibabaEinvoiceCreateResultsIncrementGetResponse 增量开票结果获取
type AlibabaEinvoiceCreateResultsIncrementGetResponse struct {
    
    /* invoice_result_list Object Array开票结果返回列表 */
    invoice_result_list InvoiceResult `json:"invoice_result_list";xml:"invoice_result_list"`
    
    /* total_count Basic符合条件的开票总数 */
    total_count int64 `json:"total_count";xml:"total_count"`
    
}
