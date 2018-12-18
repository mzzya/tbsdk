package tbsdk

// TaobaoItemcatsAuthorizeGetRequest 查询B商家被授权品牌列表、类目列表和 c 商家新品类目列表
type TaobaoItemcatsAuthorizeGetRequest struct {
    
    /* fields required需要返回的字段。目前支持有：
brand.vid, brand.name, 
item_cat.cid, item_cat.name, item_cat.status,item_cat.sort_order,item_cat.parent_cid,item_cat.is_parent,
xinpin_item_cat.cid, 
xinpin_item_cat.name, 
xinpin_item_cat.status,
xinpin_item_cat.sort_order,
xinpin_item_cat.parent_cid,
xinpin_item_cat.is_parent */
    fields string `json:"fields";xml:"fields"`
    
}

func (req *TaobaoItemcatsAuthorizeGetRequest) GetAPIName() string {
	return "taobao.itemcats.authorize.get"
}

// TaobaoItemcatsAuthorizeGetResponse 查询B商家被授权品牌列表、类目列表和 c 商家新品类目列表
type TaobaoItemcatsAuthorizeGetResponse struct {
    
    /* seller_authorize Object里面有3个数组：
Brand[]品牌列表,
ItemCat[] 类目列表
XinpinItemCat[] 针对于C卖家新品类目列表 */
    seller_authorize SellerAuthorize `json:"seller_authorize";xml:"seller_authorize"`
    
}
