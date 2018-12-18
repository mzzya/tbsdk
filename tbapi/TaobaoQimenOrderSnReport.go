package tbsdk

// TaobaoQimenOrderSnReportRequest WMS调用奇门的接口,在出库、发货、入库等场景下，ERP和WMS之间同步操作的SN列表
type TaobaoQimenOrderSnReportRequest struct {
    
    /* currentPage required当前页(从1开始) */
    currentPage int64 `json:"currentPage";xml:"currentPage"`
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* items optional商品列表 */
    items Items `json:"items";xml:"items"`
    
    /* order optional单据列表 */
    order Order `json:"order";xml:"order"`
    
    /* pageSize required每页记录的条数 */
    pageSize int64 `json:"pageSize";xml:"pageSize"`
    
    /* totalPage required总页数 */
    totalPage int64 `json:"totalPage";xml:"totalPage"`
    
}

func (req *TaobaoQimenOrderSnReportRequest) GetAPIName() string {
	return "taobao.qimen.order.sn.report"
}

// TaobaoQimenOrderSnReportResponse WMS调用奇门的接口,在出库、发货、入库等场景下，ERP和WMS之间同步操作的SN列表
type TaobaoQimenOrderSnReportResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
}
