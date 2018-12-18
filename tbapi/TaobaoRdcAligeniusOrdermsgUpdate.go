package tbsdk

// TaobaoRdcAligeniusOrdermsgUpdateRequest 用于订单消息处理状态回传
type TaobaoRdcAligeniusOrdermsgUpdateRequest struct {
    
    /* oid required子订单（消息中传的子订单） */
    oid int64 `json:"oid";xml:"oid"`
    
    /* status required处理状态，1=成功，2=处理失败 */
    status int64 `json:"status";xml:"status"`
    
    /* tid required主订单（消息中传的主订单） */
    tid int64 `json:"tid";xml:"tid"`
    
}

func (req *TaobaoRdcAligeniusOrdermsgUpdateRequest) GetAPIName() string {
	return "taobao.rdc.aligenius.ordermsg.update"
}

// TaobaoRdcAligeniusOrdermsgUpdateResponse 用于订单消息处理状态回传
type TaobaoRdcAligeniusOrdermsgUpdateResponse struct {
    
    /* result Objectresult */
    result Result `json:"result";xml:"result"`
    
}
