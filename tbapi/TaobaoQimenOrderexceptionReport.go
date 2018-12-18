package tbsdk

// TaobaoQimenOrderexceptionReportRequest WMS调用奇门的接口,当WMS接收到ERP的发货指令时，由于种种原因（5.1.5说明了各种异常场景）可能无法完成发货。WMS通过调用此接口，通知ERP具体异常情况
type TaobaoQimenOrderexceptionReportRequest struct {
    
    /* createTime optional奇门仓储字段 */
    createTime string `json:"createTime";xml:"createTime"`
    
    /* deliveryOrderCode optional奇门仓储字段 */
    deliveryOrderCode string `json:"deliveryOrderCode";xml:"deliveryOrderCode"`
    
    /* deliveryOrderId optional奇门仓储字段 */
    deliveryOrderId string `json:"deliveryOrderId";xml:"deliveryOrderId"`
    
    /* expressCode optional奇门仓储字段 */
    expressCode string `json:"expressCode";xml:"expressCode"`
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* logisticsCode optional奇门仓储字段 */
    logisticsCode string `json:"logisticsCode";xml:"logisticsCode"`
    
    /* messageDesc optional奇门仓储字段 */
    messageDesc string `json:"messageDesc";xml:"messageDesc"`
    
    /* messageId optional奇门仓储字段 */
    messageId string `json:"messageId";xml:"messageId"`
    
    /* messageType optional奇门仓储字段 */
    messageType string `json:"messageType";xml:"messageType"`
    
    /* orderLines optional奇门仓储字段 */
    orderLines OrderLine `json:"orderLines";xml:"orderLines"`
    
    /* orderType optional奇门仓储字段 */
    orderType string `json:"orderType";xml:"orderType"`
    
    /* remark optional奇门仓储字段 */
    remark string `json:"remark";xml:"remark"`
    
    /* warehouseCode optional奇门仓储字段 */
    warehouseCode string `json:"warehouseCode";xml:"warehouseCode"`
    
}

func (req *TaobaoQimenOrderexceptionReportRequest) GetAPIName() string {
	return "taobao.qimen.orderexception.report"
}

// TaobaoQimenOrderexceptionReportResponse WMS调用奇门的接口,当WMS接收到ERP的发货指令时，由于种种原因（5.1.5说明了各种异常场景）可能无法完成发货。WMS通过调用此接口，通知ERP具体异常情况
type TaobaoQimenOrderexceptionReportResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
}
