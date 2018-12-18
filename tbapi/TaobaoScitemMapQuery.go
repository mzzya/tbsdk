package tbsdk

// TaobaoScitemMapQueryRequest 查找IC商品或分销商品与后端商品的关联信息。skuId如果不传就查找该itemId下所有的sku
type TaobaoScitemMapQueryRequest struct {
    
    /* item_id requiredmap_type为1：前台ic商品id
map_type为2：分销productid */
    item_id int64 `json:"item_id";xml:"item_id"`
    
    /* sku_id optionalmap_type为1：前台ic商品skuid 
map_type为2：分销商品的skuid */
    sku_id int64 `json:"sku_id";xml:"sku_id"`
    
}

func (req *TaobaoScitemMapQueryRequest) GetAPIName() string {
	return "taobao.scitem.map.query"
}

// TaobaoScitemMapQueryResponse 查找IC商品或分销商品与后端商品的关联信息。skuId如果不传就查找该itemId下所有的sku
type TaobaoScitemMapQueryResponse struct {
    
    /* sc_item_maps Object Array后端商品映射列表 */
    sc_item_maps ScItemMap `json:"sc_item_maps";xml:"sc_item_maps"`
    
}
