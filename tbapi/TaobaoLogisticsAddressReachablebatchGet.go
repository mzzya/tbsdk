package tbsdk

// TaobaoLogisticsAddressReachablebatchGetRequest 批量判定服务是否可达
type TaobaoLogisticsAddressReachablebatchGetRequest struct {
    
    /* address_list optional筛单用户输入地址结构 */
    address_list AddressReachable `json:"address_list";xml:"address_list"`
    
}

func (req *TaobaoLogisticsAddressReachablebatchGetRequest) GetAPIName() string {
	return "taobao.logistics.address.reachablebatch.get"
}

// TaobaoLogisticsAddressReachablebatchGetResponse 批量判定服务是否可达
type TaobaoLogisticsAddressReachablebatchGetResponse struct {
    
    /* reachable_results Object Array物流是否可达结果列表 */
    reachable_results AddressReachableTopResult `json:"reachable_results";xml:"reachable_results"`
    
}
