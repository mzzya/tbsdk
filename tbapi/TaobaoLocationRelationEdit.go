package tbsdk

// TaobaoLocationRelationEditRequest 地点关联关系增量编辑
type TaobaoLocationRelationEditRequest struct {
    
    /* location_relation_list required关系对象列表 */
    location_relation_list LocationRelationDto `json:"location_relation_list";xml:"location_relation_list"`
    
}

func (req *TaobaoLocationRelationEditRequest) GetAPIName() string {
	return "taobao.location.relation.edit"
}

// TaobaoLocationRelationEditResponse 地点关联关系增量编辑
type TaobaoLocationRelationEditResponse struct {
    
    /* error_msg Basic错误信息 */
    error_msg string `json:"error_msg";xml:"error_msg"`
    
    /* errorcode Basic错误码 */
    errorcode string `json:"errorcode";xml:"errorcode"`
    
    /* is_success Basic是否成功 */
    is_success bool `json:"is_success";xml:"is_success"`
    
}
