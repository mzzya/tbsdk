package tbsdk

// TaobaoWlbWmsCainiaoBillQueryRequest 查询单据列表
type TaobaoWlbWmsCainiaoBillQueryRequest struct {
    
    /* end_modified_time required起始时间，此字段检索订单最后修改时间， 格式 yyyy-MM-dd HH:mm:ss。 */
    end_modified_time Date `json:"end_modified_time";xml:"end_modified_time"`
    
    /* order_type optional订单类型 201 销售出库 501 退货入库 502 换货出库 503 补发出库904 普通入库 903 普通出库单 306 B2B入库单 305 B2B出库单 601 采购入库 901 退供出库单 701 盘点出库 702 盘点入库 711 库存异动单 */
    order_type string `json:"order_type";xml:"order_type"`
    
    /* page_no optional页码。（大于0的整数。默认为1） */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size optional每页条数。（每页条数不超过50条。默认为50） */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* start_modified_time required结束时间，此字段检索订单最后修改时间， 格式 yyyy-MM-dd HH:mm:ss。 */
    start_modified_time Date `json:"start_modified_time";xml:"start_modified_time"`
    
}

func (req *TaobaoWlbWmsCainiaoBillQueryRequest) GetAPIName() string {
	return "taobao.wlb.wms.cainiao.bill.query"
}

// TaobaoWlbWmsCainiaoBillQueryResponse 查询单据列表
type TaobaoWlbWmsCainiaoBillQueryResponse struct {
    
    /* order_info_list Object Array订单列表信息 */
    order_info_list CainiaoBillQueryOrderinfolist `json:"order_info_list";xml:"order_info_list"`
    
    /* total_count Basic总条数 */
    total_count int64 `json:"total_count";xml:"total_count"`
    
}
