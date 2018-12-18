package tbsdk

// TaobaoQimenWavenumReportRequest WMS调用奇门的接口,该接口用来给ERP或者OMS回传波次号及对应的发货单号，以支持商家货票同行、波次内包裹的处理等需求
type TaobaoQimenWavenumReportRequest struct {
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* orders optional发货单号 */
    orders Order `json:"orders";xml:"orders"`
    
    /* waveNum required波次号 */
    waveNum string `json:"waveNum";xml:"waveNum"`
    
}

func (req *TaobaoQimenWavenumReportRequest) GetAPIName() string {
	return "taobao.qimen.wavenum.report"
}

// TaobaoQimenWavenumReportResponse WMS调用奇门的接口,该接口用来给ERP或者OMS回传波次号及对应的发货单号，以支持商家货票同行、波次内包裹的处理等需求
type TaobaoQimenWavenumReportResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
}
