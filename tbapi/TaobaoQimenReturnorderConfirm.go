package tbsdk

// TaobaoQimenReturnorderConfirmRequest taobao.qimen.returnorder.confirm
type TaobaoQimenReturnorderConfirmRequest struct {
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* orderLines optional订单信息 */
    orderLines OrderLine `json:"orderLines";xml:"orderLines"`
    
    /* returnOrder optional退货单信息 */
    returnOrder ReturnOrder `json:"returnOrder";xml:"returnOrder"`
    
}

func (req *TaobaoQimenReturnorderConfirmRequest) GetAPIName() string {
	return "taobao.qimen.returnorder.confirm"
}

// TaobaoQimenReturnorderConfirmResponse taobao.qimen.returnorder.confirm
type TaobaoQimenReturnorderConfirmResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
}
