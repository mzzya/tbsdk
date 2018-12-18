package tbsdk

// TaobaoQimenEntryorderCreateRequest ERP调用接口，创建入库单;
type TaobaoQimenEntryorderCreateRequest struct {
    
    /* entryOrder optional入库单信息 */
    entryOrder EntryOrder `json:"entryOrder";xml:"entryOrder"`
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* items optional商品信息 */
    items Item `json:"items";xml:"items"`
    
    /* orderLines optional入库单详情 */
    orderLines OrderLine `json:"orderLines";xml:"orderLines"`
    
}

func (req *TaobaoQimenEntryorderCreateRequest) GetAPIName() string {
	return "taobao.qimen.entryorder.create"
}

// TaobaoQimenEntryorderCreateResponse ERP调用接口，创建入库单;
type TaobaoQimenEntryorderCreateResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* entryOrderId Basic仓储系统入库单编码 */
    entryOrderId string `json:"entryOrderId";xml:"entryOrderId"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
}
