package tbsdk

// TaobaoFenxiaoDistributorItemsGetRequest 供应商查询分销商商品下载记录。
type TaobaoFenxiaoDistributorItemsGetRequest struct {
    
    /* distributor_id special分销商ID 。 */
    distributor_id int64 `json:"distributor_id";xml:"distributor_id"`
    
    /* end_modified optional设置结束时间,空为不设置。 */
    end_modified Date `json:"end_modified";xml:"end_modified"`
    
    /* page_no optional页码（大于0的整数，默认1） */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size optional每页记录数（默认20，最大50） */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* product_id special产品ID */
    product_id int64 `json:"product_id";xml:"product_id"`
    
    /* start_modified optional设置开始时间。空为不设置。 */
    start_modified Date `json:"start_modified";xml:"start_modified"`
    
}

func (req *TaobaoFenxiaoDistributorItemsGetRequest) GetAPIName() string {
	return "taobao.fenxiao.distributor.items.get"
}

// TaobaoFenxiaoDistributorItemsGetResponse 供应商查询分销商商品下载记录。
type TaobaoFenxiaoDistributorItemsGetResponse struct {
    
    /* records Object Array下载记录对象 */
    records FenxiaoItemRecord `json:"records";xml:"records"`
    
    /* total_results Basic查询结果记录数 */
    total_results int64 `json:"total_results";xml:"total_results"`
    
}
