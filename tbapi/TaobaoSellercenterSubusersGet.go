package tbsdk

// TaobaoSellercenterSubusersGetRequest 根据主账号nick查询该账号下所有的子账号列表，只能查询属于自己的账号信息 (主账号以及所属子账号)
type TaobaoSellercenterSubusersGetRequest struct {
    
    /* nick required表示卖家昵称 */
    nick string `json:"nick";xml:"nick"`
    
}

func (req *TaobaoSellercenterSubusersGetRequest) GetAPIName() string {
	return "taobao.sellercenter.subusers.get"
}

// TaobaoSellercenterSubusersGetResponse 根据主账号nick查询该账号下所有的子账号列表，只能查询属于自己的账号信息 (主账号以及所属子账号)
type TaobaoSellercenterSubusersGetResponse struct {
    
    /* subusers Object Array子账号基本信息列表。具体信息为id、子账号用户名、主账号id、主账号昵称、当前状态值、是否分流 */
    subusers SubUserInfo `json:"subusers";xml:"subusers"`
    
}
