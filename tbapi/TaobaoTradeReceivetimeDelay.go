package tbsdk

// TaobaoTradeReceivetimeDelayRequest 延长交易收货时间
type TaobaoTradeReceivetimeDelayRequest struct {
    
    /* days required延长收货的天数，可选值为：3, 5, 7, 10。 */
    days int64 `json:"days";xml:"days"`
    
    /* tid required主订单号 */
    tid int64 `json:"tid";xml:"tid"`
    
}

func (req *TaobaoTradeReceivetimeDelayRequest) GetAPIName() string {
	return "taobao.trade.receivetime.delay"
}

// TaobaoTradeReceivetimeDelayResponse 延长交易收货时间
type TaobaoTradeReceivetimeDelayResponse struct {
    
    /* trade Object更新后的交易数据，只包括tid和modified两个字段。 */
    trade Trade `json:"trade";xml:"trade"`
    
}
