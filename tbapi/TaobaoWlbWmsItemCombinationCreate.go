package tbsdk

// TaobaoWlbWmsItemCombinationCreateRequest 创建组合商品与子商品关系
type TaobaoWlbWmsItemCombinationCreateRequest struct {
    
    /* item_id required组合商品ID */
    item_id int64 `json:"item_id";xml:"item_id"`
    
    /* sub_item_list required子货品列表 */
    sub_item_list SubItemList `json:"sub_item_list";xml:"sub_item_list"`
    
}

func (req *TaobaoWlbWmsItemCombinationCreateRequest) GetAPIName() string {
	return "taobao.wlb.wms.item.combination.create"
}

// TaobaoWlbWmsItemCombinationCreateResponse 创建组合商品与子商品关系
type TaobaoWlbWmsItemCombinationCreateResponse struct {
    
    /* wl_error_code Basic错误编码 */
    wl_error_code string `json:"wl_error_code";xml:"wl_error_code"`
    
    /* wl_error_msg Basic错误信息 */
    wl_error_msg string `json:"wl_error_msg";xml:"wl_error_msg"`
    
    /* wl_success Basic是否成功 */
    wl_success bool `json:"wl_success";xml:"wl_success"`
    
}
