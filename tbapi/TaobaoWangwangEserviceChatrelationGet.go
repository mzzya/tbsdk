package tbsdk

// TaobaoWangwangEserviceChatrelationGetRequest 获取指定时间区间内的聊天关系（查询账号，和谁，在哪天说过话）。如：
A 和 B 在2016-09-01 和 2016-09-02 都说过话。以A为查询账号，则该接口将返回：
2016-09-01， B
2016-09-02， B
type TaobaoWangwangEserviceChatrelationGetRequest struct {
    
    /* chat_relation_request required请求参数 */
    chat_relation_request ChatRelationRequest `json:"chat_relation_request";xml:"chat_relation_request"`
    
}

func (req *TaobaoWangwangEserviceChatrelationGetRequest) GetAPIName() string {
	return "taobao.wangwang.eservice.chatrelation.get"
}

// TaobaoWangwangEserviceChatrelationGetResponse 获取指定时间区间内的聊天关系（查询账号，和谁，在哪天说过话）。如：
A 和 B 在2016-09-01 和 2016-09-02 都说过话。以A为查询账号，则该接口将返回：
2016-09-01， B
2016-09-02， B
type TaobaoWangwangEserviceChatrelationGetResponse struct {
    
    /* result Objectresult */
    result ChatRelationResult `json:"result";xml:"result"`
    
}
