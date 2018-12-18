package tbsdk

// QimenTaobaoCrmOrderSyncRequest CRM对于会员数据分析需要基于会员的购买行为，购买什么商品，价格，是否退换，使用什么优惠以及会员基础信息等等。这些信息通过销售订单数据来获取。
type QimenTaobaoCrmOrderSyncRequest struct {
    
    /* customerid requiredcustomerid */
    customerid string `json:"customerid";xml:"customerid"`
    
    /* data optional订单数据 */
    data Data `json:"data";xml:"data"`
    
    /* data_type optional数据类型（zx_order直销订单;fx_order分销订单） */
    data_type string `json:"data_type";xml:"data_type"`
    
    /* from_node_id optional数据来源ID */
    from_node_id string `json:"from_node_id";xml:"from_node_id"`
    
    /* from_type optional数据来源类型（taobao淘宝;JD京东;yihaodian一号店;dangdang当当;suning苏宁易购;amazon亚马逊;yinati银泰;mogujie蘑菇街;alibaba阿里巴巴;vop唯品会;meilishuo美丽说;youzan有赞;weixin微信;other其他） */
    from_type string `json:"from_type";xml:"from_type"`
    
    /* msg_id optional全局唯一任务编号 */
    msg_id string `json:"msg_id";xml:"msg_id"`
    
}

func (req *QimenTaobaoCrmOrderSyncRequest) GetAPIName() string {
	return "qimen.taobao.crm.order.sync"
}

// QimenTaobaoCrmOrderSyncResponse CRM对于会员数据分析需要基于会员的购买行为，购买什么商品，价格，是否退换，使用什么优惠以及会员基础信息等等。这些信息通过销售订单数据来获取。
type QimenTaobaoCrmOrderSyncResponse struct {
    
    /* code Basic0成功(其他失败) */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果 */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
}
