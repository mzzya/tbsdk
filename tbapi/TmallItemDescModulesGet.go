package tbsdk

// TmallItemDescModulesGetRequest 商品描述模块信息获取，包括运营设定的类目级别的模块信息以及用户自定义模块数量约束。
type TmallItemDescModulesGetRequest struct {
    
    /* cat_id required淘宝后台发布商品的叶子类目id，可通过taobao.itemcats.get查到。api 访问地址http://api.taobao.com/apidoc/api.htm?spm=0.0.0.0.CFhhk4&path=cid:3-apiId:122 */
    cat_id int64 `json:"cat_id";xml:"cat_id"`
    
    /* usr_id required商家主帐号id */
    usr_id string `json:"usr_id";xml:"usr_id"`
    
}

func (req *TmallItemDescModulesGetRequest) GetAPIName() string {
	return "tmall.item.desc.modules.get"
}

// TmallItemDescModulesGetResponse 商品描述模块信息获取，包括运营设定的类目级别的模块信息以及用户自定义模块数量约束。
type TmallItemDescModulesGetResponse struct {
    
    /* modular_desc_info Object返回描述模块信息 */
    modular_desc_info ModularDescInfo `json:"modular_desc_info";xml:"modular_desc_info"`
    
}
