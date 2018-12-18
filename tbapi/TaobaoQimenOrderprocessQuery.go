package tbsdk

// TaobaoQimenOrderprocessQueryRequest ERP调用订单流水查询接口
type TaobaoQimenOrderprocessQueryRequest struct {
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* orderCode required单据号 */
    orderCode string `json:"orderCode";xml:"orderCode"`
    
    /* orderId optional仓储系统单据号 */
    orderId string `json:"orderId";xml:"orderId"`
    
    /* orderSourceCode optional交易单号 */
    orderSourceCode string `json:"orderSourceCode";xml:"orderSourceCode"`
    
    /* orderType optional单据类型(JYCK=一般交易出库单;HHCK=换货出库;BFCK=补发出库;PTCK=普通出库单;DBCK=调拨出库;QTCK=其他出库;B2BRK=B2B入库;B2BCK=B2B出库;CGRK=采购入库;DBRK=调拨入库;QTRK=其他入库;XTRK=销退入库;HHRK=换货入库;CNJG=仓内加工单) */
    orderType string `json:"orderType";xml:"orderType"`
    
    /* ownerCode optional货主编码 */
    ownerCode string `json:"ownerCode";xml:"ownerCode"`
    
    /* remark optional备注 */
    remark string `json:"remark";xml:"remark"`
    
    /* warehouseCode optional仓库编码 */
    warehouseCode string `json:"warehouseCode";xml:"warehouseCode"`
    
}

func (req *TaobaoQimenOrderprocessQueryRequest) GetAPIName() string {
	return "taobao.qimen.orderprocess.query"
}

// TaobaoQimenOrderprocessQueryResponse ERP调用订单流水查询接口
type TaobaoQimenOrderprocessQueryResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
    /* orderProcess Object订单处理流程 */
    orderProcess OrderProcess `json:"orderProcess";xml:"orderProcess"`
    
}
