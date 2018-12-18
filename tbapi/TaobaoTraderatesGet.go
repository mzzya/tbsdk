package tbsdk

// TaobaoTraderatesGetRequest 搜索评价信息，只能获取距今180天内的评价记录(只支持查询卖家给出或得到的评价)。
type TaobaoTraderatesGetRequest struct {
    
    /* end_date optional评价结束时间。如果只输入结束时间，那么全部返回所有评价数据。 */
    end_date Date `json:"end_date";xml:"end_date"`
    
    /* fields required需返回的字段列表。可选值：TradeRate 结构中的所有字段，多个字段之间用“,”分隔 */
    fields string `json:"fields";xml:"fields"`
    
    /* num_iid optional商品的数字ID */
    num_iid int64 `json:"num_iid";xml:"num_iid"`
    
    /* page_no optional页码。取值范围:大于零的整数最大限制为200; 默认值:1 */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size optional每页获取条数。默认值40，最小值1，最大值150。 */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* peer_nick optional评价对方昵称 */
    peer_nick string `json:"peer_nick";xml:"peer_nick"`
    
    /* rate_type required评价类型。可选值:get(得到),give(给出) */
    rate_type string `json:"rate_type";xml:"rate_type"`
    
    /* result optional评价结果。可选值:good(好评),neutral(中评),bad(差评) */
    result string `json:"result";xml:"result"`
    
    /* role required评价者角色即评价的发起方。可选值:seller(卖家),buyer(买家)。 当 give buyer 以买家身份给卖家的评价； 当 get seller 以买家身份得到卖家给的评价； 当 give seller 以卖家身份给买家的评价； 当 get buyer 以卖家身份得到买家给的评价。 */
    role string `json:"role";xml:"role"`
    
    /* start_date optional评价开始时。如果只输入开始时间，那么能返回开始时间之后的评价数据。 */
    start_date Date `json:"start_date";xml:"start_date"`
    
    /* tid optional交易订单id，可以是父订单id号，也可以是子订单id号 */
    tid int64 `json:"tid";xml:"tid"`
    
    /* use_has_next optional是否启用has_next的分页方式，如果指定true,则返回的结果中不包含总记录数，但是会新增一个是否存在下一页的的字段，通过此种方式获取评价信息，效率在原有的基础上有80%的提升。 */
    use_has_next bool `json:"use_has_next";xml:"use_has_next"`
    
}

func (req *TaobaoTraderatesGetRequest) GetAPIName() string {
	return "taobao.traderates.get"
}

// TaobaoTraderatesGetResponse 搜索评价信息，只能获取距今180天内的评价记录(只支持查询卖家给出或得到的评价)。
type TaobaoTraderatesGetResponse struct {
    
    /* has_next Basic当使用use_has_next时返回信息，如果还有下一页则返回true */
    has_next bool `json:"has_next";xml:"has_next"`
    
    /* total_results Basic搜索到的评价总数。相同的查询时间段条件下，最大只能获取总共1500条评价记录。所以当评价大于等于1500时 ISV可以通过start_date及end_date来进行拆分，以保证可以查询到全部数据 */
    total_results int64 `json:"total_results";xml:"total_results"`
    
    /* trade_rates Object Array评价列表。返回的TradeRate包含的具体信息为入参fields请求的字段信息 */
    trade_rates TradeRate `json:"trade_rates";xml:"trade_rates"`
    
}
