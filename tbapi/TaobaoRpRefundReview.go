package tbsdk

// TaobaoRpRefundReviewRequest 审核退款单，标志是否可用于批量退款，目前仅支持天猫订单。
type TaobaoRpRefundReviewRequest struct {
    
    /* message required审核留言 */
    message string `json:"message";xml:"message"`
    
    /* operator required审核人姓名 */
    operator string `json:"operator";xml:"operator"`
    
    /* refund_id required退款单编号 */
    refund_id int64 `json:"refund_id";xml:"refund_id"`
    
    /* refund_phase required退款阶段，可选值：售中：onsale，售后：aftersale */
    refund_phase string `json:"refund_phase";xml:"refund_phase"`
    
    /* refund_version required退款最后更新时间，以时间戳的方式表示 */
    refund_version int64 `json:"refund_version";xml:"refund_version"`
    
    /* result required审核是否可用于批量退款，可选值：true（审核通过），false（审核不通过或反审核） */
    result bool `json:"result";xml:"result"`
    
}

func (req *TaobaoRpRefundReviewRequest) GetAPIName() string {
	return "taobao.rp.refund.review"
}

// TaobaoRpRefundReviewResponse 审核退款单，标志是否可用于批量退款，目前仅支持天猫订单。
type TaobaoRpRefundReviewResponse struct {
    
    /* is_success Basicsuccess */
    is_success bool `json:"is_success";xml:"is_success"`
    
}
