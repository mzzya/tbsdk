package tbsdk

// TaobaoOmniorderStoreCollectconfigGetRequest 查询门店自提配置内容
type TaobaoOmniorderStoreCollectconfigGetRequest struct {
    
    /* activity optional是否是活动期 */
    activity bool `json:"activity";xml:"activity"`
    
    /* store_id required门店ID */
    store_id int64 `json:"store_id";xml:"store_id"`
    
}

func (req *TaobaoOmniorderStoreCollectconfigGetRequest) GetAPIName() string {
	return "taobao.omniorder.store.collectconfig.get"
}

// TaobaoOmniorderStoreCollectconfigGetResponse 查询门店自提配置内容
type TaobaoOmniorderStoreCollectconfigGetResponse struct {
    
    /* result Objectresult */
    result Result `json:"result";xml:"result"`
    
}
