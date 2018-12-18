package tbsdk

// TaobaoLogisticsCompaniesGetRequest 查询淘宝网合作的物流公司信息，用于发货接口。
type TaobaoLogisticsCompaniesGetRequest struct {
    
    /* fields required需返回的字段列表。可选值:LogisticCompany 结构中的所有字段;多个字段间用","逗号隔开.
如:id,code,name,reg_mail_no
<br><font color='red'>说明：</font>
<br>id：物流公司ID
<br>code：物流公司code
<br>name：物流公司名称
<br>reg_mail_no：物流公司对应的运单规则 */
    fields string `json:"fields";xml:"fields"`
    
    /* is_recommended optional是否查询推荐物流公司.可选值:true,false.如果不提供此参数,将会返回所有支持电话联系的物流公司. */
    is_recommended bool `json:"is_recommended";xml:"is_recommended"`
    
    /* order_mode optional推荐物流公司的下单方式.可选值:offline(电话联系/自己联系),online(在线下单),all(即电话联系又在线下单). 此参数仅仅用于is_recommended 为ture时。就是说对于推荐物流公司才可用.如果不选择此参数将会返回推荐物流中支持电话联系的物流公司. */
    order_mode string `json:"order_mode";xml:"order_mode"`
    
}

func (req *TaobaoLogisticsCompaniesGetRequest) GetAPIName() string {
	return "taobao.logistics.companies.get"
}

// TaobaoLogisticsCompaniesGetResponse 查询淘宝网合作的物流公司信息，用于发货接口。
type TaobaoLogisticsCompaniesGetResponse struct {
    
    /* logistics_companies Object Array物流公司信息。返回的LogisticCompany包含的具体信息为入参fields请求的字段信息。 */
    logistics_companies LogisticsCompany `json:"logistics_companies";xml:"logistics_companies"`
    
}
