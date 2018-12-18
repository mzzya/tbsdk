package tbsdk

// TaobaoFenxiaoProductcatAddRequest 新增产品线
type TaobaoFenxiaoProductcatAddRequest struct {
    
    /* agent_cost_percent required代销默认采购价比例，注意：100.00%，则输入为10000 */
    agent_cost_percent int64 `json:"agent_cost_percent";xml:"agent_cost_percent"`
    
    /* dealer_cost_percent required经销默认采购价比例，注意：100.00%，则输入为10000 */
    dealer_cost_percent int64 `json:"dealer_cost_percent";xml:"dealer_cost_percent"`
    
    /* name required产品线名称 */
    name string `json:"name";xml:"name"`
    
    /* retail_high_percent required最高零售价比例，注意：100.00%，则输入为10000 */
    retail_high_percent int64 `json:"retail_high_percent";xml:"retail_high_percent"`
    
    /* retail_low_percent required最低零售价比例，注意：100.00%，则输入为10000 */
    retail_low_percent int64 `json:"retail_low_percent";xml:"retail_low_percent"`
    
}

func (req *TaobaoFenxiaoProductcatAddRequest) GetAPIName() string {
	return "taobao.fenxiao.productcat.add"
}

// TaobaoFenxiaoProductcatAddResponse 新增产品线
type TaobaoFenxiaoProductcatAddResponse struct {
    
    /* is_success Basic操作是否成功 */
    is_success bool `json:"is_success";xml:"is_success"`
    
    /* product_line_id Basic产品线ID */
    product_line_id int64 `json:"product_line_id";xml:"product_line_id"`
    
}
