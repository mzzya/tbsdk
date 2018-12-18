package tbsdk

// TaobaoJushitaJdpUserAddRequest 提供给接入数据推送的应用添加数据推送服务的用户
type TaobaoJushitaJdpUserAddRequest struct {
    
    /* history_days optional推送历史数据天数，只能为90天内，包含90天。当此参数不填时，表示以页面中应用配置的历史天数为准；如果为0表示这个用户不推送历史数据；其它表示推送的历史天数。 */
    history_days int64 `json:"history_days";xml:"history_days"`
    
    /* host_ip optional已废弃，新版订单同步服务不要使用。同步用户数据的机器IP,必须是界面配置的IP。 */
    host_ip string `json:"host_ip";xml:"host_ip"`
    
    /* host_name optional已废弃，新版订单同步服务不要使用。绑定类型，用户数据同步的机器名称。 */
    host_name string `json:"host_name";xml:"host_name"`
    
    /* rds_name requiredRDS实例名称 */
    rds_name string `json:"rds_name";xml:"rds_name"`
    
    /* topics optional已废弃，使用页面中应用的配置。用户同步的数据类型,多个用','号分割。可选值：trade,refund,item */
    topics string `json:"topics";xml:"topics"`
    
}

func (req *TaobaoJushitaJdpUserAddRequest) GetAPIName() string {
	return "taobao.jushita.jdp.user.add"
}

// TaobaoJushitaJdpUserAddResponse 提供给接入数据推送的应用添加数据推送服务的用户
type TaobaoJushitaJdpUserAddResponse struct {
    
    /* is_success Basic是否添加成功 */
    is_success bool `json:"is_success";xml:"is_success"`
    
}
