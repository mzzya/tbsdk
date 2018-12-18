package tbsdk

// TaobaoItemAnchorGetRequest 根据类目id和宝贝描述规范化打标类型获取该类目可用的宝贝描述模块中的锚点
type TaobaoItemAnchorGetRequest struct {
    
    /* cat_id required对应类目编号 */
    cat_id int64 `json:"cat_id";xml:"cat_id"`
    
    /* _type required宝贝模板类型是人工打标还是自动打标：人工打标为1，自动打标为0.人工和自动打标为-1. */
    _type int64 `json:"type";xml:"type"`
    
}

func (req *TaobaoItemAnchorGetRequest) GetAPIName() string {
	return "taobao.item.anchor.get"
}

// TaobaoItemAnchorGetResponse 根据类目id和宝贝描述规范化打标类型获取该类目可用的宝贝描述模块中的锚点
type TaobaoItemAnchorGetResponse struct {
    
    /* anchor_modules Object Array宝贝描述规范化可使用打标模块的锚点信息 */
    anchor_modules IdsModule `json:"anchor_modules";xml:"anchor_modules"`
    
    /* total_results Basic返回的宝贝描述模板结果数目 */
    total_results int64 `json:"total_results";xml:"total_results"`
    
}
