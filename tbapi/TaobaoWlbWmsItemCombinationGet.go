package tbsdk

// TaobaoWlbWmsItemCombinationGetRequest 查询组合商品的组合关系
type TaobaoWlbWmsItemCombinationGetRequest struct {
    
    /* itemid required货品Id */
    itemid int64 `json:"itemid";xml:"itemid"`
    
}

func (req *TaobaoWlbWmsItemCombinationGetRequest) GetAPIName() string {
	return "taobao.wlb.wms.item.combination.get"
}

// TaobaoWlbWmsItemCombinationGetResponse 查询组合商品的组合关系
type TaobaoWlbWmsItemCombinationGetResponse struct {
    
    /* result Object接口返回结果 */
    result Result `json:"result";xml:"result"`
    
}
