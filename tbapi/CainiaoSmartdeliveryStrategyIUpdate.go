package tbsdk

// CainiaoSmartdeliveryStrategyIUpdateRequest 更新智能发货引擎发货策略设置
type CainiaoSmartdeliveryStrategyIUpdateRequest struct {
    
    /* delivery_strategy_info required发货策略信息 */
    delivery_strategy_info DeliveryStrategyInfo `json:"delivery_strategy_info";xml:"delivery_strategy_info"`
    
}

func (req *CainiaoSmartdeliveryStrategyIUpdateRequest) GetAPIName() string {
	return "cainiao.smartdelivery.strategy.i.update"
}

// CainiaoSmartdeliveryStrategyIUpdateResponse 更新智能发货引擎发货策略设置
type CainiaoSmartdeliveryStrategyIUpdateResponse struct {
    
    /* successful Basic设置是否成功 */
    successful bool `json:"successful";xml:"successful"`
    
}
