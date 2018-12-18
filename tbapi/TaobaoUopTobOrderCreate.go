package tbsdk

// TaobaoUopTobOrderCreateRequest ToB仓储发货
type TaobaoUopTobOrderCreateRequest struct {
    
    /* delivery_order optionalERP出库对象 */
    delivery_order DeliveryOrder `json:"delivery_order";xml:"delivery_order"`
    
}

func (req *TaobaoUopTobOrderCreateRequest) GetAPIName() string {
	return "taobao.uop.tob.order.create"
}

// TaobaoUopTobOrderCreateResponse ToB仓储发货
type TaobaoUopTobOrderCreateResponse struct {
    
    /* delivery_orders Object Array订单 */
    delivery_orders Deliveryorder `json:"delivery_orders";xml:"delivery_orders"`
    
    /* flag Basicflag */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basicmessage */
    message string `json:"message";xml:"message"`
    
}
