package tbsdk

// TaobaoTmcUserGetRequest 查询指定用户开通的消息通道和组
type TaobaoTmcUserGetRequest struct {
    
    /* fields required需返回的字段列表，多个字段以半角逗号分隔。可选值：TmcUser结构体中的所有字段，一定要返回topic。 */
    fields string `json:"fields";xml:"fields"`
    
    /* nick required用户昵称 */
    nick string `json:"nick";xml:"nick"`
    
    /* user_platform optional用户所属的平台类型，tbUIC:淘宝用户; icbu: icbu用户;ae:ae用户 */
    user_platform string `json:"user_platform";xml:"user_platform"`
    
}

func (req *TaobaoTmcUserGetRequest) GetAPIName() string {
	return "taobao.tmc.user.get"
}

// TaobaoTmcUserGetResponse 查询指定用户开通的消息通道和组
type TaobaoTmcUserGetResponse struct {
    
    /* tmc_user Object开通的用户数据 */
    tmc_user TmcUser `json:"tmc_user";xml:"tmc_user"`
    
}
