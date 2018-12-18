package tbsdk

// CainiaoWaybillIiProductRequest 商家可以查询物流商的产品类型和服务能力。
type CainiaoWaybillIiProductRequest struct {
    
    /* cp_code required快递公司code */
    cp_code string `json:"cp_code";xml:"cp_code"`
    
}

func (req *CainiaoWaybillIiProductRequest) GetAPIName() string {
	return "cainiao.waybill.ii.product"
}

// CainiaoWaybillIiProductResponse 商家可以查询物流商的产品类型和服务能力。
type CainiaoWaybillIiProductResponse struct {
    
    /* product_types Object Array返回值 */
    product_types WaybillProductType `json:"product_types";xml:"product_types"`
    
}
