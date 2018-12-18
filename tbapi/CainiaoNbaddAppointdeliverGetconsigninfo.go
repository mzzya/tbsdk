package tbsdk

// CainiaoNbaddAppointdeliverGetconsigninfoRequest 获取支持定时派送服务发货信息
type CainiaoNbaddAppointdeliverGetconsigninfoRequest struct {
    
    /* trade_order_id required淘宝交易订单id */
    trade_order_id int64 `json:"trade_order_id";xml:"trade_order_id"`
    
}

func (req *CainiaoNbaddAppointdeliverGetconsigninfoRequest) GetAPIName() string {
	return "cainiao.nbadd.appointdeliver.getconsigninfo"
}

// CainiaoNbaddAppointdeliverGetconsigninfoResponse 获取支持定时派送服务发货信息
type CainiaoNbaddAppointdeliverGetconsigninfoResponse struct {
    
    /* is_success Basic调用是否成功，正常情况下都会成功 */
    is_success bool `json:"is_success";xml:"is_success"`
    
    /* result Object发货信息 */
    result ConsignSupportInfoDTO `json:"result";xml:"result"`
    
    /* result_code Basic错误编码 */
    result_code string `json:"result_code";xml:"result_code"`
    
    /* result_desc Basic错误描述 */
    result_desc string `json:"result_desc";xml:"result_desc"`
    
}
