package tbsdk

// TaobaoWlbWmsConsignBillGetRequest 获取销售订单发货信息
type TaobaoWlbWmsConsignBillGetRequest struct {
    
    /* cn_order_code special菜鸟订单编码,cnOrderCode与orderCode必须有一个值不可为空 */
    cn_order_code string `json:"cn_order_code";xml:"cn_order_code"`
    
    /* order_code specialERP订单编码,cnOrderCode与orderCode必须有一个值不可为空 */
    order_code string `json:"order_code";xml:"order_code"`
    
}

func (req *TaobaoWlbWmsConsignBillGetRequest) GetAPIName() string {
	return "taobao.wlb.wms.consign.bill.get"
}

// TaobaoWlbWmsConsignBillGetResponse 获取销售订单发货信息
type TaobaoWlbWmsConsignBillGetResponse struct {
    
    /* consign_send_info_list Object Array商品信息列表 */
    consign_send_info_list Consignsendinfolist `json:"consign_send_info_list";xml:"consign_send_info_list"`
    
}
