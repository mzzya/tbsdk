package tbsdk

// TaobaoOmniorderDtdConsumeRequest 该接口根据传入的码及订单信息，如果码与订单一致，则对门店自送服务进行核销。
type TaobaoOmniorderDtdConsumeRequest struct {
    
    /* param_door2door_consume_request optional核销信息 */
    param_door2door_consume_request Door2doorConsumeRequest `json:"param_door2door_consume_request";xml:"param_door2door_consume_request"`
    
}

func (req *TaobaoOmniorderDtdConsumeRequest) GetAPIName() string {
	return "taobao.omniorder.dtd.consume"
}

// TaobaoOmniorderDtdConsumeResponse 该接口根据传入的码及订单信息，如果码与订单一致，则对门店自送服务进行核销。
type TaobaoOmniorderDtdConsumeResponse struct {
    
    /* message Basic错误西溪 */
    message string `json:"message";xml:"message"`
    
    /* result_code Basic错误码，为0表示成功，非0表示失败 */
    result_code string `json:"result_code";xml:"result_code"`
    
}
