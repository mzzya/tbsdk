package tbsdk

// TmallExchangeMessageAddRequest 卖家创建换货留言
type TmallExchangeMessageAddRequest struct {
    
    /* content required留言内容 */
    content string `json:"content";xml:"content"`
    
    /* dispute_id required换货单号ID */
    dispute_id int64 `json:"dispute_id";xml:"dispute_id"`
    
    /* fields required返回字段。目前支持id,refund_id,owner_id,owner_nick,owner_role,content,pic_urls,created,message_type */
    fields string `json:"fields";xml:"fields"`
    
    /* message_pics optional凭证图片列表 */
    message_pics byte[] `json:"message_pics";xml:"message_pics"`
    
}

func (req *TmallExchangeMessageAddRequest) GetAPIName() string {
	return "tmall.exchange.message.add"
}

// TmallExchangeMessageAddResponse 卖家创建换货留言
type TmallExchangeMessageAddResponse struct {
    
    /* result Object返回结果 */
    result ResultSet `json:"result";xml:"result"`
    
}
