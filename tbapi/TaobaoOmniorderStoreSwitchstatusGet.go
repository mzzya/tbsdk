package tbsdk

// TaobaoOmniorderStoreSwitchstatusGetRequest 查询门店发货、门店自提状态
type TaobaoOmniorderStoreSwitchstatusGetRequest struct {
    
    /* seller_id required卖家ID */
    seller_id int64 `json:"seller_id";xml:"seller_id"`
    
    /* store_id required门店ID */
    store_id int64 `json:"store_id";xml:"store_id"`
    
}

func (req *TaobaoOmniorderStoreSwitchstatusGetRequest) GetAPIName() string {
	return "taobao.omniorder.store.switchstatus.get"
}

// TaobaoOmniorderStoreSwitchstatusGetResponse 查询门店发货、门店自提状态
type TaobaoOmniorderStoreSwitchstatusGetResponse struct {
    
    /* result Objectresult */
    result Result `json:"result";xml:"result"`
    
}
