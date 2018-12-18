package tbsdk

// TaobaoLogisticsAddressRemoveRequest 用此接口删除卖家地址库
type TaobaoLogisticsAddressRemoveRequest struct {
    
    /* contact_id required地址库ID */
    contact_id int64 `json:"contact_id";xml:"contact_id"`
    
}

func (req *TaobaoLogisticsAddressRemoveRequest) GetAPIName() string {
	return "taobao.logistics.address.remove"
}

// TaobaoLogisticsAddressRemoveResponse 用此接口删除卖家地址库
type TaobaoLogisticsAddressRemoveResponse struct {
    
    /* address_result Object只返回修改日期modify_date */
    address_result AddressResult `json:"address_result";xml:"address_result"`
    
}
