package tbsdk

// CainiaoSmartdeliveryPriceofferIQueryRequest 查询智能发货引擎商家价格信息
type CainiaoSmartdeliveryPriceofferIQueryRequest struct {
    
    /* query_cp_price_info_request required请求参数 */
    query_cp_price_info_request QueryCpPriceInfoRequest `json:"query_cp_price_info_request";xml:"query_cp_price_info_request"`
    
}

func (req *CainiaoSmartdeliveryPriceofferIQueryRequest) GetAPIName() string {
	return "cainiao.smartdelivery.priceoffer.i.query"
}

// CainiaoSmartdeliveryPriceofferIQueryResponse 查询智能发货引擎商家价格信息
type CainiaoSmartdeliveryPriceofferIQueryResponse struct {
    
    /* cp_price_info_list Object Array返回结果列表 */
    cp_price_info_list CpPriceInfo `json:"cp_price_info_list";xml:"cp_price_info_list"`
    
}
