package tbsdk

// TaobaoTraderateImprImprwordByfeedidGetRequest 根据卖家nick和交易id（如果是子订单，输入子订单id），查询到该条评价的大家印象标签
type TaobaoTraderateImprImprwordByfeedidGetRequest struct {
    
    /* child_trade_id required交易订单id号（如果包含子订单，请使用子订单id号） */
    child_trade_id int64 `json:"child_trade_id";xml:"child_trade_id"`
    
}

func (req *TaobaoTraderateImprImprwordByfeedidGetRequest) GetAPIName() string {
	return "taobao.traderate.impr.imprword.byfeedid.get"
}

// TaobaoTraderateImprImprwordByfeedidGetResponse 根据卖家nick和交易id（如果是子订单，输入子订单id），查询到该条评价的大家印象标签
type TaobaoTraderateImprImprwordByfeedidGetResponse struct {
    
    /* impr_feed Object根据子订单和买家昵称找到的评价和印象词结果 */
    impr_feed ImprFeedIdDO `json:"impr_feed";xml:"impr_feed"`
    
}
