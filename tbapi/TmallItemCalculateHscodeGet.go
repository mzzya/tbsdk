package tbsdk

// TmallItemCalculateHscodeGetRequest 算法获取hscode
type TmallItemCalculateHscodeGetRequest struct {
    
    /* item_id optional商品id */
    item_id int64 `json:"item_id";xml:"item_id"`
    
}

func (req *TmallItemCalculateHscodeGetRequest) GetAPIName() string {
	return "tmall.item.calculate.hscode.get"
}

// TmallItemCalculateHscodeGetResponse 算法获取hscode
type TmallItemCalculateHscodeGetResponse struct {
    
    /* results Basic Array算法返回预测的hscode数据 */
    results map[string]interface{} `json:"results";xml:"results"`
    
}
