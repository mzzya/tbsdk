package tbsdk

// TaobaoQimenOrderCallbackRequest 配送拦截
type TaobaoQimenOrderCallbackRequest struct {
    
    /* deliveryOrderCode optional奇门仓储字段,C123,string(50),, */
    deliveryOrderCode string `json:"deliveryOrderCode";xml:"deliveryOrderCode"`
    
    /* expressCode optional运单号 */
    expressCode string `json:"expressCode";xml:"expressCode"`
    
    /* orderId optional奇门仓储字段,C123,string(50),, */
    orderId string `json:"orderId";xml:"orderId"`
    
    /* ownerCode optional奇门仓储字段,C123,string(50),, */
    ownerCode string `json:"ownerCode";xml:"ownerCode"`
    
    /* warehouseCode optional奇门仓储字段,C123,string(50),, */
    warehouseCode string `json:"warehouseCode";xml:"warehouseCode"`
    
}

func (req *TaobaoQimenOrderCallbackRequest) GetAPIName() string {
	return "taobao.qimen.order.callback"
}

// TaobaoQimenOrderCallbackResponse 配送拦截
type TaobaoQimenOrderCallbackResponse struct {
    
    /* code Basic响应码,0,string(50),, */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure,success,string(10),必填, */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息,invalid appkey,string(100),, */
    message string `json:"message";xml:"message"`
    
}
