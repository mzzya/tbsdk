package tbsdk

// TaobaoQimenReplenishplanCreateRequest ERP调用奇门的接口,通知WMS创建补货计划
type TaobaoQimenReplenishplanCreateRequest struct {
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* gmtDeadTime required最晚入库时间(YYYY-MM-DD HH:MM:SS) */
    gmtDeadTime string `json:"gmtDeadTime";xml:"gmtDeadTime"`
    
    /* items optional单据信息 */
    items ReplenishplanCreateItem `json:"items";xml:"items"`
    
    /* orderCode required外部系统单号(ERP单号) */
    orderCode string `json:"orderCode";xml:"orderCode"`
    
    /* userId required商家ID */
    userId string `json:"userId";xml:"userId"`
    
    /* warehouseCode required仓库编码 */
    warehouseCode string `json:"warehouseCode";xml:"warehouseCode"`
    
}

func (req *TaobaoQimenReplenishplanCreateRequest) GetAPIName() string {
	return "taobao.qimen.replenishplan.create"
}

// TaobaoQimenReplenishplanCreateResponse ERP调用奇门的接口,通知WMS创建补货计划
type TaobaoQimenReplenishplanCreateResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
}
