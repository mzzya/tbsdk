package tbsdk

// TaobaoOmniorderGuideDataGetRequest 获取全渠道导购产品，目前包括随心购、随身购扫码、加购和交易数据。
type TaobaoOmniorderGuideDataGetRequest struct {
    
    /* page_no required页码，从1开始 */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size required每页数量，不能大于1000 */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* start_time required拉取数据开始时间 */
    start_time Date `json:"start_time";xml:"start_time"`
    
    /* _type requireddetail_smg_scan: 扫码购扫码明细;detail_smg_cart: 扫码购加购明细;detail_smg_order: 扫码购订单明细;detail_sxg_search: 随心购搜索明细;detail_sxg_view_item: 随心购商品浏览明细;detail_sxg_cart: 随心购脚骨明细;detail_sxg_order: 随心购订单明细 */
    _type string `json:"type";xml:"type"`
    
}

func (req *TaobaoOmniorderGuideDataGetRequest) GetAPIName() string {
	return "taobao.omniorder.guide.data.get"
}

// TaobaoOmniorderGuideDataGetResponse 获取全渠道导购产品，目前包括随心购、随身购扫码、加购和交易数据。
type TaobaoOmniorderGuideDataGetResponse struct {
    
    /* data_list Basic Array拉取的数据数组，如果为空，表示数据拉取完毕。拉取的数据字段包括打点时间、商家id、商品id和门店id等，传入的类型不同，返回的字段有所不同，可以根据具体类型的返回结果具体处理 */
    data_list map[string]interface{} `json:"data_list";xml:"data_list"`
    
}
