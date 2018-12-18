package tbsdk

// TaobaoRdcAligeniusAccountValidateRequest 提供应对接AG的erp系统查询其旗下的商家是否为AG商家
type TaobaoRdcAligeniusAccountValidateRequest struct {
    
}

func (req *TaobaoRdcAligeniusAccountValidateRequest) GetAPIName() string {
	return "taobao.rdc.aligenius.account.validate"
}

// TaobaoRdcAligeniusAccountValidateResponse 提供应对接AG的erp系统查询其旗下的商家是否为AG商家
type TaobaoRdcAligeniusAccountValidateResponse struct {
    
    /* result Objectresult */
    result Result `json:"result";xml:"result"`
    
}
