package tbsdk

// TaobaoQimenReturnorderCreateRequest ERP调用奇门的接口,创建退货单信息;该接口和入库单的区别就是该接口是从入库单接口中单独剥离出来的，专门处理退货引起的入 库操作
type TaobaoQimenReturnorderCreateRequest struct {
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* itemList optionalitemList */
    itemList Item `json:"itemList";xml:"itemList"`
    
    /* orderLines optional订单信息 */
    orderLines OrderLine `json:"orderLines";xml:"orderLines"`
    
    /* returnOrder optional退货单信息 */
    returnOrder ReturnOrder `json:"returnOrder";xml:"returnOrder"`
    
}

func (req *TaobaoQimenReturnorderCreateRequest) GetAPIName() string {
	return "taobao.qimen.returnorder.create"
}

// TaobaoQimenReturnorderCreateResponse ERP调用奇门的接口,创建退货单信息;该接口和入库单的区别就是该接口是从入库单接口中单独剥离出来的，专门处理退货引起的入 库操作
type TaobaoQimenReturnorderCreateResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
    /* returnOrderId Basic仓储系统退货单编码 */
    returnOrderId string `json:"returnOrderId";xml:"returnOrderId"`
    
}
