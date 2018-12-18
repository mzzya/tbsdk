package tbsdk

// TmallExchangeGetRequest 获取单笔换货详情
type TmallExchangeGetRequest struct {
    
    /* dispute_id required换货单号ID */
    dispute_id int64 `json:"dispute_id";xml:"dispute_id"`
    
    /* fields required返回字段。目前支持dispute_id, bizorder_id, num, buyer_nick, status, created, modified, reason, title, buyer_logistic_no, seller_logistic_no, bought_sku, exchange_sku, buyer_address, address, buyer_phone, buyer_logistic_name, seller_logistic_name, alipay_no, buyer_name, seller_nick */
    fields string `json:"fields";xml:"fields"`
    
}

func (req *TmallExchangeGetRequest) GetAPIName() string {
	return "tmall.exchange.get"
}

// TmallExchangeGetResponse 获取单笔换货详情
type TmallExchangeGetResponse struct {
    
    /* result Object返回结果 */
    result ExchangeBaseResponse `json:"result";xml:"result"`
    
}
