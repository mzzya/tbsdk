package tbsdk

// TmallExchangeRefuseRequest 卖家拒绝换货申请
type TmallExchangeRefuseRequest struct {
    
    /* dispute_id required换货单号ID */
    dispute_id int64 `json:"dispute_id";xml:"dispute_id"`
    
    /* fields required返回字段。目前支持dispute_id, bizorder_id, modified, status */
    fields string `json:"fields";xml:"fields"`
    
    /* leave_message optional拒绝换货申请时的留言 */
    leave_message string `json:"leave_message";xml:"leave_message"`
    
    /* leave_message_pics optional凭证图片 */
    leave_message_pics byte[] `json:"leave_message_pics";xml:"leave_message_pics"`
    
    /* seller_refuse_reason_id required换货原因对应ID */
    seller_refuse_reason_id int64 `json:"seller_refuse_reason_id";xml:"seller_refuse_reason_id"`
    
}

func (req *TmallExchangeRefuseRequest) GetAPIName() string {
	return "tmall.exchange.refuse"
}

// TmallExchangeRefuseResponse 卖家拒绝换货申请
type TmallExchangeRefuseResponse struct {
    
    /* result Object返回结果 */
    result ExchangeBaseResponse `json:"result";xml:"result"`
    
}
