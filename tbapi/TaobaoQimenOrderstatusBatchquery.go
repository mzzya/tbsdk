package tbsdk

// TaobaoQimenOrderstatusBatchqueryRequest ERP调用奇门的接口,查询订单在仓库的状态
type TaobaoQimenOrderstatusBatchqueryRequest struct {
    
    /* currentPage required当前第几页(从1开始) */
    currentPage int64 `json:"currentPage";xml:"currentPage"`
    
    /* endTime optional订单最后操作时间(查询截止时间点;格式:YYYY-MM-DD hh:mm:ss) */
    endTime string `json:"endTime";xml:"endTime"`
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* orderType optional单据类型(JYCK=一般交易出库单;HHCK=换货出库;BFCK=补发出库;PTCK=普通出库单;DBCK=调拨出库;QTCK=其他出库;B2BRK=B2B入库;B2BCK=B2B出库;CGRK=采购入库;DBRK=调拨入库;QTRK=其他入库;XTRK=销退入库;HHRK= 换货入库;CNJG= 仓内加工单) */
    orderType string `json:"orderType";xml:"orderType"`
    
    /* ownerCode required货主编码 */
    ownerCode string `json:"ownerCode";xml:"ownerCode"`
    
    /* pageSize required页面大小(建议不超过100条) */
    pageSize int64 `json:"pageSize";xml:"pageSize"`
    
    /* startTime required订单最后操作时间(查询起始时间点;格式:YYYY-MM-DD hh:mm:ss) */
    startTime string `json:"startTime";xml:"startTime"`
    
    /* warehouseCode optional仓库编码 */
    warehouseCode string `json:"warehouseCode";xml:"warehouseCode"`
    
}

func (req *TaobaoQimenOrderstatusBatchqueryRequest) GetAPIName() string {
	return "taobao.qimen.orderstatus.batchquery"
}

// TaobaoQimenOrderstatusBatchqueryResponse ERP调用奇门的接口,查询订单在仓库的状态
type TaobaoQimenOrderstatusBatchqueryResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
    /* orders Object Array单据信息 */
    orders Order `json:"orders";xml:"orders"`
    
    /* totalPage Basic总页数 */
    totalPage int64 `json:"totalPage";xml:"totalPage"`
    
}
