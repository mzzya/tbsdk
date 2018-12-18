package tbsdk

// TaobaoSellercenterRoleInfoGetRequest 获取指定角色的信息。只能查询属于自己的角色信息 (主账号或者某个主账号的子账号登陆，只能查询属于该主账号的角色信息)
type TaobaoSellercenterRoleInfoGetRequest struct {
    
    /* role_id required角色id */
    role_id int64 `json:"role_id";xml:"role_id"`
    
}

func (req *TaobaoSellercenterRoleInfoGetRequest) GetAPIName() string {
	return "taobao.sellercenter.role.info.get"
}

// TaobaoSellercenterRoleInfoGetResponse 获取指定角色的信息。只能查询属于自己的角色信息 (主账号或者某个主账号的子账号登陆，只能查询属于该主账号的角色信息)
type TaobaoSellercenterRoleInfoGetResponse struct {
    
    /* role Object角色具体信息 */
    role Role `json:"role";xml:"role"`
    
}
