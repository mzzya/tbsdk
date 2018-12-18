package tbsdk

// TaobaoQianniuKefuevalGetRequest 获取买家对客服的服务评价
type TaobaoQianniuKefuevalGetRequest struct {
    
    /* btime required开始时间，格式yyyyMMddHHmmss */
    btime string `json:"btime";xml:"btime"`
    
    /* etime required结束时间,格式yyyyMMddHHmmss */
    etime string `json:"etime";xml:"etime"`
    
    /* query_ids required客服的nick，多个用逗号分隔，不要超过10个，带cntaobao的长nick */
    query_ids string `json:"query_ids";xml:"query_ids"`
    
}

func (req *TaobaoQianniuKefuevalGetRequest) GetAPIName() string {
	return "taobao.qianniu.kefueval.get"
}

// TaobaoQianniuKefuevalGetResponse 获取买家对客服的服务评价
type TaobaoQianniuKefuevalGetResponse struct {
    
    /* result_count BasicresultCount */
    result_count int64 `json:"result_count";xml:"result_count"`
    
    /* staff_eval_details Object ArraystaffEvalDetails */
    staff_eval_details EvalDetail `json:"staff_eval_details";xml:"staff_eval_details"`
    
}
