package tbsdk

// TaobaoSubuserInfoUpdateRequest 修改指定账户子账号的基本信息（通过主账号登陆只能修改属于该主账号的子账号基本信息）
type TaobaoSubuserInfoUpdateRequest struct {
    
    /* is_disable_subaccount optional是否停用子账号 true:表示停用该子账号false:表示开启该子账号 */
    is_disable_subaccount bool `json:"is_disable_subaccount";xml:"is_disable_subaccount"`
    
    /* is_dispatch optional子账号是否参与分流 true:参与分流 false:不参与分流 */
    is_dispatch bool `json:"is_dispatch";xml:"is_dispatch"`
    
    /* sub_id required子账号ID */
    sub_id int64 `json:"sub_id";xml:"sub_id"`
    
}

func (req *TaobaoSubuserInfoUpdateRequest) GetAPIName() string {
	return "taobao.subuser.info.update"
}

// TaobaoSubuserInfoUpdateResponse 修改指定账户子账号的基本信息（通过主账号登陆只能修改属于该主账号的子账号基本信息）
type TaobaoSubuserInfoUpdateResponse struct {
    
    /* is_success Basic操作是否成功 true:操作成功; false:操作失败 */
    is_success bool `json:"is_success";xml:"is_success"`
    
}
