package tbsdk

// TaobaoQimenCombineitemQueryRequest 组合货品关系查询
type TaobaoQimenCombineitemQueryRequest struct {
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* itemId optional奇门仓储字段 */
    itemId string `json:"itemId";xml:"itemId"`
    
    /* ownerCode optional奇门仓储字段 */
    ownerCode string `json:"ownerCode";xml:"ownerCode"`
    
}

func (req *TaobaoQimenCombineitemQueryRequest) GetAPIName() string {
	return "taobao.qimen.combineitem.query"
}

// TaobaoQimenCombineitemQueryResponse 组合货品关系查询
type TaobaoQimenCombineitemQueryResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* combItems Object Array奇门仓储字段 */
    combItems CombItem `json:"combItems";xml:"combItems"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
}
