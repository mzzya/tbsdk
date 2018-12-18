package tbsdk

// TaobaoWlbWmsInventoryQueryRequest 支持按汇总（不分批次和渠道的总的库存数量）、渠道、批次三类方式查询商品实时库存
type TaobaoWlbWmsInventoryQueryRequest struct {
    
    /* batch_code optional库存批次号，type=2时字段传值有效。 商品设置为批次管理时，商品产生批次库存。当商品为批次管理时，此字段不传值，返回所有批次库存信息。 */
    batch_code string `json:"batch_code";xml:"batch_code"`
    
    /* channel_code optional渠道编码，type=3时字段传值有效。（TB 淘系， OTHERS 其他） */
    channel_code string `json:"channel_code";xml:"channel_code"`
    
    /* due_date optional失效日期，type=2时字段传值有效。 */
    due_date Date `json:"due_date";xml:"due_date"`
    
    /* inventory_type optional库存类型。 (1 正品 101 残次 102 机损 103 箱损 201 冻结库存 301 在途库存 ) */
    inventory_type int64 `json:"inventory_type";xml:"inventory_type"`
    
    /* item_id optional菜鸟商品ID */
    item_id string `json:"item_id";xml:"item_id"`
    
    /* page_no optional分页的第几页 */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size optional每页多少条，最大50条 */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* produce_date optional生产日期，type=2时字段传值有效。 */
    produce_date Date `json:"produce_date";xml:"produce_date"`
    
    /* store_code optional仓库编码 */
    store_code string `json:"store_code";xml:"store_code"`
    
    /* _type optional库存查询类型 1-	汇总库存，不区分批次和渠道 2-	批次库存，库存按商品批次维度划分 3-	渠道库存，库存按渠道维度划分 （当前业务不支持批次库存和渠道库存共存，批次库存无渠道属性，渠道库存无批次属性） */
    _type int64 `json:"type";xml:"type"`
    
}

func (req *TaobaoWlbWmsInventoryQueryRequest) GetAPIName() string {
	return "taobao.wlb.wms.inventory.query"
}

// TaobaoWlbWmsInventoryQueryResponse 支持按汇总（不分批次和渠道的总的库存数量）、渠道、批次三类方式查询商品实时库存
type TaobaoWlbWmsInventoryQueryResponse struct {
    
    /* item_list Object Array商品详情列表 */
    item_list WmsInventoryQueryItemlist `json:"item_list";xml:"item_list"`
    
    /* total_count Basic总数 */
    total_count int64 `json:"total_count";xml:"total_count"`
    
    /* wl_error_code Basic错误代码 */
    wl_error_code string `json:"wl_error_code";xml:"wl_error_code"`
    
    /* wl_error_msg Basic错误信息 */
    wl_error_msg string `json:"wl_error_msg";xml:"wl_error_msg"`
    
    /* wl_success Basic是否成功 */
    wl_success bool `json:"wl_success";xml:"wl_success"`
    
}
