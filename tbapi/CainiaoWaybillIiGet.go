package tbsdk

// CainiaoWaybillIiGetRequest 菜鸟电子面单的云打印申请电子面单号的方法
type CainiaoWaybillIiGetRequest struct {
    
    /* param_waybill_cloud_print_apply_new_request required入参信息 */
    param_waybill_cloud_print_apply_new_request WaybillCloudPrintApplyNewRequest `json:"param_waybill_cloud_print_apply_new_request";xml:"param_waybill_cloud_print_apply_new_request"`
    
}

func (req *CainiaoWaybillIiGetRequest) GetAPIName() string {
	return "cainiao.waybill.ii.get"
}

// CainiaoWaybillIiGetResponse 菜鸟电子面单的云打印申请电子面单号的方法
type CainiaoWaybillIiGetResponse struct {
    
    /* modules Object Array系统自动生成 */
    modules WaybillCloudPrintResponse `json:"modules";xml:"modules"`
    
}
