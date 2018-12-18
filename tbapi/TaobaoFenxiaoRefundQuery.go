package tbsdk

// TaobaoFenxiaoRefundQueryRequest 供应商按查询条件批量查询代销采购退款
type TaobaoFenxiaoRefundQueryRequest struct {
    
    /* end_date required代销采购退款最迟修改时间。与start_date的最大时间间隔不能超过30天 */
    end_date Date `json:"end_date";xml:"end_date"`
    
    /* page_no optional页码（大于0的整数。无值或小于1的值按默认值1计） */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size optional每页条数（大于0但小于等于50的整数。无值或大于50或小于1的值按默认值50计） */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* query_seller_refund optional是否查询下游买家的退款信息 */
    query_seller_refund bool `json:"query_seller_refund";xml:"query_seller_refund"`
    
    /* start_date required代销采购退款单最早修改时间 */
    start_date Date `json:"start_date";xml:"start_date"`
    
}

func (req *TaobaoFenxiaoRefundQueryRequest) GetAPIName() string {
	return "taobao.fenxiao.refund.query"
}

// TaobaoFenxiaoRefundQueryResponse 供应商按查询条件批量查询代销采购退款
type TaobaoFenxiaoRefundQueryResponse struct {
    
    /* refund_list Object Array代销采购退款列表 */
    refund_list RefundDetail `json:"refund_list";xml:"refund_list"`
    
    /* total_results Basic按查询条件查到的记录总数 */
    total_results int64 `json:"total_results";xml:"total_results"`
    
}
