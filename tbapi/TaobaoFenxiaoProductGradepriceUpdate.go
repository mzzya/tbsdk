package tbsdk

// TaobaoFenxiaoProductGradepriceUpdateRequest 供应商可以针对产品不同的sku，指定对应交易类型（代销or经销）方式下，设定折扣方式（按等级or指定分销商）以及对应优惠后的采购价格
type TaobaoFenxiaoProductGradepriceUpdateRequest struct {
    
    /* ids required会员等级的id或者分销商id，例如：”1001,2001,1002” */
    ids int64 `json:"ids";xml:"ids"`
    
    /* prices required优惠价格,大小为0到100000000之间的整数或两位小数，例：优惠价格为：100元2角5分，传入的参数应写成：100.25 */
    prices string `json:"prices";xml:"prices"`
    
    /* product_id required产品Id */
    product_id int64 `json:"product_id";xml:"product_id"`
    
    /* sku_id optionalskuId，如果产品有skuId,必须要输入skuId;没有skuId的时候不必选 */
    sku_id int64 `json:"sku_id";xml:"sku_id"`
    
    /* target_type required选择折扣方式：GRADE（按等级进行设置）;DISCITUTOR(按分销商进行设置）。例如"GRADE,DISTRIBUTOR" */
    target_type string `json:"target_type";xml:"target_type"`
    
    /* trade_type optional交易类型： AGENT(代销)、DEALER(经销)，ALL(代销和经销) */
    trade_type string `json:"trade_type";xml:"trade_type"`
    
}

func (req *TaobaoFenxiaoProductGradepriceUpdateRequest) GetAPIName() string {
	return "taobao.fenxiao.product.gradeprice.update"
}

// TaobaoFenxiaoProductGradepriceUpdateResponse 供应商可以针对产品不同的sku，指定对应交易类型（代销or经销）方式下，设定折扣方式（按等级or指定分销商）以及对应优惠后的采购价格
type TaobaoFenxiaoProductGradepriceUpdateResponse struct {
    
    /* is_success Basic返回操作结果：成功或失败 */
    is_success bool `json:"is_success";xml:"is_success"`
    
}
