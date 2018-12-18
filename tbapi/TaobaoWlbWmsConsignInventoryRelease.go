package tbsdk

// TaobaoWlbWmsConsignInventoryReleaseRequest ERP释放预占用库存
type TaobaoWlbWmsConsignInventoryReleaseRequest struct {
    
    /* content optional库存释放 */
    content Content `json:"content";xml:"content"`
    
}

func (req *TaobaoWlbWmsConsignInventoryReleaseRequest) GetAPIName() string {
	return "taobao.wlb.wms.consign.inventory.release"
}

// TaobaoWlbWmsConsignInventoryReleaseResponse ERP释放预占用库存
type TaobaoWlbWmsConsignInventoryReleaseResponse struct {
    
    /* result Object接口返回model */
    result Result `json:"result";xml:"result"`
    
}
