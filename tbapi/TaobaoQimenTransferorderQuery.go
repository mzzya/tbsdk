package tbsdk

// TaobaoQimenTransferorderQueryRequest 调拨单查询
type TaobaoQimenTransferorderQueryRequest struct {
    
    /* erpOrderCode optionalERP单号 */
    erpOrderCode string `json:"erpOrderCode";xml:"erpOrderCode"`
    
    /* ownerCode required111 */
    ownerCode string `json:"ownerCode";xml:"ownerCode"`
    
    /* transferOrderCode required调拨单号 */
    transferOrderCode string `json:"transferOrderCode";xml:"transferOrderCode"`
    
}

func (req *TaobaoQimenTransferorderQueryRequest) GetAPIName() string {
	return "taobao.qimen.transferorder.query"
}

// TaobaoQimenTransferorderQueryResponse 调拨单查询
type TaobaoQimenTransferorderQueryResponse struct {
    
    /* code Basic响应码,0,string(50),, */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure,success,string(10),必填, */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息,invalid appkey,string(100),, */
    message string `json:"message";xml:"message"`
    
    /* transferOrderDetail Object调拨单细节 */
    transferOrderDetail TransferOrderDetail `json:"transferOrderDetail";xml:"transferOrderDetail"`
    
}
