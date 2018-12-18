package tbsdk

// CainiaoSmartdeliveryStrategyWarehouseIDeleteRequest 删除智能发货引擎仓策略
type CainiaoSmartdeliveryStrategyWarehouseIDeleteRequest struct {
    
    /* warehouse_id optional仓id */
    warehouse_id int64 `json:"warehouse_id";xml:"warehouse_id"`
    
}

func (req *CainiaoSmartdeliveryStrategyWarehouseIDeleteRequest) GetAPIName() string {
	return "cainiao.smartdelivery.strategy.warehouse.i.delete"
}

// CainiaoSmartdeliveryStrategyWarehouseIDeleteResponse 删除智能发货引擎仓策略
type CainiaoSmartdeliveryStrategyWarehouseIDeleteResponse struct {
    
    /* is_delete_success Basicdata */
    is_delete_success bool `json:"is_delete_success";xml:"is_delete_success"`
    
}
