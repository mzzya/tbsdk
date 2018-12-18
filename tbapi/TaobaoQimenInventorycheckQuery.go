package tbsdk

// TaobaoQimenInventorycheckQueryRequest ERP调用奇门的接口,查询库存盘点情况
type TaobaoQimenInventorycheckQueryRequest struct {
    
    /* checkOrderCode required盘点单编码 */
    checkOrderCode string `json:"checkOrderCode";xml:"checkOrderCode"`
    
    /* checkOrderId optional仓储系统的盘点单编码 */
    checkOrderId string `json:"checkOrderId";xml:"checkOrderId"`
    
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

func (req *TaobaoQimenInventorycheckQueryRequest) GetAPIName() string {
	return "taobao.qimen.inventorycheck.query"
}

// TaobaoQimenInventorycheckQueryResponse ERP调用奇门的接口,查询库存盘点情况
type TaobaoQimenInventorycheckQueryResponse struct {
    
    /* checkOrderCode Basic盘点单编码 */
    checkOrderCode string `json:"checkOrderCode";xml:"checkOrderCode"`
    
    /* checkOrderId Basic仓储系统的盘点单编码 */
    checkOrderId string `json:"checkOrderId";xml:"checkOrderId"`
    
    /* checkTime Basic盘点时间(YYYY-MM-DD HH:MM:SS) */
    checkTime string `json:"checkTime";xml:"checkTime"`
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* items Object Array商品库存列表 */
    items Item `json:"items";xml:"items"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
    /* ownerCode Basic货主编码 */
    ownerCode string `json:"ownerCode";xml:"ownerCode"`
    
    /* remark Basic备注 */
    remark string `json:"remark";xml:"remark"`
    
    /* totalLines BasicorderLines总条数 */
    totalLines int64 `json:"totalLines";xml:"totalLines"`
    
    /* warehouseCode Basic仓库编码 */
    warehouseCode string `json:"warehouseCode";xml:"warehouseCode"`
    
}
