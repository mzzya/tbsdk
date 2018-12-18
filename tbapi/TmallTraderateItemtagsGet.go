package tbsdk

// TmallTraderateItemtagsGetRequest 通过商品ID获取标签详细信息
type TmallTraderateItemtagsGetRequest struct {
    
    /* item_id required商品ID */
    item_id int64 `json:"item_id";xml:"item_id"`
    
}

func (req *TmallTraderateItemtagsGetRequest) GetAPIName() string {
	return "tmall.traderate.itemtags.get"
}

// TmallTraderateItemtagsGetResponse 通过商品ID获取标签详细信息
type TmallTraderateItemtagsGetResponse struct {
    
    /* tags Object Array标签列表 */
    tags TmallRateTagDetail `json:"tags";xml:"tags"`
    
}
