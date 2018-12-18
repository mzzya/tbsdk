package tbsdk

// TaobaoWlbOrderJzwithinsConsignRequest 为支持家装类目的商家，对绑定家装物流服务的订单可以在商家的ERP中发货、批量发货，因此开发带安装服务商的发货接口
type TaobaoWlbOrderJzwithinsConsignRequest struct {
    
    /* ins_partner optional物流服务商信息 */
    ins_partner JzPartnerNew `json:"ins_partner";xml:"ins_partner"`
    
    /* jz_consign_args required家装物流发货参数 */
    jz_consign_args JzConsignArgsNew `json:"jz_consign_args";xml:"jz_consign_args"`
    
    /* tid required淘宝交易订单号 */
    tid int64 `json:"tid";xml:"tid"`
    
    /* tms_partner required物流服务商信息 */
    tms_partner JzPartnerNew `json:"tms_partner";xml:"tms_partner"`
    
}

func (req *TaobaoWlbOrderJzwithinsConsignRequest) GetAPIName() string {
	return "taobao.wlb.order.jzwithins.consign"
}

// TaobaoWlbOrderJzwithinsConsignResponse 为支持家装类目的商家，对绑定家装物流服务的订单可以在商家的ERP中发货、批量发货，因此开发带安装服务商的发货接口
type TaobaoWlbOrderJzwithinsConsignResponse struct {
    
    /* is_success Basic发货成功或者失败 */
    is_success bool `json:"is_success";xml:"is_success"`
    
    /* result_info Basic发货返回信息，如果发货错误则报出对应错误 */
    result_info string `json:"result_info";xml:"result_info"`
    
}
