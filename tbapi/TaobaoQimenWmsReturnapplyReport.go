package tbsdk

// TaobaoQimenWmsReturnapplyReportRequest 退货异常包裹单通知接口
type TaobaoQimenWmsReturnapplyReportRequest struct {
    
    /* request optional请求对象 */
    request Request `json:"request";xml:"request"`
    
}

func (req *TaobaoQimenWmsReturnapplyReportRequest) GetAPIName() string {
	return "taobao.qimen.wms.returnapply.report"
}

// TaobaoQimenWmsReturnapplyReportResponse 退货异常包裹单通知接口
type TaobaoQimenWmsReturnapplyReportResponse struct {
    
    /* response Object响应对象 */
    response Response `json:"response";xml:"response"`
    
}
