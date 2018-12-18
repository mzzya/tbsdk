package tbsdk

// TaobaoQimenInventoryreserveCreateRequest ERP调用奇门的接口,查询发货库存预占用信息
type TaobaoQimenInventoryreserveCreateRequest struct {
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* itemInventories optional交易订单信息 */
    itemInventories ItemInventory `json:"itemInventories";xml:"itemInventories"`
    
    /* maxWarehouseNum optional最大仓库数(当一个仓库不满足库存时.是否允许占用多个仓库库存.如果允许.需要指定最大分仓数量.但不能拆分订单明细. 0：不限个数.只要满足发货.不限分占几个仓库 1：默认值.只能单仓发 >1:最大 占用数量) */
    maxWarehouseNum int64 `json:"maxWarehouseNum";xml:"maxWarehouseNum"`
    
    /* orderCode requiredERP订单编码 */
    orderCode string `json:"orderCode";xml:"orderCode"`
    
    /* orderSource required订单来源(213 天猫、201 淘宝、214 京东、202 1688 阿里中文站、203 苏宁在线、204 亚马逊中国、205 当当、208 1号店、207 唯品会、209 国美在线、210 拍拍、206 易贝ebay、211 聚美优品、212 乐蜂 网、215 邮乐、216 凡客、217 优购、218 银泰、219 易讯、221 聚尚网、222 蘑菇街、223 POS门店、301 其他) */
    orderSource int64 `json:"orderSource";xml:"orderSource"`
    
    /* ownerCode required货主编码 */
    ownerCode string `json:"ownerCode";xml:"ownerCode"`
    
    /* receiverInfo optional收件者信息 */
    receiverInfo ReceiverInfo `json:"receiverInfo";xml:"receiverInfo"`
    
    /* warehouseCode optional仓库编码 */
    warehouseCode string `json:"warehouseCode";xml:"warehouseCode"`
    
}

func (req *TaobaoQimenInventoryreserveCreateRequest) GetAPIName() string {
	return "taobao.qimen.inventoryreserve.create"
}

// TaobaoQimenInventoryreserveCreateResponse ERP调用奇门的接口,查询发货库存预占用信息
type TaobaoQimenInventoryreserveCreateResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* isRetry Basic是否需要重试(Y/N (默认为N)返回Y时建议系统自动重试) */
    isRetry string `json:"isRetry";xml:"isRetry"`
    
    /* itemInventories Object Array */
    itemInventories ItemInventory `json:"itemInventories";xml:"itemInventories"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
    /* orderCode BasicERP订单编码 */
    orderCode string `json:"orderCode";xml:"orderCode"`
    
}
