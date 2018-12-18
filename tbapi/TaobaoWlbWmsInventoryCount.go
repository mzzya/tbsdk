package tbsdk

// TaobaoWlbWmsInventoryCountRequest 损益单回传
type TaobaoWlbWmsInventoryCountRequest struct {
    
    /* content optional损益单回传信息 */
    content WlbWmsInventoryCount `json:"content";xml:"content"`
    
}

func (req *TaobaoWlbWmsInventoryCountRequest) GetAPIName() string {
	return "taobao.wlb.wms.inventory.count"
}

// TaobaoWlbWmsInventoryCountResponse 损益单回传
type TaobaoWlbWmsInventoryCountResponse struct {
    
    /* result Objectresult */
    result WlbWmsInventoryCountResp `json:"result";xml:"result"`
    
}
