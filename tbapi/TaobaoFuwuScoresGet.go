package tbsdk

// TaobaoFuwuScoresGetRequest 根据日期、查询appkey对应服务评价，每次调用只能查询某一天服务评价信息，可设置分页查询，页大小最大为100，非实时接口，延迟时间为30分钟
type TaobaoFuwuScoresGetRequest struct {
    
    /* current_page required当前页 */
    current_page int64 `json:"current_page";xml:"current_page"`
    
    /* date required评价日期，查询某一天的评价 */
    date Date `json:"date";xml:"date"`
    
    /* page_size optional每页获取条数。默认值40，最小值1，最大值100。 */
    page_size int64 `json:"page_size";xml:"page_size"`
    
}

func (req *TaobaoFuwuScoresGetRequest) GetAPIName() string {
	return "taobao.fuwu.scores.get"
}

// TaobaoFuwuScoresGetResponse 根据日期、查询appkey对应服务评价，每次调用只能查询某一天服务评价信息，可设置分页查询，页大小最大为100，非实时接口，延迟时间为30分钟
type TaobaoFuwuScoresGetResponse struct {
    
    /* score_result Object Array评价流水记录 */
    score_result ScoreResult `json:"score_result";xml:"score_result"`
    
}
