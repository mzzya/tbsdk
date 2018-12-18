package tbsdk

// TaobaoQimenStockoutConfirmRequest 货品出库后，WMS将状态回传给ERP
type TaobaoQimenStockoutConfirmRequest struct {
    
    /* deliveryOrder optionaldeliveryOrder */
    deliveryOrder DeliveryOrder `json:"deliveryOrder";xml:"deliveryOrder"`
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* orderLines optionalorderLines */
    orderLines OrderLine `json:"orderLines";xml:"orderLines"`
    
    /* packages optionalpackages */
    packages Package `json:"packages";xml:"packages"`
    
}

func (req *TaobaoQimenStockoutConfirmRequest) GetAPIName() string {
	return "taobao.qimen.stockout.confirm"
}

// TaobaoQimenStockoutConfirmResponse 货品出库后，WMS将状态回传给ERP
type TaobaoQimenStockoutConfirmResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
}
