package tbsdk

// TaobaoQimenSingleitemQueryRequest 商品查询接口
type TaobaoQimenSingleitemQueryRequest struct {
    
    /* itemCode optional商品编码,S1234,string(50),必填, */
    itemCode string `json:"itemCode";xml:"itemCode"`
    
    /* itemId optional仓储系统商品编码,C123,string(50),必填, */
    itemId string `json:"itemId";xml:"itemId"`
    
    /* ownerCode optional货主编码,H123,string(50),必填, */
    ownerCode string `json:"ownerCode";xml:"ownerCode"`
    
}

func (req *TaobaoQimenSingleitemQueryRequest) GetAPIName() string {
	return "taobao.qimen.singleitem.query"
}

// TaobaoQimenSingleitemQueryResponse 商品查询接口
type TaobaoQimenSingleitemQueryResponse struct {
    
    /* code Basic响应码,0,string(50),, */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure,success,string(10),必填, */
    flag string `json:"flag";xml:"flag"`
    
    /* item Objectitem */
    item Item `json:"item";xml:"item"`
    
    /* message Basic响应信息,invalid appkey,string(100),, */
    message string `json:"message";xml:"message"`
    
}
