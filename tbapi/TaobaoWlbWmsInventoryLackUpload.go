package tbsdk

// TaobaoWlbWmsInventoryLackUploadRequest 缺货通知
type TaobaoWlbWmsInventoryLackUploadRequest struct {
    
    /* content optional缺货通知信息 */
    content WlbWmsInventoryLackUpload `json:"content";xml:"content"`
    
}

func (req *TaobaoWlbWmsInventoryLackUploadRequest) GetAPIName() string {
	return "taobao.wlb.wms.inventory.lack.upload"
}

// TaobaoWlbWmsInventoryLackUploadResponse 缺货通知
type TaobaoWlbWmsInventoryLackUploadResponse struct {
    
    /* result Object缺货回告 */
    result WlbWmsInventoryLackUploadResp `json:"result";xml:"result"`
    
}
