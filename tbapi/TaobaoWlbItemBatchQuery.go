package tbsdk

// TaobaoWlbItemBatchQueryRequest 根据用户id，item id list和store code来查询商品库存信息和批次信息
type TaobaoWlbItemBatchQueryRequest struct {
    
    /* item_ids required需要查询的商品ID列表，以字符串表示，ID间以;隔开 */
    item_ids string `json:"item_ids";xml:"item_ids"`
    
    /* page_no optional分页查询参数，指定查询页数，默认为1 */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size optional分页查询参数，每页查询数量，默认20，最大值50,大于50时按照50条查询 */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* store_code optional仓库编号 */
    store_code string `json:"store_code";xml:"store_code"`
    
}

func (req *TaobaoWlbItemBatchQueryRequest) GetAPIName() string {
	return "taobao.wlb.item.batch.query"
}

// TaobaoWlbItemBatchQueryResponse 根据用户id，item id list和store code来查询商品库存信息和批次信息
type TaobaoWlbItemBatchQueryResponse struct {
    
    /* item_inventory_batch_list Object Array商品库存及批次信息查询结果 */
    item_inventory_batch_list WlbItemBatchInventory `json:"item_inventory_batch_list";xml:"item_inventory_batch_list"`
    
    /* total_count Basic返回结果记录的数量 */
    total_count int64 `json:"total_count";xml:"total_count"`
    
}
