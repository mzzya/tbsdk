package tbsdk

// TaobaoRpReturngoodsAgreeRequest 卖家同意退货，支持淘宝和天猫的订单。
type TaobaoRpReturngoodsAgreeRequest struct {
    
    /* address optional卖家提供的退货地址，淘宝退款为必填项。 */
    address string `json:"address";xml:"address"`
    
    /* mobile optional卖家手机，淘宝退款为必填项。 */
    mobile string `json:"mobile";xml:"mobile"`
    
    /* name optional卖家姓名，淘宝退款为必填项。 */
    name string `json:"name";xml:"name"`
    
    /* post optional卖家提供的退货地址的邮编，淘宝退款为必填项。 */
    post string `json:"post";xml:"post"`
    
    /* post_fee_bear_role optional邮费承担方，买家承担值为1，卖家承担值为0 */
    post_fee_bear_role int64 `json:"post_fee_bear_role";xml:"post_fee_bear_role"`
    
    /* refund_id required退款编号 */
    refund_id int64 `json:"refund_id";xml:"refund_id"`
    
    /* refund_phase optional售中：onsale，售后：aftersale，天猫退款为必填项。 */
    refund_phase string `json:"refund_phase";xml:"refund_phase"`
    
    /* refund_version optional退款版本号，天猫退款为必填项。 */
    refund_version int64 `json:"refund_version";xml:"refund_version"`
    
    /* remark optional卖家退货留言，天猫退款为必填项。 */
    remark string `json:"remark";xml:"remark"`
    
    /* seller_address_id optional卖家收货地址编号，天猫淘宝退款都为必填项。 */
    seller_address_id int64 `json:"seller_address_id";xml:"seller_address_id"`
    
    /* tel optional卖家座机，淘宝退款为必填项。 */
    tel string `json:"tel";xml:"tel"`
    
}

func (req *TaobaoRpReturngoodsAgreeRequest) GetAPIName() string {
	return "taobao.rp.returngoods.agree"
}

// TaobaoRpReturngoodsAgreeResponse 卖家同意退货，支持淘宝和天猫的订单。
type TaobaoRpReturngoodsAgreeResponse struct {
    
    /* is_success Basic操作成功 */
    is_success bool `json:"is_success";xml:"is_success"`
    
}
