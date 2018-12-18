package tbsdk

// TaobaoWlbWaybillIGetRequest 商家根据订单信息，实时、批量获取指定物流服务商的电子面单号。
type TaobaoWlbWaybillIGetRequest struct {
    
    /* waybill_apply_new_request required面单申请 */
    waybill_apply_new_request WaybillApplyNewRequest `json:"waybill_apply_new_request";xml:"waybill_apply_new_request"`
    
}

func (req *TaobaoWlbWaybillIGetRequest) GetAPIName() string {
	return "taobao.wlb.waybill.i.get"
}

// TaobaoWlbWaybillIGetResponse 商家根据订单信息，实时、批量获取指定物流服务商的电子面单号。
type TaobaoWlbWaybillIGetResponse struct {
    
    /* waybill_apply_new_cols Object Array面单申请接口返回信息 */
    waybill_apply_new_cols WaybillApplyNewInfo `json:"waybill_apply_new_cols";xml:"waybill_apply_new_cols"`
    
}
