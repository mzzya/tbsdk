package tbsdk

// TaobaoQimenOrderCancelRequest ERP调用奇门的接口,取消创建单据操作。场景介绍：ERP主动发起取消某些创建的单据。如入库单、出库单、退货单等；所有的场景
type TaobaoQimenOrderCancelRequest struct {
    
    /* cancelReason optional取消原因 */
    cancelReason string `json:"cancelReason";xml:"cancelReason"`
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* orderCode required单据编码 */
    orderCode string `json:"orderCode";xml:"orderCode"`
    
    /* orderId optional仓储系统单据编码 */
    orderId string `json:"orderId";xml:"orderId"`
    
    /* orderType optional单据类型(JYCK=一般交易出库单;HHCK= 换货出库;BFCK=补发出库;PTCK=普通出库单;DBCK=调拨出库;B2BRK=B2B入库;B2BCK=B2B出库;QTCK=其他出库;SCRK=生产入库;LYRK=领用入库;CCRK=残次品入库;CGRK=采购入库;DBRK= 调拨入库;QTRK=其他入库;XTRK= 销退入库;THRK=退货入库;HHRK= 换货入库;CNJG= 仓内加工单;CGTH=采购退货出库单) */
    orderType string `json:"orderType";xml:"orderType"`
    
    /* ownerCode optional货主编码 */
    ownerCode string `json:"ownerCode";xml:"ownerCode"`
    
    /* remark optional备注 */
    remark string `json:"remark";xml:"remark"`
    
    /* warehouseCode required仓库编码(统仓统配等无需ERP指定仓储编码的情况填OTHER) */
    warehouseCode string `json:"warehouseCode";xml:"warehouseCode"`
    
}

func (req *TaobaoQimenOrderCancelRequest) GetAPIName() string {
	return "taobao.qimen.order.cancel"
}

// TaobaoQimenOrderCancelResponse ERP调用奇门的接口,取消创建单据操作。场景介绍：ERP主动发起取消某些创建的单据。如入库单、出库单、退货单等；所有的场景
type TaobaoQimenOrderCancelResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
}
