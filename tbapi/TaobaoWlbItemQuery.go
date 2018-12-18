package tbsdk

// TaobaoWlbItemQueryRequest 根据状态、卖家、SKU等信息查询商品列表
type TaobaoWlbItemQueryRequest struct {
    
    /* is_sku optional是否是最小库存单元，只有最小库存单元的商品才可以有库存,值只能给"true","false"来表示;  若值不在范围内，则按true处理 */
    is_sku string `json:"is_sku";xml:"is_sku"`
    
    /* item_code optional商家编码 */
    item_code string `json:"item_code";xml:"item_code"`
    
    /* item_type optionalITEM类型(只允许输入以下英文或空)  NORMAL 0:普通商品;  COMBINE 1:是否是组合商品  DISTRIBUTION 2:是否是分销商品(货主是别人)  若值不在范围内，则按NORMAL处理 */
    item_type string `json:"item_type";xml:"item_type"`
    
    /* name optional商品名称 */
    name string `json:"name";xml:"name"`
    
    /* page_no optional当前页 */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size optional分页记录个数，如果用户输入的记录数大于50，则一页显示50条记录 */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* parent_id optional父ID,只有is_sku=1时才能有父ID，商品也可以没有付商品 */
    parent_id int64 `json:"parent_id";xml:"parent_id"`
    
    /* status optional只能输入以下值或空：  ITEM_STATUS_VALID -- 1 表示 有效；  ITEM_STATUS_LOCK -- 2 表示锁住。  若值不在范围内，按ITEM_STATUS_VALID处理 */
    status string `json:"status";xml:"status"`
    
    /* title optional商品前台销售名字 */
    title string `json:"title";xml:"title"`
    
}

func (req *TaobaoWlbItemQueryRequest) GetAPIName() string {
	return "taobao.wlb.item.query"
}

// TaobaoWlbItemQueryResponse 根据状态、卖家、SKU等信息查询商品列表
type TaobaoWlbItemQueryResponse struct {
    
    /* item_list Object Array商品信息列表 */
    item_list WlbItem `json:"item_list";xml:"item_list"`
    
    /* total_count Basic结果总数 */
    total_count int64 `json:"total_count";xml:"total_count"`
    
}
