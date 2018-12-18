package tbsdk

// TaobaoWlbItemCombinationGetRequest 根据商品id查询商品组合关系
type TaobaoWlbItemCombinationGetRequest struct {
    
    /* item_id required要查询的组合商品id */
    item_id int64 `json:"item_id";xml:"item_id"`
    
}

func (req *TaobaoWlbItemCombinationGetRequest) GetAPIName() string {
	return "taobao.wlb.item.combination.get"
}

// TaobaoWlbItemCombinationGetResponse 根据商品id查询商品组合关系
type TaobaoWlbItemCombinationGetResponse struct {
    
    /* item_id_list Basic Array组合子商品id列表 */
    item_id_list int64 `json:"item_id_list";xml:"item_id_list"`
    
}
