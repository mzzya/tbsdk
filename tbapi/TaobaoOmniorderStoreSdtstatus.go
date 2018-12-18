package tbsdk

// TaobaoOmniorderStoreSdtstatusRequest 提供给商家查询运力单的状态。
type TaobaoOmniorderStoreSdtstatusRequest struct {
    
    /* package_id optional菜鸟裹裹的包裹ID */
    package_id int64 `json:"package_id";xml:"package_id"`
    
}

func (req *TaobaoOmniorderStoreSdtstatusRequest) GetAPIName() string {
	return "taobao.omniorder.store.sdtstatus"
}

// TaobaoOmniorderStoreSdtstatusResponse 提供给商家查询运力单的状态。
type TaobaoOmniorderStoreSdtstatusResponse struct {
    
    /* result Objectresult */
    result Result `json:"result";xml:"result"`
    
}
