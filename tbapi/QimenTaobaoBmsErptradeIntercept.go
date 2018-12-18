package tbsdk

// QimenTaobaoBmsErptradeInterceptRequest 调用ERP接口，拦截订单
type QimenTaobaoBmsErptradeInterceptRequest struct {
    
    /* customerId required货主ID */
    customerId string `json:"customerId";xml:"customerId"`
    
    /* request optional请求主体 */
    request BmsTaobaoOrderIntercepteRequest `json:"request";xml:"request"`
    
}

func (req *QimenTaobaoBmsErptradeInterceptRequest) GetAPIName() string {
	return "qimen.taobao.bms.erptrade.intercept"
}

// QimenTaobaoBmsErptradeInterceptResponse 调用ERP接口，拦截订单
type QimenTaobaoBmsErptradeInterceptResponse struct {
    
    /* response ObjectResponse */
    response Response `json:"response";xml:"response"`
    
}
