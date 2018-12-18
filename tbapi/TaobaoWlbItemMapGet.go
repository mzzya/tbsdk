package tbsdk

// TaobaoWlbItemMapGetRequest 根据物流宝商品ID查询商品映射关系
type TaobaoWlbItemMapGetRequest struct {
    
    /* item_id required要查询映射关系的物流宝商品id */
    item_id int64 `json:"item_id";xml:"item_id"`
    
}

func (req *TaobaoWlbItemMapGetRequest) GetAPIName() string {
	return "taobao.wlb.item.map.get"
}

// TaobaoWlbItemMapGetResponse 根据物流宝商品ID查询商品映射关系
type TaobaoWlbItemMapGetResponse struct {
    
    /* is_success Basic是否成功 */
    is_success bool `json:"is_success";xml:"is_success"`
    
    /* out_entity_item_list Object Array外部商品实体 */
    out_entity_item_list OutEntityItem `json:"out_entity_item_list";xml:"out_entity_item_list"`
    
}
