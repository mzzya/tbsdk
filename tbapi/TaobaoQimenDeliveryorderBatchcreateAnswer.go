package tbsdk

// TaobaoQimenDeliveryorderBatchcreateAnswerRequest WMS调用接口，用于异步化的批量发货单创建结果通知。（如菜鸟发货单批量创建结果的返回）
type TaobaoQimenDeliveryorderBatchcreateAnswerRequest struct {
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* orders optional发货单列表 */
    orders Order `json:"orders";xml:"orders"`
    
}

func (req *TaobaoQimenDeliveryorderBatchcreateAnswerRequest) GetAPIName() string {
	return "taobao.qimen.deliveryorder.batchcreate.answer"
}

// TaobaoQimenDeliveryorderBatchcreateAnswerResponse WMS调用接口，用于异步化的批量发货单创建结果通知。（如菜鸟发货单批量创建结果的返回）
type TaobaoQimenDeliveryorderBatchcreateAnswerResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
}
