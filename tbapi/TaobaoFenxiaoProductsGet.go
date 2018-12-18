package tbsdk

// TaobaoFenxiaoProductsGetRequest 查询供应商的产品数据。

    * 入参传入pids将优先查询，即只按这个条件查询。
    *入参传入sku_number将优先查询(没有传入pids)，即只按这个条件查询(最多显示50条)
    * 入参fields传skus将查询sku的数据，不传该参数默认不查询，返回产品的其它信息。
    * 入参fields传入images将查询多图数据，不传只返回主图数据。
    * 入参fields仅对传入pids生效（只有按ID查询时，才能查询额外的数据）
    * 查询结果按照产品发布时间倒序，即时间近的数据在前。
type TaobaoFenxiaoProductsGetRequest struct {
    
    /* end_modified optional结束修改时间 */
    end_modified Date `json:"end_modified";xml:"end_modified"`
    
    /* fields optional指定查询额外的信息，可选值：skus（sku数据）、images（多图），多个可选值用逗号分割。 */
    fields string `json:"fields";xml:"fields"`
    
    /* is_authz optional查询产品列表时，查询入参“是否需要授权”
yes:需要授权 
no:不需要授权 */
    is_authz string `json:"is_authz";xml:"is_authz"`
    
    /* item_ids optional可根据导入的商品ID列表查询，优先级次于产品ID、sku_numbers，高于其他分页查询条件。最大限制20，用逗号分割，例如：“1001,1002,1003,1004,1005” */
    item_ids string `json:"item_ids";xml:"item_ids"`
    
    /* outer_id optional商家编码 */
    outer_id string `json:"outer_id";xml:"outer_id"`
    
    /* page_no optional页码（大于0的整数，默认1） */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size optional每页记录数（默认20，最大50） */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* pids optional产品ID列表（最大限制30），用逗号分割，例如：“1001,1002,1003,1004,1005” */
    pids string `json:"pids";xml:"pids"`
    
    /* productcat_id optional产品线ID */
    productcat_id int64 `json:"productcat_id";xml:"productcat_id"`
    
    /* sku_number optionalsku商家编码 */
    sku_number string `json:"sku_number";xml:"sku_number"`
    
    /* start_modified optional开始修改时间 */
    start_modified Date `json:"start_modified";xml:"start_modified"`
    
}

func (req *TaobaoFenxiaoProductsGetRequest) GetAPIName() string {
	return "taobao.fenxiao.products.get"
}

// TaobaoFenxiaoProductsGetResponse 查询供应商的产品数据。

    * 入参传入pids将优先查询，即只按这个条件查询。
    *入参传入sku_number将优先查询(没有传入pids)，即只按这个条件查询(最多显示50条)
    * 入参fields传skus将查询sku的数据，不传该参数默认不查询，返回产品的其它信息。
    * 入参fields传入images将查询多图数据，不传只返回主图数据。
    * 入参fields仅对传入pids生效（只有按ID查询时，才能查询额外的数据）
    * 查询结果按照产品发布时间倒序，即时间近的数据在前。
type TaobaoFenxiaoProductsGetResponse struct {
    
    /* products Object Array产品对象记录集。返回 FenxiaoProduct 包含的字段信息。 */
    products FenxiaoProduct `json:"products";xml:"products"`
    
    /* total_results Basic查询结果记录数 */
    total_results int64 `json:"total_results";xml:"total_results"`
    
}
