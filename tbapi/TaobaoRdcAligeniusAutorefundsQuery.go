package tbsdk

// TaobaoRdcAligeniusAutorefundsQueryRequest （此接口后期将不再维护，请勿使用）供第三方商家查询授权给自己的所有退款款订单的退款信息
type TaobaoRdcAligeniusAutorefundsQueryRequest struct {
    
    /* end_time required查询数据时间段结束时间 */
    end_time Date `json:"end_time";xml:"end_time"`
    
    /* page_no required页数 */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size required每页数据数 */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* start_time required查询数据时间段开始时间 */
    start_time Date `json:"start_time";xml:"start_time"`
    
}

func (req *TaobaoRdcAligeniusAutorefundsQueryRequest) GetAPIName() string {
	return "taobao.rdc.aligenius.autorefunds.query"
}

// TaobaoRdcAligeniusAutorefundsQueryResponse （此接口后期将不再维护，请勿使用）供第三方商家查询授权给自己的所有退款款订单的退款信息
type TaobaoRdcAligeniusAutorefundsQueryResponse struct {
    
    /* result Objectresult */
    result Result `json:"result";xml:"result"`
    
}
