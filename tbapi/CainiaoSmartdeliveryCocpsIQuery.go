package tbsdk

// CainiaoSmartdeliveryCocpsIQueryRequest 获取电子面单订购关系中智能发货引擎支持的合作物流公司
type CainiaoSmartdeliveryCocpsIQueryRequest struct {
    
    /* send_address optional发货地址 */
    send_address Address `json:"send_address";xml:"send_address"`
    
}

func (req *CainiaoSmartdeliveryCocpsIQueryRequest) GetAPIName() string {
	return "cainiao.smartdelivery.cocps.i.query"
}

// CainiaoSmartdeliveryCocpsIQueryResponse 获取电子面单订购关系中智能发货引擎支持的合作物流公司
type CainiaoSmartdeliveryCocpsIQueryResponse struct {
    
    /* smart_delivery_collaborate_cps_info_list Object Array返回结果 */
    smart_delivery_collaborate_cps_info_list SmartDeliveryCollaborateCpsInfo `json:"smart_delivery_collaborate_cps_info_list";xml:"smart_delivery_collaborate_cps_info_list"`
    
}
