package tbsdk

// TaobaoFenxiaoDiscountsGetRequest 查询折扣信息
type TaobaoFenxiaoDiscountsGetRequest struct {
    
    /* discount_id optional折扣ID */
    discount_id int64 `json:"discount_id";xml:"discount_id"`
    
    /* ext_fields optional指定查询额外的信息，可选值：DETAIL（查询折扣详情），多个可选值用逗号分割。（只允许指定折扣ID情况下使用） */
    ext_fields string `json:"ext_fields";xml:"ext_fields"`
    
}

func (req *TaobaoFenxiaoDiscountsGetRequest) GetAPIName() string {
	return "taobao.fenxiao.discounts.get"
}

// TaobaoFenxiaoDiscountsGetResponse 查询折扣信息
type TaobaoFenxiaoDiscountsGetResponse struct {
    
    /* discounts Object Array折扣数据结构 */
    discounts Discount `json:"discounts";xml:"discounts"`
    
    /* total_results Basic记录数 */
    total_results int64 `json:"total_results";xml:"total_results"`
    
}
