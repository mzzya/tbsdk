package tbsdk

// TaobaoWlbWaybillIFullupdateRequest 商家更新电子面单号对应的订单信息。
<br/>a.涉及到订单信息（如拆/合单、取消订单等）、修改发货地的逻辑时候，需要使用CANCEL+GET组合操作。一般的面单信息更新使用该接口；
<br/>b.该接口是全量修改接口，对于不做更新的字段也要把原有的字段值传进去，否则做为修改为空处理。
type TaobaoWlbWaybillIFullupdateRequest struct {
    
    /* waybill_apply_full_update_request required更新面单信息请求 */
    waybill_apply_full_update_request WaybillApplyFullUpdateRequest `json:"waybill_apply_full_update_request";xml:"waybill_apply_full_update_request"`
    
}

func (req *TaobaoWlbWaybillIFullupdateRequest) GetAPIName() string {
	return "taobao.wlb.waybill.i.fullupdate"
}

// TaobaoWlbWaybillIFullupdateResponse 商家更新电子面单号对应的订单信息。
<br/>a.涉及到订单信息（如拆/合单、取消订单等）、修改发货地的逻辑时候，需要使用CANCEL+GET组合操作。一般的面单信息更新使用该接口；
<br/>b.该接口是全量修改接口，对于不做更新的字段也要把原有的字段值传进去，否则做为修改为空处理。
type TaobaoWlbWaybillIFullupdateResponse struct {
    
    /* waybill_apply_update_info Object更新接口出参 */
    waybill_apply_update_info WaybillApplyUpdateInfo `json:"waybill_apply_update_info";xml:"waybill_apply_update_info"`
    
}
