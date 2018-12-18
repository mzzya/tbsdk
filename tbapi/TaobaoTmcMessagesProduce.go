package tbsdk

// TaobaoTmcMessagesProduceRequest 批量发送消息
type TaobaoTmcMessagesProduceRequest struct {
    
    /* messages requiredtmc消息列表, 最多50条，元素结构与taobao.tmc.message.produce一致，用json表示的消息列表。例如：[{"content": "{\"tid\":1234554321,\"status\":\"X_LOGISTICS_PRINTED\",\"action_time\":\"2014-08-08 18:24:00\",\"seller_nick\": \"向阳aa\",\"operator\":\"小张\"}","topic": "taobao_jds_TradeTrace"},{"content": "{\"tid\":1234554321,\"status\":\"X_LOGISTICS_PRINTED\",\"action_time\":\"2014-08-08 18:24:00\",\"seller_nick\": \"向阳aa\",\"operator\":\"小张\"}","topic": "taobao_jds_TradeTrace"}] */
    messages TmcPublishMessage `json:"messages";xml:"messages"`
    
}

func (req *TaobaoTmcMessagesProduceRequest) GetAPIName() string {
	return "taobao.tmc.messages.produce"
}

// TaobaoTmcMessagesProduceResponse 批量发送消息
type TaobaoTmcMessagesProduceResponse struct {
    
    /* is_all_success Basic是否全部成功 */
    is_all_success bool `json:"is_all_success";xml:"is_all_success"`
    
    /* results Object Array发送结果，与发送时的参数顺序一致。如果is_all_success为true时，不用校验result是否成功 */
    results TmcProduceResult `json:"results";xml:"results"`
    
}
