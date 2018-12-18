package tbsdk

// TaobaoQimenSnReportRequest WMS调用奇门的接口,在仓库出库单后, 把SN信息回传给ERP
type TaobaoQimenSnReportRequest struct {
    
    /* currentPage required当前页(从1开始) */
    currentPage int64 `json:"currentPage";xml:"currentPage"`
    
    /* deliveryOrder optional发货单信息 */
    deliveryOrder DeliveryOrder `json:"deliveryOrder";xml:"deliveryOrder"`
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* items optional商品列表 */
    items Item `json:"items";xml:"items"`
    
    /* pageSize required每页记录的条数 */
    pageSize int64 `json:"pageSize";xml:"pageSize"`
    
    /* totalPage required总页数 */
    totalPage int64 `json:"totalPage";xml:"totalPage"`
    
}

func (req *TaobaoQimenSnReportRequest) GetAPIName() string {
	return "taobao.qimen.sn.report"
}

// TaobaoQimenSnReportResponse WMS调用奇门的接口,在仓库出库单后, 把SN信息回传给ERP
type TaobaoQimenSnReportResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
}
