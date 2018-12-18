package tbsdk

// CainiaoBmsOrderCreateRequest 通过接口，在菜鸟商家工作台创建订单，并通过商家工作台进行发货管理。
type CainiaoBmsOrderCreateRequest struct {
    
    /* buyer_message optional买家留言 */
    buyer_message string `json:"buyer_message";xml:"buyer_message"`
    
    /* buyer_nick optional买家名称" */
    buyer_nick string `json:"buyer_nick";xml:"buyer_nick"`
    
    /* cod_fee optionalCOD服务费 */
    cod_fee int64 `json:"cod_fee";xml:"cod_fee"`
    
    /* created required创建时间 */
    created Date `json:"created";xml:"created"`
    
    /* discount_fee optional优惠金额，单位分 */
    discount_fee int64 `json:"discount_fee";xml:"discount_fee"`
    
    /* invoice_amount optional发票金额，单位分 */
    invoice_amount int64 `json:"invoice_amount";xml:"invoice_amount"`
    
    /* invoice_title optional发票抬头 */
    invoice_title string `json:"invoice_title";xml:"invoice_title"`
    
    /* invoice_type optional发票类型(1 电子、2 纸质) */
    invoice_type int64 `json:"invoice_type";xml:"invoice_type"`
    
    /* is_cod optional是否COD */
    is_cod bool `json:"is_cod";xml:"is_cod"`
    
    /* is_invoice optional是否打印发票 */
    is_invoice bool `json:"is_invoice";xml:"is_invoice"`
    
    /* items optionaldemo */
    items Items `json:"items";xml:"items"`
    
    /* order_amount required订单总金额，单位分 */
    order_amount int64 `json:"order_amount";xml:"order_amount"`
    
    /* order_type optional订单类型，默认0，普通销售订单，3:B2B单 */
    order_type string `json:"order_type";xml:"order_type"`
    
    /* paied_amount required支付金额，单位分 */
    paied_amount int64 `json:"paied_amount";xml:"paied_amount"`
    
    /* pay_time required支付时间 */
    pay_time Date `json:"pay_time";xml:"pay_time"`
    
    /* post_fee optional快递费，单位分 */
    post_fee int64 `json:"post_fee";xml:"post_fee"`
    
    /* receiver_address required收货地址 */
    receiver_address string `json:"receiver_address";xml:"receiver_address"`
    
    /* receiver_city required收货市 */
    receiver_city string `json:"receiver_city";xml:"receiver_city"`
    
    /* receiver_country optional收货人国籍 */
    receiver_country string `json:"receiver_country";xml:"receiver_country"`
    
    /* receiver_district optional收货区 */
    receiver_district string `json:"receiver_district";xml:"receiver_district"`
    
    /* receiver_mobile optional收货人手机，手机与电话不可同时为空 */
    receiver_mobile string `json:"receiver_mobile";xml:"receiver_mobile"`
    
    /* receiver_name required收货人名称 */
    receiver_name string `json:"receiver_name";xml:"receiver_name"`
    
    /* receiver_phone optional收货人电话，手机与电话不可同时为空 */
    receiver_phone string `json:"receiver_phone";xml:"receiver_phone"`
    
    /* receiver_state required收货省 */
    receiver_state string `json:"receiver_state";xml:"receiver_state"`
    
    /* receiver_town optional街道 */
    receiver_town string `json:"receiver_town";xml:"receiver_town"`
    
    /* receiver_zip required收货邮编 */
    receiver_zip string `json:"receiver_zip";xml:"receiver_zip"`
    
    /* seller_memo optional卖家备注 */
    seller_memo string `json:"seller_memo";xml:"seller_memo"`
    
    /* shop_code required店铺信息 */
    shop_code string `json:"shop_code";xml:"shop_code"`
    
    /* trade_id required交易单号 */
    trade_id string `json:"trade_id";xml:"trade_id"`
    
    /* wait_pay_amount required应收金额，单位分 */
    wait_pay_amount int64 `json:"wait_pay_amount";xml:"wait_pay_amount"`
    
}

func (req *CainiaoBmsOrderCreateRequest) GetAPIName() string {
	return "cainiao.bms.order.create"
}

// CainiaoBmsOrderCreateResponse 通过接口，在菜鸟商家工作台创建订单，并通过商家工作台进行发货管理。
type CainiaoBmsOrderCreateResponse struct {
    
}
