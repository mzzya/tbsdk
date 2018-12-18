package tbsdk

// TaobaoTraderateImprImprwordsGetRequest 根据淘宝后台类目的一级类目和叶子类目
type TaobaoTraderateImprImprwordsGetRequest struct {
    
    /* cat_leaf_id optional淘宝叶子类目id */
    cat_leaf_id int64 `json:"cat_leaf_id";xml:"cat_leaf_id"`
    
    /* cat_root_id required淘宝一级类目id */
    cat_root_id int64 `json:"cat_root_id";xml:"cat_root_id"`
    
}

func (req *TaobaoTraderateImprImprwordsGetRequest) GetAPIName() string {
	return "taobao.traderate.impr.imprwords.get"
}

// TaobaoTraderateImprImprwordsGetResponse 根据淘宝后台类目的一级类目和叶子类目
type TaobaoTraderateImprImprwordsGetResponse struct {
    
    /* impr_words Basic Array返回类目下所有大家印象的标签 */
    impr_words string `json:"impr_words";xml:"impr_words"`
    
}
