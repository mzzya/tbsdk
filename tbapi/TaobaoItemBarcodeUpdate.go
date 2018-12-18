package tbsdk

// TaobaoItemBarcodeUpdateRequest 通过该接口，将商品以及SKU上得条形码信息补全
type TaobaoItemBarcodeUpdateRequest struct {
    
    /* isforce optional是否强制保存商品条码。
true：强制保存
false ：需要执行条码库校验 */
    isforce bool `json:"isforce";xml:"isforce"`
    
    /* item_barcode optional商品条形码，如果不用更新，可选择不填 */
    item_barcode string `json:"item_barcode";xml:"item_barcode"`
    
    /* item_id required被更新商品的ID */
    item_id int64 `json:"item_id";xml:"item_id"`
    
    /* sku_barcodes optionalSKU维度的条形码，和sku_ids字段一一对应，中间以英文逗号分隔 */
    sku_barcodes string `json:"sku_barcodes";xml:"sku_barcodes"`
    
    /* sku_ids optional被更新SKU的ID列表，中间以英文逗号进行分隔。如果没有SKU或者不需要更新SKU的条形码，不需要设置 */
    sku_ids string `json:"sku_ids";xml:"sku_ids"`
    
    /* src optional访问来源，这字段提供给千牛扫码枪用，
其他调用方，不需要填写 */
    src string `json:"src";xml:"src"`
    
}

func (req *TaobaoItemBarcodeUpdateRequest) GetAPIName() string {
	return "taobao.item.barcode.update"
}

// TaobaoItemBarcodeUpdateResponse 通过该接口，将商品以及SKU上得条形码信息补全
type TaobaoItemBarcodeUpdateResponse struct {
    
    /* item Object商品结构里的num_iid，modified */
    item Item `json:"item";xml:"item"`
    
}
