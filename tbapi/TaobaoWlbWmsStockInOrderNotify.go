package tbsdk

// TaobaoWlbWmsStockInOrderNotifyRequest 入库通知单
type TaobaoWlbWmsStockInOrderNotifyRequest struct {
    
    /* expect_end_time optional预期送达结束时间 */
    expect_end_time string `json:"expect_end_time";xml:"expect_end_time"`
    
    /* expect_start_time optional预期送达开始时间 */
    expect_start_time string `json:"expect_start_time";xml:"expect_start_time"`
    
    /* extend_fields optional扩展属性, key-value结构，格式要求： 以英文分号“;”分隔每组key-value，以英文冒号“:”分隔key与value。如果value中带有分号，需要转成下划线“_”，如果带有冒号，需要转成中划线“-” */
    extend_fields string `json:"extend_fields";xml:"extend_fields"`
    
    /* inbound_type_desc optional可选择性文本透传至WMS，比如加工归还、委外归还、借出归还、内部归还等 */
    inbound_type_desc string `json:"inbound_type_desc";xml:"inbound_type_desc"`
    
    /* order_code required入库单据编码 */
    order_code string `json:"order_code";xml:"order_code"`
    
    /* order_create_time required单据创建时间 */
    order_create_time Date `json:"order_create_time";xml:"order_create_time"`
    
    /* order_flag optional订单标记以逗号分隔：  9:上门退货入库 13: 退货时是否收取发票，默认不收取（即没13为多选项，如1,2,8,9） */
    order_flag string `json:"order_flag";xml:"order_flag"`
    
    /* order_item_list required系统自动生成 */
    order_item_list Orderitemlistwlbwmsstockinordernotifywl `json:"order_item_list";xml:"order_item_list"`
    
    /* order_type required单据类型 601普通入库单、501销退入库单、302 调拨入库单、904其他入库单、306 B2B入库 */
    order_type int64 `json:"order_type";xml:"order_type"`
    
    /* prev_order_code optional来源单据号，如销售退货时填充原销售订单号 */
    prev_order_code string `json:"prev_order_code";xml:"prev_order_code"`
    
    /* receiver_info optional系统自动生成 */
    receiver_info Receiverinfowlbwmsstockinordernotifywl `json:"receiver_info";xml:"receiver_info"`
    
    /* remark optional备注，销退入库订单需要留言备注填充到此字段 */
    remark string `json:"remark";xml:"remark"`
    
    /* return_reason optional销退时请提供退货的原因 */
    return_reason string `json:"return_reason";xml:"return_reason"`
    
    /* sender_info optional系统自动生成 */
    sender_info Senderinfowlbwmsstockinordernotifywl `json:"sender_info";xml:"sender_info"`
    
    /* store_code required仓库编码 */
    store_code string `json:"store_code";xml:"store_code"`
    
    /* supplier_code optional供应商编码，往来单位编码 */
    supplier_code string `json:"supplier_code";xml:"supplier_code"`
    
    /* supplier_name optional供应商名称 ，往来单位名称 */
    supplier_name string `json:"supplier_name";xml:"supplier_name"`
    
    /* tms_order_code optional运单号&托运单号 1)	入库单支持多运单号录入 2)	销退场景下如果是拒收（非妥投运单）由ERP填充原运单号 */
    tms_order_code string `json:"tms_order_code";xml:"tms_order_code"`
    
    /* tms_service_code optional配送公司编码，拒收（非妥投）的销退订单，由ERP填充原单配送公司编码 */
    tms_service_code string `json:"tms_service_code";xml:"tms_service_code"`
    
    /* tms_service_name optional快递公司名称 */
    tms_service_name string `json:"tms_service_name";xml:"tms_service_name"`
    
}

func (req *TaobaoWlbWmsStockInOrderNotifyRequest) GetAPIName() string {
	return "taobao.wlb.wms.stock.in.order.notify"
}

// TaobaoWlbWmsStockInOrderNotifyResponse 入库通知单
type TaobaoWlbWmsStockInOrderNotifyResponse struct {
    
    /* order_code Basic仓储订单编码 */
    order_code string `json:"order_code";xml:"order_code"`
    
    /* wl_error_code Basic错误编码 */
    wl_error_code string `json:"wl_error_code";xml:"wl_error_code"`
    
    /* wl_error_msg Basic错误详细 */
    wl_error_msg string `json:"wl_error_msg";xml:"wl_error_msg"`
    
    /* wl_success Basic是否成功 */
    wl_success bool `json:"wl_success";xml:"wl_success"`
    
}
