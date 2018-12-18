package tbsdk

// CainiaoSmartdeliveryPriceofferIUpdateRequest 智能发货引擎更新价格信息模板
type CainiaoSmartdeliveryPriceofferIUpdateRequest struct {
    
    /* cp_price_info required物流公司价格信息 */
    cp_price_info CpPriceInfo `json:"cp_price_info";xml:"cp_price_info"`
    
}

func (req *CainiaoSmartdeliveryPriceofferIUpdateRequest) GetAPIName() string {
	return "cainiao.smartdelivery.priceoffer.i.update"
}

// CainiaoSmartdeliveryPriceofferIUpdateResponse 智能发货引擎更新价格信息模板
type CainiaoSmartdeliveryPriceofferIUpdateResponse struct {
    
    /* successful Basic设置是否成功 */
    successful bool `json:"successful";xml:"successful"`
    
}
