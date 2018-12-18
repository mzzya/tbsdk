package tbsdk

// TaobaoOmniorderDtdResendRequest 该接口触发对门店自送发码短信进行重发，码内容不变，接受码的手机号也不变。每个码限制每日重发一次，总共重发5次
type TaobaoOmniorderDtdResendRequest struct {
    
    /* main_order_id required淘宝主订单ID */
    main_order_id int64 `json:"main_order_id";xml:"main_order_id"`
    
}

func (req *TaobaoOmniorderDtdResendRequest) GetAPIName() string {
	return "taobao.omniorder.dtd.resend"
}

// TaobaoOmniorderDtdResendResponse 该接口触发对门店自送发码短信进行重发，码内容不变，接受码的手机号也不变。每个码限制每日重发一次，总共重发5次
type TaobaoOmniorderDtdResendResponse struct {
    
    /* message Basic错误信息 */
    message string `json:"message";xml:"message"`
    
    /* result_code Basic错误码，为0表示成功，非0表示失败 */
    result_code string `json:"result_code";xml:"result_code"`
    
}
