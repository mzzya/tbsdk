package tbsdk

// QimenTaobaoIcpOrderStockinordecanceltoerpRequest 一盘货入库单取消消息通知
type QimenTaobaoIcpOrderStockinordecanceltoerpRequest struct {
    
    /* customerId required货主ID */
    customerId string `json:"customerId";xml:"customerId"`
    
    /* entryOrderlist optional订单列表 */
    entryOrderlist EntryOrderlist `json:"entryOrderlist";xml:"entryOrderlist"`
    
}

func (req *QimenTaobaoIcpOrderStockinordecanceltoerpRequest) GetAPIName() string {
	return "qimen.taobao.icp.order.stockinordecanceltoerp"
}

// QimenTaobaoIcpOrderStockinordecanceltoerpResponse 一盘货入库单取消消息通知
type QimenTaobaoIcpOrderStockinordecanceltoerpResponse struct {
    
    /* response Object */
    response Response `json:"response";xml:"response"`
    
}
