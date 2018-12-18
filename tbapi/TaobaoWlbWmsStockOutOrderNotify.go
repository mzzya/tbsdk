package tbsdk

// TaobaoWlbWmsStockOutOrderNotifyRequest 出库单通知
type TaobaoWlbWmsStockOutOrderNotifyRequest struct {
    
    /* car_no optional车牌号 */
    car_no string `json:"car_no";xml:"car_no"`
    
    /* carriers_name optional承运商名称 */
    carriers_name string `json:"carriers_name";xml:"carriers_name"`
    
    /* extend_fields optional拓展属性 */
    extend_fields string `json:"extend_fields";xml:"extend_fields"`
    
    /* order_code requiredERP单据ID */
    order_code string `json:"order_code";xml:"order_code"`
    
    /* order_create_time required订单创建时间 */
    order_create_time Date `json:"order_create_time";xml:"order_create_time"`
    
    /* order_item_list optional订单商品信息列表 */
    order_item_list Orderitemlistwlbwmsstockoutordernotify `json:"order_item_list";xml:"order_item_list"`
    
    /* order_type required单据类型 301 调拨出库单、901普通出库单、903 其他出库单 305 B2B出库 */
    order_type int64 `json:"order_type";xml:"order_type"`
    
    /* outbound_type_desc optionalERP可选择性文本透传至WMS */
    outbound_type_desc string `json:"outbound_type_desc";xml:"outbound_type_desc"`
    
    /* pick_call optional取件人电话 */
    pick_call string `json:"pick_call";xml:"pick_call"`
    
    /* pick_id optional取件人身份证ID */
    pick_id string `json:"pick_id";xml:"pick_id"`
    
    /* pick_name optional取件人姓名 */
    pick_name string `json:"pick_name";xml:"pick_name"`
    
    /* prev_order_code optional前物流订单号，如退货入库单可能会用到 */
    prev_order_code string `json:"prev_order_code";xml:"prev_order_code"`
    
    /* receiver_info optional收件人信息 */
    receiver_info Receiverwlbwmsstockoutordernotify `json:"receiver_info";xml:"receiver_info"`
    
    /* remark optional备注 */
    remark string `json:"remark";xml:"remark"`
    
    /* send_time optional要求出库日期 */
    send_time Date `json:"send_time";xml:"send_time"`
    
    /* sender_info optional发货方信息 */
    sender_info Senderwlbwmsstockoutordernotify `json:"sender_info";xml:"sender_info"`
    
    /* store_code required仓储编码 */
    store_code string `json:"store_code";xml:"store_code"`
    
    /* transport_mode optional出库方式 */
    transport_mode string `json:"transport_mode";xml:"transport_mode"`
    
}

func (req *TaobaoWlbWmsStockOutOrderNotifyRequest) GetAPIName() string {
	return "taobao.wlb.wms.stock.out.order.notify"
}

// TaobaoWlbWmsStockOutOrderNotifyResponse 出库单通知
type TaobaoWlbWmsStockOutOrderNotifyResponse struct {
    
    /* order_code Basic仓储订单编码 */
    order_code string `json:"order_code";xml:"order_code"`
    
    /* wl_error_code Basic错误编码 */
    wl_error_code string `json:"wl_error_code";xml:"wl_error_code"`
    
    /* wl_error_msg Basic错误详细 */
    wl_error_msg string `json:"wl_error_msg";xml:"wl_error_msg"`
    
    /* wl_success Basic是否成功 */
    wl_success bool `json:"wl_success";xml:"wl_success"`
    
}
