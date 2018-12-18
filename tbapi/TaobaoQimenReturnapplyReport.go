package tbsdk

// TaobaoQimenReturnapplyReportRequest 退货异常包裹单通知接口
type TaobaoQimenReturnapplyReportRequest struct {
    
    /* request optional */
    request Request `json:"request";xml:"request"`
    
}

func (req *TaobaoQimenReturnapplyReportRequest) GetAPIName() string {
	return "taobao.qimen.returnapply.report"
}

// TaobaoQimenReturnapplyReportResponse 退货异常包裹单通知接口
type TaobaoQimenReturnapplyReportResponse struct {
    
    /* response Object */
    response Response `json:"response";xml:"response"`
    
}
