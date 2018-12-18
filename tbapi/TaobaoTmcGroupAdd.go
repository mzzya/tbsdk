package tbsdk

// TaobaoTmcGroupAddRequest 为已开通用户添加用户分组
type TaobaoTmcGroupAddRequest struct {
    
    /* group_name required分组名称，同一个应用下需要保证唯一性，最长32个字符。添加分组后，消息通道会为用户的消息分配独立分组，但之前的消息还是存储于默认分组中。不能以default开头，default开头为系统默认组。 */
    group_name string `json:"group_name";xml:"group_name"`
    
    /* nicks required用户昵称列表，以半角逗号分隔，支持子账号，支持增量添加用户 */
    nicks string `json:"nicks";xml:"nicks"`
    
    /* user_platform optional用户所属于的平台类型，tbUIC:淘宝用户; icbu: icbu用户;ae:ae用户 */
    user_platform string `json:"user_platform";xml:"user_platform"`
    
}

func (req *TaobaoTmcGroupAddRequest) GetAPIName() string {
	return "taobao.tmc.group.add"
}

// TaobaoTmcGroupAddResponse 为已开通用户添加用户分组
type TaobaoTmcGroupAddResponse struct {
    
    /* created Basic创建时间 */
    created Date `json:"created";xml:"created"`
    
    /* group_name Basic分组名称 */
    group_name string `json:"group_name";xml:"group_name"`
    
}
