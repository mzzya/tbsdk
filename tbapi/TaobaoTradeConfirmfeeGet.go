package tbsdk

// TaobaoTradeConfirmfeeGetRequest 获取交易确认收货费用，可以获取主订单或子订单的确认收货费用
type TaobaoTradeConfirmfeeGetRequest struct {
    
    /* tid required交易主订单或子订单ID */
    tid int64 `json:"tid";xml:"tid"`
    
}

func (req *TaobaoTradeConfirmfeeGetRequest) GetAPIName() string {
	return "taobao.trade.confirmfee.get"
}

// TaobaoTradeConfirmfeeGetResponse 获取交易确认收货费用，可以获取主订单或子订单的确认收货费用
type TaobaoTradeConfirmfeeGetResponse struct {
    
    /* trade_confirm_fee Object获取到的交易确认收货费用 */
    trade_confirm_fee TradeConfirmFee `json:"trade_confirm_fee";xml:"trade_confirm_fee"`
    
}
