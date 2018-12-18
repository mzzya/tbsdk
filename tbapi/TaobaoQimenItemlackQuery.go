package tbsdk

// TaobaoQimenItemlackQueryRequest ERP调用奇门的接口,查询库存商品缺货情况
type TaobaoQimenItemlackQueryRequest struct {
    
    /* deliveryOrderCode required出库单号 */
    deliveryOrderCode string `json:"deliveryOrderCode";xml:"deliveryOrderCode"`
    
    /* deliveryOrderId optional仓储系统出库单号 */
    deliveryOrderId string `json:"deliveryOrderId";xml:"deliveryOrderId"`
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* ownerCode optional货主编码 */
    ownerCode string `json:"ownerCode";xml:"ownerCode"`
    
    /* page required当前页(从1开始) */
    page int64 `json:"page";xml:"page"`
    
    /* pageSize required每页orderLine条数(最多100条) */
    pageSize int64 `json:"pageSize";xml:"pageSize"`
    
    /* warehouseCode optional仓库编码 */
    warehouseCode string `json:"warehouseCode";xml:"warehouseCode"`
    
}

func (req *TaobaoQimenItemlackQueryRequest) GetAPIName() string {
	return "taobao.qimen.itemlack.query"
}

// TaobaoQimenItemlackQueryResponse ERP调用奇门的接口,查询库存商品缺货情况
type TaobaoQimenItemlackQueryResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* createTime Basic缺货回告创建时间(YYYY-MM-DD HH:mm:ss) */
    createTime string `json:"createTime";xml:"createTime"`
    
    /* deliveryOrderCode BasicERP的发货单编码 */
    deliveryOrderCode string `json:"deliveryOrderCode";xml:"deliveryOrderCode"`
    
    /* deliveryOrderId Basic仓库系统的发货单编码 */
    deliveryOrderId string `json:"deliveryOrderId";xml:"deliveryOrderId"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* items Object Array缺货商品列表 */
    items Item `json:"items";xml:"items"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
    /* totalLines BasicorderLines总条数 */
    totalLines int64 `json:"totalLines";xml:"totalLines"`
    
    /* warehouseCode Basic仓库编码 */
    warehouseCode string `json:"warehouseCode";xml:"warehouseCode"`
    
}
