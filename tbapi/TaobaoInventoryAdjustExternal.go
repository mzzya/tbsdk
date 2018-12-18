package tbsdk

// TaobaoInventoryAdjustExternalRequest 商家非交易调整库存，调拨出库、盘点等时调用
type TaobaoInventoryAdjustExternalRequest struct {
    
    /* biz_type optional外部订单类型, BALANCE:盘点、NON_TAOBAO_TRADE:非淘宝交易、ALLOCATE:调拨、OTHERS:其他 */
    biz_type string `json:"biz_type";xml:"biz_type"`
    
    /* biz_unique_code required商家外部定单号 */
    biz_unique_code string `json:"biz_unique_code";xml:"biz_unique_code"`
    
    /* items required商品初始库存信息： [{"scItemId":"商品后端ID，如果有传scItemCode,参数可以为0","scItemCode":"商品商家编码","inventoryType":"库存类型  1：正常,”direction”: 1: 盘盈 -1: 盘亏,参数可选,"quantity":"数量(正数)"}] */
    items string `json:"items";xml:"items"`
    
    /* occupy_operate_code optional库存占用返回的操作码.operate_type 为OUTBOUND时，如果是确认事先进行过的库存占用，需要传入当时返回的操作码，并且明细必须与申请时保持一致 */
    occupy_operate_code string `json:"occupy_operate_code";xml:"occupy_operate_code"`
    
    /* operate_time optional业务操作时间 */
    operate_time string `json:"operate_time";xml:"operate_time"`
    
    /* operate_type optionaltest */
    operate_type string `json:"operate_type";xml:"operate_type"`
    
    /* reduce_type optionaltest */
    reduce_type string `json:"reduce_type";xml:"reduce_type"`
    
    /* store_code required商家仓库编码 */
    store_code string `json:"store_code";xml:"store_code"`
    
}

func (req *TaobaoInventoryAdjustExternalRequest) GetAPIName() string {
	return "taobao.inventory.adjust.external"
}

// TaobaoInventoryAdjustExternalResponse 商家非交易调整库存，调拨出库、盘点等时调用
type TaobaoInventoryAdjustExternalResponse struct {
    
    /* operate_code Basic操作返回码 */
    operate_code string `json:"operate_code";xml:"operate_code"`
    
    /* tip_infos Object Array提示信息 */
    tip_infos TipInfo `json:"tip_infos";xml:"tip_infos"`
    
}
