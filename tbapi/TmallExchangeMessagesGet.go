package tbsdk

// TmallExchangeMessagesGetRequest 查询换货订单留言列表
type TmallExchangeMessagesGetRequest struct {
    
    /* dispute_id required换货单号ID */
    dispute_id int64 `json:"dispute_id";xml:"dispute_id"`
    
    /* fields required返回的字段。具体包括：id,refund_id,owner_id,owner_nick,owner_role,content,pic_urls,created,message_type */
    fields string `json:"fields";xml:"fields"`
    
    /* operator_roles optional留言创建角色。具体包括：卖家主账户(1)、卖家子账户(2)、小二(3)、买家(4)、系统(5)、系统超时(6) */
    operator_roles int64 `json:"operator_roles";xml:"operator_roles"`
    
    /* page_no optional页码 */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size optional每页条数 */
    page_size int64 `json:"page_size";xml:"page_size"`
    
}

func (req *TmallExchangeMessagesGetRequest) GetAPIName() string {
	return "tmall.exchange.messages.get"
}

// TmallExchangeMessagesGetResponse 查询换货订单留言列表
type TmallExchangeMessagesGetResponse struct {
    
    /* result Object返回结果 */
    result RefundMessageResult `json:"result";xml:"result"`
    
}
