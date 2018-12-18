package tbsdk

// TaobaoQimenTagItemsQueryRequest 调用该接口，查询打了某个标的商品列表。说明：该接口调用后，返回值的时间较长，建议不要经常调用。
type TaobaoQimenTagItemsQueryRequest struct {
    
    /* remark optional备注，string（500） */
    remark string `json:"remark";xml:"remark"`
    
    /* tag_type required打标值，string（50），TBKU=同步库存标，MDZT=门店自提标，必填 */
    tag_type string `json:"tag_type";xml:"tag_type"`
    
}

func (req *TaobaoQimenTagItemsQueryRequest) GetAPIName() string {
	return "taobao.qimen.tag.items.query"
}

// TaobaoQimenTagItemsQueryResponse 调用该接口，查询打了某个标的商品列表。说明：该接口调用后，返回值的时间较长，建议不要经常调用。
type TaobaoQimenTagItemsQueryResponse struct {
    
    /* flag Basicflag */
    flag string `json:"flag";xml:"flag"`
    
    /* item_ids Basic ArrayitemIds */
    item_ids int64 `json:"item_ids";xml:"item_ids"`
    
    /* message Basicmessage */
    message string `json:"message";xml:"message"`
    
    /* tag_type BasictagType */
    tag_type string `json:"tag_type";xml:"tag_type"`
    
}
