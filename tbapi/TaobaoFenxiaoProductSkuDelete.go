package tbsdk

// TaobaoFenxiaoProductSkuDeleteRequest 根据sku properties删除sku数据
type TaobaoFenxiaoProductSkuDeleteRequest struct {
    
    /* product_id required产品id */
    product_id int64 `json:"product_id";xml:"product_id"`
    
    /* properties requiredsku属性 */
    properties string `json:"properties";xml:"properties"`
    
}

func (req *TaobaoFenxiaoProductSkuDeleteRequest) GetAPIName() string {
	return "taobao.fenxiao.product.sku.delete"
}

// TaobaoFenxiaoProductSkuDeleteResponse 根据sku properties删除sku数据
type TaobaoFenxiaoProductSkuDeleteResponse struct {
    
    /* created Basic操作时间 */
    created Date `json:"created";xml:"created"`
    
    /* result Basic操作结果 */
    result bool `json:"result";xml:"result"`
    
}
