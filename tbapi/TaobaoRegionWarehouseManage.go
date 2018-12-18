package tbsdk

// TaobaoRegionWarehouseManageRequest 编辑仓库覆盖范围
type TaobaoRegionWarehouseManageRequest struct {
    
    /* regions required可映射三级地址,例: 广东省 */
    regions string `json:"regions";xml:"regions"`
    
    /* store_code required仓库编码 */
    store_code string `json:"store_code";xml:"store_code"`
    
}

func (req *TaobaoRegionWarehouseManageRequest) GetAPIName() string {
	return "taobao.region.warehouse.manage"
}

// TaobaoRegionWarehouseManageResponse 编辑仓库覆盖范围
type TaobaoRegionWarehouseManageResponse struct {
    
    /* result Object返回结果 */
    result BaseResult `json:"result";xml:"result"`
    
}
