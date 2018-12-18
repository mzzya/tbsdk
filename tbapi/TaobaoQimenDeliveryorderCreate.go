package tbsdk

// TaobaoQimenDeliveryorderCreateRequest taobao.qimen.deliveryorder.create
type TaobaoQimenDeliveryorderCreateRequest struct {
    
    /* deliveryOrder optional发货单信息 */
    deliveryOrder DeliveryOrder `json:"deliveryOrder";xml:"deliveryOrder"`
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* items optional */
    items Item `json:"items";xml:"items"`
    
    /* orderLines optional订单列表 */
    orderLines OrderLine `json:"orderLines";xml:"orderLines"`
    
}

func (req *TaobaoQimenDeliveryorderCreateRequest) GetAPIName() string {
	return "taobao.qimen.deliveryorder.create"
}

// TaobaoQimenDeliveryorderCreateResponse taobao.qimen.deliveryorder.create
type TaobaoQimenDeliveryorderCreateResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* createTime Basic订单创建时间(YYYY-MM-DD HH:MM:SS) */
    createTime string `json:"createTime";xml:"createTime"`
    
    /* deliveryOrderId Basic出库单仓储系统编码 */
    deliveryOrderId string `json:"deliveryOrderId";xml:"deliveryOrderId"`
    
    /* deliveryOrders Object Array发货单信息 */
    deliveryOrders DeliveryOrder `json:"deliveryOrders";xml:"deliveryOrders"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* logisticsCode Basic物流公司编码(统仓统配使用) */
    logisticsCode string `json:"logisticsCode";xml:"logisticsCode"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
    /* warehouseCode Basic仓库编码(统仓统配使用) */
    warehouseCode string `json:"warehouseCode";xml:"warehouseCode"`
    
}
