package tbsdk

// CainiaoWaybillIiCancelRequest 面单号有误需要取消的时候，调用该接口取消获取的电子面单。
type CainiaoWaybillIiCancelRequest struct {
    
    /* cp_code required快递公司code */
    cp_code string `json:"cp_code";xml:"cp_code"`
    
    /* waybill_code required电子面单号 */
    waybill_code string `json:"waybill_code";xml:"waybill_code"`
    
}

func (req *CainiaoWaybillIiCancelRequest) GetAPIName() string {
	return "cainiao.waybill.ii.cancel"
}

// CainiaoWaybillIiCancelResponse 面单号有误需要取消的时候，调用该接口取消获取的电子面单。
type CainiaoWaybillIiCancelResponse struct {
    
    /* cancel_result Basic调用取消是否成功 */
    cancel_result bool `json:"cancel_result";xml:"cancel_result"`
    
}
