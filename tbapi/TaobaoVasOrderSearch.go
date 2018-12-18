package tbsdk

// TaobaoVasOrderSearchRequest 用于ISV查询自己名下的应用及收费项目的订单记录（已付款订单）。
建议用于查询前一日的历史记录，不适合用作实时数据查询。
现在只能查询90天以内的数据
该接口限制每分钟所有appkey调用总和只能有800次。
type TaobaoVasOrderSearchRequest struct {
    
    /* article_code required应用收费代码，从合作伙伴后台（my.open.taobao.com）-收费管理-收费项目列表 能够获得该应用的收费代码 */
    article_code string `json:"article_code";xml:"article_code"`
    
    /* biz_order_id optional订单号 */
    biz_order_id int64 `json:"biz_order_id";xml:"biz_order_id"`
    
    /* biz_type optional订单类型，1=新订 2=续订 3=升级 4=后台赠送 5=后台自动续订 6=订单审核后生成订购关系（暂时用不到） 空=全部 */
    biz_type int64 `json:"biz_type";xml:"biz_type"`
    
    /* end_created optional订单创建时间（订购时间）结束值 */
    end_created Date `json:"end_created";xml:"end_created"`
    
    /* item_code optional收费项目代码，从合作伙伴后台（my.open.taobao.com）-收费管理-收费项目列表 能够获得收费项目代码 */
    item_code string `json:"item_code";xml:"item_code"`
    
    /* nick optional淘宝会员名 */
    nick string `json:"nick";xml:"nick"`
    
    /* order_id optional子订单号 */
    order_id int64 `json:"order_id";xml:"order_id"`
    
    /* page_no optional页码 */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size optional一页包含的记录数 */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* start_created optional订单创建时间（订购时间）起始值（当start_created和end_created都不填写时，默认返回最近90天的数据） */
    start_created Date `json:"start_created";xml:"start_created"`
    
}

func (req *TaobaoVasOrderSearchRequest) GetAPIName() string {
	return "taobao.vas.order.search"
}

// TaobaoVasOrderSearchResponse 用于ISV查询自己名下的应用及收费项目的订单记录（已付款订单）。
建议用于查询前一日的历史记录，不适合用作实时数据查询。
现在只能查询90天以内的数据
该接口限制每分钟所有appkey调用总和只能有800次。
type TaobaoVasOrderSearchResponse struct {
    
    /* article_biz_orders Object Array商品订单对象 */
    article_biz_orders ArticleBizOrder `json:"article_biz_orders";xml:"article_biz_orders"`
    
    /* total_item Basic总记录数 */
    total_item int64 `json:"total_item";xml:"total_item"`
    
}
