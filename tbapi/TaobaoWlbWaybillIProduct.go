package tbsdk

// TaobaoWlbWaybillIProductRequest 商家可以查询物流商的产品类型和服务能力。
type TaobaoWlbWaybillIProductRequest struct {
    
    /* waybill_product_type_request required查询物流商电子面单产品类型入参 */
    waybill_product_type_request WaybillProductTypeRequest `json:"waybill_product_type_request";xml:"waybill_product_type_request"`
    
}

func (req *TaobaoWlbWaybillIProductRequest) GetAPIName() string {
	return "taobao.wlb.waybill.i.product"
}

// TaobaoWlbWaybillIProductResponse 商家可以查询物流商的产品类型和服务能力。
type TaobaoWlbWaybillIProductResponse struct {
    
    /* product_types Object Array产品类型返回 */
    product_types WaybillProductType `json:"product_types";xml:"product_types"`
    
}
