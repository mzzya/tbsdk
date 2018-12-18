package tbsdk

// TaobaoFenxiaoProductToChannelImportRequest 支持供应商将已有产品导入到某个渠道销售
type TaobaoFenxiaoProductToChannelImportRequest struct {
    
    /* channel required要导入的渠道[21 零售PLUS]目前仅支持此渠道 */
    channel int64 `json:"channel";xml:"channel"`
    
    /* product_id required要导入的产品id */
    product_id int64 `json:"product_id";xml:"product_id"`
    
}

func (req *TaobaoFenxiaoProductToChannelImportRequest) GetAPIName() string {
	return "taobao.fenxiao.product.to.channel.import"
}

// TaobaoFenxiaoProductToChannelImportResponse 支持供应商将已有产品导入到某个渠道销售
type TaobaoFenxiaoProductToChannelImportResponse struct {
    
}
