package tbsdk

// TaobaoInventoryMerchantAdjustRequest 货品库存商家端调整 ，入库，出库，盘点
type TaobaoInventoryMerchantAdjustRequest struct {
    
    /* inventory_check required调整库存对象 */
    inventory_check InventoryCheckDto `json:"inventory_check";xml:"inventory_check"`
    
}

func (req *TaobaoInventoryMerchantAdjustRequest) GetAPIName() string {
	return "taobao.inventory.merchant.adjust"
}

// TaobaoInventoryMerchantAdjustResponse 货品库存商家端调整 ，入库，出库，盘点
type TaobaoInventoryMerchantAdjustResponse struct {
    
    /* result Objectresult */
    result SingleResult `json:"result";xml:"result"`
    
}
