package tbsdk

// TaobaoRdcAligeniusWarehouseResendUpdateRequest 补发单状态回传接口
type TaobaoRdcAligeniusWarehouseResendUpdateRequest struct {
    
    /* param0 required参数 */
    param0 UpdateResendStatusDto `json:"param0";xml:"param0"`
    
}

func (req *TaobaoRdcAligeniusWarehouseResendUpdateRequest) GetAPIName() string {
	return "taobao.rdc.aligenius.warehouse.resend.update"
}

// TaobaoRdcAligeniusWarehouseResendUpdateResponse 补发单状态回传接口
type TaobaoRdcAligeniusWarehouseResendUpdateResponse struct {
    
    /* result Objectresult */
    result Result `json:"result";xml:"result"`
    
}
