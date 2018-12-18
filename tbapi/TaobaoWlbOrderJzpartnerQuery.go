package tbsdk

// TaobaoWlbOrderJzpartnerQueryRequest 为支持家装类目的商家，对绑定家装物流服务的订单可以在商家的ERP中发货、批量发货，因此开发根据服务类型查询所有的服务商列表的接口
type TaobaoWlbOrderJzpartnerQueryRequest struct {
    
    /* service_type optionalserviceType表示查询所有的支持服务类型的服务商。 家装干线服务     11 家装干支服务     12 家装干支装服务   13 卫浴大件干线     14 卫浴大件干支     15 卫浴大件安装     16 地板干线         17 地板干支         18 地板安装         19 灯具安装         20 卫浴小件安装     21 （注：同一个服务商针对不同类型的serviceType是具有不同的tpCode的） */
    service_type int64 `json:"service_type";xml:"service_type"`
    
    /* taobao_trade_id optional淘宝交易订单号，如果不填写Tid则必须填写serviceType。如果填写Tid，则表明只需要查询对应订单的服务商。 */
    taobao_trade_id int64 `json:"taobao_trade_id";xml:"taobao_trade_id"`
    
}

func (req *TaobaoWlbOrderJzpartnerQueryRequest) GetAPIName() string {
	return "taobao.wlb.order.jzpartner.query"
}

// TaobaoWlbOrderJzpartnerQueryResponse 为支持家装类目的商家，对绑定家装物流服务的订单可以在商家的ERP中发货、批量发货，因此开发根据服务类型查询所有的服务商列表的接口
type TaobaoWlbOrderJzpartnerQueryResponse struct {
    
    /* install_list Object Array安装服务商列表 */
    install_list PartnerNew `json:"install_list";xml:"install_list"`
    
    /* is_success Basic接口查询成功或者失败 */
    is_success bool `json:"is_success";xml:"is_success"`
    
    /* result_info Basic查询返回信息，如果失败，存储错误信息 */
    result_info string `json:"result_info";xml:"result_info"`
    
    /* server_list Object Array物流配送服务商对象列表 */
    server_list PartnerNew `json:"server_list";xml:"server_list"`
    
}
