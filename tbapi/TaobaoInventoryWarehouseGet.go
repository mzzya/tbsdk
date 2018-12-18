package tbsdk

// TaobaoInventoryWarehouseGetRequest 获取商家仓信息
type TaobaoInventoryWarehouseGetRequest struct {
    
    /* store_code required仓库编码 */
    store_code string `json:"store_code";xml:"store_code"`
    
}

func (req *TaobaoInventoryWarehouseGetRequest) GetAPIName() string {
	return "taobao.inventory.warehouse.get"
}

// TaobaoInventoryWarehouseGetResponse 获取商家仓信息
type TaobaoInventoryWarehouseGetResponse struct {
    
    /* result Objectresult */
    result BaseResult `json:"result";xml:"result"`
    
}
