package tbsdk

// TaobaoFenxiaoDistributorProductsGetRequest 分销商查询供应商产品信息
type TaobaoFenxiaoDistributorProductsGetRequest struct {
    
    /* download_status optionaldownload_status */
    download_status string `json:"download_status";xml:"download_status"`
    
    /* end_time optional结束时间 */
    end_time Date `json:"end_time";xml:"end_time"`
    
    /* fields optional指定查询额外的信息，可选值：skus（sku数据）、images（多图），多个可选值用逗号分割。 */
    fields string `json:"fields";xml:"fields"`
    
    /* item_ids optional根据商品ID列表查询，优先级次于产品ID列表，高于其他分页查询条件。如果商品不是分销商品，自动过滤。最大限制20，用逗号分割，例如：“1001,1002,1003,1004,1005” */
    item_ids int64 `json:"item_ids";xml:"item_ids"`
    
    /* order_by optionalorder_by */
    order_by string `json:"order_by";xml:"order_by"`
    
    /* page_no optional页码（大于0的整数，默认1） */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size optional每页记录数（默认20，最大50） */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* pids optional产品ID列表（最大限制30），用逗号分割，例如：“1001,1002,1003,1004,1005” */
    pids int64 `json:"pids";xml:"pids"`
    
    /* productcat_id optional产品线ID */
    productcat_id int64 `json:"productcat_id";xml:"productcat_id"`
    
    /* start_time optional开始修改时间 */
    start_time Date `json:"start_time";xml:"start_time"`
    
    /* supplier_nick optional供应商nick，分页查询时，必传 */
    supplier_nick string `json:"supplier_nick";xml:"supplier_nick"`
    
    /* time_type optionaltime_type */
    time_type string `json:"time_type";xml:"time_type"`
    
    /* trade_type optionaltrade_type */
    trade_type string `json:"trade_type";xml:"trade_type"`
    
}

func (req *TaobaoFenxiaoDistributorProductsGetRequest) GetAPIName() string {
	return "taobao.fenxiao.distributor.products.get"
}

// TaobaoFenxiaoDistributorProductsGetResponse 分销商查询供应商产品信息
type TaobaoFenxiaoDistributorProductsGetResponse struct {
    
    /* has_next Basic是否存在下一页 */
    has_next bool `json:"has_next";xml:"has_next"`
    
    /* products Object Array产品对象记录集。返回 FenxiaoProduct 包含的字段信息。 */
    products FenxiaoProduct `json:"products";xml:"products"`
    
}
