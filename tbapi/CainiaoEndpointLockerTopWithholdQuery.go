package tbsdk

// CainiaoEndpointLockerTopWithholdQueryRequest 查询是否有代扣欠款，是否签署代扣协议。
type CainiaoEndpointLockerTopWithholdQueryRequest struct {
    
    /* company_code required柜子公司编码 */
    company_code string `json:"company_code";xml:"company_code"`
    
    /* open_user_id required开放用户Id */
    open_user_id string `json:"open_user_id";xml:"open_user_id"`
    
    /* order_type required柜子业务：0-取件业务，1-寄件业务，2-派样业务，3-小件员约柜（在约期内仅能使用一次）4-小件员租柜（在约期内可反复使用） */
    order_type int64 `json:"order_type";xml:"order_type"`
    
}

func (req *CainiaoEndpointLockerTopWithholdQueryRequest) GetAPIName() string {
	return "cainiao.endpoint.locker.top.withhold.query"
}

// CainiaoEndpointLockerTopWithholdQueryResponse 查询是否有代扣欠款，是否签署代扣协议。
type CainiaoEndpointLockerTopWithholdQueryResponse struct {
    
    /* result Objectresponse */
    result SingleResult `json:"result";xml:"result"`
    
}
