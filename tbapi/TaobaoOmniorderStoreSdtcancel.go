package tbsdk

// TaobaoOmniorderStoreSdtcancelRequest 通知速店通取消取号
type TaobaoOmniorderStoreSdtcancelRequest struct {
    
    /* package_id required取号返回的packageId */
    package_id int64 `json:"package_id";xml:"package_id"`
    
}

func (req *TaobaoOmniorderStoreSdtcancelRequest) GetAPIName() string {
	return "taobao.omniorder.store.sdtcancel"
}

// TaobaoOmniorderStoreSdtcancelResponse 通知速店通取消取号
type TaobaoOmniorderStoreSdtcancelResponse struct {
    
    /* result Object返回结果 */
    result Result `json:"result";xml:"result"`
    
}
