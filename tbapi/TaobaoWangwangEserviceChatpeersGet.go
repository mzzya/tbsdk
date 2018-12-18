package tbsdk

// TaobaoWangwangEserviceChatpeersGetRequest 获取聊天对象列表，只能获取最近1个月内的数据且查询时间段<=7天,只支持xml返回 。
type TaobaoWangwangEserviceChatpeersGetRequest struct {
    
}

func (req *TaobaoWangwangEserviceChatpeersGetRequest) GetAPIName() string {
	return "taobao.wangwang.eservice.chatpeers.get"
}

// TaobaoWangwangEserviceChatpeersGetResponse 获取聊天对象列表，只能获取最近1个月内的数据且查询时间段<=7天,只支持xml返回 。
type TaobaoWangwangEserviceChatpeersGetResponse struct {
    
    /* chatpeers Object Array聊天对象ID列表 */
    chatpeers Chatpeer `json:"chatpeers";xml:"chatpeers"`
    
    /* count Basic成员数目。 */
    count int64 `json:"count";xml:"count"`
    
    /* ret Basic返回码： 
10000:成功； 

60000：时间非法，包括1)时间段超过7天,或2)起始时间距今超过30天，或3)时间格式不对； 

50000：聊天用户ID不是该店铺的帐号； 

40000：系统错误，包括必填参数为空。 */
    ret int64 `json:"ret";xml:"ret"`
    
}
