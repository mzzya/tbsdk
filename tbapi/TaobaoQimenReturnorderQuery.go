package tbsdk

// TaobaoQimenReturnorderQueryRequest ERP调用奇门的接口，查询退货入库信息
type TaobaoQimenReturnorderQueryRequest struct {
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* ownerCode optional货主编码 */
    ownerCode string `json:"ownerCode";xml:"ownerCode"`
    
    /* page required当前页(从1开始) */
    page int64 `json:"page";xml:"page"`
    
    /* pageSize required每页orderLine条数(最多100条) */
    pageSize int64 `json:"pageSize";xml:"pageSize"`
    
    /* returnOrderCode required退货单编码 */
    returnOrderCode string `json:"returnOrderCode";xml:"returnOrderCode"`
    
    /* returnOrderId optional仓库系统订单编码 */
    returnOrderId string `json:"returnOrderId";xml:"returnOrderId"`
    
    /* warehouseCode optional仓库编码 */
    warehouseCode string `json:"warehouseCode";xml:"warehouseCode"`
    
}

func (req *TaobaoQimenReturnorderQueryRequest) GetAPIName() string {
	return "taobao.qimen.returnorder.query"
}

// TaobaoQimenReturnorderQueryResponse ERP调用奇门的接口，查询退货入库信息
type TaobaoQimenReturnorderQueryResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
    /* orderLines Object Array订单信息 */
    orderLines OrderLine `json:"orderLines";xml:"orderLines"`
    
    /* returnOrder Object退货单信息 */
    returnOrder ReturnOrder `json:"returnOrder";xml:"returnOrder"`
    
}
