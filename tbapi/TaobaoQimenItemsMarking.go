package tbsdk

// TaobaoQimenItemsMarkingRequest 调用该接口，对商品进行XXXX标的打标、去标的动作。
type TaobaoQimenItemsMarkingRequest struct {
    
    /* action_type required操作类型，string（50），ADD=打标，DELETE=去标，必填 */
    action_type string `json:"action_type";xml:"action_type"`
    
    /* item_ids required线上商品ID，long，必填 */
    item_ids int64 `json:"item_ids";xml:"item_ids"`
    
    /* remark optional备注，string（500） */
    remark string `json:"remark";xml:"remark"`
    
    /* tag_type required打标值，string（50），TBKU=同步库存标，MDZT=门店自提标，必填 */
    tag_type string `json:"tag_type";xml:"tag_type"`
    
}

func (req *TaobaoQimenItemsMarkingRequest) GetAPIName() string {
	return "taobao.qimen.items.marking"
}

// TaobaoQimenItemsMarkingResponse 调用该接口，对商品进行XXXX标的打标、去标的动作。
type TaobaoQimenItemsMarkingResponse struct {
    
    /* flag Basicflag */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basicmessage */
    message string `json:"message";xml:"message"`
    
}
