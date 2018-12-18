package tbsdk

// TaobaoRefundMessagesGetRequest 查询退款留言/凭证列表
type TaobaoRefundMessagesGetRequest struct {
    
    /* fields required需返回的字段列表。可选值：RefundMessage结构体中的所有字段，以半角逗号(,)分隔。 */
    fields string `json:"fields";xml:"fields"`
    
    /* page_no optional页码 */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size optional每页条数 */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* refund_id required退款单号 */
    refund_id int64 `json:"refund_id";xml:"refund_id"`
    
    /* refund_phase optional退款阶段，可选值：onsale（售中），aftersale（售后），天猫退款为必传。 */
    refund_phase string `json:"refund_phase";xml:"refund_phase"`
    
}

func (req *TaobaoRefundMessagesGetRequest) GetAPIName() string {
	return "taobao.refund.messages.get"
}

// TaobaoRefundMessagesGetResponse 查询退款留言/凭证列表
type TaobaoRefundMessagesGetResponse struct {
    
    /* refund_messages Object Array查询到的退款留言/凭证列表 */
    refund_messages RefundMessage `json:"refund_messages";xml:"refund_messages"`
    
    /* total_results Basic查询到的退款留言/凭证总数 */
    total_results int64 `json:"total_results";xml:"total_results"`
    
}
