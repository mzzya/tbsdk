package tbsdk

// QimenTaobaoUopTmsupdatemessagetoerpRequest ERP配消息回告（拒签、签收、揽收等）
type QimenTaobaoUopTmsupdatemessagetoerpRequest struct {
    
    /* request optional */
    request Struct `json:"request";xml:"request"`
    
}

func (req *QimenTaobaoUopTmsupdatemessagetoerpRequest) GetAPIName() string {
	return "qimen.taobao.uop.tmsupdatemessagetoerp"
}

// QimenTaobaoUopTmsupdatemessagetoerpResponse ERP配消息回告（拒签、签收、揽收等）
type QimenTaobaoUopTmsupdatemessagetoerpResponse struct {
    
    /* response Objectresponse */
    response Response `json:"response";xml:"response"`
    
}
