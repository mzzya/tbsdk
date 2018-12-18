package tbsdk

// TaobaoRdcAligeniusOrderEventUpdateRequest 供erp回传订单变更状态事件
type TaobaoRdcAligeniusOrderEventUpdateRequest struct {
    
    /* param0 optional参数 */
    param0 OrderEventDto `json:"param0";xml:"param0"`
    
}

func (req *TaobaoRdcAligeniusOrderEventUpdateRequest) GetAPIName() string {
	return "taobao.rdc.aligenius.order.event.update"
}

// TaobaoRdcAligeniusOrderEventUpdateResponse 供erp回传订单变更状态事件
type TaobaoRdcAligeniusOrderEventUpdateResponse struct {
    
    /* fail_code Basic错误码 */
    fail_code string `json:"fail_code";xml:"fail_code"`
    
    /* fail_info Basic错误描述 */
    fail_info string `json:"fail_info";xml:"fail_info"`
    
    /* success_flag Basic是否成功 */
    success_flag bool `json:"success_flag";xml:"success_flag"`
    
}
