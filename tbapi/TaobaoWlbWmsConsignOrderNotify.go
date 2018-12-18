package tbsdk

// TaobaoWlbWmsConsignOrderNotifyRequest 发货订单通知
type TaobaoWlbWmsConsignOrderNotifyRequest struct {
    
    /* alipay_no optional废弃，支付宝交易号 */
    alipay_no string `json:"alipay_no";xml:"alipay_no"`
    
    /* ar_amount optional订单应收金额，消费者还需要付的金额，单位分 */
    ar_amount int64 `json:"ar_amount";xml:"ar_amount"`
    
    /* car_no optional废弃，车牌号 */
    car_no string `json:"car_no";xml:"car_no"`
    
    /* carriers_name optional废弃，承运商名称 */
    carriers_name string `json:"carriers_name";xml:"carriers_name"`
    
    /* deliver_requirements optional配送要求 */
    deliver_requirements Deliverrequirementswlbwmsconsignordernotify `json:"deliver_requirements";xml:"deliver_requirements"`
    
    /* discount_amount optional订单优惠金额，整单优惠金额，单位分 */
    discount_amount int64 `json:"discount_amount";xml:"discount_amount"`
    
    /* extend_fields optional拓展属性 */
    extend_fields string `json:"extend_fields";xml:"extend_fields"`
    
    /* got_amount optional订单已付金额，消费者已经支付的金额，单位分 */
    got_amount int64 `json:"got_amount";xml:"got_amount"`
    
    /* invoice_info_list optional发票信息列表 */
    invoice_info_list Invoicelistwlbwmsconsignordernotify `json:"invoice_info_list";xml:"invoice_info_list"`
    
    /* order_amount optional订单总金额,=总商品金额-订单优惠金额+快递费用，单位分 */
    order_amount int64 `json:"order_amount";xml:"order_amount"`
    
    /* order_code requiredERP订单号 */
    order_code string `json:"order_code";xml:"order_code"`
    
    /* order_create_time optional订单创建时间,ERP创建订单时间 */
    order_create_time Date `json:"order_create_time";xml:"order_create_time"`
    
    /* order_examination_time optional订单审核时间,ERP创建支付时间 */
    order_examination_time Date `json:"order_examination_time";xml:"order_examination_time"`
    
    /* order_flag optional订单标识 (1: cod –货到付款，4:invoiceinfo-需要发票) */
    order_flag string `json:"order_flag";xml:"order_flag"`
    
    /* order_item_list optional订单商品信息列表 */
    order_item_list Orderitemlistwlbwmsconsignordernotify `json:"order_item_list";xml:"order_item_list"`
    
    /* order_pay_time optional订单支付时间 */
    order_pay_time Date `json:"order_pay_time";xml:"order_pay_time"`
    
    /* order_priority optional废弃，订单优先级0 -10，根据订单作业优先级设置，数字越大，优先级越高 */
    order_priority int64 `json:"order_priority";xml:"order_priority"`
    
    /* order_shop_create_time optional下单时间，订单在交易平台创建时间 */
    order_shop_create_time Date `json:"order_shop_create_time";xml:"order_shop_create_time"`
    
    /* order_source optional订单来源（213 天猫，201 淘宝，214 京东，202 1688 阿里中文站 ，203 苏宁在线，204 亚马逊中国，205 当当，208 1号店，207 唯品会，209 国美在线，210 拍拍，206 易贝ebay，211 聚美优品，212 乐蜂网，215 邮乐，216 凡客，217 优购，218 银泰，219 易讯，221 聚尚网，222 蘑菇街，223 POS门店，301 其他） */
    order_source int64 `json:"order_source";xml:"order_source"`
    
    /* order_sub_source optional废弃，交易内部来源，文本透传 WAP(手机);HITAO(嗨淘);TOP(TOP平台);TAOBAO(普通淘宝);JHS(聚划算) 一笔订单可能同时有以上多个标记，则以逗号分隔 */
    order_sub_source string `json:"order_sub_source";xml:"order_sub_source"`
    
    /* order_type required单据类型 201 一般交易出库单 202 B2B交易出库单 502 换货出库单 503 补发出库单 */
    order_type int64 `json:"order_type";xml:"order_type"`
    
    /* picker_call optional废弃，取件人电话 */
    picker_call string `json:"picker_call";xml:"picker_call"`
    
    /* picker_id optional废弃，取件人身份证 */
    picker_id string `json:"picker_id";xml:"picker_id"`
    
    /* picker_name optional废弃，取件人姓名 */
    picker_name string `json:"picker_name";xml:"picker_name"`
    
    /* postfee optional快递费用，单位分 */
    postfee int64 `json:"postfee";xml:"postfee"`
    
    /* prev_order_code optional前物流订单号，订单类型为502 换货出库单 503 补发出库单时，需求传入此内容 */
    prev_order_code string `json:"prev_order_code";xml:"prev_order_code"`
    
    /* receiver_info optional收件人信息 */
    receiver_info Receiverwlbwmsconsignordernotify `json:"receiver_info";xml:"receiver_info"`
    
    /* remark optional备注 */
    remark string `json:"remark";xml:"remark"`
    
    /* sender_info optional发货方信息 */
    sender_info Senderwlbwmsconsignordernotify `json:"sender_info";xml:"sender_info"`
    
    /* service_fee optionalCOD服务费，单位分 */
    service_fee int64 `json:"service_fee";xml:"service_fee"`
    
    /* store_code optional仓库编码，此字段为空时，由菜鸟路由仓库发货 */
    store_code string `json:"store_code";xml:"store_code"`
    
    /* tms_service_code optional快递公司编码，此字段为空时，由菜鸟选择快递配送 */
    tms_service_code string `json:"tms_service_code";xml:"tms_service_code"`
    
    /* tms_service_name optional快递公司名称 */
    tms_service_name string `json:"tms_service_name";xml:"tms_service_name"`
    
    /* transport_mode optional废弃，出库方式（自提，快递，销毁） */
    transport_mode string `json:"transport_mode";xml:"transport_mode"`
    
}

func (req *TaobaoWlbWmsConsignOrderNotifyRequest) GetAPIName() string {
	return "taobao.wlb.wms.consign.order.notify"
}

// TaobaoWlbWmsConsignOrderNotifyResponse 发货订单通知
type TaobaoWlbWmsConsignOrderNotifyResponse struct {
    
    /* consign_order_list Object Array系统自动生成 */
    consign_order_list Consignorderlist `json:"consign_order_list";xml:"consign_order_list"`
    
    /* order_code Basic订单编码 */
    order_code string `json:"order_code";xml:"order_code"`
    
    /* wl_error_code Basic错误编码 */
    wl_error_code string `json:"wl_error_code";xml:"wl_error_code"`
    
    /* wl_error_msg Basic错误详细 */
    wl_error_msg string `json:"wl_error_msg";xml:"wl_error_msg"`
    
    /* wl_success Basic是否成功 */
    wl_success bool `json:"wl_success";xml:"wl_success"`
    
}
