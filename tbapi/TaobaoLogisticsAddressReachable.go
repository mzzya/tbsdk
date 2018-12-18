package tbsdk

// TaobaoLogisticsAddressReachableRequest 根据输入的目标地址，判断服务是否可达。
现已支持筛单的快递公司共15家：中国邮政、EMS、国通、汇通、快捷、全峰、优速、圆通、宅急送、中通、顺丰、天天、韵达、德邦快递、申通
type TaobaoLogisticsAddressReachableRequest struct {
    
    /* address optional详细地址 */
    address string `json:"address";xml:"address"`
    
    /* area_code optional标准区域编码(三级行政区编码或是四级行政区)，可以通过taobao.areas.get获取，如北京市朝阳区为110105 */
    area_code string `json:"area_code";xml:"area_code"`
    
    /* partner_ids required物流公司编码ID，可以从这个接口获取所有物流公司的标准编码taobao.logistics.companies.get，可以传入多个值，以英文逗号分隔，如“1000000952,1000000953” */
    partner_ids string `json:"partner_ids";xml:"partner_ids"`
    
    /* service_type required服务对应的数字编码，如揽派范围对应88 */
    service_type int64 `json:"service_type";xml:"service_type"`
    
    /* source_area_code optional发货地，标准区域编码(四级行政)，可以通过taobao.areas.get获取，如浙江省杭州市余杭区闲林街道为 330110011 */
    source_area_code string `json:"source_area_code";xml:"source_area_code"`
    
}

func (req *TaobaoLogisticsAddressReachableRequest) GetAPIName() string {
	return "taobao.logistics.address.reachable"
}

// TaobaoLogisticsAddressReachableResponse 根据输入的目标地址，判断服务是否可达。
现已支持筛单的快递公司共15家：中国邮政、EMS、国通、汇通、快捷、全峰、优速、圆通、宅急送、中通、顺丰、天天、韵达、德邦快递、申通
type TaobaoLogisticsAddressReachableResponse struct {
    
    /* reachable_result_list Object Array地址可达返回结果，每个TP对应一个 */
    reachable_result_list AddressReachableResult `json:"reachable_result_list";xml:"reachable_result_list"`
    
}
