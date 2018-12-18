package tbsdk

// TaobaoWlbNotifyMessageConfirmRequest 确认物流宝可执行消息
type TaobaoWlbNotifyMessageConfirmRequest struct {
    
    /* message_id required物流宝通知消息的id，通过taobao.wlb.notify.message.page.get接口得到的WlbMessage数据结构中的id字段 */
    message_id string `json:"message_id";xml:"message_id"`
    
}

func (req *TaobaoWlbNotifyMessageConfirmRequest) GetAPIName() string {
	return "taobao.wlb.notify.message.confirm"
}

// TaobaoWlbNotifyMessageConfirmResponse 确认物流宝可执行消息
type TaobaoWlbNotifyMessageConfirmResponse struct {
    
    /* gmt_modified Basic物流宝消息确认时间 */
    gmt_modified Date `json:"gmt_modified";xml:"gmt_modified"`
    
}
