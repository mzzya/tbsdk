package tbsdk

// TaobaoItemQuantityUpdateRequest 提供按照全量或增量形式修改宝贝/SKU库存的功能
type TaobaoItemQuantityUpdateRequest struct {
    
    /* num_iid required商品数字ID，必填参数 */
    num_iid int64 `json:"num_iid";xml:"num_iid"`
    
    /* outer_id optionalSKU的商家编码，可选参数。如果不填则默认修改宝贝的库存，如果填了则按照商家编码搜索出对应的SKU并修改库存。当sku_id和本字段都填写时以sku_id为准搜索对应SKU */
    outer_id string `json:"outer_id";xml:"outer_id"`
    
    /* quantity required库存修改值，必选。当全量更新库存时，quantity必须为大于等于0的正整数；当增量更新库存时，quantity为整数，可小于等于0。若增量更新时传入的库存为负数，则负数与实际库存之和不能小于0。比如当前实际库存为1，传入增量更新quantity=-1，库存改为0 */
    quantity int64 `json:"quantity";xml:"quantity"`
    
    /* sku_id optional要操作的SKU的数字ID，可选。如果不填默认修改宝贝的库存，如果填上则修改该SKU的库存 */
    sku_id int64 `json:"sku_id";xml:"sku_id"`
    
    /* _type optional库存更新方式，可选。1为全量更新，2为增量更新。如果不填，默认为全量更新 */
    _type int64 `json:"type";xml:"type"`
    
}

func (req *TaobaoItemQuantityUpdateRequest) GetAPIName() string {
	return "taobao.item.quantity.update"
}

// TaobaoItemQuantityUpdateResponse 提供按照全量或增量形式修改宝贝/SKU库存的功能
type TaobaoItemQuantityUpdateResponse struct {
    
    /* item Objectiid、numIid、num和modified，skus中每个sku的skuId、quantity和modified */
    item Item `json:"item";xml:"item"`
    
}
