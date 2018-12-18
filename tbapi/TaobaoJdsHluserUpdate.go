package tbsdk

// TaobaoJdsHluserUpdateRequest 订单全链路用户信息修改，比如是否开放买家端展示
type TaobaoJdsHluserUpdateRequest struct {
    
    /* open_for_buyer optional回流信息是否开通买家端展示,可选值open,close */
    open_for_buyer string `json:"open_for_buyer";xml:"open_for_buyer"`
    
    /* open_nodes optional为空,则默认是X_TO_SYSTEM,X_WAIT_ALLOCATION,X_OUT_WAREHOUSE 可选值是X_DOWNLOADED,X_TO_SYSTEM,X_SERVICE_AUDITED,X_FINANCE_AUDITED,X_ALLOCATION_NOTIFIED,X_WAIT_ALLOCATION,X_SORT_PRINTED,X_SEND_PRINTED,X_LOGISTICS_PRINTED,X_SORTED,X_EXAMINED,X_PACKAGED,X_WEIGHED,X_OUT_WAREHOUS */
    open_nodes string `json:"open_nodes";xml:"open_nodes"`
    
}

func (req *TaobaoJdsHluserUpdateRequest) GetAPIName() string {
	return "taobao.jds.hluser.update"
}

// TaobaoJdsHluserUpdateResponse 订单全链路用户信息修改，比如是否开放买家端展示
type TaobaoJdsHluserUpdateResponse struct {
    
    /* result Basic是否成功 */
    result bool `json:"result";xml:"result"`
    
}
