package tbsdk

// TaobaoWlbOutInventoryChangeNotifyRequest 拥有自有仓的企业物流用户通过该接口把自有仓的库存通知到物流宝，由物流宝维护该库存，控制前台显示库存的准确性。
type TaobaoWlbOutInventoryChangeNotifyRequest struct {
    
    /* change_count required库存变化数量 */
    change_count int64 `json:"change_count";xml:"change_count"`
    
    /* item_id required物流宝商品id或前台宝贝id（由type类型决定） */
    item_id int64 `json:"item_id";xml:"item_id"`
    
    /* op_type requiredOUT--出库 IN--入库 */
    op_type string `json:"op_type";xml:"op_type"`
    
    /* order_source_code optional订单号，如果source为TAOBAO_TRADE,order_source_code必须为淘宝交易号 */
    order_source_code string `json:"order_source_code";xml:"order_source_code"`
    
    /* out_biz_code required库存变化唯一标识，用于去重，一个外部唯一编码唯一标识一次库存变化。 */
    out_biz_code string `json:"out_biz_code";xml:"out_biz_code"`
    
    /* result_count required本次库存变化后库存余额 */
    result_count int64 `json:"result_count";xml:"result_count"`
    
    /* source required（1）OTHER： 其他 （2）TAOBAO_TRADE： 淘宝交易 （3）OTHER_TRADE：其他交易 （4）ALLCOATE： 调拨 （5）CHECK:盘点 （6）PURCHASE:采购 */
    source string `json:"source";xml:"source"`
    
    /* store_code optional目前非必须，系统会选择默认值 */
    store_code string `json:"store_code";xml:"store_code"`
    
    /* _type requiredWLB_ITEM--物流宝商品 IC_ITEM--淘宝商品 IC_SKU--淘宝sku */
    _type string `json:"type";xml:"type"`
    
}

func (req *TaobaoWlbOutInventoryChangeNotifyRequest) GetAPIName() string {
	return "taobao.wlb.out.inventory.change.notify"
}

// TaobaoWlbOutInventoryChangeNotifyResponse 拥有自有仓的企业物流用户通过该接口把自有仓的库存通知到物流宝，由物流宝维护该库存，控制前台显示库存的准确性。
type TaobaoWlbOutInventoryChangeNotifyResponse struct {
    
    /* gmt_modified Basic库存变化通知成功时间 */
    gmt_modified Date `json:"gmt_modified";xml:"gmt_modified"`
    
}
