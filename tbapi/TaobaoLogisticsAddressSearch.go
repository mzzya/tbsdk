package tbsdk

// TaobaoLogisticsAddressSearchRequest 通过此接口查询卖家地址库，
type TaobaoLogisticsAddressSearchRequest struct {
    
    /* rdef optional可选，参数列表如下：<br><font color='red'>no_def:查询非默认地址<br>get_def:查询默认取货地址，也即对应卖家后台地址库中发货地址（send_def暂不起作用）<br>cancel_def:查询默认退货地址<br>缺省此参数时，查询所有当前用户的地址库</font> */
    rdef string `json:"rdef";xml:"rdef"`
    
}

func (req *TaobaoLogisticsAddressSearchRequest) GetAPIName() string {
	return "taobao.logistics.address.search"
}

// TaobaoLogisticsAddressSearchResponse 通过此接口查询卖家地址库，
type TaobaoLogisticsAddressSearchResponse struct {
    
    /* addresses Object Array一组地址库数据， */
    addresses AddressResult `json:"addresses";xml:"addresses"`
    
}
