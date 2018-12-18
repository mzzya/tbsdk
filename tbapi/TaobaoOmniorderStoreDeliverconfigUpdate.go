package tbsdk

// TaobaoOmniorderStoreDeliverconfigUpdateRequest 修改门店发货配置内容
type TaobaoOmniorderStoreDeliverconfigUpdateRequest struct {
    
    /* store_deliver_config required卖家发货配置 */
    store_deliver_config StoreDeliverConfig `json:"store_deliver_config";xml:"store_deliver_config"`
    
    /* store_id required门店ID */
    store_id int64 `json:"store_id";xml:"store_id"`
    
}

func (req *TaobaoOmniorderStoreDeliverconfigUpdateRequest) GetAPIName() string {
	return "taobao.omniorder.store.deliverconfig.update"
}

// TaobaoOmniorderStoreDeliverconfigUpdateResponse 修改门店发货配置内容
type TaobaoOmniorderStoreDeliverconfigUpdateResponse struct {
    
    /* result Objectresult */
    result Result `json:"result";xml:"result"`
    
}
