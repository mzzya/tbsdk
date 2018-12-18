package tbsdk

// TaobaoRdcAligeniusDistributionLogisticsCancelRequest 截配状态回传接口
type TaobaoRdcAligeniusDistributionLogisticsCancelRequest struct {
    
    /* param0 optional参数 */
    param0 CancelDistributionDto `json:"param0";xml:"param0"`
    
}

func (req *TaobaoRdcAligeniusDistributionLogisticsCancelRequest) GetAPIName() string {
	return "taobao.rdc.aligenius.distribution.logistics.cancel"
}

// TaobaoRdcAligeniusDistributionLogisticsCancelResponse 截配状态回传接口
type TaobaoRdcAligeniusDistributionLogisticsCancelResponse struct {
    
    /* result Objectresult */
    result Result `json:"result";xml:"result"`
    
}
