package tbsdk

// TaobaoOmniorderStoreDeliverconfigGetRequest 查询门店发货配置内容
type TaobaoOmniorderStoreDeliverconfigGetRequest struct {
    
    /* activity optional是否是活动期 */
    activity bool `json:"activity";xml:"activity"`
    
    /* store_id required门店ID */
    store_id int64 `json:"store_id";xml:"store_id"`
    
}

func (req *TaobaoOmniorderStoreDeliverconfigGetRequest) GetAPIName() string {
	return "taobao.omniorder.store.deliverconfig.get"
}

// TaobaoOmniorderStoreDeliverconfigGetResponse 查询门店发货配置内容
type TaobaoOmniorderStoreDeliverconfigGetResponse struct {
    
    /* result Objectresult */
    result Result `json:"result";xml:"result"`
    
}
