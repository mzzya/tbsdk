package tbsdk

// TaobaoRdcAligeniusSendgoodsCancelRequest 提供商家在仅退款中发送取消发货状态
type TaobaoRdcAligeniusSendgoodsCancelRequest struct {
    
    /* param required请求参数 */
    param CancelGoodsDto `json:"param";xml:"param"`
    
}

func (req *TaobaoRdcAligeniusSendgoodsCancelRequest) GetAPIName() string {
	return "taobao.rdc.aligenius.sendgoods.cancel"
}

// TaobaoRdcAligeniusSendgoodsCancelResponse 提供商家在仅退款中发送取消发货状态
type TaobaoRdcAligeniusSendgoodsCancelResponse struct {
    
    /* result Objectresult */
    result Result `json:"result";xml:"result"`
    
}
