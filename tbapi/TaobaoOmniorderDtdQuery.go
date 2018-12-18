package tbsdk

// TaobaoOmniorderDtdQueryRequest 门店自送根据核销码码查询订单信息
type TaobaoOmniorderDtdQueryRequest struct {
    
    /* code required核销码 */
    code string `json:"code";xml:"code"`
    
}

func (req *TaobaoOmniorderDtdQueryRequest) GetAPIName() string {
	return "taobao.omniorder.dtd.query"
}

// TaobaoOmniorderDtdQueryResponse 门店自送根据核销码码查询订单信息
type TaobaoOmniorderDtdQueryResponse struct {
    
    /* data Objectdata */
    data Door2doorQueryResult `json:"data";xml:"data"`
    
    /* message Basic错误信息 */
    message string `json:"message";xml:"message"`
    
    /* result_code Basic错误码，为0表示成功，非0表示失败 */
    result_code string `json:"result_code";xml:"result_code"`
    
}
