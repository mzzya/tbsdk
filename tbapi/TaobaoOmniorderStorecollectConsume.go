package tbsdk

// TaobaoOmniorderStorecollectConsumeRequest 全渠道门店自提核销订单
type TaobaoOmniorderStorecollectConsumeRequest struct {
    
    /* code required核销码 */
    code string `json:"code";xml:"code"`
    
    /* main_order_id required淘宝主订单ID */
    main_order_id int64 `json:"main_order_id";xml:"main_order_id"`
    
    /* operator optional核销操作人信息 */
    operator string `json:"operator";xml:"operator"`
    
}

func (req *TaobaoOmniorderStorecollectConsumeRequest) GetAPIName() string {
	return "taobao.omniorder.storecollect.consume"
}

// TaobaoOmniorderStorecollectConsumeResponse 全渠道门店自提核销订单
type TaobaoOmniorderStorecollectConsumeResponse struct {
    
    /* err_code Basic0表示成功 */
    err_code string `json:"err_code";xml:"err_code"`
    
    /* err_msg Basic核销错误信息 */
    err_msg string `json:"err_msg";xml:"err_msg"`
    
}
