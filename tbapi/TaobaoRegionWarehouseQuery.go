package tbsdk

// TaobaoRegionWarehouseQueryRequest 查询仓库覆盖范围
type TaobaoRegionWarehouseQueryRequest struct {
    
    /* store_code required仓库编码 */
    store_code string `json:"store_code";xml:"store_code"`
    
}

func (req *TaobaoRegionWarehouseQueryRequest) GetAPIName() string {
	return "taobao.region.warehouse.query"
}

// TaobaoRegionWarehouseQueryResponse 查询仓库覆盖范围
type TaobaoRegionWarehouseQueryResponse struct {
    
    /* result Objectresult */
    result BaseResult `json:"result";xml:"result"`
    
}
