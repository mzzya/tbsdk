package tbsdk

// QimenTaobaoBmsErptradeTransferconsignRequest BMS调用ERP订单菜鸟仓&商家仓互转接口
type QimenTaobaoBmsErptradeTransferconsignRequest struct {
    
    /* customerId required货主ID */
    customerId string `json:"customerId";xml:"customerId"`
    
    /* request optional请求体 */
    request BmsErptradeTransferConsignRequest `json:"request";xml:"request"`
    
}

func (req *QimenTaobaoBmsErptradeTransferconsignRequest) GetAPIName() string {
	return "qimen.taobao.bms.erptrade.transferconsign"
}

// QimenTaobaoBmsErptradeTransferconsignResponse BMS调用ERP订单菜鸟仓&商家仓互转接口
type QimenTaobaoBmsErptradeTransferconsignResponse struct {
    
    /* response Object */
    response Response `json:"response";xml:"response"`
    
}
