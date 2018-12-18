package tbsdk

// QimenTaobaoIcpOrderStockoutordermessagetoerpRequest 出库单信息推送接口
type QimenTaobaoIcpOrderStockoutordermessagetoerpRequest struct {
    
    /* customerId required货主ID */
    customerId string `json:"customerId";xml:"customerId"`
    
    /* entryOutOrderlist optional出库单记录集 */
    entryOutOrderlist EntryOutOrderlist `json:"entryOutOrderlist";xml:"entryOutOrderlist"`
    
}

func (req *QimenTaobaoIcpOrderStockoutordermessagetoerpRequest) GetAPIName() string {
	return "qimen.taobao.icp.order.stockoutordermessagetoerp"
}

// QimenTaobaoIcpOrderStockoutordermessagetoerpResponse 出库单信息推送接口
type QimenTaobaoIcpOrderStockoutordermessagetoerpResponse struct {
    
    /* result Object返回结果 */
    result Result `json:"result";xml:"result"`
    
}
