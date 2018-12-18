package tbsdk

// TaobaoAreasGetRequest 查询标准地址区域代码信息。可以直接参考最新的行政区域代码：
<a href="http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2016/index.html">http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2016/index.html</a>
type TaobaoAreasGetRequest struct {
    
    /* fields required需返回的字段列表.可选值:Area 结构中的所有字段;多个字段之间用","分隔.如:id,type,name,parent_id,zip. */
    fields string `json:"fields";xml:"fields"`
    
}

func (req *TaobaoAreasGetRequest) GetAPIName() string {
	return "taobao.areas.get"
}

// TaobaoAreasGetResponse 查询标准地址区域代码信息。可以直接参考最新的行政区域代码：
<a href="http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2016/index.html">http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2016/index.html</a>
type TaobaoAreasGetResponse struct {
    
    /* areas Object Array地址区域信息列表.返回的Area包含的具体信息为入参fields请求的字段信息 。 */
    areas Area `json:"areas";xml:"areas"`
    
}
