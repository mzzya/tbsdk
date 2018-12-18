package tbsdk

// QimenTaobaoQimenEntryorderCallbackRequest ERP通过该接口回传预约入库单对应的商家仓出库单状态
type QimenTaobaoQimenEntryorderCallbackRequest struct {
    
    /* appkey required应用在奇门申请的appkey */
    appkey string `json:"appkey";xml:"appkey"`
    
    /* entryorderlist required入库单信息列表 */
    entryorderlist Struct `json:"entryorderlist";xml:"entryorderlist"`
    
    /* userId required货主在奇门授权的ID */
    userId string `json:"userId";xml:"userId"`
    
}

func (req *QimenTaobaoQimenEntryorderCallbackRequest) GetAPIName() string {
	return "qimen.taobao.qimen.entryorder.callback"
}

// QimenTaobaoQimenEntryorderCallbackResponse ERP通过该接口回传预约入库单对应的商家仓出库单状态
type QimenTaobaoQimenEntryorderCallbackResponse struct {
    
    /* code Basic错误码:CD001 */
    code string `json:"code";xml:"code"`
    
    /* flag Basic回传结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic异常信息 */
    message string `json:"message";xml:"message"`
    
}
