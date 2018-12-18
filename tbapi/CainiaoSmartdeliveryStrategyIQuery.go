package tbsdk

// CainiaoSmartdeliveryStrategyIQueryRequest 查询智能发货引擎发货策略设置
type CainiaoSmartdeliveryStrategyIQueryRequest struct {
    
    /* send_address optional发货地址 */
    send_address Address `json:"send_address";xml:"send_address"`
    
}

func (req *CainiaoSmartdeliveryStrategyIQueryRequest) GetAPIName() string {
	return "cainiao.smartdelivery.strategy.i.query"
}

// CainiaoSmartdeliveryStrategyIQueryResponse 查询智能发货引擎发货策略设置
type CainiaoSmartdeliveryStrategyIQueryResponse struct {
    
    /* delivery_strategy_info_list Object Array返回结果列表 */
    delivery_strategy_info_list DeliveryStrategyInfo `json:"delivery_strategy_info_list";xml:"delivery_strategy_info_list"`
    
}
