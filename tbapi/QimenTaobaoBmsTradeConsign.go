package tbsdk

// QimenTaobaoBmsTradeConsignRequest BMS通知ERP交易单整单出库接口
type QimenTaobaoBmsTradeConsignRequest struct {
    
    /* customerId required货主ID */
    customerId string `json:"customerId";xml:"customerId"`
    
    /* request optional请求实体 */
    request BmsTradeConsignRequest `json:"request";xml:"request"`
    
}

func (req *QimenTaobaoBmsTradeConsignRequest) GetAPIName() string {
	return "qimen.taobao.bms.trade.consign"
}

// QimenTaobaoBmsTradeConsignResponse BMS通知ERP交易单整单出库接口
type QimenTaobaoBmsTradeConsignResponse struct {
    
    /* response Object */
    response Response `json:"response";xml:"response"`
    
}
