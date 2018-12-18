package tbsdk

// TaobaoQimenStockchangeReportRequest WMS调用奇门的接口,将库存异动信息信息回传给ERP
type TaobaoQimenStockchangeReportRequest struct {
    
    /* currentPage optional奇门仓储字段,说明,string(50),, */
    currentPage string `json:"currentPage";xml:"currentPage"`
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* items optionalitem */
    items Item `json:"items";xml:"items"`
    
    /* orderCode optional奇门仓储字段,说明,string(50),, */
    orderCode string `json:"orderCode";xml:"orderCode"`
    
    /* orderType optional奇门仓储字段,说明,string(50),, */
    orderType string `json:"orderType";xml:"orderType"`
    
    /* ownerCode optional奇门仓储字段,说明,string(50),, */
    ownerCode string `json:"ownerCode";xml:"ownerCode"`
    
    /* pageSize optional奇门仓储字段,说明,string(50),, */
    pageSize string `json:"pageSize";xml:"pageSize"`
    
    /* remark optional奇门仓储字段,说明,string(50),, */
    remark string `json:"remark";xml:"remark"`
    
    /* totalPage optional奇门仓储字段,说明,string(50),, */
    totalPage string `json:"totalPage";xml:"totalPage"`
    
    /* warehouseCode optional奇门仓储字段,说明,string(50),, */
    warehouseCode string `json:"warehouseCode";xml:"warehouseCode"`
    
}

func (req *TaobaoQimenStockchangeReportRequest) GetAPIName() string {
	return "taobao.qimen.stockchange.report"
}

// TaobaoQimenStockchangeReportResponse WMS调用奇门的接口,将库存异动信息信息回传给ERP
type TaobaoQimenStockchangeReportResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
}
