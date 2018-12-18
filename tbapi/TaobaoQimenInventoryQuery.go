package tbsdk

// TaobaoQimenInventoryQueryRequest ERP调用奇门的接口,查询商品的库存量
type TaobaoQimenInventoryQueryRequest struct {
    
    /* criteriaList optional查询准则 */
    criteriaList Criteria `json:"criteriaList";xml:"criteriaList"`
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* remark optional备注 */
    remark string `json:"remark";xml:"remark"`
    
}

func (req *TaobaoQimenInventoryQueryRequest) GetAPIName() string {
	return "taobao.qimen.inventory.query"
}

// TaobaoQimenInventoryQueryResponse ERP调用奇门的接口,查询商品的库存量
type TaobaoQimenInventoryQueryResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* items Object Array商品的库存信息列表 */
    items Item `json:"items";xml:"items"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
}
