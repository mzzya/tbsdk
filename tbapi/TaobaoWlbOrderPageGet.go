package tbsdk

// TaobaoWlbOrderPageGetRequest 分页查询物流宝订单
type TaobaoWlbOrderPageGetRequest struct {
    
    /* end_time optional查询截止时间 */
    end_time Date `json:"end_time";xml:"end_time"`
    
    /* order_code optional物流订单编号 */
    order_code string `json:"order_code";xml:"order_code"`
    
    /* order_status optionalTMS拒签：-100 接收方拒签：-200 */
    order_status int64 `json:"order_status";xml:"order_status"`
    
    /* order_sub_type optional订单子类型： （1）OTHER： 其他 （2）TAOBAO_TRADE： 淘宝交易 （3）OTHER_TRADE：其他交易 （4）ALLCOATE： 调拨 （5）CHECK: 盘点单 （6）PURCHASE: 采购单 */
    order_sub_type string `json:"order_sub_type";xml:"order_sub_type"`
    
    /* order_type optional订单类型: （1）NORMAL_OUT ：正常出库 （2）NORMAL_IN：正常入库 （3）RETURN_IN：退货入库 （4）EXCHANGE_OUT：换货出库 */
    order_type string `json:"order_type";xml:"order_type"`
    
    /* page_no optional分页的第几页 */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size optional每页多少条 */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* start_time optional查询开始时间 */
    start_time Date `json:"start_time";xml:"start_time"`
    
}

func (req *TaobaoWlbOrderPageGetRequest) GetAPIName() string {
	return "taobao.wlb.order.page.get"
}

// TaobaoWlbOrderPageGetResponse 分页查询物流宝订单
type TaobaoWlbOrderPageGetResponse struct {
    
    /* order_list Object Array物流宝订单对象 */
    order_list WlbOrder `json:"order_list";xml:"order_list"`
    
    /* total_count Basic总条数 */
    total_count int64 `json:"total_count";xml:"total_count"`
    
}
