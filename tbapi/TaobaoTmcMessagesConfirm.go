package tbsdk

// TaobaoTmcMessagesConfirmRequest 确认消费消息的状态
type TaobaoTmcMessagesConfirmRequest struct {
    
    /* f_message_ids optional处理失败的消息ID列表--已废弃，无需传此字段 */
    f_message_ids int64 `json:"f_message_ids";xml:"f_message_ids"`
    
    /* group_name optional分组名称，不传代表默认分组 */
    group_name string `json:"group_name";xml:"group_name"`
    
    /* s_message_ids required处理成功的消息ID列表 最大 200个ID */
    s_message_ids int64 `json:"s_message_ids";xml:"s_message_ids"`
    
}

func (req *TaobaoTmcMessagesConfirmRequest) GetAPIName() string {
	return "taobao.tmc.messages.confirm"
}

// TaobaoTmcMessagesConfirmResponse 确认消费消息的状态
type TaobaoTmcMessagesConfirmResponse struct {
    
    /* is_success Basic是否成功 */
    is_success bool `json:"is_success";xml:"is_success"`
    
}
