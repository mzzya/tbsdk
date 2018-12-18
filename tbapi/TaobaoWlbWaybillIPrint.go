package tbsdk

// TaobaoWlbWaybillIPrintRequest 打印面单前的校验接口，判断面单号信息与订单信息是否匹配。
type TaobaoWlbWaybillIPrintRequest struct {
    
    /* waybill_apply_print_check_request required打印请求 */
    waybill_apply_print_check_request WaybillApplyPrintCheckRequest `json:"waybill_apply_print_check_request";xml:"waybill_apply_print_check_request"`
    
}

func (req *TaobaoWlbWaybillIPrintRequest) GetAPIName() string {
	return "taobao.wlb.waybill.i.print"
}

// TaobaoWlbWaybillIPrintResponse 打印面单前的校验接口，判断面单号信息与订单信息是否匹配。
type TaobaoWlbWaybillIPrintResponse struct {
    
    /* waybill_apply_print_check_infos Object Array面单打印信息 */
    waybill_apply_print_check_infos WaybillApplyPrintCheckInfo `json:"waybill_apply_print_check_infos";xml:"waybill_apply_print_check_infos"`
    
}
