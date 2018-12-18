package tbsdk

// TaobaoWlbOrderCancelRequest 取消物流宝订单
type TaobaoWlbOrderCancelRequest struct {
    
    /* wlb_order_code required物流宝订单编号 */
    wlb_order_code string `json:"wlb_order_code";xml:"wlb_order_code"`
    
}

func (req *TaobaoWlbOrderCancelRequest) GetAPIName() string {
	return "taobao.wlb.order.cancel"
}

// TaobaoWlbOrderCancelResponse 取消物流宝订单
type TaobaoWlbOrderCancelResponse struct {
    
    /* error_code_list Basic错误编码列表 */
    error_code_list string `json:"error_code_list";xml:"error_code_list"`
    
    /* modify_time Basic修改时间，只有在取消成功的情况下，才可以做 */
    modify_time Date `json:"modify_time";xml:"modify_time"`
    
}
