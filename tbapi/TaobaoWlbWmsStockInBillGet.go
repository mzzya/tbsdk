package tbsdk

// TaobaoWlbWmsStockInBillGetRequest 获取入库单信息
type TaobaoWlbWmsStockInBillGetRequest struct {
    
    /* cn_order_code special菜鸟订单编码，查询单个订单时orderCode与cnOrderCode必须有一个参数值不为空，两个参数都赋值时，以cnOrderCode值检索数据 */
    cn_order_code string `json:"cn_order_code";xml:"cn_order_code"`
    
    /* order_code specialERP订单编码，查询单个订单时orderCode与cnOrderCode必须有一个参数值不为空，两个参数都赋值时，以cnOrderCode值检索数据 */
    order_code string `json:"order_code";xml:"order_code"`
    
}

func (req *TaobaoWlbWmsStockInBillGetRequest) GetAPIName() string {
	return "taobao.wlb.wms.stock.in.bill.get"
}

// TaobaoWlbWmsStockInBillGetResponse 获取入库单信息
type TaobaoWlbWmsStockInBillGetResponse struct {
    
    /* stock_in_info Object入库单信息 */
    stock_in_info CainiaoStockInBillStockininfo `json:"stock_in_info";xml:"stock_in_info"`
    
}
