package tbsdk

// TaobaoRdcAligeniusOrderReturngoodsNotifyRequest 退货单创建结果反馈
type TaobaoRdcAligeniusOrderReturngoodsNotifyRequest struct {
    
    /* parent_order_id optional主订单号 */
    parent_order_id int64 `json:"parent_order_id";xml:"parent_order_id"`
    
    /* refund_return_notes optional退货单创建结果列表 */
    refund_return_notes RefundReturnNotes `json:"refund_return_notes";xml:"refund_return_notes"`
    
}

func (req *TaobaoRdcAligeniusOrderReturngoodsNotifyRequest) GetAPIName() string {
	return "taobao.rdc.aligenius.order.returngoods.notify"
}

// TaobaoRdcAligeniusOrderReturngoodsNotifyResponse 退货单创建结果反馈
type TaobaoRdcAligeniusOrderReturngoodsNotifyResponse struct {
    
    /* result Basicsuccess */
    result bool `json:"result";xml:"result"`
    
    /* result_code BasicerrorCode */
    result_code string `json:"result_code";xml:"result_code"`
    
    /* result_info BasicerrorInfo */
    result_info string `json:"result_info";xml:"result_info"`
    
}
