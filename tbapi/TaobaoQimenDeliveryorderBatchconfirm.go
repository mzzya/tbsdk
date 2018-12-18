package tbsdk

// TaobaoQimenDeliveryorderBatchconfirmRequest taobao.qimen.deliveryorder.batchconfirm
type TaobaoQimenDeliveryorderBatchconfirmRequest struct {
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* orders optional发货单列表 */
    orders Order `json:"orders";xml:"orders"`
    
}

func (req *TaobaoQimenDeliveryorderBatchconfirmRequest) GetAPIName() string {
	return "taobao.qimen.deliveryorder.batchconfirm"
}

// TaobaoQimenDeliveryorderBatchconfirmResponse taobao.qimen.deliveryorder.batchconfirm
type TaobaoQimenDeliveryorderBatchconfirmResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
}
