package tbsdk

// TaobaoWlbWaybillICancelRequest 面单号有误需要取消的时候，调用该接口取消获取的电子面单。
type TaobaoWlbWaybillICancelRequest struct {
    
    /* waybill_apply_cancel_request required取消接口入参 */
    waybill_apply_cancel_request WaybillApplyCancelRequest `json:"waybill_apply_cancel_request";xml:"waybill_apply_cancel_request"`
    
}

func (req *TaobaoWlbWaybillICancelRequest) GetAPIName() string {
	return "taobao.wlb.waybill.i.cancel"
}

// TaobaoWlbWaybillICancelResponse 面单号有误需要取消的时候，调用该接口取消获取的电子面单。
type TaobaoWlbWaybillICancelResponse struct {
    
    /* cancel_result Basic调用取消是否成功 */
    cancel_result bool `json:"cancel_result";xml:"cancel_result"`
    
}
