package tbsdk

// TaobaoInventoryAdjustTradeRequest 商家交易调整库存，淘宝交易、B2B经销等
type TaobaoInventoryAdjustTradeRequest struct {
    
    /* biz_unique_code required商家外部定单号 */
    biz_unique_code string `json:"biz_unique_code";xml:"biz_unique_code"`
    
    /* items required商品初始库存信息： [{ "TBOrderCode”:”淘宝交易号”,"TBSubOrderCode ":"淘宝子交易单号,赠品可以不填","”isGift”:”TRUE或者FALSE,是否赠品”,storeCode":"商家仓库编码"," scItemId ":"商品后端ID","scItemCode":"商品商家编码"," originScItemId ":"原商品ID","inventoryType":"","quantity":"111","isComplete":"TRUE或者FALSE，是否全部确认出库"}] */
    items string `json:"items";xml:"items"`
    
    /* operate_time required业务操作时间 */
    operate_time Date `json:"operate_time";xml:"operate_time"`
    
    /* tb_order_type required订单类型：B2C、B2B */
    tb_order_type string `json:"tb_order_type";xml:"tb_order_type"`
    
}

func (req *TaobaoInventoryAdjustTradeRequest) GetAPIName() string {
	return "taobao.inventory.adjust.trade"
}

// TaobaoInventoryAdjustTradeResponse 商家交易调整库存，淘宝交易、B2B经销等
type TaobaoInventoryAdjustTradeResponse struct {
    
    /* operate_code Basic操作返回码 */
    operate_code string `json:"operate_code";xml:"operate_code"`
    
    /* tip_infos Object Array提示信息 */
    tip_infos TipInfo `json:"tip_infos";xml:"tip_infos"`
    
}
