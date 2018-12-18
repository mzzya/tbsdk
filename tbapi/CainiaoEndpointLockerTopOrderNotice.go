package tbsdk

// CainiaoEndpointLockerTopOrderNoticeRequest 合作公司对订单手动触发短信，有次数限制
type CainiaoEndpointLockerTopOrderNoticeRequest struct {
    
    /* mail_no required运单号 */
    mail_no string `json:"mail_no";xml:"mail_no"`
    
    /* order_code required合作公司唯一订单编号 */
    order_code string `json:"order_code";xml:"order_code"`
    
    /* scene_code required场景编号：0：重发短信，1：催取短信 */
    scene_code int64 `json:"scene_code";xml:"scene_code"`
    
    /* station_id required站点ID */
    station_id string `json:"station_id";xml:"station_id"`
    
}

func (req *CainiaoEndpointLockerTopOrderNoticeRequest) GetAPIName() string {
	return "cainiao.endpoint.locker.top.order.notice"
}

// CainiaoEndpointLockerTopOrderNoticeResponse 合作公司对订单手动触发短信，有次数限制
type CainiaoEndpointLockerTopOrderNoticeResponse struct {
    
    /* result Objectresult */
    result SingleResult `json:"result";xml:"result"`
    
}
