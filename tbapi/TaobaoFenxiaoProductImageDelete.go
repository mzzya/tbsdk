package tbsdk

// TaobaoFenxiaoProductImageDeleteRequest 产品图片删除，只删除图片信息，不真正删除图片
type TaobaoFenxiaoProductImageDeleteRequest struct {
    
    /* position required图片位置 */
    position int64 `json:"position";xml:"position"`
    
    /* product_id required产品ID */
    product_id int64 `json:"product_id";xml:"product_id"`
    
    /* properties optionalproperties表示sku图片的属性。key:value形式，key是pid，value是vid。如果position是0的话，则properties需要是必传项 */
    properties string `json:"properties";xml:"properties"`
    
}

func (req *TaobaoFenxiaoProductImageDeleteRequest) GetAPIName() string {
	return "taobao.fenxiao.product.image.delete"
}

// TaobaoFenxiaoProductImageDeleteResponse 产品图片删除，只删除图片信息，不真正删除图片
type TaobaoFenxiaoProductImageDeleteResponse struct {
    
    /* created Basic操作时间 */
    created Date `json:"created";xml:"created"`
    
    /* result Basic操作结果 */
    result bool `json:"result";xml:"result"`
    
}
