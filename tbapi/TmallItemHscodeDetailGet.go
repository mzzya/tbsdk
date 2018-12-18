package tbsdk

// TmallItemHscodeDetailGetRequest 通过hscode获取计量单位和销售单位
type TmallItemHscodeDetailGetRequest struct {
    
    /* hscode requiredhscode */
    hscode string `json:"hscode";xml:"hscode"`
    
}

func (req *TmallItemHscodeDetailGetRequest) GetAPIName() string {
	return "tmall.item.hscode.detail.get"
}

// TmallItemHscodeDetailGetResponse 通过hscode获取计量单位和销售单位
type TmallItemHscodeDetailGetResponse struct {
    
    /* results Basic Array返回的计量单位和销售单位 */
    results map[string]interface{} `json:"results";xml:"results"`
    
}
