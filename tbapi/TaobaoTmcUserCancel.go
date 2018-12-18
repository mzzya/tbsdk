package tbsdk

// TaobaoTmcUserCancelRequest 取消用户的消息服务
type TaobaoTmcUserCancelRequest struct {
    
    /* nick required用户昵称 */
    nick string `json:"nick";xml:"nick"`
    
    /* user_platform optional用户所属的平台类型，tbUIC:淘宝用户; icbu: icbu用户;ae:ae用户 */
    user_platform string `json:"user_platform";xml:"user_platform"`
    
}

func (req *TaobaoTmcUserCancelRequest) GetAPIName() string {
	return "taobao.tmc.user.cancel"
}

// TaobaoTmcUserCancelResponse 取消用户的消息服务
type TaobaoTmcUserCancelResponse struct {
    
    /* is_success Basic是否成功,如果为false并且没有错误码，表示删除的用户不存在。 */
    is_success bool `json:"is_success";xml:"is_success"`
    
}
