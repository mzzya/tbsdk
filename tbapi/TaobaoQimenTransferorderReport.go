package tbsdk

// TaobaoQimenTransferorderReportRequest 调拨单通知
type TaobaoQimenTransferorderReportRequest struct {
    
    /* confirmInTime optional确认入库时间 */
    confirmInTime string `json:"confirmInTime";xml:"confirmInTime"`
    
    /* confirmOutTime optional确认出库时间 */
    confirmOutTime string `json:"confirmOutTime";xml:"confirmOutTime"`
    
    /* createTime optional调拨单创建时间 */
    createTime string `json:"createTime";xml:"createTime"`
    
    /* erpOrderCode optionalerpOrderCode */
    erpOrderCode string `json:"erpOrderCode";xml:"erpOrderCode"`
    
    /* fromWarehouseCode optional调拨出库仓编码 */
    fromWarehouseCode string `json:"fromWarehouseCode";xml:"fromWarehouseCode"`
    
    /* items optional项目集 */
    items Items `json:"items";xml:"items"`
    
    /* orderStatus optionalorderStatus */
    orderStatus string `json:"orderStatus";xml:"orderStatus"`
    
    /* ownerCode optional111 */
    ownerCode string `json:"ownerCode";xml:"ownerCode"`
    
    /* toWarehouseCode optional调拨入库仓编码 */
    toWarehouseCode string `json:"toWarehouseCode";xml:"toWarehouseCode"`
    
    /* transferInOrderCode optional调拨入库单号 */
    transferInOrderCode string `json:"transferInOrderCode";xml:"transferInOrderCode"`
    
    /* transferOrderCode optional调拨单号,0,string(50),必填, */
    transferOrderCode string `json:"transferOrderCode";xml:"transferOrderCode"`
    
    /* transferOutOrderCode optional调拨出库单号 */
    transferOutOrderCode string `json:"transferOutOrderCode";xml:"transferOutOrderCode"`
    
}

func (req *TaobaoQimenTransferorderReportRequest) GetAPIName() string {
	return "taobao.qimen.transferorder.report"
}

// TaobaoQimenTransferorderReportResponse 调拨单通知
type TaobaoQimenTransferorderReportResponse struct {
    
    /* code Basic响应码,0,string(50),, */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure,success,string(10),必填, */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息,invalid appkey,string(100),, */
    message string `json:"message";xml:"message"`
    
}
