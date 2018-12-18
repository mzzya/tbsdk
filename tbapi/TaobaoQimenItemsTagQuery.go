package tbsdk

// TaobaoQimenItemsTagQueryRequest 调用该接口，查询某个/某批商品上的标
type TaobaoQimenItemsTagQueryRequest struct {
    
    /* item_ids required线上淘宝商品ID，long，必填 */
    item_ids int64 `json:"item_ids";xml:"item_ids"`
    
}

func (req *TaobaoQimenItemsTagQueryRequest) GetAPIName() string {
	return "taobao.qimen.items.tag.query"
}

// TaobaoQimenItemsTagQueryResponse 调用该接口，查询某个/某批商品上的标
type TaobaoQimenItemsTagQueryResponse struct {
    
    /* flag Basicflag */
    flag string `json:"flag";xml:"flag"`
    
    /* item_tags Object ArrayitemTags */
    item_tags ItemTag `json:"item_tags";xml:"item_tags"`
    
    /* message Basicmessage */
    message string `json:"message";xml:"message"`
    
}
