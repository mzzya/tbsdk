package tbsdk

// TaobaoTmcQueueGetRequest 根据appkey和groupName获取消息队列积压情况
type TaobaoTmcQueueGetRequest struct {
    
    /* group_name requiredTMC组名 */
    group_name string `json:"group_name";xml:"group_name"`
    
}

func (req *TaobaoTmcQueueGetRequest) GetAPIName() string {
	return "taobao.tmc.queue.get"
}

// TaobaoTmcQueueGetResponse 根据appkey和groupName获取消息队列积压情况
type TaobaoTmcQueueGetResponse struct {
    
    /* datas Object Array队列详细信息 */
    datas TmcQueueInfo `json:"datas";xml:"datas"`
    
}
