package tbsdk

// TmallExchangeConsigngoodsRequest 卖家发货
type TmallExchangeConsigngoodsRequest struct {
    
    /* dispute_id required换货单号ID */
    dispute_id int64 `json:"dispute_id";xml:"dispute_id"`
    
    /* fields required返回字段 */
    fields string `json:"fields";xml:"fields"`
    
    /* logistics_company_name required卖家发货的快递公司 */
    logistics_company_name string `json:"logistics_company_name";xml:"logistics_company_name"`
    
    /* logistics_no required卖家发货的物流单号 */
    logistics_no string `json:"logistics_no";xml:"logistics_no"`
    
    /* logistics_type required卖家发货的物流类型，100表示平邮，200表示快递 */
    logistics_type int64 `json:"logistics_type";xml:"logistics_type"`
    
}

func (req *TmallExchangeConsigngoodsRequest) GetAPIName() string {
	return "tmall.exchange.consigngoods"
}

// TmallExchangeConsigngoodsResponse 卖家发货
type TmallExchangeConsigngoodsResponse struct {
    
    /* result Objectresult */
    result RefundBaseResponse `json:"result";xml:"result"`
    
}
