package tbsdk

// TaobaoOmniorderStoreSdtconsignRequest ISV取完单号后通知菜鸟裹裹发货
type TaobaoOmniorderStoreSdtconsignRequest struct {
    
    /* package_id required取号接口返回的包裹id */
    package_id string `json:"package_id";xml:"package_id"`
    
    /* tag_code optional发货标签号 */
    tag_code string `json:"tag_code";xml:"tag_code"`
    
}

func (req *TaobaoOmniorderStoreSdtconsignRequest) GetAPIName() string {
	return "taobao.omniorder.store.sdtconsign"
}

// TaobaoOmniorderStoreSdtconsignResponse ISV取完单号后通知菜鸟裹裹发货
type TaobaoOmniorderStoreSdtconsignResponse struct {
    
    /* data Objectdata */
    data SdtConsignResponse `json:"data";xml:"data"`
    
    /* err_code Basic异常码 0 为正常，否则异常 */
    err_code string `json:"err_code";xml:"err_code"`
    
    /* message Basic异常信息 */
    message string `json:"message";xml:"message"`
    
}
