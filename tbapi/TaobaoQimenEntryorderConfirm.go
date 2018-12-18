package tbsdk

// TaobaoQimenEntryorderConfirmRequest WMS调用接口，回传入库单信息;
type TaobaoQimenEntryorderConfirmRequest struct {
    
    /* entryOrder optional入库单信息 */
    entryOrder EntryOrder `json:"entryOrder";xml:"entryOrder"`
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* items optional商品信息列表 */
    items Item `json:"items";xml:"items"`
    
    /* orderLines optional订单信息 */
    orderLines OrderLine `json:"orderLines";xml:"orderLines"`
    
}

func (req *TaobaoQimenEntryorderConfirmRequest) GetAPIName() string {
	return "taobao.qimen.entryorder.confirm"
}

// TaobaoQimenEntryorderConfirmResponse WMS调用接口，回传入库单信息;
type TaobaoQimenEntryorderConfirmResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
}
