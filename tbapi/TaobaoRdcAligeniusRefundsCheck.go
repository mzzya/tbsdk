package tbsdk

// TaobaoRdcAligeniusRefundsCheckRequest 根据退款信息，对退款单进行审核
type TaobaoRdcAligeniusRefundsCheckRequest struct {
    
    /* param required入参 */
    param RefundCheckDto `json:"param";xml:"param"`
    
}

func (req *TaobaoRdcAligeniusRefundsCheckRequest) GetAPIName() string {
	return "taobao.rdc.aligenius.refunds.check"
}

// TaobaoRdcAligeniusRefundsCheckResponse 根据退款信息，对退款单进行审核
type TaobaoRdcAligeniusRefundsCheckResponse struct {
    
    /* result Objectresult */
    result Result `json:"result";xml:"result"`
    
}
