package tbsdk

// TaobaoTaeBillGetRequest 查询单笔账单明细
type TaobaoTaeBillGetRequest struct {
    
    /* account_id required虚拟账户科目编号 */
    account_id int64 `json:"account_id";xml:"account_id"`
    
    /* bid optional账单编号 */
    bid int64 `json:"bid";xml:"bid"`
    
    /* fields required传入需要返回的字段 */
    fields string `json:"fields";xml:"fields"`
    
    /* id required账单编号 */
    id int64 `json:"id";xml:"id"`
    
}

func (req *TaobaoTaeBillGetRequest) GetAPIName() string {
	return "taobao.tae.bill.get"
}

// TaobaoTaeBillGetResponse 查询单笔账单明细
type TaobaoTaeBillGetResponse struct {
    
    /* bill Object账单明细 */
    bill BillDto `json:"bill";xml:"bill"`
    
}
