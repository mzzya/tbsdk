package tbsdk

// CainiaoEndpointLockerTopOrderWithholdRequest 提供代扣，允许有一笔欠款。
type CainiaoEndpointLockerTopOrderWithholdRequest struct {
    
    /* company_code required柜子公司编码 */
    company_code string `json:"company_code";xml:"company_code"`
    
    /* extra optional扩展字段 */
    extra string `json:"extra";xml:"extra"`
    
    /* gui_id required柜子id */
    gui_id string `json:"gui_id";xml:"gui_id"`
    
    /* mail_no required运单号 */
    mail_no string `json:"mail_no";xml:"mail_no"`
    
    /* open_user_id required开放用户id */
    open_user_id string `json:"open_user_id";xml:"open_user_id"`
    
    /* order_code required柜子订单编码 */
    order_code string `json:"order_code";xml:"order_code"`
    
    /* order_type required订单类型(0-取件业务，1-寄件业务，2-派样业务) */
    order_type int64 `json:"order_type";xml:"order_type"`
    
    /* total_fee required代扣金额（全额），单位：分 */
    total_fee int64 `json:"total_fee";xml:"total_fee"`
    
}

func (req *CainiaoEndpointLockerTopOrderWithholdRequest) GetAPIName() string {
	return "cainiao.endpoint.locker.top.order.withhold"
}

// CainiaoEndpointLockerTopOrderWithholdResponse 提供代扣，允许有一笔欠款。
type CainiaoEndpointLockerTopOrderWithholdResponse struct {
    
    /* result Objectresult */
    result SingleResult `json:"result";xml:"result"`
    
}
