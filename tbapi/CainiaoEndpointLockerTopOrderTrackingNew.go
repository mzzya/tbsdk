package tbsdk

// CainiaoEndpointLockerTopOrderTrackingNewRequest 用于合作公司通知入站、出站信息到菜鸟，共定义了5个操作，1个投件操作，3个取件操作，1个超时提醒。
type CainiaoEndpointLockerTopOrderTrackingNewRequest struct {
    
    /* track_info required回传信息 */
    track_info CollectTrackingInfo `json:"track_info";xml:"track_info"`
    
}

func (req *CainiaoEndpointLockerTopOrderTrackingNewRequest) GetAPIName() string {
	return "cainiao.endpoint.locker.top.order.tracking.new"
}

// CainiaoEndpointLockerTopOrderTrackingNewResponse 用于合作公司通知入站、出站信息到菜鸟，共定义了5个操作，1个投件操作，3个取件操作，1个超时提醒。
type CainiaoEndpointLockerTopOrderTrackingNewResponse struct {
    
    /* result Objectresult */
    result SingleResult `json:"result";xml:"result"`
    
}
