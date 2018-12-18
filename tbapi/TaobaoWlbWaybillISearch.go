package tbsdk

// TaobaoWlbWaybillISearchRequest 获取发货地&CP开通状态&账户的使用情况
type TaobaoWlbWaybillISearchRequest struct {
    
    /* waybill_apply_request required查询网点信息 */
    waybill_apply_request WaybillApplyRequest `json:"waybill_apply_request";xml:"waybill_apply_request"`
    
}

func (req *TaobaoWlbWaybillISearchRequest) GetAPIName() string {
	return "taobao.wlb.waybill.i.search"
}

// TaobaoWlbWaybillISearchResponse 获取发货地&CP开通状态&账户的使用情况
type TaobaoWlbWaybillISearchResponse struct {
    
    /* subscribtions Object Array订购关系 */
    subscribtions WaybillApplySubscriptionInfo `json:"subscribtions";xml:"subscribtions"`
    
}
