package tbsdk

// TaobaoScitemOutercodeGetRequest 根据outerCode查询商品
type TaobaoScitemOutercodeGetRequest struct {
    
    /* outer_code required商品编码 */
    outer_code string `json:"outer_code";xml:"outer_code"`
    
}

func (req *TaobaoScitemOutercodeGetRequest) GetAPIName() string {
	return "taobao.scitem.outercode.get"
}

// TaobaoScitemOutercodeGetResponse 根据outerCode查询商品
type TaobaoScitemOutercodeGetResponse struct {
    
    /* sc_item Object后台商品 */
    sc_item ScItem `json:"sc_item";xml:"sc_item"`
    
}
