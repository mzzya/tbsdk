package tbsdk

// TaobaoInventoryIpcInventorydetailGetRequest 查询库存明细
type TaobaoInventoryIpcInventorydetailGetRequest struct {
    
    /* biz_order_id optional主订单号，可选 */
    biz_order_id int64 `json:"biz_order_id";xml:"biz_order_id"`
    
    /* biz_sub_order_id optional子订单号，可选 */
    biz_sub_order_id int64 `json:"biz_sub_order_id";xml:"biz_sub_order_id"`
    
    /* page_index required当前页数 */
    page_index int64 `json:"page_index";xml:"page_index"`
    
    /* page_size required一页显示多少条 */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* sc_item_id required仓储商品id */
    sc_item_id int64 `json:"sc_item_id";xml:"sc_item_id"`
    
    /* status_query required1:查询预扣  4：查询占用 */
    status_query int64 `json:"status_query";xml:"status_query"`
    
}

func (req *TaobaoInventoryIpcInventorydetailGetRequest) GetAPIName() string {
	return "taobao.inventory.ipc.inventorydetail.get"
}

// TaobaoInventoryIpcInventorydetailGetResponse 查询库存明细
type TaobaoInventoryIpcInventorydetailGetResponse struct {
    
    /* inventory_details Object Array库存明细列表 */
    inventory_details IpcInventoryDetailDo `json:"inventory_details";xml:"inventory_details"`
    
}
