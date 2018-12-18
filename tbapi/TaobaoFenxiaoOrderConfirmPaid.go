package tbsdk

// TaobaoFenxiaoOrderConfirmPaidRequest 供应商确认收款（非支付宝交易）。
type TaobaoFenxiaoOrderConfirmPaidRequest struct {
    
    /* confirm_remark optional确认支付信息（字数小于100） */
    confirm_remark string `json:"confirm_remark";xml:"confirm_remark"`
    
    /* purchase_order_id required采购单编号。 */
    purchase_order_id int64 `json:"purchase_order_id";xml:"purchase_order_id"`
    
}

func (req *TaobaoFenxiaoOrderConfirmPaidRequest) GetAPIName() string {
	return "taobao.fenxiao.order.confirm.paid"
}

// TaobaoFenxiaoOrderConfirmPaidResponse 供应商确认收款（非支付宝交易）。
type TaobaoFenxiaoOrderConfirmPaidResponse struct {
    
    /* is_success Basic确认结果成功与否 */
    is_success bool `json:"is_success";xml:"is_success"`
    
}
