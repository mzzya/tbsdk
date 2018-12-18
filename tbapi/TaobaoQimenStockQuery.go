package tbsdk

// TaobaoQimenStockQueryRequest ERP调用奇门的接口,查询商品的库存量
type TaobaoQimenStockQueryRequest struct {
    
    /* batchCode optional批次编码 */
    batchCode string `json:"batchCode";xml:"batchCode"`
    
    /* expireDate optional商品过期日期(YYYY-MM-DD) */
    expireDate string `json:"expireDate";xml:"expireDate"`
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* inventoryType optional库存类型(ZP=正品;CC=残次;JS=机损;XS=箱损;ZT=在途库存;默认为查所有类型的库存) */
    inventoryType string `json:"inventoryType";xml:"inventoryType"`
    
    /* itemCode required商品编码 */
    itemCode string `json:"itemCode";xml:"itemCode"`
    
    /* itemId optional仓储系统商品ID */
    itemId string `json:"itemId";xml:"itemId"`
    
    /* ownerCode optional货主编码 */
    ownerCode string `json:"ownerCode";xml:"ownerCode"`
    
    /* page required当前页(从1开始) */
    page int64 `json:"page";xml:"page"`
    
    /* pageSize required每页条数(最多100条) */
    pageSize int64 `json:"pageSize";xml:"pageSize"`
    
    /* productDate optional商品生产日期(YYYY-MM-DD) */
    productDate string `json:"productDate";xml:"productDate"`
    
    /* remark optional备注 */
    remark string `json:"remark";xml:"remark"`
    
    /* warehouseCode optional仓库编码 */
    warehouseCode string `json:"warehouseCode";xml:"warehouseCode"`
    
}

func (req *TaobaoQimenStockQueryRequest) GetAPIName() string {
	return "taobao.qimen.stock.query"
}

// TaobaoQimenStockQueryResponse ERP调用奇门的接口,查询商品的库存量
type TaobaoQimenStockQueryResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* items Object Array商品的库存信息列表 */
    items Item `json:"items";xml:"items"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
    /* totalCount Basic总数 */
    totalCount int64 `json:"totalCount";xml:"totalCount"`
    
}
