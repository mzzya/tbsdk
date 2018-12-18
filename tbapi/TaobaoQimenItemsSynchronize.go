package tbsdk

// TaobaoQimenItemsSynchronizeRequest ERP调用奇门的接口,批量同步商品信息给WMS
type TaobaoQimenItemsSynchronizeRequest struct {
    
    /* actionType required操作类型(两种类型：add|update) */
    actionType string `json:"actionType";xml:"actionType"`
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* items optional同步的商品信息 */
    items Item `json:"items";xml:"items"`
    
    /* ownerCode required货主编码 */
    ownerCode string `json:"ownerCode";xml:"ownerCode"`
    
    /* warehouseCode required仓库编码(统仓统配等无需ERP指定仓储编码的情况填OTHER) */
    warehouseCode string `json:"warehouseCode";xml:"warehouseCode"`
    
}

func (req *TaobaoQimenItemsSynchronizeRequest) GetAPIName() string {
	return "taobao.qimen.items.synchronize"
}

// TaobaoQimenItemsSynchronizeResponse ERP调用奇门的接口,批量同步商品信息给WMS
type TaobaoQimenItemsSynchronizeResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* items Object Array商品同步批量接口中单商品信息 */
    items BatchItemSynItem `json:"items";xml:"items"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
}
