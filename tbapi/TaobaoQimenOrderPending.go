package tbsdk

// TaobaoQimenOrderPendingRequest ERP调用奇门的接口,挂起某些创建的单据;场景介绍：ERP主动发起挂起（恢复）某些创建的单据，如入库单、出库单、退货单等
type TaobaoQimenOrderPendingRequest struct {
    
    /* actionType required操作类型(pending=挂起;restore=恢复) */
    actionType string `json:"actionType";xml:"actionType"`
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* orderCode required单据编码 */
    orderCode string `json:"orderCode";xml:"orderCode"`
    
    /* orderId optional仓储系统单据编码 */
    orderId string `json:"orderId";xml:"orderId"`
    
    /* orderType optional单据类型(JYCK=一般交易出库单;HHCK=换货出库;BFCK=补发出库;PTCK=普通出库单;DBCK=调拨出库;QTCK=其他出库;B2BRK=B2B入库;B2BCK=B2B出库;CGRK=采购入库;DBRK=调拨入库;QTRK=其他入库;XTRK=销退入库;HHRK=换货入库;CNJG=仓内加工单) */
    orderType string `json:"orderType";xml:"orderType"`
    
    /* ownerCode optional货主编码 */
    ownerCode string `json:"ownerCode";xml:"ownerCode"`
    
    /* reason optional挂起/恢复原因 */
    reason string `json:"reason";xml:"reason"`
    
    /* warehouseCode required仓库编码(统仓统配等无需ERP指定仓储编码的情况填OTHER) */
    warehouseCode string `json:"warehouseCode";xml:"warehouseCode"`
    
}

func (req *TaobaoQimenOrderPendingRequest) GetAPIName() string {
	return "taobao.qimen.order.pending"
}

// TaobaoQimenOrderPendingResponse ERP调用奇门的接口,挂起某些创建的单据;场景介绍：ERP主动发起挂起（恢复）某些创建的单据，如入库单、出库单、退货单等
type TaobaoQimenOrderPendingResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
}
