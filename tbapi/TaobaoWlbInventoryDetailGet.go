package tbsdk

// TaobaoWlbInventoryDetailGetRequest 查询库存详情，通过商品ID获取发送请求的卖家的库存详情
type TaobaoWlbInventoryDetailGetRequest struct {
    
    /* inventory_type_list optional库存类型列表，值包括：
VENDIBLE--可销售库存
FREEZE--冻结库存
ONWAY--在途库存
DEFECT--残次品库存
ENGINE_DAMAGE--机损
BOX_DAMAGE--箱损
EXPIRATION--过保 */
    inventory_type_list string `json:"inventory_type_list";xml:"inventory_type_list"`
    
    /* item_id required商品ID */
    item_id int64 `json:"item_id";xml:"item_id"`
    
    /* store_code optional仓库编码 */
    store_code string `json:"store_code";xml:"store_code"`
    
}

func (req *TaobaoWlbInventoryDetailGetRequest) GetAPIName() string {
	return "taobao.wlb.inventory.detail.get"
}

// TaobaoWlbInventoryDetailGetResponse 查询库存详情，通过商品ID获取发送请求的卖家的库存详情
type TaobaoWlbInventoryDetailGetResponse struct {
    
    /* inventory_list Object Array库存详情对象。其中包括货主ID，仓库编码，库存，库存类型等属性 */
    inventory_list WlbInventory `json:"inventory_list";xml:"inventory_list"`
    
    /* item_id Basic入参的item_id */
    item_id int64 `json:"item_id";xml:"item_id"`
    
}
