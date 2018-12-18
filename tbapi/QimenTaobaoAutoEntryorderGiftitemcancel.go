package tbsdk

// QimenTaobaoAutoEntryorderGiftitemcancelRequest 该接口一次只能回传一个主交易号的BMS增加的货品取消信息
type QimenTaobaoAutoEntryorderGiftitemcancelRequest struct {
    
    /* customerId requiredcustomerId */
    customerId string `json:"customerId";xml:"customerId"`
    
    /* request optional请求 */
    request ErpBmsOrderGiftCancelRequest `json:"request";xml:"request"`
    
}

func (req *QimenTaobaoAutoEntryorderGiftitemcancelRequest) GetAPIName() string {
	return "qimen.taobao.auto.entryorder.giftitemcancel"
}

// QimenTaobaoAutoEntryorderGiftitemcancelResponse 该接口一次只能回传一个主交易号的BMS增加的货品取消信息
type QimenTaobaoAutoEntryorderGiftitemcancelResponse struct {
    
    /* response Objectresponse */
    response Response `json:"response";xml:"response"`
    
}
