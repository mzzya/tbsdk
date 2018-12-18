package tbsdk

// TaobaoQimenItemlackReportRequest WMS调用奇门的接口,将商家在库某商品缺货的信息回传给ERP
type TaobaoQimenItemlackReportRequest struct {
    
    /* createTime required缺货回告创建时间(YYYY-MM-DD HH:mm:ss) */
    createTime string `json:"createTime";xml:"createTime"`
    
    /* deliveryOrderCode requiredERP的发货单编码 */
    deliveryOrderCode string `json:"deliveryOrderCode";xml:"deliveryOrderCode"`
    
    /* deliveryOrderId optional仓库系统的发货单编码 */
    deliveryOrderId string `json:"deliveryOrderId";xml:"deliveryOrderId"`
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* items optional缺货商品列表 */
    items Item `json:"items";xml:"items"`
    
    /* outBizCode required外部业务编码(消息ID;用于去重;ISV对于同一请求;分配一个唯一性的编码。用来保证因为网络等原因导致重复传输;请求不 会被重复处理) */
    outBizCode string `json:"outBizCode";xml:"outBizCode"`
    
    /* remark optional备注 */
    remark string `json:"remark";xml:"remark"`
    
    /* warehouseCode required仓库编码 */
    warehouseCode string `json:"warehouseCode";xml:"warehouseCode"`
    
}

func (req *TaobaoQimenItemlackReportRequest) GetAPIName() string {
	return "taobao.qimen.itemlack.report"
}

// TaobaoQimenItemlackReportResponse WMS调用奇门的接口,将商家在库某商品缺货的信息回传给ERP
type TaobaoQimenItemlackReportResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
}
