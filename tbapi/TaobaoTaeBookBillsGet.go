package tbsdk

// TaobaoTaeBookBillsGetRequest tae查询虚拟账户明细数据
type TaobaoTaeBookBillsGetRequest struct {
    
    /* account_id optional虚拟账户科目编号 */
    account_id int64 `json:"account_id";xml:"account_id"`
    
    /* end_time required记账结束时间,end_time与start_time相差不能超过30天 */
    end_time Date `json:"end_time";xml:"end_time"`
    
    /* fields required需要返回的字段:bid,account_id,journal_type,amount,book_time,description,gmt_create,gmt_modified ,如果不是以上字段将自动忽略 */
    fields string `json:"fields";xml:"fields"`
    
    /* journal_types optional明细流水类型:流水类型:101、可用金充值；102、可用金扣除；103、冻结；104、解冻；105、冻结金充值；106、冻结金扣除 */
    journal_types int64 `json:"journal_types";xml:"journal_types"`
    
    /* page_no optional页码,传入值为1代表第一页,传入值为2代表第二页,依此类推.默认返回的数据是从第一页开始 */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size optional每页大小,建议40~100,不能超过100 */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* start_time required记账开始时间 */
    start_time Date `json:"start_time";xml:"start_time"`
    
}

func (req *TaobaoTaeBookBillsGetRequest) GetAPIName() string {
	return "taobao.tae.book.bills.get"
}

// TaobaoTaeBookBillsGetResponse tae查询虚拟账户明细数据
type TaobaoTaeBookBillsGetResponse struct {
    
    /* bills Object Array虚拟账户账单列表 */
    bills TopAcctCashJourDto `json:"bills";xml:"bills"`
    
    /* has_next Basic是否有下一页 */
    has_next bool `json:"has_next";xml:"has_next"`
    
    /* total_results Basic当前查询的结果数,非总数 */
    total_results int64 `json:"total_results";xml:"total_results"`
    
}
