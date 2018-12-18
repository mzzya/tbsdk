package tbsdk

// TaobaoWlbWmsReturnBillGetRequest 通过订单号获取单个销退单收货信息。
type TaobaoWlbWmsReturnBillGetRequest struct {
    
    /* cn_order_code special菜鸟订单编码，查询单个订单时orderCode与cnOrderCode必须有一个参数值不为空，两个参数都赋值时，以cnOrderCode值检索数据 */
    cn_order_code string `json:"cn_order_code";xml:"cn_order_code"`
    
    /* order_code specialERP订单编码，查询单个订单时orderCode与cnOrderCode必须有一个参数值不为空，两个参数都赋值时，以cnOrderCode值检索数据 */
    order_code string `json:"order_code";xml:"order_code"`
    
}

func (req *TaobaoWlbWmsReturnBillGetRequest) GetAPIName() string {
	return "taobao.wlb.wms.return.bill.get"
}

// TaobaoWlbWmsReturnBillGetResponse 通过订单号获取单个销退单收货信息。
type TaobaoWlbWmsReturnBillGetResponse struct {
    
    /* return_order_info Object回退订单信息 */
    return_order_info CainiaoReturnBillReturnorderinfo `json:"return_order_info";xml:"return_order_info"`
    
}
