package tbsdk

// CainiaoSmartdeliveryStrategyWarehouseIQueryRequest 智能发货引擎仓维度策略查询
type CainiaoSmartdeliveryStrategyWarehouseIQueryRequest struct {
    
    /* param_query_delivery_strategy_request optional查询请求参数 */
    param_query_delivery_strategy_request QueryDeliveryStrategyRequest `json:"param_query_delivery_strategy_request";xml:"param_query_delivery_strategy_request"`
    
}

func (req *CainiaoSmartdeliveryStrategyWarehouseIQueryRequest) GetAPIName() string {
	return "cainiao.smartdelivery.strategy.warehouse.i.query"
}

// CainiaoSmartdeliveryStrategyWarehouseIQueryResponse 智能发货引擎仓维度策略查询
type CainiaoSmartdeliveryStrategyWarehouseIQueryResponse struct {
    
    /* delivery_strategy_info_list Object Array返回结果列表 */
    delivery_strategy_info_list DeliveryStrategyInfo `json:"delivery_strategy_info_list";xml:"delivery_strategy_info_list"`
    
}
