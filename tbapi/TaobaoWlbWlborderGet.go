package tbsdk

// TaobaoWlbWlborderGetRequest 根据物流宝订单编号查询物流宝订单概要信息
type TaobaoWlbWlborderGetRequest struct {
    
    /* wlb_order_code required物流宝订单编码 */
    wlb_order_code string `json:"wlb_order_code";xml:"wlb_order_code"`
    
}

func (req *TaobaoWlbWlborderGetRequest) GetAPIName() string {
	return "taobao.wlb.wlborder.get"
}

// TaobaoWlbWlborderGetResponse 根据物流宝订单编号查询物流宝订单概要信息
type TaobaoWlbWlborderGetResponse struct {
    
    /* wlb_order Object物流宝订单对象 */
    wlb_order WlbOrder `json:"wlb_order";xml:"wlb_order"`
    
}
