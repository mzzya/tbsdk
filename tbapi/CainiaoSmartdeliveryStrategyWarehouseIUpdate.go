package tbsdk

// CainiaoSmartdeliveryStrategyWarehouseIUpdateRequest 智能发货引擎发货策略设置仓维度
type CainiaoSmartdeliveryStrategyWarehouseIUpdateRequest struct {
    
    /* delivery_strategy_set_request required智能发货设置请求参数 */
    delivery_strategy_set_request DeliveryStrategySetRequest `json:"delivery_strategy_set_request";xml:"delivery_strategy_set_request"`
    
}

func (req *CainiaoSmartdeliveryStrategyWarehouseIUpdateRequest) GetAPIName() string {
	return "cainiao.smartdelivery.strategy.warehouse.i.update"
}

// CainiaoSmartdeliveryStrategyWarehouseIUpdateResponse 智能发货引擎发货策略设置仓维度
type CainiaoSmartdeliveryStrategyWarehouseIUpdateResponse struct {
    
    /* warehouse_info Object仓信息 */
    warehouse_info WarehouseDto `json:"warehouse_info";xml:"warehouse_info"`
    
}
