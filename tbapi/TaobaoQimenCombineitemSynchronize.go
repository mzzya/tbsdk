package tbsdk

// TaobaoQimenCombineitemSynchronizeRequest ERP调用奇门的接口,将商品信息同步给WMS
type TaobaoQimenCombineitemSynchronizeRequest struct {
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* itemCode required组合商品的ERP编码 */
    itemCode string `json:"itemCode";xml:"itemCode"`
    
    /* itemId optionaltemp */
    itemId string `json:"itemId";xml:"itemId"`
    
    /* items optional组合商品接口中的单商品信息 */
    items Item `json:"items";xml:"items"`
    
    /* ownerCode required货主编码 */
    ownerCode string `json:"ownerCode";xml:"ownerCode"`
    
    /* warehouseCode optional仓库编码 */
    warehouseCode string `json:"warehouseCode";xml:"warehouseCode"`
    
}

func (req *TaobaoQimenCombineitemSynchronizeRequest) GetAPIName() string {
	return "taobao.qimen.combineitem.synchronize"
}

// TaobaoQimenCombineitemSynchronizeResponse ERP调用奇门的接口,将商品信息同步给WMS
type TaobaoQimenCombineitemSynchronizeResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
}
