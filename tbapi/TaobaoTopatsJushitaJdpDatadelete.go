package tbsdk

// TaobaoTopatsJushitaJdpDatadeleteRequest 异步删除rds库数据推送表的数据
type TaobaoTopatsJushitaJdpDatadeleteRequest struct {
    
    /* end_date required删除数据时间段的结束修改时间，格式为：yyyy-MM-dd HH:mm:ss，结束时间必须为前天的23:59:59秒以前，根据是业务的modified时间 */
    end_date Date `json:"end_date";xml:"end_date"`
    
    /* rds_name optional数据库实例名，当删除用户的绑定关系已经不在推送里时，此参数必填 */
    rds_name string `json:"rds_name";xml:"rds_name"`
    
    /* start_date required删除数据时间段的起始修改时间，格式为：yyyy-MM-dd HH:mm:ss,根据是业务的modified时间 */
    start_date Date `json:"start_date";xml:"start_date"`
    
    /* sync_types required推送的数据类型，可选值为：tb_trade(淘宝交易)，tb_item(淘宝商品)，tb_refund(淘宝退款)，fx_trade(分销订单)，同时删除多种类型以分号分隔，如："tb_trade;tb_item;tb_refund;fx_trade" */
    sync_types string `json:"sync_types";xml:"sync_types"`
    
    /* user_nick required用户昵称，必填 */
    user_nick string `json:"user_nick";xml:"user_nick"`
    
}

func (req *TaobaoTopatsJushitaJdpDatadeleteRequest) GetAPIName() string {
	return "taobao.topats.jushita.jdp.datadelete"
}

// TaobaoTopatsJushitaJdpDatadeleteResponse 异步删除rds库数据推送表的数据
type TaobaoTopatsJushitaJdpDatadeleteResponse struct {
    
    /* result Objectresult */
    result ResultDo `json:"result";xml:"result"`
    
}
