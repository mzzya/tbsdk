package tbsdk

// TaobaoTmcGroupsGetRequest 获取自定义用户分组列表
type TaobaoTmcGroupsGetRequest struct {
    
    /* group_names optional要查询分组的名称，多个分组用半角逗号分隔，不传代表查询所有分组信息，但不会返回组下面的用户信息。如果应用没有设置分组则返回空。组名不能以default开头，default开头是系统默认的组。 */
    group_names string `json:"group_names";xml:"group_names"`
    
    /* page_no optional页码 */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size optional每页返回多少个分组 */
    page_size int64 `json:"page_size";xml:"page_size"`
    
}

func (req *TaobaoTmcGroupsGetRequest) GetAPIName() string {
	return "taobao.tmc.groups.get"
}

// TaobaoTmcGroupsGetResponse 获取自定义用户分组列表
type TaobaoTmcGroupsGetResponse struct {
    
    /* groups Object Arraydasdasd */
    groups TmcGroup `json:"groups";xml:"groups"`
    
    /* total_results Basic分组总数 */
    total_results int64 `json:"total_results";xml:"total_results"`
    
}
