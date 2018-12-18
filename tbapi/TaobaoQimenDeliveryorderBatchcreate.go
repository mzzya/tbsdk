package tbsdk

// TaobaoQimenDeliveryorderBatchcreateRequest ERP调用接口，将发货信息批量推送给WMS
type TaobaoQimenDeliveryorderBatchcreateRequest struct {
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* orders optional订单信息 */
    orders Order `json:"orders";xml:"orders"`
    
}

func (req *TaobaoQimenDeliveryorderBatchcreateRequest) GetAPIName() string {
	return "taobao.qimen.deliveryorder.batchcreate"
}

// TaobaoQimenDeliveryorderBatchcreateResponse ERP调用接口，将发货信息批量推送给WMS
type TaobaoQimenDeliveryorderBatchcreateResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
    /* orders Object Array订单详情 */
    orders Order `json:"orders";xml:"orders"`
    
}
