package tbsdk

// TaobaoOcTradesBytagGetRequest 根据标签查询订单编号
type TaobaoOcTradesBytagGetRequest struct {
    
    /* page optional当前页 */
    page int64 `json:"page";xml:"page"`
    
    /* page_size optional分页大小 */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* tag_name required标签名称 */
    tag_name string `json:"tag_name";xml:"tag_name"`
    
    /* tag_type required标签类型，1官方，2自定义 */
    tag_type int64 `json:"tag_type";xml:"tag_type"`
    
}

func (req *TaobaoOcTradesBytagGetRequest) GetAPIName() string {
	return "taobao.oc.trades.bytag.get"
}

// TaobaoOcTradesBytagGetResponse 根据标签查询订单编号
type TaobaoOcTradesBytagGetResponse struct {
    
    /* tids Basic Array打了该标签的订单编号列表 */
    tids int64 `json:"tids";xml:"tids"`
    
    /* totals Basic总数 */
    totals int64 `json:"totals";xml:"totals"`
    
}
