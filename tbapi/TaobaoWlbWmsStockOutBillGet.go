package tbsdk

// TaobaoWlbWmsStockOutBillGetRequest 通过订单号获取单个出库单发货信息
type TaobaoWlbWmsStockOutBillGetRequest struct {
    
    /* cn_order_code special菜鸟订单编码，查询单个订单时orderCode与cnOrderCode必须有一个参数值不为空，两个参数都赋值时，以cnOrderCode值检索数据 */
    cn_order_code string `json:"cn_order_code";xml:"cn_order_code"`
    
    /* order_code specialERP订单编码，查询单个订单时orderCode与cnOrderCode必须有一个参数值不为空，两个参数都赋值时，以cnOrderCode值检索数据 */
    order_code string `json:"order_code";xml:"order_code"`
    
}

func (req *TaobaoWlbWmsStockOutBillGetRequest) GetAPIName() string {
	return "taobao.wlb.wms.stock.out.bill.get"
}

// TaobaoWlbWmsStockOutBillGetResponse 通过订单号获取单个出库单发货信息
type TaobaoWlbWmsStockOutBillGetResponse struct {
    
    /* stock_out_info Object出库信息 */
    stock_out_info CainiaoStockOutBillStockoutinfo `json:"stock_out_info";xml:"stock_out_info"`
    
}
