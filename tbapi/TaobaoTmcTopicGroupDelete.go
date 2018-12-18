package tbsdk

// TaobaoTmcTopicGroupDeleteRequest 删除根据topic名称路由消息到不同的分组关系
type TaobaoTmcTopicGroupDeleteRequest struct {
    
    /* group_id optional消息分组Id，一般不用填写，如果分组已经被删除，则根据问题排查工具返回的ID删除路由关系 */
    group_id int64 `json:"group_id";xml:"group_id"`
    
    /* group_name required消息分组名 */
    group_name string `json:"group_name";xml:"group_name"`
    
    /* topics required消息topic名称，多个以逗号(,)分割 */
    topics string `json:"topics";xml:"topics"`
    
}

func (req *TaobaoTmcTopicGroupDeleteRequest) GetAPIName() string {
	return "taobao.tmc.topic.group.delete"
}

// TaobaoTmcTopicGroupDeleteResponse 删除根据topic名称路由消息到不同的分组关系
type TaobaoTmcTopicGroupDeleteResponse struct {
    
    /* result Basictrue */
    result bool `json:"result";xml:"result"`
    
}
