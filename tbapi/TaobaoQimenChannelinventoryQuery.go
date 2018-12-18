package tbsdk

// TaobaoQimenChannelinventoryQueryRequest 渠道库存查询
type TaobaoQimenChannelinventoryQueryRequest struct {
    
    /* channelCodes optional奇门仓储字段 */
    channelCodes string `json:"channelCodes";xml:"channelCodes"`
    
    /* itemCodes optional奇门仓储字段 */
    itemCodes string `json:"itemCodes";xml:"itemCodes"`
    
    /* ownerCode optional奇门仓储字段,C123,string(50),, */
    ownerCode string `json:"ownerCode";xml:"ownerCode"`
    
    /* warehouseCodes optional奇门仓储字段 */
    warehouseCodes string `json:"warehouseCodes";xml:"warehouseCodes"`
    
}

func (req *TaobaoQimenChannelinventoryQueryRequest) GetAPIName() string {
	return "taobao.qimen.channelinventory.query"
}

// TaobaoQimenChannelinventoryQueryResponse 渠道库存查询
type TaobaoQimenChannelinventoryQueryResponse struct {
    
    /* code Basic响应码,0,string(50),, */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure,success,string(10),必填, */
    flag string `json:"flag";xml:"flag"`
    
    /* itemInventories Object ArrayitemInventories */
    itemInventories ItemInventory `json:"itemInventories";xml:"itemInventories"`
    
    /* message Basic响应信息,invalid appkey,string(100),, */
    message string `json:"message";xml:"message"`
    
}
