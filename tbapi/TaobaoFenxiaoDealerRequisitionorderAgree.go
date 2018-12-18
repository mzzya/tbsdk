package tbsdk

// TaobaoFenxiaoDealerRequisitionorderAgreeRequest 供应商或分销商通过采购申请/经销采购单审核
type TaobaoFenxiaoDealerRequisitionorderAgreeRequest struct {
    
    /* dealer_order_id required采购申请/经销采购单编号 */
    dealer_order_id int64 `json:"dealer_order_id";xml:"dealer_order_id"`
    
}

func (req *TaobaoFenxiaoDealerRequisitionorderAgreeRequest) GetAPIName() string {
	return "taobao.fenxiao.dealer.requisitionorder.agree"
}

// TaobaoFenxiaoDealerRequisitionorderAgreeResponse 供应商或分销商通过采购申请/经销采购单审核
type TaobaoFenxiaoDealerRequisitionorderAgreeResponse struct {
    
    /* is_success Basic操作是否成功。true：成功；false：失败。 */
    is_success bool `json:"is_success";xml:"is_success"`
    
}
