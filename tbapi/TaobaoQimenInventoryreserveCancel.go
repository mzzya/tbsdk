package tbsdk

// TaobaoQimenInventoryreserveCancelRequest 库存预占取消
type TaobaoQimenInventoryreserveCancelRequest struct {
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* itemInventories optional奇门仓储字段 */
    itemInventories ItemInventory `json:"itemInventories";xml:"itemInventories"`
    
    /* orderCode optional奇门仓储字段 */
    orderCode string `json:"orderCode";xml:"orderCode"`
    
    /* orderSource optional奇门仓储字段 */
    orderSource string `json:"orderSource";xml:"orderSource"`
    
    /* ownerCode optional奇门仓储字段 */
    ownerCode string `json:"ownerCode";xml:"ownerCode"`
    
}

func (req *TaobaoQimenInventoryreserveCancelRequest) GetAPIName() string {
	return "taobao.qimen.inventoryreserve.cancel"
}

// TaobaoQimenInventoryreserveCancelResponse 库存预占取消
type TaobaoQimenInventoryreserveCancelResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* isRetry Basic奇门仓储字段 */
    isRetry string `json:"isRetry";xml:"isRetry"`
    
    /* itemInventories Object Array奇门仓储字段 */
    itemInventories ItemInventory `json:"itemInventories";xml:"itemInventories"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
    /* orderCode Basic奇门仓储字段 */
    orderCode string `json:"orderCode";xml:"orderCode"`
    
}
