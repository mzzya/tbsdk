package tbsdk

// TaobaoRdcAligeniusWarehouseReverseUpdateRequest 销退单状态回传
type TaobaoRdcAligeniusWarehouseReverseUpdateRequest struct {
    
    /* param0 required参数 */
    param0 UpdateReverseStatusDto `json:"param0";xml:"param0"`
    
}

func (req *TaobaoRdcAligeniusWarehouseReverseUpdateRequest) GetAPIName() string {
	return "taobao.rdc.aligenius.warehouse.reverse.update"
}

// TaobaoRdcAligeniusWarehouseReverseUpdateResponse 销退单状态回传
type TaobaoRdcAligeniusWarehouseReverseUpdateResponse struct {
    
    /* result Objectresult */
    result Result `json:"result";xml:"result"`
    
}
