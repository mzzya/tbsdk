package tbsdk

// TaobaoFenxiaoProductSkusGetRequest 产品sku查询
type TaobaoFenxiaoProductSkusGetRequest struct {
    
    /* product_id required产品ID */
    product_id int64 `json:"product_id";xml:"product_id"`
    
}

func (req *TaobaoFenxiaoProductSkusGetRequest) GetAPIName() string {
	return "taobao.fenxiao.product.skus.get"
}

// TaobaoFenxiaoProductSkusGetResponse 产品sku查询
type TaobaoFenxiaoProductSkusGetResponse struct {
    
    /* skus Object Arraysku信息 */
    skus FenxiaoSku `json:"skus";xml:"skus"`
    
    /* total_results Basic记录数量 */
    total_results int64 `json:"total_results";xml:"total_results"`
    
}
