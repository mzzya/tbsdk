package tbsdk

// QimenTaobaoIcpOrderStockoutordecanceltoerpRequest 出入库单取消推送接口
type QimenTaobaoIcpOrderStockoutordecanceltoerpRequest struct {
    
    /* customerId required出库单所属货主ID */
    customerId string `json:"customerId";xml:"customerId"`
    
    /* entryOrderlist optional出入库单信息（一单） */
    entryOrderlist EntryOrderlist `json:"entryOrderlist";xml:"entryOrderlist"`
    
}

func (req *QimenTaobaoIcpOrderStockoutordecanceltoerpRequest) GetAPIName() string {
	return "qimen.taobao.icp.order.stockoutordecanceltoerp"
}

// QimenTaobaoIcpOrderStockoutordecanceltoerpResponse 出入库单取消推送接口
type QimenTaobaoIcpOrderStockoutordecanceltoerpResponse struct {
    
    /* response Object */
    response Response `json:"response";xml:"response"`
    
}
