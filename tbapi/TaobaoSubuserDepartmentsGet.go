package tbsdk

// TaobaoSubuserDepartmentsGetRequest 获取指定账户的所有部门列表，其实包括有每个部门的ID、父部门ID、部门名称（通过主账号登陆只能查询属于该主账号下的所有部门信息）。
type TaobaoSubuserDepartmentsGetRequest struct {
    
    /* user_nick required主账号用户名 */
    user_nick string `json:"user_nick";xml:"user_nick"`
    
}

func (req *TaobaoSubuserDepartmentsGetRequest) GetAPIName() string {
	return "taobao.subuser.departments.get"
}

// TaobaoSubuserDepartmentsGetResponse 获取指定账户的所有部门列表，其实包括有每个部门的ID、父部门ID、部门名称（通过主账号登陆只能查询属于该主账号下的所有部门信息）。
type TaobaoSubuserDepartmentsGetResponse struct {
    
    /* departments Object Array部门信息 */
    departments Department `json:"departments";xml:"departments"`
    
}
