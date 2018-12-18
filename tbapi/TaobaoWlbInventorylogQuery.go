package tbsdk

// TaobaoWlbInventorylogQueryRequest 通过商品ID等几个条件来分页查询库存变更记录
type TaobaoWlbInventorylogQueryRequest struct {
    
    /* gmt_end optional结束修改时间,小于等于该时间 */
    gmt_end Date `json:"gmt_end";xml:"gmt_end"`
    
    /* gmt_start optional起始修改时间,大于等于该时间 */
    gmt_start Date `json:"gmt_start";xml:"gmt_start"`
    
    /* item_id optional商品ID */
    item_id int64 `json:"item_id";xml:"item_id"`
    
    /* op_type optional库存操作作类型(可以为空) CHU_KU 1-出库 RU_KU 2-入库 FREEZE 3-冻结 THAW 4-解冻 CHECK_FREEZE 5-冻结确认 CHANGE_KU 6-库存类型变更 若值不在范围内，则按CHU_KU处理 */
    op_type string `json:"op_type";xml:"op_type"`
    
    /* op_user_id optional可指定授权的用户来查询 */
    op_user_id int64 `json:"op_user_id";xml:"op_user_id"`
    
    /* order_code optional单号 */
    order_code string `json:"order_code";xml:"order_code"`
    
    /* page_no optional当前页 */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size optional分页记录个数 */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* store_code optional仓库编码 */
    store_code string `json:"store_code";xml:"store_code"`
    
}

func (req *TaobaoWlbInventorylogQueryRequest) GetAPIName() string {
	return "taobao.wlb.inventorylog.query"
}

// TaobaoWlbInventorylogQueryResponse 通过商品ID等几个条件来分页查询库存变更记录
type TaobaoWlbInventorylogQueryResponse struct {
    
    /* inventory_log_list Object Array库存变更记录 */
    inventory_log_list WlbItemInventoryLog `json:"inventory_log_list";xml:"inventory_log_list"`
    
    /* total_count Basic记录数 */
    total_count int64 `json:"total_count";xml:"total_count"`
    
}
