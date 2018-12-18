package tbsdk

// CainiaoWaybillIiSearchRequest 获取发货地&CP开通状态&账户的使用情况
type CainiaoWaybillIiSearchRequest struct {
    
    /* cp_code optional物流公司code */
    cp_code string `json:"cp_code";xml:"cp_code"`
    
}

func (req *CainiaoWaybillIiSearchRequest) GetAPIName() string {
	return "cainiao.waybill.ii.search"
}

// CainiaoWaybillIiSearchResponse 获取发货地&CP开通状态&账户的使用情况
type CainiaoWaybillIiSearchResponse struct {
    
    /* waybill_apply_subscription_cols Object ArrayCP网点信息及对应的商家的发货信息 */
    waybill_apply_subscription_cols WaybillApplySubscriptionInfo `json:"waybill_apply_subscription_cols";xml:"waybill_apply_subscription_cols"`
    
}
