package tbsdk

// TaobaoTmcMessagesConsumeRequest 消费多条消息。消费时如果没有返回消息，建议做控制，不要一直调api，浪费应用的流量。如对程序做好优化，若没有消息则，sleep 100ms 等。
type TaobaoTmcMessagesConsumeRequest struct {
    
    /* group_name optional用户分组名称，不传表示消费默认分组，如果应用没有设置用户分组，传入分组名称将会返回错误 */
    group_name string `json:"group_name";xml:"group_name"`
    
    /* quantity optional每次批量消费消息的条数 */
    quantity int64 `json:"quantity";xml:"quantity"`
    
}

func (req *TaobaoTmcMessagesConsumeRequest) GetAPIName() string {
	return "taobao.tmc.messages.consume"
}

// TaobaoTmcMessagesConsumeResponse 消费多条消息。消费时如果没有返回消息，建议做控制，不要一直调api，浪费应用的流量。如对程序做好优化，若没有消息则，sleep 100ms 等。
type TaobaoTmcMessagesConsumeResponse struct {
    
    /* messages Object Array消息列表 */
    messages TmcMessage `json:"messages";xml:"messages"`
    
}
