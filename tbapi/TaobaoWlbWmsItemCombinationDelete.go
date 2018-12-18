package tbsdk

// TaobaoWlbWmsItemCombinationDeleteRequest 删除货品组合关系
type TaobaoWlbWmsItemCombinationDeleteRequest struct {
    
    /* item_id optional组合货品ID */
    item_id int64 `json:"item_id";xml:"item_id"`
    
}

func (req *TaobaoWlbWmsItemCombinationDeleteRequest) GetAPIName() string {
	return "taobao.wlb.wms.item.combination.delete"
}

// TaobaoWlbWmsItemCombinationDeleteResponse 删除货品组合关系
type TaobaoWlbWmsItemCombinationDeleteResponse struct {
    
    /* wl_error_code Basic错误编码 */
    wl_error_code string `json:"wl_error_code";xml:"wl_error_code"`
    
    /* wl_error_msg Basic错误信息 */
    wl_error_msg string `json:"wl_error_msg";xml:"wl_error_msg"`
    
    /* wl_success Basic是否成功 */
    wl_success bool `json:"wl_success";xml:"wl_success"`
    
}
