package tbsdk

// TaobaoRefundRefusereasonGetRequest 获取商家拒绝原因列表
type TaobaoRefundRefusereasonGetRequest struct {
    
    /* fields required返回参数 */
    fields string `json:"fields";xml:"fields"`
    
    /* refund_id required退款编号 */
    refund_id int64 `json:"refund_id";xml:"refund_id"`
    
    /* refund_phase required售中或售后 */
    refund_phase string `json:"refund_phase";xml:"refund_phase"`
    
}

func (req *TaobaoRefundRefusereasonGetRequest) GetAPIName() string {
	return "taobao.refund.refusereason.get"
}

// TaobaoRefundRefusereasonGetResponse 获取商家拒绝原因列表
type TaobaoRefundRefusereasonGetResponse struct {
    
    /* has_next Basic是否存在下一页 */
    has_next bool `json:"has_next";xml:"has_next"`
    
    /* reasons Object Array卖家拒绝原因对象 */
    reasons Reason `json:"reasons";xml:"reasons"`
    
    /* total_results Basic原因个数 */
    total_results int64 `json:"total_results";xml:"total_results"`
    
}
