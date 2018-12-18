package tbsdk

// TaobaoSubuserDutysGetRequest 通过主账号Nick获取该账户下的所有职务信息，职务信息中包括职务ID、职务名称以及职务等级（通过主账号登陆只能获取属于该主账号下的职务信息）
type TaobaoSubuserDutysGetRequest struct {
    
    /* user_nick required主账号用户名 */
    user_nick string `json:"user_nick";xml:"user_nick"`
    
}

func (req *TaobaoSubuserDutysGetRequest) GetAPIName() string {
	return "taobao.subuser.dutys.get"
}

// TaobaoSubuserDutysGetResponse 通过主账号Nick获取该账户下的所有职务信息，职务信息中包括职务ID、职务名称以及职务等级（通过主账号登陆只能获取属于该主账号下的职务信息）
type TaobaoSubuserDutysGetResponse struct {
    
    /* dutys Object Array职务信息 */
    dutys Duty `json:"dutys";xml:"dutys"`
    
}
