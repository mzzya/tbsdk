package tbsdk

// TaobaoTmcTopicGroupAddRequest 根据topic名称路由消息到不同的分组。（前提：发送方未指定分组名）
如果是需要授权的消息，分组路由先判断用户分组路由(使用taobao.tmc.group.add添加的路由)，用户分组路由不存在时，才会判断topic分组路由
type TaobaoTmcTopicGroupAddRequest struct {
    
    /* group_name required消息分组名，如果不存在，会自动创建 */
    group_name string `json:"group_name";xml:"group_name"`
    
    /* topics required消息topic名称，多个以逗号(,)分割 */
    topics string `json:"topics";xml:"topics"`
    
}

func (req *TaobaoTmcTopicGroupAddRequest) GetAPIName() string {
	return "taobao.tmc.topic.group.add"
}

// TaobaoTmcTopicGroupAddResponse 根据topic名称路由消息到不同的分组。（前提：发送方未指定分组名）
如果是需要授权的消息，分组路由先判断用户分组路由(使用taobao.tmc.group.add添加的路由)，用户分组路由不存在时，才会判断topic分组路由
type TaobaoTmcTopicGroupAddResponse struct {
    
    /* result Basictrue */
    result bool `json:"result";xml:"result"`
    
}
