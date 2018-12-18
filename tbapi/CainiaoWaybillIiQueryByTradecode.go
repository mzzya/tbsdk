package tbsdk

// CainiaoWaybillIiQueryByTradecodeRequest 通过订单号查看面单的信息
type CainiaoWaybillIiQueryByTradecodeRequest struct {
    
    /* param_list optional订单号列表 */
    param_list WaybillDetailQueryByBizSubCodeRequest `json:"param_list";xml:"param_list"`
    
}

func (req *CainiaoWaybillIiQueryByTradecodeRequest) GetAPIName() string {
	return "cainiao.waybill.ii.query.by.tradecode"
}

// CainiaoWaybillIiQueryByTradecodeResponse 通过订单号查看面单的信息
type CainiaoWaybillIiQueryByTradecodeResponse struct {
    
    /* modules Object Array查询返回值 */
    modules WaybillCloudPrintWithResultDescResponse `json:"modules";xml:"modules"`
    
}
