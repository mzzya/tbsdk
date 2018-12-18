package tbsdk

// TaobaoInventoryWarehouseManageRequest 创建商家仓或者更新商家仓信息
type TaobaoInventoryWarehouseManageRequest struct {
    
    /* ware_house_dto required仓库信息 */
    ware_house_dto WareHouseDto `json:"ware_house_dto";xml:"ware_house_dto"`
    
}

func (req *TaobaoInventoryWarehouseManageRequest) GetAPIName() string {
	return "taobao.inventory.warehouse.manage"
}

// TaobaoInventoryWarehouseManageResponse 创建商家仓或者更新商家仓信息
type TaobaoInventoryWarehouseManageResponse struct {
    
    /* result Objectresult */
    result Result `json:"result";xml:"result"`
    
}
