package tbsdk

// TaobaoQimenWarehouseinfoQueryRequest 货主仓库资源查询
type TaobaoQimenWarehouseinfoQueryRequest struct {
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* ownerCode optional奇门仓储字段 */
    ownerCode string `json:"ownerCode";xml:"ownerCode"`
    
}

func (req *TaobaoQimenWarehouseinfoQueryRequest) GetAPIName() string {
	return "taobao.qimen.warehouseinfo.query"
}

// TaobaoQimenWarehouseinfoQueryResponse 货主仓库资源查询
type TaobaoQimenWarehouseinfoQueryResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
    /* ownerCode Basic奇门仓储字段 */
    ownerCode string `json:"ownerCode";xml:"ownerCode"`
    
    /* ownerName Basic奇门仓储字段 */
    ownerName string `json:"ownerName";xml:"ownerName"`
    
    /* warehouseInfos Object Array奇门仓储字段 */
    warehouseInfos WarehouseInfo `json:"warehouseInfos";xml:"warehouseInfos"`
    
}
