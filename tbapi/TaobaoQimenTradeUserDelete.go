package tbsdk

// TaobaoQimenTradeUserDeleteRequest 删除奇门订单链路用户
type TaobaoQimenTradeUserDeleteRequest struct {
    
}

func (req *TaobaoQimenTradeUserDeleteRequest) GetAPIName() string {
	return "taobao.qimen.trade.user.delete"
}

// TaobaoQimenTradeUserDeleteResponse 删除奇门订单链路用户
type TaobaoQimenTradeUserDeleteResponse struct {
    
    /* modal Basicmodal */
    modal bool `json:"modal";xml:"modal"`
    
}
