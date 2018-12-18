package tbsdk

// CainiaoSmartdeliveryCpIModifyRequest 商家修改智能发货引擎推荐的cp
type CainiaoSmartdeliveryCpIModifyRequest struct {
    
    /* modify_smart_delivery_cp_request required修改智选CP请求 */
    modify_smart_delivery_cp_request ModifySmartDeliveryCpRequest `json:"modify_smart_delivery_cp_request";xml:"modify_smart_delivery_cp_request"`
    
}

func (req *CainiaoSmartdeliveryCpIModifyRequest) GetAPIName() string {
	return "cainiao.smartdelivery.cp.i.modify"
}

// CainiaoSmartdeliveryCpIModifyResponse 商家修改智能发货引擎推荐的cp
type CainiaoSmartdeliveryCpIModifyResponse struct {
    
    /* modify_smart_delivery_cp_response Object更新智能发货智选cp返回结果 */
    modify_smart_delivery_cp_response ModifySmartDeliveryCpResponse `json:"modify_smart_delivery_cp_response";xml:"modify_smart_delivery_cp_response"`
    
}
