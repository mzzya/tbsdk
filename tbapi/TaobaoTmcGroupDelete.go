package tbsdk

// TaobaoTmcGroupDeleteRequest 删除指定的分组或分组下的用户
type TaobaoTmcGroupDeleteRequest struct {
    
    /* group_name required分组名称，分组删除后，用户的消息将会存储于默认分组中。警告：由于分组已经删除，用户之前未消费的消息将无法再获取。不能以default开头，default开头为系统默认组。 */
    group_name string `json:"group_name";xml:"group_name"`
    
    /* nicks optional用户列表，不传表示删除整个分组，如果用户全部删除后，也会自动删除整个分组 */
    nicks string `json:"nicks";xml:"nicks"`
    
    /* user_platform optional用户所属于的平台类型，tbUIC:淘宝用户; icbu: icbu用户;ae:ae用户 */
    user_platform string `json:"user_platform";xml:"user_platform"`
    
}

func (req *TaobaoTmcGroupDeleteRequest) GetAPIName() string {
	return "taobao.tmc.group.delete"
}

// TaobaoTmcGroupDeleteResponse 删除指定的分组或分组下的用户
type TaobaoTmcGroupDeleteResponse struct {
    
    /* is_success Basic是否成功 */
    is_success bool `json:"is_success";xml:"is_success"`
    
}
