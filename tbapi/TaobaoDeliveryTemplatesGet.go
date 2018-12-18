package tbsdk

// TaobaoDeliveryTemplatesGetRequest 根据用户ID获取用户下所有模板
type TaobaoDeliveryTemplatesGetRequest struct {
    
    /* fields required需返回的字段列表。 <br/> 
<B>
可选值:示例中的值;各字段之间用","隔开
</B>
<br/><br/> 
<font color=blue>
template_id：查询模板ID <br/> 
template_name:查询模板名称 <br/> 
assumer：查询服务承担放 <br/> 
valuation:查询计价规则 <br/> 
supports:查询增值服务列表 <br/> 
created:查询模板创建时间 <br/> 
modified:查询修改时间<br/> 
consign_area_id:运费模板上设置的卖家发货地址最小级别ID<br/> 
address:卖家地址信息
</font>
<br/><br/> 
<font color=bule>
query_cod:查询货到付款运费信息<br/> 
query_post:查询平邮运费信息 <br/> 
query_express:查询快递公司运费信息 <br/> 
query_ems:查询EMS运费信息<br/> 
query_bzsd:查询普通保障速递运费信息<br/> 
query_wlb:查询物流宝保障速递运费信息<br/> 
query_furniture:查询家装物流运费信息
<font color=blue>
<br/><br/>
<font color=red>不能有空格</font> */
    fields string `json:"fields";xml:"fields"`
    
}

func (req *TaobaoDeliveryTemplatesGetRequest) GetAPIName() string {
	return "taobao.delivery.templates.get"
}

// TaobaoDeliveryTemplatesGetResponse 根据用户ID获取用户下所有模板
type TaobaoDeliveryTemplatesGetResponse struct {
    
    /* delivery_templates Object Array运费模板列表 */
    delivery_templates DeliveryTemplate `json:"delivery_templates";xml:"delivery_templates"`
    
    /* total_results Basic获得到符合条件的结果总数 */
    total_results int64 `json:"total_results";xml:"total_results"`
    
}
