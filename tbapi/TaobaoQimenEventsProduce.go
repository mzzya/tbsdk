package tbsdk

// TaobaoQimenEventsProduceRequest 批量发送消息
type TaobaoQimenEventsProduceRequest struct {
    
    /* messages required奇门事件列表, 最多50条 */
    messages QimenEvent `json:"messages";xml:"messages"`
    
}

func (req *TaobaoQimenEventsProduceRequest) GetAPIName() string {
	return "taobao.qimen.events.produce"
}

// TaobaoQimenEventsProduceResponse 批量发送消息
type TaobaoQimenEventsProduceResponse struct {
    
    /* is_all_success Basic是否全部成功 */
    is_all_success bool `json:"is_all_success";xml:"is_all_success"`
    
    /* results Object Array发送结果，与发送时的参数顺序一致。如果is_all_success为true时，不用校验result是否成功 */
    results QimenResult `json:"results";xml:"results"`
    
}
