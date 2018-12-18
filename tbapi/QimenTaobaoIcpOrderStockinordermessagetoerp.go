package tbsdk

// QimenTaobaoIcpOrderStockinordermessagetoerpRequest 供应链创建的入库单推送商家ERP
type QimenTaobaoIcpOrderStockinordermessagetoerpRequest struct {
    
    /* customerId required货主ID */
    customerId string `json:"customerId";xml:"customerId"`
    
    /* entryOrderlist optional入库单记录集 */
    entryOrderlist EntryOrderlist `json:"entryOrderlist";xml:"entryOrderlist"`
    
}

func (req *QimenTaobaoIcpOrderStockinordermessagetoerpRequest) GetAPIName() string {
	return "qimen.taobao.icp.order.stockinordermessagetoerp"
}

// QimenTaobaoIcpOrderStockinordermessagetoerpResponse 供应链创建的入库单推送商家ERP
type QimenTaobaoIcpOrderStockinordermessagetoerpResponse struct {
    
    /* result Object返回结果 */
    result Result `json:"result";xml:"result"`
    
}
