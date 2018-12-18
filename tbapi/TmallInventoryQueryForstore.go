package tbsdk

// TmallInventoryQueryForstoreRequest 商家查询后端商品仓库库存
type TmallInventoryQueryForstoreRequest struct {
    
    /* param_list required查询列表 */
    param_list InventoryQueryForStoreRequest `json:"param_list";xml:"param_list"`
    
}

func (req *TmallInventoryQueryForstoreRequest) GetAPIName() string {
	return "tmall.inventory.query.forstore"
}

// TmallInventoryQueryForstoreResponse 商家查询后端商品仓库库存
type TmallInventoryQueryForstoreResponse struct {
    
    /* errorcode Basic错误code */
    errorcode string `json:"errorcode";xml:"errorcode"`
    
    /* errormessage Basic错误信息 */
    errormessage string `json:"errormessage";xml:"errormessage"`
    
    /* issuccess Basic整体成功或失败 */
    issuccess bool `json:"issuccess";xml:"issuccess"`
    
    /* result Object查询结果 */
    result InventoryQueryResult `json:"result";xml:"result"`
    
}
