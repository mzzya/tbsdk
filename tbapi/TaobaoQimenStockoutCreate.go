package tbsdk

// TaobaoQimenStockoutCreateRequest ERP调用奇门接口，创建出库单信息
type TaobaoQimenStockoutCreateRequest struct {
    
    /* deliveryOrder optional出库单信息 */
    deliveryOrder DeliveryOrder `json:"deliveryOrder";xml:"deliveryOrder"`
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* items optional */
    items Item `json:"items";xml:"items"`
    
    /* orderLines optional单据信息 */
    orderLines OrderLine `json:"orderLines";xml:"orderLines"`
    
}

func (req *TaobaoQimenStockoutCreateRequest) GetAPIName() string {
	return "taobao.qimen.stockout.create"
}

// TaobaoQimenStockoutCreateResponse ERP调用奇门接口，创建出库单信息
type TaobaoQimenStockoutCreateResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* createTime Basic订单创建时间(YYYY-MM-DD HH:MM:SS) */
    createTime string `json:"createTime";xml:"createTime"`
    
    /* deliveryOrderId Basic出库单仓储系统编码 */
    deliveryOrderId string `json:"deliveryOrderId";xml:"deliveryOrderId"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
}
