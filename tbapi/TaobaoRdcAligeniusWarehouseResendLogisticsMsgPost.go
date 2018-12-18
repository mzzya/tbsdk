package tbsdk

// TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostRequest 补发单erp物流信息回传平台
type TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostRequest struct {
    
    /* param0 required参数 */
    param0 SendResendLogisticsMsgDto `json:"param0";xml:"param0"`
    
}

func (req *TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostRequest) GetAPIName() string {
	return "taobao.rdc.aligenius.warehouse.resend.logistics.msg.post"
}

// TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostResponse 补发单erp物流信息回传平台
type TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostResponse struct {
    
    /* result Objectresult */
    result Result `json:"result";xml:"result"`
    
}
