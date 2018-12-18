package tbsdk

// TaobaoQimenDeliveryorderConfirmRequest taobao.qimen.deliveryorder.confirm
type TaobaoQimenDeliveryorderConfirmRequest struct {
    
    /* deliveryOrder optional发货单信息 */
    deliveryOrder DeliveryOrder `json:"deliveryOrder";xml:"deliveryOrder"`
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* orderLines optional单据列表 */
    orderLines OrderLine `json:"orderLines";xml:"orderLines"`
    
    /* packages optional包裹信息 */
    packages Package `json:"packages";xml:"packages"`
    
}

func (req *TaobaoQimenDeliveryorderConfirmRequest) GetAPIName() string {
	return "taobao.qimen.deliveryorder.confirm"
}

// TaobaoQimenDeliveryorderConfirmResponse taobao.qimen.deliveryorder.confirm
type TaobaoQimenDeliveryorderConfirmResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
}
