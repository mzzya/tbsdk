package tbsdk

// TaobaoOcTradetagsGetRequest 根据订单查询订单标签。<br/>
返回的tag说明:1为官方标，2为自定义标，3为主站只读标签。<br/>
官方标签和自定义标签请看taobao.oc.tradetag.attach 接口说明<br/>
主站只读标签请看:http://open.taobao.com/doc/detail.htm?id=102865<br/>
type TaobaoOcTradetagsGetRequest struct {
    
    /* history optional是否查询历史标签 */
    history int64 `json:"history";xml:"history"`
    
    /* tag_names optional不填，则不做标签名称过滤 */
    tag_names string `json:"tag_names";xml:"tag_names"`
    
    /* tag_types optional不填，则默认只查询1,2。1为官方标，2为自定义标，3为主站只读标签 */
    tag_types string `json:"tag_types";xml:"tag_types"`
    
    /* tid required交易主订单id */
    tid int64 `json:"tid";xml:"tid"`
    
}

func (req *TaobaoOcTradetagsGetRequest) GetAPIName() string {
	return "taobao.oc.tradetags.get"
}

// TaobaoOcTradetagsGetResponse 根据订单查询订单标签。<br/>
返回的tag说明:1为官方标，2为自定义标，3为主站只读标签。<br/>
官方标签和自定义标签请看taobao.oc.tradetag.attach 接口说明<br/>
主站只读标签请看:http://open.taobao.com/doc/detail.htm?id=102865<br/>
type TaobaoOcTradetagsGetResponse struct {
    
    /* trade_tags Object Array返回结果 */
    trade_tags TradeTagRelationDo `json:"trade_tags";xml:"trade_tags"`
    
}
