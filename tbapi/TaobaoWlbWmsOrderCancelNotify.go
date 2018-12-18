package tbsdk

// TaobaoWlbWmsOrderCancelNotifyRequest 单据取消接口
type TaobaoWlbWmsOrderCancelNotifyRequest struct {
    
    /* order_code required订单类型 */
    order_code string `json:"order_code";xml:"order_code"`
    
    /* order_type required单据类型 601普通入库单、501销退入库单、302 调拨入库单、904其他入库单、301 调拨出库单、901普通出库单、903 其他出库单、201 一般交易出库单 202 B2B交易出库单 502 换货出库单 503 补发出库单、1102 仓内加工作业单、 701 盘亏单、702盘盈单、711 库存状态调整单 */
    order_type string `json:"order_type";xml:"order_type"`
    
    /* store_code required仓库编码 */
    store_code string `json:"store_code";xml:"store_code"`
    
}

func (req *TaobaoWlbWmsOrderCancelNotifyRequest) GetAPIName() string {
	return "taobao.wlb.wms.order.cancel.notify"
}

// TaobaoWlbWmsOrderCancelNotifyResponse 单据取消接口
type TaobaoWlbWmsOrderCancelNotifyResponse struct {
    
    /* wl_error_code Basic错误编码 */
    wl_error_code string `json:"wl_error_code";xml:"wl_error_code"`
    
    /* wl_error_msg Basic错误详细 */
    wl_error_msg string `json:"wl_error_msg";xml:"wl_error_msg"`
    
    /* wl_success Basic是否成功 */
    wl_success bool `json:"wl_success";xml:"wl_success"`
    
}
