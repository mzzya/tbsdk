package tbsdk

// TaobaoLogisticsExpressModifyAppointRequest 商家通过此api操作修改物流单，交易单的收货人地址、收货人联系方式、预约配送日期
type TaobaoLogisticsExpressModifyAppointRequest struct {
    
    /* express_modify_appoint_top_request required改约请求对象 */
    express_modify_appoint_top_request ExpressModifyAppointTopRequestDto `json:"express_modify_appoint_top_request";xml:"express_modify_appoint_top_request"`
    
}

func (req *TaobaoLogisticsExpressModifyAppointRequest) GetAPIName() string {
	return "taobao.logistics.express.modify.appoint"
}

// TaobaoLogisticsExpressModifyAppointResponse 商家通过此api操作修改物流单，交易单的收货人地址、收货人联系方式、预约配送日期
type TaobaoLogisticsExpressModifyAppointResponse struct {
    
    /* result Object调用结果 */
    result SingleResultDto `json:"result";xml:"result"`
    
}
