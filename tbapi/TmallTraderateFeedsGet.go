package tbsdk

// TmallTraderateFeedsGetRequest 通过子订单ID获取天猫订单对应的评价，追评，以及对应的语义标签
type TmallTraderateFeedsGetRequest struct {
    
    /* child_trade_id required交易子订单ID */
    child_trade_id int64 `json:"child_trade_id";xml:"child_trade_id"`
    
}

func (req *TmallTraderateFeedsGetRequest) GetAPIName() string {
	return "tmall.traderate.feeds.get"
}

// TmallTraderateFeedsGetResponse 通过子订单ID获取天猫订单对应的评价，追评，以及对应的语义标签
type TmallTraderateFeedsGetResponse struct {
    
    /* tmall_rate_info Object返回评价信息 */
    tmall_rate_info TmallRateInfo `json:"tmall_rate_info";xml:"tmall_rate_info"`
    
}
