package tbsdk

// TaobaoQimenSingleitemSynchronizeRequest ERP调用奇门的接口,同步商品信息给WMS
type TaobaoQimenSingleitemSynchronizeRequest struct {
    
    /* actionType required操作类型(两种类型：add|update) */
    actionType string `json:"actionType";xml:"actionType"`
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* item optional商品信息 */
    item Item `json:"item";xml:"item"`
    
    /* ownerCode required货主编码 */
    ownerCode string `json:"ownerCode";xml:"ownerCode"`
    
    /* supplierCode optional供应商编码 */
    supplierCode string `json:"supplierCode";xml:"supplierCode"`
    
    /* supplierName optional供应商名称 */
    supplierName string `json:"supplierName";xml:"supplierName"`
    
    /* warehouseCode required仓库编码(统仓统配等无需ERP指定仓储编码的情况填OTHER) */
    warehouseCode string `json:"warehouseCode";xml:"warehouseCode"`
    
}

func (req *TaobaoQimenSingleitemSynchronizeRequest) GetAPIName() string {
	return "taobao.qimen.singleitem.synchronize"
}

// TaobaoQimenSingleitemSynchronizeResponse ERP调用奇门的接口,同步商品信息给WMS
type TaobaoQimenSingleitemSynchronizeResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* itemId Basic仓储系统商品Id(当这个字段不为空的时候;所有erp传输的时候都碰到itemid必传) */
    itemId string `json:"itemId";xml:"itemId"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
}
