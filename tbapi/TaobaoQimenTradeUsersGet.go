package tbsdk

// TaobaoQimenTradeUsersGetRequest 获取已开通奇门订单服务的用户列表
type TaobaoQimenTradeUsersGetRequest struct {
    
    /* page_index required每页的数量 */
    page_index int64 `json:"page_index";xml:"page_index"`
    
    /* page_size required页数 */
    page_size int64 `json:"page_size";xml:"page_size"`
    
}

func (req *TaobaoQimenTradeUsersGetRequest) GetAPIName() string {
	return "taobao.qimen.trade.users.get"
}

// TaobaoQimenTradeUsersGetResponse 获取已开通奇门订单服务的用户列表
type TaobaoQimenTradeUsersGetResponse struct {
    
    /* total_count BasictotalCount */
    total_count int64 `json:"total_count";xml:"total_count"`
    
    /* users Object Arraymodal */
    users QimenUser `json:"users";xml:"users"`
    
}
