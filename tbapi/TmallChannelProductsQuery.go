package tbsdk

// TmallChannelProductsQueryRequest 渠道中心，供应商查询其产品数据，返回同时符合所有查询条件的产品信息
type TmallChannelProductsQueryRequest struct {
    
    /* page_num required页码数，从1开始 */
    page_num int64 `json:"page_num";xml:"page_num"`
    
    /* page_size required分页大小 */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* product_ids optional产品Id */
    product_ids int64 `json:"product_ids";xml:"product_ids"`
    
    /* product_line_id optional产品线Id */
    product_line_id int64 `json:"product_line_id";xml:"product_line_id"`
    
    /* product_number optional商家产品编码 */
    product_number string `json:"product_number";xml:"product_number"`
    
    /* sku_number optional商家SKU编码 */
    sku_number string `json:"sku_number";xml:"sku_number"`
    
}

func (req *TmallChannelProductsQueryRequest) GetAPIName() string {
	return "tmall.channel.products.query"
}

// TmallChannelProductsQueryResponse 渠道中心，供应商查询其产品数据，返回同时符合所有查询条件的产品信息
type TmallChannelProductsQueryResponse struct {
    
    /* result Objectresult */
    result PageResultDto `json:"result";xml:"result"`
    
}
