package tbsdk

// TaobaoQimenOrderprocessReportRequest WMS调用奇门的接口,将订单在仓库的状态回传给ERP；场景说明：仓库仓内操作状态回传给ERP, 比如打包操作完成时, 回传一个打 包完成的状态给到ERP, ERP自行决定如何处理。
type TaobaoQimenOrderprocessReportRequest struct {
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* order optional订单信息 */
    order Order `json:"order";xml:"order"`
    
    /* process optional订单处理信息 */
    process Process `json:"process";xml:"process"`
    
    /* remark optional备注 */
    remark string `json:"remark";xml:"remark"`
    
}

func (req *TaobaoQimenOrderprocessReportRequest) GetAPIName() string {
	return "taobao.qimen.orderprocess.report"
}

// TaobaoQimenOrderprocessReportResponse WMS调用奇门的接口,将订单在仓库的状态回传给ERP；场景说明：仓库仓内操作状态回传给ERP, 比如打包操作完成时, 回传一个打 包完成的状态给到ERP, ERP自行决定如何处理。
type TaobaoQimenOrderprocessReportResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
}
