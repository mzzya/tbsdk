package tbsdk

// TaobaoItemsOnsaleGetRequest 获取当前用户作为卖家的出售中的商品列表，并能根据传入的搜索条件对出售中的商品列表进行过滤 
只能获得商品的部分信息，商品的详细信息请通过taobao.item.get获取
type TaobaoItemsOnsaleGetRequest struct {
    
    /* auction_type optional商品类型：a-拍卖,b-一口价 */
    auction_type string `json:"auction_type";xml:"auction_type"`
    
    /* cid optional商品类目ID。ItemCat中的cid字段。可以通过taobao.itemcats.get取到 */
    cid int64 `json:"cid";xml:"cid"`
    
    /* end_modified optional结束的修改时间 */
    end_modified Date `json:"end_modified";xml:"end_modified"`
    
    /* fields required需返回的字段列表。可选值：Item商品结构体中的以下字段： approve_status,num_iid,title,nick,type,cid,pic_url,num,props,valid_thru,list_time,price,has_discount,has_invoice,has_warranty,has_showcase,modified,delist_time,postage_id,seller_cids,outer_id,sold_quantity ；字段之间用“,”分隔。不支持其他字段，如果需要获取其他字段数据，调用taobao.item.seller.get 获取。 */
    fields string `json:"fields";xml:"fields"`
    
    /* has_discount optional是否参与会员折扣。可选值：true，false。默认不过滤该条件 */
    has_discount bool `json:"has_discount";xml:"has_discount"`
    
    /* has_showcase optional是否橱窗推荐。 可选值：true，false。默认不过滤该条件 */
    has_showcase bool `json:"has_showcase";xml:"has_showcase"`
    
    /* is_combine optional组合商品 */
    is_combine bool `json:"is_combine";xml:"is_combine"`
    
    /* is_cspu optional是否挂接了达尔文标准产品体系 */
    is_cspu bool `json:"is_cspu";xml:"is_cspu"`
    
    /* is_ex optional商品是否在外部网店显示 */
    is_ex bool `json:"is_ex";xml:"is_ex"`
    
    /* is_taobao optional商品是否在淘宝显示 */
    is_taobao bool `json:"is_taobao";xml:"is_taobao"`
    
    /* order_by optional排序方式。格式为column:asc/desc ，column可选值:list_time(上架时间),delist_time(下架时间),num(商品数量)，modified(最近修改时间)，sold_quantity（商品销量）,;默认上架时间降序(即最新上架排在前面)。如按照上架时间降序排序方式为list_time:desc */
    order_by string `json:"order_by";xml:"order_by"`
    
    /* page_no optional页码。取值范围:大于零的整数。默认值为1,即默认返回第一页数据。用此接口获取数据时，当翻页获取的条数（page_no*page_size）超过10万,为了保护后台搜索引擎，接口将报错。所以请大家尽可能的细化自己的搜索条件，例如根据修改时间分段获取商品 */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size optional每页条数。取值范围:大于零的整数;最大值：200；默认值：40。用此接口获取数据时，当翻页获取的条数（page_no*page_size）超过2万,为了保护后台搜索引擎，接口将报错。所以请大家尽可能的细化自己的搜索条件，例如根据修改时间分段获取商品 */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* q optional搜索字段。搜索商品的title。 */
    q string `json:"q";xml:"q"`
    
    /* seller_cids optional卖家店铺内自定义类目ID。多个之间用“,”分隔。可以根据taobao.sellercats.list.get获得.(<font color="red">注：目前最多支持32个ID号传入</font>) */
    seller_cids string `json:"seller_cids";xml:"seller_cids"`
    
    /* start_modified optional起始的修改时间 */
    start_modified Date `json:"start_modified";xml:"start_modified"`
    
}

func (req *TaobaoItemsOnsaleGetRequest) GetAPIName() string {
	return "taobao.items.onsale.get"
}

// TaobaoItemsOnsaleGetResponse 获取当前用户作为卖家的出售中的商品列表，并能根据传入的搜索条件对出售中的商品列表进行过滤 
只能获得商品的部分信息，商品的详细信息请通过taobao.item.get获取
type TaobaoItemsOnsaleGetResponse struct {
    
    /* items Object Array搜索到的商品列表，具体字段根据设定的fields决定，不包括desc字段 */
    items Item `json:"items";xml:"items"`
    
    /* total_results Basic搜索到符合条件的结果总数 */
    total_results int64 `json:"total_results";xml:"total_results"`
    
}
