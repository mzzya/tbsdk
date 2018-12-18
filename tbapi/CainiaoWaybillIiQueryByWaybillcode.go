package tbsdk

// CainiaoWaybillIiQueryByWaybillcodeRequest 通过面单号查看面单号的当前状态，如签收、发货、失效等。
type CainiaoWaybillIiQueryByWaybillcodeRequest struct {
    
    /* param_list optional系统自动生成 */
    param_list WaybillDetailQueryByWaybillCodeRequest `json:"param_list";xml:"param_list"`
    
}

func (req *CainiaoWaybillIiQueryByWaybillcodeRequest) GetAPIName() string {
	return "cainiao.waybill.ii.query.by.waybillcode"
}

// CainiaoWaybillIiQueryByWaybillcodeResponse 通过面单号查看面单号的当前状态，如签收、发货、失效等。
type CainiaoWaybillIiQueryByWaybillcodeResponse struct {
    
    /* modules Object Array查询返回值 */
    modules WaybillCloudPrintWithResultDescResponse `json:"modules";xml:"modules"`
    
}
