package tbsdk

// TaobaoQimenCombineitemDeleteRequest 组合货品删除
type TaobaoQimenCombineitemDeleteRequest struct {
    
    /* itemId optional奇门仓储字段,C123,string(50),, */
    itemId string `json:"itemId";xml:"itemId"`
    
    /* ownerCode optional奇门仓储字段,C123,string(50),, */
    ownerCode string `json:"ownerCode";xml:"ownerCode"`
    
}

func (req *TaobaoQimenCombineitemDeleteRequest) GetAPIName() string {
	return "taobao.qimen.combineitem.delete"
}

// TaobaoQimenCombineitemDeleteResponse 组合货品删除
type TaobaoQimenCombineitemDeleteResponse struct {
    
    /* code Basic响应码,0,string(50),, */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure,success,string(10),必填, */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息,invalid appkey,string(100),, */
    message string `json:"message";xml:"message"`
    
}
