package tbsdk

// TaobaoUserSellerGetRequest 查询卖家用户信息（只能查询有店铺的用户） 只能卖家类应用调用。
type TaobaoUserSellerGetRequest struct {
    
    /* fields required需要返回的字段列表，可选值为返回示例值中的可以看到的字段 */
    fields string `json:"fields";xml:"fields"`
    
}

func (req *TaobaoUserSellerGetRequest) GetAPIName() string {
	return "taobao.user.seller.get"
}

// TaobaoUserSellerGetResponse 查询卖家用户信息（只能查询有店铺的用户） 只能卖家类应用调用。
type TaobaoUserSellerGetResponse struct {
    
    /* user Object用户信息 */
    user User `json:"user";xml:"user"`
    
}
