package tbsdk

// CainiaoMerchantInventoryAdjustRequest 商家仓库存调整接口，目前仅支持全量更新
type CainiaoMerchantInventoryAdjustRequest struct {
    
    /* adjust_request required商家仓编辑库存 */
    adjust_request MerStoreInvAdjustDto `json:"adjust_request";xml:"adjust_request"`
    
    /* app_name required调用方应用名 */
    app_name string `json:"app_name";xml:"app_name"`
    
    /* operation optional操作 */
    operation string `json:"operation";xml:"operation"`
    
}

func (req *CainiaoMerchantInventoryAdjustRequest) GetAPIName() string {
	return "cainiao.merchant.inventory.adjust"
}

// CainiaoMerchantInventoryAdjustResponse 商家仓库存调整接口，目前仅支持全量更新
type CainiaoMerchantInventoryAdjustResponse struct {
    
    /* result Objectresult */
    result SingleResultDto `json:"result";xml:"result"`
    
}
