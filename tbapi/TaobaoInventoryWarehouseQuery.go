package tbsdk

// TaobaoInventoryWarehouseQueryRequest 分页查询商家仓信息
type TaobaoInventoryWarehouseQueryRequest struct {
    
    /* page_no required页码 */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size required页大小 */
    page_size int64 `json:"page_size";xml:"page_size"`
    
}

func (req *TaobaoInventoryWarehouseQueryRequest) GetAPIName() string {
	return "taobao.inventory.warehouse.query"
}

// TaobaoInventoryWarehouseQueryResponse 分页查询商家仓信息
type TaobaoInventoryWarehouseQueryResponse struct {
    
    /* result Objectresult */
    result PaginationResult `json:"result";xml:"result"`
    
}
