package tbsdk

// TaobaoWlbWmsConsignInventoryOccupancyRequest ERP发货库存预占用
type TaobaoWlbWmsConsignInventoryOccupancyRequest struct {
    
    /* content optional库存占用 */
    content Content `json:"content";xml:"content"`
    
}

func (req *TaobaoWlbWmsConsignInventoryOccupancyRequest) GetAPIName() string {
	return "taobao.wlb.wms.consign.inventory.occupancy"
}

// TaobaoWlbWmsConsignInventoryOccupancyResponse ERP发货库存预占用
type TaobaoWlbWmsConsignInventoryOccupancyResponse struct {
    
    /* is_retry Basic返回失败时，是否需求重试，为true时，建议系统自动重试 */
    is_retry bool `json:"is_retry";xml:"is_retry"`
    
    /* item_inventory_list Object Array库存占用明细列表 */
    item_inventory_list Iteminventorylist `json:"item_inventory_list";xml:"item_inventory_list"`
    
    /* order_code BasicERP订单编码 */
    order_code string `json:"order_code";xml:"order_code"`
    
    /* wl_error_code Basic错误码 */
    wl_error_code string `json:"wl_error_code";xml:"wl_error_code"`
    
    /* wl_error_msg Basic错误信息 */
    wl_error_msg string `json:"wl_error_msg";xml:"wl_error_msg"`
    
    /* wl_success Basic是否成功 */
    wl_success string `json:"wl_success";xml:"wl_success"`
    
}
