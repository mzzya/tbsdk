package tbsdk

// TaobaoWlbItemMapGetByExtentityRequest 根据外部实体类型和实体id查询映射的物流宝商品id
type TaobaoWlbItemMapGetByExtentityRequest struct {
    
    /* ext_entity_id required外部实体类型对应的商品id */
    ext_entity_id int64 `json:"ext_entity_id";xml:"ext_entity_id"`
    
    /* ext_entity_type required外部实体类型： IC_ITEM--ic商品 IC_SKU--ic销售单元 */
    ext_entity_type string `json:"ext_entity_type";xml:"ext_entity_type"`
    
}

func (req *TaobaoWlbItemMapGetByExtentityRequest) GetAPIName() string {
	return "taobao.wlb.item.map.get.by.extentity"
}

// TaobaoWlbItemMapGetByExtentityResponse 根据外部实体类型和实体id查询映射的物流宝商品id
type TaobaoWlbItemMapGetByExtentityResponse struct {
    
    /* is_success Basic是否成功 */
    is_success bool `json:"is_success";xml:"is_success"`
    
    /* item_id Basic物流宝商品id */
    item_id int64 `json:"item_id";xml:"item_id"`
    
}
