package tbsdk

// TaobaoJushitaJdpUsersGetRequest 获取开通的订单同步服务的用户，含有rds的路由关系
type TaobaoJushitaJdpUsersGetRequest struct {
    
    /* end_modified optional此参数一般不用传，用于查询最后更改时间在某个时间段内的用户 */
    end_modified Date `json:"end_modified";xml:"end_modified"`
    
    /* page_no optional当前页数 */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size optional每页记录数，默认200 */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* start_modified optional此参数一般不用传，用于查询最后更改时间在某个时间段内的用户 */
    start_modified Date `json:"start_modified";xml:"start_modified"`
    
    /* target_appkey optional普通isv不能传入此参数 */
    target_appkey string `json:"target_appkey";xml:"target_appkey"`
    
    /* user_id optional如果传了user_id表示单条查询 */
    user_id int64 `json:"user_id";xml:"user_id"`
    
}

func (req *TaobaoJushitaJdpUsersGetRequest) GetAPIName() string {
	return "taobao.jushita.jdp.users.get"
}

// TaobaoJushitaJdpUsersGetResponse 获取开通的订单同步服务的用户，含有rds的路由关系
type TaobaoJushitaJdpUsersGetResponse struct {
    
    /* total_results Basic总记录数 */
    total_results int64 `json:"total_results";xml:"total_results"`
    
    /* users Object Array用户列表 */
    users JdpUser `json:"users";xml:"users"`
    
}
