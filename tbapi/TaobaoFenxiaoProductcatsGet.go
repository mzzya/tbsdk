package tbsdk

// TaobaoFenxiaoProductcatsGetRequest 查询供应商的所有产品线数据。根据登陆用户来查询，不需要其他入参
type TaobaoFenxiaoProductcatsGetRequest struct {
    
    /* fields optional返回字段列表 */
    fields string `json:"fields";xml:"fields"`
    
}

func (req *TaobaoFenxiaoProductcatsGetRequest) GetAPIName() string {
	return "taobao.fenxiao.productcats.get"
}

// TaobaoFenxiaoProductcatsGetResponse 查询供应商的所有产品线数据。根据登陆用户来查询，不需要其他入参
type TaobaoFenxiaoProductcatsGetResponse struct {
    
    /* productcats Object Array产品线列表。返回 ProductCat 包含的字段信息。 */
    productcats ProductCat `json:"productcats";xml:"productcats"`
    
    /* total_results Basic查询结果记录数 */
    total_results int64 `json:"total_results";xml:"total_results"`
    
}
