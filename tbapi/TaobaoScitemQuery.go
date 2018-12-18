package tbsdk

// TaobaoScitemQueryRequest 查询后端商品
type TaobaoScitemQueryRequest struct {
    
    /* bar_code optional条形码 */
    bar_code string `json:"bar_code";xml:"bar_code"`
    
    /* item_name optional商品名称 */
    item_name string `json:"item_name";xml:"item_name"`
    
    /* item_type optionalITEM类型(只允许输入以下英文或空) NORMAL 0:普通商品; COMBINE 1:是否是组合商品 DISTRIBUTION */
    item_type int64 `json:"item_type";xml:"item_type"`
    
    /* outer_code optional商家给商品的一个编码 */
    outer_code string `json:"outer_code";xml:"outer_code"`
    
    /* page_index optional当前页码数 */
    page_index int64 `json:"page_index";xml:"page_index"`
    
    /* page_size optional分页记录个数，如果用户输入的记录数大于50，则一页显示50条记录 */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* wms_code optional仓库编码 */
    wms_code string `json:"wms_code";xml:"wms_code"`
    
}

func (req *TaobaoScitemQueryRequest) GetAPIName() string {
	return "taobao.scitem.query"
}

// TaobaoScitemQueryResponse 查询后端商品
type TaobaoScitemQueryResponse struct {
    
    /* query_pagination Object分页 */
    query_pagination QueryPagination `json:"query_pagination";xml:"query_pagination"`
    
    /* sc_item_list Object ArrayList<ScItemDO> */
    sc_item_list ScItem `json:"sc_item_list";xml:"sc_item_list"`
    
    /* total_page Basic商品条数 */
    total_page int64 `json:"total_page";xml:"total_page"`
    
}
