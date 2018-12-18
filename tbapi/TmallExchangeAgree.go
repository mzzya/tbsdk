package tbsdk

// TmallExchangeAgreeRequest 卖家同意换货申请
type TmallExchangeAgreeRequest struct {
    
    /* address_id required收货地址id，如需获取请调用该top接口：taobao.logistics.address.search，对应属性为contact_id */
    address_id int64 `json:"address_id";xml:"address_id"`
    
    /* complete_address optional详细收货地址 */
    complete_address string `json:"complete_address";xml:"complete_address"`
    
    /* dispute_id required换货单号ID */
    dispute_id int64 `json:"dispute_id";xml:"dispute_id"`
    
    /* fields required返回字段。当前支持的有 dispute_id, bizorder_id, modified, status */
    fields string `json:"fields";xml:"fields"`
    
    /* leave_message optional卖家留言 */
    leave_message string `json:"leave_message";xml:"leave_message"`
    
    /* leave_message_pics optional上传图片举证 */
    leave_message_pics byte[] `json:"leave_message_pics";xml:"leave_message_pics"`
    
    /* mobile optional收货人手机号 */
    mobile string `json:"mobile";xml:"mobile"`
    
    /* post optional邮政编码 */
    post string `json:"post";xml:"post"`
    
}

func (req *TmallExchangeAgreeRequest) GetAPIName() string {
	return "tmall.exchange.agree"
}

// TmallExchangeAgreeResponse 卖家同意换货申请
type TmallExchangeAgreeResponse struct {
    
    /* result Object返回结果 */
    result ExchangeBaseResponse `json:"result";xml:"result"`
    
}
