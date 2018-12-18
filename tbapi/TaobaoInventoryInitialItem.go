package tbsdk

// TaobaoInventoryInitialItemRequest 商家仓商品初始化在各个仓中库存
type TaobaoInventoryInitialItemRequest struct {
    
    /* sc_item_id required后端商品id */
    sc_item_id int64 `json:"sc_item_id";xml:"sc_item_id"`
    
    /* store_inventorys required商品初始库存信息： [{"storeCode":"必选,商家仓库编号","inventoryType":"可选，库存类型 1：正常,2：损坏,3：冻结,10：质押,11-20:用户自定义,默认为1","quantity":"必选,数量"}] */
    store_inventorys string `json:"store_inventorys";xml:"store_inventorys"`
    
}

func (req *TaobaoInventoryInitialItemRequest) GetAPIName() string {
	return "taobao.inventory.initial.item"
}

// TaobaoInventoryInitialItemResponse 商家仓商品初始化在各个仓中库存
type TaobaoInventoryInitialItemResponse struct {
    
    /* tip_infos Object Array提示信息 */
    tip_infos TipInfo `json:"tip_infos";xml:"tip_infos"`
    
}
