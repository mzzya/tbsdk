package tbsdk

// CainiaoEndpointLockerTopOrderNoticesendQueryRequest 合作公司查询消息发送的接口，判断是否裹裹发送消息
type CainiaoEndpointLockerTopOrderNoticesendQueryRequest struct {
    
    /* getter_phone optional收件人手机号 */
    getter_phone string `json:"getter_phone";xml:"getter_phone"`
    
    /* mail_no required运单号 */
    mail_no string `json:"mail_no";xml:"mail_no"`
    
    /* station_id required站点id */
    station_id string `json:"station_id";xml:"station_id"`
    
}

func (req *CainiaoEndpointLockerTopOrderNoticesendQueryRequest) GetAPIName() string {
	return "cainiao.endpoint.locker.top.order.noticesend.query"
}

// CainiaoEndpointLockerTopOrderNoticesendQueryResponse 合作公司查询消息发送的接口，判断是否裹裹发送消息
type CainiaoEndpointLockerTopOrderNoticesendQueryResponse struct {
    
    /* result Object返回结果 */
    result SingleResult `json:"result";xml:"result"`
    
}
