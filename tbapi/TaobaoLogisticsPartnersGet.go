package tbsdk

// TaobaoLogisticsPartnersGetRequest 查询物流公司信息（可以查询目的地可不可达情况）
type TaobaoLogisticsPartnersGetRequest struct {
    
    /* goods_value optional货物价格.只有当选择货到付款此参数才会有效 */
    goods_value string `json:"goods_value";xml:"goods_value"`
    
    /* is_need_carriage optional是否需要揽收资费信息，默认false。在此值为false时，返回内容中将无carriage。在未设置source_id或target_id的情况下，无法查询揽收资费信息。自己联系无揽收资费记录。 */
    is_need_carriage bool `json:"is_need_carriage";xml:"is_need_carriage"`
    
    /* service_type optional服务类型，根据此参数可查出提供相应服务类型的物流公司信息(物流公司状态正常)，可选值：cod(货到付款)、online(在线下单)、 offline(自己联系)、limit(限时物流)。然后再根据source_id,target_id,goods_value这三个条件来过滤物流公司. 目前输入自己联系服务类型将会返回空，因为自己联系并没有具体的服务信息，所以不会有记录。 */
    service_type string `json:"service_type";xml:"service_type"`
    
    /* source_id optional物流公司揽货地地区码（必须是区、县一级的）.参考:http://www.stats.gov.cn/tjsj/tjbz/xzqhdm/201401/t20140116_501070.html或者调用 taobao.areas.get 获取 */
    source_id string `json:"source_id";xml:"source_id"`
    
    /* target_id optional物流公司派送地地区码（必须是区、县一级的）.参考:http://www.stats.gov.cn/tjsj/tjbz/xzqhdm/201401/t20140116_501070.html或者调用 taobao.areas.get 获取 */
    target_id string `json:"target_id";xml:"target_id"`
    
}

func (req *TaobaoLogisticsPartnersGetRequest) GetAPIName() string {
	return "taobao.logistics.partners.get"
}

// TaobaoLogisticsPartnersGetResponse 查询物流公司信息（可以查询目的地可不可达情况）
type TaobaoLogisticsPartnersGetResponse struct {
    
    /* logistics_partners Object Array查询揽送范围之内的物流公司信息 */
    logistics_partners LogisticsPartner `json:"logistics_partners";xml:"logistics_partners"`
    
}
