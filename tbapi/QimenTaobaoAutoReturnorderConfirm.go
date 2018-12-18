package tbsdk

// QimenTaobaoAutoReturnorderConfirmRequest 菜鸟自动创建销退单的入库确认回传
type QimenTaobaoAutoReturnorderConfirmRequest struct {
    
    /* customerId requiredcustomerId */
    customerId string `json:"customerId";xml:"customerId"`
    
    /* request optional */
    request AutoReturnInOrderConfirmRequest `json:"request";xml:"request"`
    
}

func (req *QimenTaobaoAutoReturnorderConfirmRequest) GetAPIName() string {
	return "qimen.taobao.auto.returnorder.confirm"
}

// QimenTaobaoAutoReturnorderConfirmResponse 菜鸟自动创建销退单的入库确认回传
type QimenTaobaoAutoReturnorderConfirmResponse struct {
    
    /* response Objectresponse */
    response Response `json:"response";xml:"response"`
    
}
