package tbsdk

// TaobaoQimenInventoryReportRequest WMS调用奇门的接口,将库存盘点情况回传ERP
type TaobaoQimenInventoryReportRequest struct {
    
    /* adjustType required变动类型：CHECK=盘点 ADJUST=调整 */
    adjustType string `json:"adjustType";xml:"adjustType"`
    
    /* checkOrderCode required盘点单编码 */
    checkOrderCode string `json:"checkOrderCode";xml:"checkOrderCode"`
    
    /* checkOrderId optional仓储系统的盘点单编码 */
    checkOrderId string `json:"checkOrderId";xml:"checkOrderId"`
    
    /* checkTime optional盘点时间(YYYY-MM-DD HH:MM:SS) */
    checkTime string `json:"checkTime";xml:"checkTime"`
    
    /* currentPage required当前页(从1开始) */
    currentPage int64 `json:"currentPage";xml:"currentPage"`
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* items optional商品库存信息列表 */
    items Item `json:"items";xml:"items"`
    
    /* orderType optionalorderType */
    orderType string `json:"orderType";xml:"orderType"`
    
    /* outBizCode required外部业务编码(消息ID;用于去重;ISV对于同一请求;分配一个唯一性的编码。用来保证因为网络等原因导致重复传输;请求不 会被重复处理) */
    outBizCode string `json:"outBizCode";xml:"outBizCode"`
    
    /* ownerCode required货主编码 */
    ownerCode string `json:"ownerCode";xml:"ownerCode"`
    
    /* pageSize required每页记录的条数 */
    pageSize int64 `json:"pageSize";xml:"pageSize"`
    
    /* remark optional备注 */
    remark string `json:"remark";xml:"remark"`
    
    /* totalPage required总页数 */
    totalPage int64 `json:"totalPage";xml:"totalPage"`
    
    /* warehouseCode required仓库编码 */
    warehouseCode string `json:"warehouseCode";xml:"warehouseCode"`
    
}

func (req *TaobaoQimenInventoryReportRequest) GetAPIName() string {
	return "taobao.qimen.inventory.report"
}

// TaobaoQimenInventoryReportResponse WMS调用奇门的接口,将库存盘点情况回传ERP
type TaobaoQimenInventoryReportResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
}
