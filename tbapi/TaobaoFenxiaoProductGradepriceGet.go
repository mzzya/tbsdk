package tbsdk

// TaobaoFenxiaoProductGradepriceGetRequest 等级折扣查询
type TaobaoFenxiaoProductGradepriceGetRequest struct {
    
    /* product_id required产品id */
    product_id int64 `json:"product_id";xml:"product_id"`
    
    /* sku_id optionalskuId */
    sku_id int64 `json:"sku_id";xml:"sku_id"`
    
    /* trade_type optional经、代销模式（1：代销、2：经销） */
    trade_type int64 `json:"trade_type";xml:"trade_type"`
    
}

func (req *TaobaoFenxiaoProductGradepriceGetRequest) GetAPIName() string {
	return "taobao.fenxiao.product.gradeprice.get"
}

// TaobaoFenxiaoProductGradepriceGetResponse 等级折扣查询
type TaobaoFenxiaoProductGradepriceGetResponse struct {
    
    /* grade_discounts Object Array等级折扣列表 */
    grade_discounts GradeDiscount `json:"grade_discounts";xml:"grade_discounts"`
    
    /* is_success Basic操作是否成功 */
    is_success bool `json:"is_success";xml:"is_success"`
    
}
