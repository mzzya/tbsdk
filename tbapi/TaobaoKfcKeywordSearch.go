package tbsdk

// TaobaoKfcKeywordSearchRequest 对输入的文本信息进行禁忌关键词匹配，返回匹配的结果
type TaobaoKfcKeywordSearchRequest struct {
    
    /* apply optional应用点，分为一级应用点、二级应用点。其中一级应用点通常是指某一个系统或产品，比如淘宝的商品应用（taobao_auction）；二级应用点，是指一级应用点下的具体的分类，比如商品标题(title)、商品描述(content)。不同的二级应用可以设置不同关键词。

这里的apply参数是由一级应用点与二级应用点合起来的字符（一级应用点+"."+二级应用点），如taobao_auction.title。


通常apply参数是不需要传递的。如有特殊需求（比如特殊的过滤需求，需要自己维护一套自己词库），需传递此参数。 */
    apply string `json:"apply";xml:"apply"`
    
    /* content required需要过滤的文本信息 */
    content string `json:"content";xml:"content"`
    
    /* nick optional发布信息的淘宝会员名，可以不传 */
    nick string `json:"nick";xml:"nick"`
    
}

func (req *TaobaoKfcKeywordSearchRequest) GetAPIName() string {
	return "taobao.kfc.keyword.search"
}

// TaobaoKfcKeywordSearchResponse 对输入的文本信息进行禁忌关键词匹配，返回匹配的结果
type TaobaoKfcKeywordSearchResponse struct {
    
    /* kfc_search_result ObjectKFC 关键词过滤匹配结果 */
    kfc_search_result KfcSearchResult `json:"kfc_search_result";xml:"kfc_search_result"`
    
}
