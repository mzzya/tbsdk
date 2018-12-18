package tbsdk

// TaobaoQimenEntryorderQueryRequest ERP调用接口，查询入库单信息;
type TaobaoQimenEntryorderQueryRequest struct {
    
    /* entryOrderCode required入库单编码 */
    entryOrderCode string `json:"entryOrderCode";xml:"entryOrderCode"`
    
    /* entryOrderId optional仓储系统入库单ID */
    entryOrderId string `json:"entryOrderId";xml:"entryOrderId"`
    
    /* extOrderCode optionalextOrderCode */
    extOrderCode string `json:"extOrderCode";xml:"extOrderCode"`
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* orderType optionalorderType */
    orderType string `json:"orderType";xml:"orderType"`
    
    /* ownerCode optional货主编码 */
    ownerCode string `json:"ownerCode";xml:"ownerCode"`
    
    /* page required当前页(从1开始) */
    page int64 `json:"page";xml:"page"`
    
    /* pageSize required每页orderLine条数(最多100条) */
    pageSize int64 `json:"pageSize";xml:"pageSize"`
    
    /* remark optional备注 */
    remark string `json:"remark";xml:"remark"`
    
    /* warehouseCode optional仓库编码 */
    warehouseCode string `json:"warehouseCode";xml:"warehouseCode"`
    
}

func (req *TaobaoQimenEntryorderQueryRequest) GetAPIName() string {
	return "taobao.qimen.entryorder.query"
}

// TaobaoQimenEntryorderQueryResponse ERP调用接口，查询入库单信息;
type TaobaoQimenEntryorderQueryResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* entryOrder Object入库单信息 */
    entryOrder EntryOrder `json:"entryOrder";xml:"entryOrder"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
    /* orderLines Object Array入库单单据信息 */
    orderLines OrderLine `json:"orderLines";xml:"orderLines"`
    
    /* totalLines BasicorderLines总条数 */
    totalLines int64 `json:"totalLines";xml:"totalLines"`
    
}
