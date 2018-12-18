package tbsdk

// TaobaoQimenStockoutQueryRequest 查询出库单详情
type TaobaoQimenStockoutQueryRequest struct {
    
    /* deliveryOrderCode required出库单号 */
    deliveryOrderCode string `json:"deliveryOrderCode";xml:"deliveryOrderCode"`
    
    /* deliveryOrderId optional仓储系统出库单号 */
    deliveryOrderId string `json:"deliveryOrderId";xml:"deliveryOrderId"`
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* ownerCode optional货主编码 */
    ownerCode string `json:"ownerCode";xml:"ownerCode"`
    
    /* page required当前页 */
    page int64 `json:"page";xml:"page"`
    
    /* pageSize required每页orderLine条数(最多100条) */
    pageSize int64 `json:"pageSize";xml:"pageSize"`
    
    /* warehouseCode optional仓库编码 */
    warehouseCode string `json:"warehouseCode";xml:"warehouseCode"`
    
}

func (req *TaobaoQimenStockoutQueryRequest) GetAPIName() string {
	return "taobao.qimen.stockout.query"
}

// TaobaoQimenStockoutQueryResponse 查询出库单详情
type TaobaoQimenStockoutQueryResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* deliveryOrder Object出库单信息 */
    deliveryOrder DeliveryOrder `json:"deliveryOrder";xml:"deliveryOrder"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
    /* orderLines Object Array单据信息列表 */
    orderLines OrderLine `json:"orderLines";xml:"orderLines"`
    
    /* packages Object Array包裹信息 */
    packages Package `json:"packages";xml:"packages"`
    
    /* totalLines BasicorderLines总条数 */
    totalLines int64 `json:"totalLines";xml:"totalLines"`
    
}
