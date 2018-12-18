package tbsdk

// TaobaoInventoryInitialRequest 商家仓库存初始化接口，直接按照商家指定的商品库存数进行填充，没有单据核对，不参与库存对账
type TaobaoInventoryInitialRequest struct {
    
    /* items required商品初始库存信息： [{"scItemId":"商品后端ID，如果有传scItemCode,参数可以为0","scItemCode":"商品商家编码","inventoryType":"库存类型  1：正常,2：损坏,3：冻结,10：质押,11-20:用户自定义","quantity":"数量"}] */
    items string `json:"items";xml:"items"`
    
    /* store_code required商家仓库编码 */
    store_code string `json:"store_code";xml:"store_code"`
    
}

func (req *TaobaoInventoryInitialRequest) GetAPIName() string {
	return "taobao.inventory.initial"
}

// TaobaoInventoryInitialResponse 商家仓库存初始化接口，直接按照商家指定的商品库存数进行填充，没有单据核对，不参与库存对账
type TaobaoInventoryInitialResponse struct {
    
    /* tip_infos Object Array提示信息 */
    tip_infos TipInfo `json:"tip_infos";xml:"tip_infos"`
    
}
