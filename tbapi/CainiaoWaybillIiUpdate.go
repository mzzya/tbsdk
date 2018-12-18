package tbsdk

// CainiaoWaybillIiUpdateRequest 商家更新电子面单号对应的面单信息。
type CainiaoWaybillIiUpdateRequest struct {
    
    /* param_waybill_cloud_print_update_request required更新请求信息 */
    param_waybill_cloud_print_update_request WaybillCloudPrintUpdateRequest `json:"param_waybill_cloud_print_update_request";xml:"param_waybill_cloud_print_update_request"`
    
}

func (req *CainiaoWaybillIiUpdateRequest) GetAPIName() string {
	return "cainiao.waybill.ii.update"
}

// CainiaoWaybillIiUpdateResponse 商家更新电子面单号对应的面单信息。
type CainiaoWaybillIiUpdateResponse struct {
    
    /* print_data Basic模板内容 */
    print_data string `json:"print_data";xml:"print_data"`
    
    /* waybill_code Basic面单号 */
    waybill_code string `json:"waybill_code";xml:"waybill_code"`
    
}
