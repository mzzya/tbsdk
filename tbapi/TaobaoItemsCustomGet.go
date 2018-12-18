package tbsdk

// TaobaoItemsCustomGetRequest 跟据卖家设定的商品外部id获取商品 
这个商品对应卖家从传入的session中获取，需要session绑定
type TaobaoItemsCustomGetRequest struct {
    
    /* fields required需返回的字段列表，参考：Item商品结构体说明，其中barcode、sku.barcode等条形码字段暂不支持；多个字段之间用“,”分隔。 */
    fields string `json:"fields";xml:"fields"`
    
    /* outer_id required商品的外部商品ID，支持批量，最多不超过40个。 */
    outer_id string `json:"outer_id";xml:"outer_id"`
    
}

func (req *TaobaoItemsCustomGetRequest) GetAPIName() string {
	return "taobao.items.custom.get"
}

// TaobaoItemsCustomGetResponse 跟据卖家设定的商品外部id获取商品 
这个商品对应卖家从传入的session中获取，需要session绑定
type TaobaoItemsCustomGetResponse struct {
    
    /* items Object Array商品列表，具体返回字段以fields决定 */
    items Item `json:"items";xml:"items"`
    
}
