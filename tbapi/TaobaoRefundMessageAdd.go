package tbsdk

// TaobaoRefundMessageAddRequest 创建退款留言/凭证
type TaobaoRefundMessageAddRequest struct {
    
    /* content required留言内容。最大长度: 400个字节 */
    content string `json:"content";xml:"content"`
    
    /* image required图片（凭证）。类型: JPG,GIF,PNG;最大为: 500K */
    image byte[] `json:"image";xml:"image"`
    
    /* refund_id required退款编号。 */
    refund_id int64 `json:"refund_id";xml:"refund_id"`
    
}

func (req *TaobaoRefundMessageAddRequest) GetAPIName() string {
	return "taobao.refund.message.add"
}

// TaobaoRefundMessageAddResponse 创建退款留言/凭证
type TaobaoRefundMessageAddResponse struct {
    
    /* refund_message Object退款信息。包含id和created */
    refund_message RefundMessage `json:"refund_message";xml:"refund_message"`
    
}
