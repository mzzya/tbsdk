package tbsdk

// TmallExchangeRefusereasonGetRequest 获取拒绝换货原因列表
type TmallExchangeRefusereasonGetRequest struct {
    
    /* dispute_id required换货单号ID */
    dispute_id int64 `json:"dispute_id";xml:"dispute_id"`
    
    /* dispute_type optional换货申请类型：0-任意类型；1-售中；2-售后 */
    dispute_type int64 `json:"dispute_type";xml:"dispute_type"`
    
    /* fields required返回字段 */
    fields string `json:"fields";xml:"fields"`
    
}

func (req *TmallExchangeRefusereasonGetRequest) GetAPIName() string {
	return "tmall.exchange.refusereason.get"
}

// TmallExchangeRefusereasonGetResponse 获取拒绝换货原因列表
type TmallExchangeRefusereasonGetResponse struct {
    
    /* result Object返回结果 */
    result ResultSet `json:"result";xml:"result"`
    
}
