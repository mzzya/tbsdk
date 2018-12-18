package tbsdk

// TaobaoFenxiaoRefundGetRequest 分销商或供应商可以查询某子单的退款信息，以及下游订单的退款信息
type TaobaoFenxiaoRefundGetRequest struct {
    
    /* query_seller_refund optional是否查询下游买家的退款信息 */
    query_seller_refund bool `json:"query_seller_refund";xml:"query_seller_refund"`
    
    /* sub_order_id required要查询的退款子单的id */
    sub_order_id int64 `json:"sub_order_id";xml:"sub_order_id"`
    
}

func (req *TaobaoFenxiaoRefundGetRequest) GetAPIName() string {
	return "taobao.fenxiao.refund.get"
}

// TaobaoFenxiaoRefundGetResponse 分销商或供应商可以查询某子单的退款信息，以及下游订单的退款信息
type TaobaoFenxiaoRefundGetResponse struct {
    
    /* refund_detail Object退款详情 */
    refund_detail RefundDetail `json:"refund_detail";xml:"refund_detail"`
    
}
