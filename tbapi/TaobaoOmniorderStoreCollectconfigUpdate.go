package tbsdk

// TaobaoOmniorderStoreCollectconfigUpdateRequest 修改门店自提配置内容
type TaobaoOmniorderStoreCollectconfigUpdateRequest struct {
    
    /* store_collect_config required门店自提配置 */
    store_collect_config StoreCollectConfig `json:"store_collect_config";xml:"store_collect_config"`
    
    /* store_id required门店ID */
    store_id int64 `json:"store_id";xml:"store_id"`
    
}

func (req *TaobaoOmniorderStoreCollectconfigUpdateRequest) GetAPIName() string {
	return "taobao.omniorder.store.collectconfig.update"
}

// TaobaoOmniorderStoreCollectconfigUpdateResponse 修改门店自提配置内容
type TaobaoOmniorderStoreCollectconfigUpdateResponse struct {
    
    /* result Objectresult */
    result Result `json:"result";xml:"result"`
    
}
