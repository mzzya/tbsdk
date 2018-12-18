package tbsdk

// TaobaoFenxiaoProductImageUploadRequest 产品主图图片空间相对路径或绝对路径添加或更新，或者是图片上传。如果指定位置的图片已存在，则覆盖原有信息。如果位置为1,自动设为主图；如果位置为0，表示属性图片
type TaobaoFenxiaoProductImageUploadRequest struct {
    
    /* image special产品图片 */
    image byte[] `json:"image";xml:"image"`
    
    /* pic_path special产品主图图片空间相对路径或绝对路径 */
    pic_path string `json:"pic_path";xml:"pic_path"`
    
    /* position required图片位置，0-14之间。0：操作sku属性图片，1：主图，2-5：细节图，6-14：额外主图 */
    position int64 `json:"position";xml:"position"`
    
    /* product_id required产品ID */
    product_id int64 `json:"product_id";xml:"product_id"`
    
    /* properties optionalproperties表示sku图片的属性。key:value形式，key是pid，value是vid。如果position是0的话，则properties需要是必传项 */
    properties string `json:"properties";xml:"properties"`
    
}

func (req *TaobaoFenxiaoProductImageUploadRequest) GetAPIName() string {
	return "taobao.fenxiao.product.image.upload"
}

// TaobaoFenxiaoProductImageUploadResponse 产品主图图片空间相对路径或绝对路径添加或更新，或者是图片上传。如果指定位置的图片已存在，则覆盖原有信息。如果位置为1,自动设为主图；如果位置为0，表示属性图片
type TaobaoFenxiaoProductImageUploadResponse struct {
    
    /* created Basic操作时间 */
    created Date `json:"created";xml:"created"`
    
    /* result Basic操作是否成功 */
    result bool `json:"result";xml:"result"`
    
}
