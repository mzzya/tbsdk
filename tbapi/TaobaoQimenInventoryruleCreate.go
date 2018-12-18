package tbsdk

// TaobaoQimenInventoryruleCreateRequest 渠道间库存规则设置
type TaobaoQimenInventoryruleCreateRequest struct {
    
    /* inventoryRules optionalinventoryRules */
    inventoryRules InventoryRule `json:"inventoryRules";xml:"inventoryRules"`
    
    /* ownerCode optional奇门仓储字段,C1223,string(50),, */
    ownerCode string `json:"ownerCode";xml:"ownerCode"`
    
}

func (req *TaobaoQimenInventoryruleCreateRequest) GetAPIName() string {
	return "taobao.qimen.inventoryrule.create"
}

// TaobaoQimenInventoryruleCreateResponse 渠道间库存规则设置
type TaobaoQimenInventoryruleCreateResponse struct {
    
    /* code Basic响应码,0,string(50),, */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure,success,string(10),必填, */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息,invalid appkey,string(100),, */
    message string `json:"message";xml:"message"`
    
}
