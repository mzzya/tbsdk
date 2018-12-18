package tbsdk

// TaobaoInventoryQueryRequest 商家查询商品总体库存信息
type TaobaoInventoryQueryRequest struct {
    
    /* sc_item_codes optional后端商品的商家编码列表，控制到50个 */
    sc_item_codes string `json:"sc_item_codes";xml:"sc_item_codes"`
    
    /* sc_item_ids required后端商品ID 列表，控制到50个 */
    sc_item_ids string `json:"sc_item_ids";xml:"sc_item_ids"`
    
    /* seller_nick optional卖家昵称 */
    seller_nick string `json:"seller_nick";xml:"seller_nick"`
    
    /* store_codes optional仓库列表:GLY001^GLY002 */
    store_codes string `json:"store_codes";xml:"store_codes"`
    
}

func (req *TaobaoInventoryQueryRequest) GetAPIName() string {
	return "taobao.inventory.query"
}

// TaobaoInventoryQueryResponse 商家查询商品总体库存信息
type TaobaoInventoryQueryResponse struct {
    
    /* item_inventorys Object Array商品总体库存信息 */
    item_inventorys InventorySum `json:"item_inventorys";xml:"item_inventorys"`
    
    /* tip_infos Object Array提示信息，提示不存在的后端商品 */
    tip_infos TipInfo `json:"tip_infos";xml:"tip_infos"`
    
}
