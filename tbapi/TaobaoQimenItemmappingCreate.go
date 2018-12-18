package tbsdk

// TaobaoQimenItemmappingCreateRequest 前后端商品映射
type TaobaoQimenItemmappingCreateRequest struct {
    
    /* actionType optional奇门仓储字段,C123,string(50),必填, */
    actionType string `json:"actionType";xml:"actionType"`
    
    /* itemId optional奇门仓储字段,C123,string(50),必填, */
    itemId string `json:"itemId";xml:"itemId"`
    
    /* itemSource optional奇门仓储字段,C123,string(50),必填, */
    itemSource string `json:"itemSource";xml:"itemSource"`
    
    /* ownerCode optional奇门仓储字段,C123,string(50),必填, */
    ownerCode string `json:"ownerCode";xml:"ownerCode"`
    
    /* shopItemId optional奇门仓储字段,C123,string(50),必填, */
    shopItemId string `json:"shopItemId";xml:"shopItemId"`
    
    /* shopNick optional奇门仓储字段,C123,string(50),必填, */
    shopNick string `json:"shopNick";xml:"shopNick"`
    
    /* skuId optional奇门仓储字段,C123,string(50),必填, */
    skuId string `json:"skuId";xml:"skuId"`
    
}

func (req *TaobaoQimenItemmappingCreateRequest) GetAPIName() string {
	return "taobao.qimen.itemmapping.create"
}

// TaobaoQimenItemmappingCreateResponse 前后端商品映射
type TaobaoQimenItemmappingCreateResponse struct {
    
    /* code Basic响应码,0,string(50),, */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure,success,string(10),必填, */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息,invalid appkey,string(100),, */
    message string `json:"message";xml:"message"`
    
}
