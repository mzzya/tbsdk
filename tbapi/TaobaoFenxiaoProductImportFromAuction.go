package tbsdk

// TaobaoFenxiaoProductImportFromAuctionRequest 供应商选择关联店铺的前台宝贝，导入生成产品
type TaobaoFenxiaoProductImportFromAuctionRequest struct {
    
    /* auction_id required店铺宝贝id */
    auction_id int64 `json:"auction_id";xml:"auction_id"`
    
    /* product_line_id required产品线id */
    product_line_id int64 `json:"product_line_id";xml:"product_line_id"`
    
    /* trade_type required导入产品需要支持的交易类型：[1 代销][ 2 经销 ][3 代销和经销] */
    trade_type int64 `json:"trade_type";xml:"trade_type"`
    
}

func (req *TaobaoFenxiaoProductImportFromAuctionRequest) GetAPIName() string {
	return "taobao.fenxiao.product.import.from.auction"
}

// TaobaoFenxiaoProductImportFromAuctionResponse 供应商选择关联店铺的前台宝贝，导入生成产品
type TaobaoFenxiaoProductImportFromAuctionResponse struct {
    
    /* opt_time Basic操作时间 */
    opt_time Date `json:"opt_time";xml:"opt_time"`
    
    /* pid Basic生成的产品id */
    pid int64 `json:"pid";xml:"pid"`
    
}
