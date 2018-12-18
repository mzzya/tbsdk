package tbsdk

// TaobaoWlbWmsReturnOrderNotifyRequest 销售退货通知
type TaobaoWlbWmsReturnOrderNotifyRequest struct {
    
    /* buyer_nick optional买家昵称 */
    buyer_nick string `json:"buyer_nick";xml:"buyer_nick"`
    
    /* extend_fields optional扩展属性, key-value结构，格式要求： 以英文分号“;”分隔每组key-value，以英文冒号“:”分隔key与value。如果value中带有分号，需要转成下划线“_”，如果带有冒号，需要转成中划线“-” */
    extend_fields string `json:"extend_fields";xml:"extend_fields"`
    
    /* order_code optionalERP单据编码 */
    order_code string `json:"order_code";xml:"order_code"`
    
    /* order_create_time optionalERP订单创建时间 */
    order_create_time Date `json:"order_create_time";xml:"order_create_time"`
    
    /* order_flag optional用字符串格式来表示订单标记列表：比如VISIT^ SELLER_AFFORD^SYNC_RETURN_BILL 等，中间用“^”来隔开 ----------------------------------------  订单标记list（所有字母全部大写）： 9:VISIT-上门12: SELLER_AFFORD 是否卖家承担运费 默认是，即没 13: SYNC_RETURN_BILL，同时退回发票 */
    order_flag string `json:"order_flag";xml:"order_flag"`
    
    /* order_item_list optional商品信息列表 */
    order_item_list Orderitemlistwlbwmsreturnordernotify `json:"order_item_list";xml:"order_item_list"`
    
    /* order_source optional订单来源 201 淘宝，202 1688，203 苏宁，204 亚马逊中国，205 当当，206 ebay，207 唯品会，208 1号店，209 国美，210 拍拍，211 聚美优品，212 乐峰，214 京东，301 其他 */
    order_source string `json:"order_source";xml:"order_source"`
    
    /* order_type optional订单类型 501 销退入库 */
    order_type string `json:"order_type";xml:"order_type"`
    
    /* owner_user_id optional货主ID */
    owner_user_id string `json:"owner_user_id";xml:"owner_user_id"`
    
    /* prev_order_code optional来源单据号，销售退货时填充原发货的LBX号 */
    prev_order_code string `json:"prev_order_code";xml:"prev_order_code"`
    
    /* receiver_info optional收件人信息 */
    receiver_info Receiverinfowlbwmsreturnordernotify `json:"receiver_info";xml:"receiver_info"`
    
    /* remark optional备注 */
    remark string `json:"remark";xml:"remark"`
    
    /* return_reason optional销退时请提供退货的原因 */
    return_reason string `json:"return_reason";xml:"return_reason"`
    
    /* sender_info optional发件人信息 */
    sender_info Senderinfowlbwmsreturnordernotify `json:"sender_info";xml:"sender_info"`
    
    /* store_code optional仓库编码 */
    store_code string `json:"store_code";xml:"store_code"`
    
    /* tms_order_code optional运单号 */
    tms_order_code string `json:"tms_order_code";xml:"tms_order_code"`
    
    /* tms_service_code optional快递公司编码 */
    tms_service_code string `json:"tms_service_code";xml:"tms_service_code"`
    
    /* tms_service_name optional快递公司名称 */
    tms_service_name string `json:"tms_service_name";xml:"tms_service_name"`
    
}

func (req *TaobaoWlbWmsReturnOrderNotifyRequest) GetAPIName() string {
	return "taobao.wlb.wms.return.order.notify"
}

// TaobaoWlbWmsReturnOrderNotifyResponse 销售退货通知
type TaobaoWlbWmsReturnOrderNotifyResponse struct {
    
    /* create_time Basic订单创建时间 */
    create_time Date `json:"create_time";xml:"create_time"`
    
    /* store_order_code BasicLBX929829111 */
    store_order_code string `json:"store_order_code";xml:"store_order_code"`
    
    /* wl_error_code Basic错误编码 */
    wl_error_code string `json:"wl_error_code";xml:"wl_error_code"`
    
    /* wl_error_msg Basic错误信息 */
    wl_error_msg string `json:"wl_error_msg";xml:"wl_error_msg"`
    
    /* wl_success Basic是否成功 */
    wl_success bool `json:"wl_success";xml:"wl_success"`
    
}
