package tbsdk

// TaobaoQimenDeliveryorderQueryRequest ERP调用奇门的发货单查询接口，查询发货单详情
type TaobaoQimenDeliveryorderQueryRequest struct {
    
    /* deliveryOrderCode optional奇门仓储字段,说明,string(50),, */
    deliveryOrderCode string `json:"deliveryOrderCode";xml:"deliveryOrderCode"`
    
    /* deliveryOrderId optional奇门仓储字段,说明,string(50),, */
    deliveryOrderId string `json:"deliveryOrderId";xml:"deliveryOrderId"`
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* orderCode required发库单号 */
    orderCode string `json:"orderCode";xml:"orderCode"`
    
    /* orderId optional仓储系统发库单号 */
    orderId string `json:"orderId";xml:"orderId"`
    
    /* orderSourceCode optional交易单号 */
    orderSourceCode string `json:"orderSourceCode";xml:"orderSourceCode"`
    
    /* ownerCode optional货主编码 */
    ownerCode string `json:"ownerCode";xml:"ownerCode"`
    
    /* page required当前页 */
    page int64 `json:"page";xml:"page"`
    
    /* pageSize required每页orderLine条数(最多100条) */
    pageSize int64 `json:"pageSize";xml:"pageSize"`
    
    /* remark optional备注 */
    remark string `json:"remark";xml:"remark"`
    
    /* status optional奇门仓储字段,说明,string(50),, */
    status string `json:"status";xml:"status"`
    
    /* warehouseCode optional仓库编码 */
    warehouseCode string `json:"warehouseCode";xml:"warehouseCode"`
    
}

func (req *TaobaoQimenDeliveryorderQueryRequest) GetAPIName() string {
	return "taobao.qimen.deliveryorder.query"
}

// TaobaoQimenDeliveryorderQueryResponse ERP调用奇门的发货单查询接口，查询发货单详情
type TaobaoQimenDeliveryorderQueryResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* deliveryOrder Object发货单信息 */
    deliveryOrder DeliveryOrder `json:"deliveryOrder";xml:"deliveryOrder"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
    /* orderLines Object Array单据列表 */
    orderLines OrderLine `json:"orderLines";xml:"orderLines"`
    
    /* packages Object Array包裹信息 */
    packages Package `json:"packages";xml:"packages"`
    
    /* totalLines BasicorderLines总条数 */
    totalLines int64 `json:"totalLines";xml:"totalLines"`
    
}
