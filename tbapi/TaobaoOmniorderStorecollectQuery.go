package tbsdk

// TaobaoOmniorderStorecollectQueryRequest 全渠道门店自提根据核销码查询订单
type TaobaoOmniorderStorecollectQueryRequest struct {
    
    /* code required核销码 */
    code string `json:"code";xml:"code"`
    
}

func (req *TaobaoOmniorderStorecollectQueryRequest) GetAPIName() string {
	return "taobao.omniorder.storecollect.query"
}

// TaobaoOmniorderStorecollectQueryResponse 全渠道门店自提根据核销码查询订单
type TaobaoOmniorderStorecollectQueryResponse struct {
    
    /* result Objectresult */
    result Result `json:"result";xml:"result"`
    
}
