package tbsdk

// TaobaoQimenStoreprocessConfirmRequest WMS调用奇门的接口,回传仓内加工单创建情况
type TaobaoQimenStoreprocessConfirmRequest struct {
    
    /* actualQty optional实际作业总数量 */
    actualQty int64 `json:"actualQty";xml:"actualQty"`
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* materialitems optional加工商品列表 */
    materialitems MaterialItem `json:"materialitems";xml:"materialitems"`
    
    /* orderCompleteTime required加工单完成时间(YYYY-MM-DD HH:MM:SS) */
    orderCompleteTime string `json:"orderCompleteTime";xml:"orderCompleteTime"`
    
    /* orderType required单据类型(CNJG=仓内加工作业单) */
    orderType string `json:"orderType";xml:"orderType"`
    
    /* outBizCode required外部业务编码(一个合作伙伴中要求唯一多次确认时;每次传入要求唯一(一般传入WMS损益单据编码)) */
    outBizCode string `json:"outBizCode";xml:"outBizCode"`
    
    /* ownerCode optional货主编码 */
    ownerCode string `json:"ownerCode";xml:"ownerCode"`
    
    /* processOrderCode required加工单编码 */
    processOrderCode string `json:"processOrderCode";xml:"processOrderCode"`
    
    /* processOrderId optional仓储系统加工单ID */
    processOrderId string `json:"processOrderId";xml:"processOrderId"`
    
    /* productitems optional加工商品列表 */
    productitems ProductItem `json:"productitems";xml:"productitems"`
    
    /* remark optional备注 */
    remark string `json:"remark";xml:"remark"`
    
}

func (req *TaobaoQimenStoreprocessConfirmRequest) GetAPIName() string {
	return "taobao.qimen.storeprocess.confirm"
}

// TaobaoQimenStoreprocessConfirmResponse WMS调用奇门的接口,回传仓内加工单创建情况
type TaobaoQimenStoreprocessConfirmResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
}
