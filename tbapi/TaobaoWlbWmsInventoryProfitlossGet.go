package tbsdk

// TaobaoWlbWmsInventoryProfitlossGetRequest 通过订单列表批量获取库存损益单信息
type TaobaoWlbWmsInventoryProfitlossGetRequest struct {
    
    /* cn_order_code required菜鸟订单编码 */
    cn_order_code string `json:"cn_order_code";xml:"cn_order_code"`
    
}

func (req *TaobaoWlbWmsInventoryProfitlossGetRequest) GetAPIName() string {
	return "taobao.wlb.wms.inventory.profitloss.get"
}

// TaobaoWlbWmsInventoryProfitlossGetResponse 通过订单列表批量获取库存损益单信息
type TaobaoWlbWmsInventoryProfitlossGetResponse struct {
    
    /* profit_loss_info Object损益信息 */
    profit_loss_info CainiaoInventoryProfitlossProfitlossinfo `json:"profit_loss_info";xml:"profit_loss_info"`
    
}
