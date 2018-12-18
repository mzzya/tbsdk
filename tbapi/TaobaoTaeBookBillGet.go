package tbsdk

// TaobaoTaeBookBillGetRequest tae查询单笔虚拟账户明细
type TaobaoTaeBookBillGetRequest struct {
    
    /* account_id required虚拟账户科目编号 */
    account_id int64 `json:"account_id";xml:"account_id"`
    
    /* bid optional虚拟账户流水编号 */
    bid int64 `json:"bid";xml:"bid"`
    
    /* fields required需要返回的字段:参见BookBill结构体 */
    fields string `json:"fields";xml:"fields"`
    
    /* id required虚拟账户流水编号 */
    id int64 `json:"id";xml:"id"`
    
}

func (req *TaobaoTaeBookBillGetRequest) GetAPIName() string {
	return "taobao.tae.book.bill.get"
}

// TaobaoTaeBookBillGetResponse tae查询单笔虚拟账户明细
type TaobaoTaeBookBillGetResponse struct {
    
    /* bookbill Object虚拟账户账单 */
    bookbill TopAcctCashJourDto `json:"bookbill";xml:"bookbill"`
    
}
