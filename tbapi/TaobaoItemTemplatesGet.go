package tbsdk

// TaobaoItemTemplatesGetRequest 查询当前登录用户的店铺的宝贝详情页的模板名称
type TaobaoItemTemplatesGetRequest struct {
    
}

func (req *TaobaoItemTemplatesGetRequest) GetAPIName() string {
	return "taobao.item.templates.get"
}

// TaobaoItemTemplatesGetResponse 查询当前登录用户的店铺的宝贝详情页的模板名称
type TaobaoItemTemplatesGetResponse struct {
    
    /* item_template_list Object Array返回宝贝模板对象。包含模板id，模板name，还有模板的类别（0表示外店，1表示内店） */
    item_template_list ItemTemplate `json:"item_template_list";xml:"item_template_list"`
    
}
