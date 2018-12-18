package tbsdk

// TaobaoWlbItemGetRequest 根据商品ID获取商品信息,除了获取商品信息外还可获取商品属性信息和库存信息。
type TaobaoWlbItemGetRequest struct {
    
    /* item_id required商品ID */
    item_id int64 `json:"item_id";xml:"item_id"`
    
}

func (req *TaobaoWlbItemGetRequest) GetAPIName() string {
	return "taobao.wlb.item.get"
}

// TaobaoWlbItemGetResponse 根据商品ID获取商品信息,除了获取商品信息外还可获取商品属性信息和库存信息。
type TaobaoWlbItemGetResponse struct {
    
    /* item Object商品信息 */
    item WlbItem `json:"item";xml:"item"`
    
}
