package tbsdk

// TaobaoNextoneLogisticsWarehouseUpdateRequest 商家上传退货入仓状态给ag
type TaobaoNextoneLogisticsWarehouseUpdateRequest struct {
    
    /* refund_id required退款编号 */
    refund_id int64 `json:"refund_id";xml:"refund_id"`
    
    /* warehouse_status required退货入仓状态 1.已入仓 */
    warehouse_status int64 `json:"warehouse_status";xml:"warehouse_status"`
    
}

func (req *TaobaoNextoneLogisticsWarehouseUpdateRequest) GetAPIName() string {
	return "taobao.nextone.logistics.warehouse.update"
}

// TaobaoNextoneLogisticsWarehouseUpdateResponse 商家上传退货入仓状态给ag
type TaobaoNextoneLogisticsWarehouseUpdateResponse struct {
    
    /* err_code BasicerrorCode */
    err_code string `json:"err_code";xml:"err_code"`
    
    /* err_info BasicerrorInfo */
    err_info string `json:"err_info";xml:"err_info"`
    
    /* result_data BasicresultData */
    result_data map[string]interface{} `json:"result_data";xml:"result_data"`
    
    /* succeed Basicsuccess */
    succeed bool `json:"succeed";xml:"succeed"`
    
}
