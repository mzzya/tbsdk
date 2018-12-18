package tbsdk

// TaobaoOmniorderStoreSdtquerystationRequest 速店通查询站点信息
type TaobaoOmniorderStoreSdtquerystationRequest struct {
    
    /* param_long2 required取号时返回的packageId */
    param_long2 int64 `json:"param_long2";xml:"param_long2"`
    
}

func (req *TaobaoOmniorderStoreSdtquerystationRequest) GetAPIName() string {
	return "taobao.omniorder.store.sdtquerystation"
}

// TaobaoOmniorderStoreSdtquerystationResponse 速店通查询站点信息
type TaobaoOmniorderStoreSdtquerystationResponse struct {
    
    /* result Objectresult */
    result Result `json:"result";xml:"result"`
    
}
