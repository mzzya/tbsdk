package tbsdk

// TaobaoOpenuidGetBytradeRequest 通过订单获取对应买家的openUID,需要卖家授权
type TaobaoOpenuidGetBytradeRequest struct {
    
    /* tid required订单ID */
    tid int64 `json:"tid";xml:"tid"`
    
}

func (req *TaobaoOpenuidGetBytradeRequest) GetAPIName() string {
	return "taobao.openuid.get.bytrade"
}

// TaobaoOpenuidGetBytradeResponse 通过订单获取对应买家的openUID,需要卖家授权
type TaobaoOpenuidGetBytradeResponse struct {
    
    /* open_uid Basic当前交易tid对应买家的openuid */
    open_uid string `json:"open_uid";xml:"open_uid"`
    
}
