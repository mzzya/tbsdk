package tbsdk

// QimenTaobaoUopTobDeliveryorderConfirmRequest 2B订单出库确认回传
type QimenTaobaoUopTobDeliveryorderConfirmRequest struct {
    
    /* customerId requiredcustomerId */
    customerId string `json:"customerId";xml:"customerId"`
    
    /* request optionaltob出库确认对象 */
    request ToBDeliveryOrderConfirmRequest `json:"request";xml:"request"`
    
}

func (req *QimenTaobaoUopTobDeliveryorderConfirmRequest) GetAPIName() string {
	return "qimen.taobao.uop.tob.deliveryorder.confirm"
}

// QimenTaobaoUopTobDeliveryorderConfirmResponse 2B订单出库确认回传
type QimenTaobaoUopTobDeliveryorderConfirmResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果 */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
}
