package tbsdk

// TaobaoSkusCustomGetRequest 跟据卖家设定的Sku的外部id获取商品，如果一个outer_id对应多个Sku会返回所有符合条件的sku 
这个Sku所属卖家从传入的session中获取，需要session绑定(注：iid标签里是num_iid的值，可以用作num_iid使用)
type TaobaoSkusCustomGetRequest struct {
    
    /* fields required需返回的字段列表。可选值：Sku结构体中的所有字段；字段之间用“,”隔开 */
    fields string `json:"fields";xml:"fields"`
    
    /* outer_id requiredSku的外部商家ID */
    outer_id string `json:"outer_id";xml:"outer_id"`
    
}

func (req *TaobaoSkusCustomGetRequest) GetAPIName() string {
	return "taobao.skus.custom.get"
}

// TaobaoSkusCustomGetResponse 跟据卖家设定的Sku的外部id获取商品，如果一个outer_id对应多个Sku会返回所有符合条件的sku 
这个Sku所属卖家从传入的session中获取，需要session绑定(注：iid标签里是num_iid的值，可以用作num_iid使用)
type TaobaoSkusCustomGetResponse struct {
    
    /* skus Object ArraySku对象，具体字段以fields决定 */
    skus Sku `json:"skus";xml:"skus"`
    
}
