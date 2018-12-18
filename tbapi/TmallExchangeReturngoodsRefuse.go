package tbsdk

// TmallExchangeReturngoodsRefuseRequest 卖家拒绝买家换货申请
type TmallExchangeReturngoodsRefuseRequest struct {
    
    /* dispute_id required换货单号ID */
    dispute_id int64 `json:"dispute_id";xml:"dispute_id"`
    
    /* leave_message optional留言说明 */
    leave_message string `json:"leave_message";xml:"leave_message"`
    
    /* leave_message_pics optional凭证图片 */
    leave_message_pics byte[] `json:"leave_message_pics";xml:"leave_message_pics"`
    
    /* seller_refuse_reason_id required拒绝原因ID */
    seller_refuse_reason_id int64 `json:"seller_refuse_reason_id";xml:"seller_refuse_reason_id"`
    
}

func (req *TmallExchangeReturngoodsRefuseRequest) GetAPIName() string {
	return "tmall.exchange.returngoods.refuse"
}

// TmallExchangeReturngoodsRefuseResponse 卖家拒绝买家换货申请
type TmallExchangeReturngoodsRefuseResponse struct {
    
    /* result Object返回结果 */
    result ExchangeBaseResponse `json:"result";xml:"result"`
    
}
