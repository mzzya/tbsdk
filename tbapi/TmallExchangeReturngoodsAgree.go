package tbsdk

// TmallExchangeReturngoodsAgreeRequest 卖家确认收货
type TmallExchangeReturngoodsAgreeRequest struct {
    
    /* dispute_id required换货单号ID */
    dispute_id int64 `json:"dispute_id";xml:"dispute_id"`
    
    /* fields required返回字段。目前支持dispute_id（换货单号ID）,bizorder_id（正向交易单号ID）, modified（订单修改时间）, status（当前换货状态） */
    fields string `json:"fields";xml:"fields"`
    
}

func (req *TmallExchangeReturngoodsAgreeRequest) GetAPIName() string {
	return "tmall.exchange.returngoods.agree"
}

// TmallExchangeReturngoodsAgreeResponse 卖家确认收货
type TmallExchangeReturngoodsAgreeResponse struct {
    
    /* result Object返回结果 */
    result ExchangeBaseResponse `json:"result";xml:"result"`
    
}
