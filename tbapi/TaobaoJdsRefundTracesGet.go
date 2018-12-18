package tbsdk

// TaobaoJdsRefundTracesGetRequest 获取聚石塔数据共享的交易全链路的退款信息
type TaobaoJdsRefundTracesGetRequest struct {
    
    /* refund_id required淘宝的退款编号 */
    refund_id int64 `json:"refund_id";xml:"refund_id"`
    
    /* return_user_status optional是否返回用户状态(是否存在) */
    return_user_status bool `json:"return_user_status";xml:"return_user_status"`
    
}

func (req *TaobaoJdsRefundTracesGetRequest) GetAPIName() string {
	return "taobao.jds.refund.traces.get"
}

// TaobaoJdsRefundTracesGetResponse 获取聚石塔数据共享的交易全链路的退款信息
type TaobaoJdsRefundTracesGetResponse struct {
    
    /* traces Object Array退款跟踪列表 */
    traces RefundTrace `json:"traces";xml:"traces"`
    
    /* user_status Basic用户在全链路系统中的状态(比如是否存在) */
    user_status string `json:"user_status";xml:"user_status"`
    
}
