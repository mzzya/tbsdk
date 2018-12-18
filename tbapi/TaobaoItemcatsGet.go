package tbsdk

// TaobaoItemcatsGetRequest 获取后台供卖家发布商品的标准商品类目。
type TaobaoItemcatsGetRequest struct {
    
    /* cids special商品所属类目ID列表，用半角逗号(,)分隔 例如:(18957,19562,) (cids、parent_cid至少传一个) */
    cids int64 `json:"cids";xml:"cids"`
    
    /* datetime special无效字段，暂无法使用。时间戳（格式:yyyy-MM-dd HH:mm:ss）如果该字段没有传，则取当前所有的类目信息,如果传了parent_cid或者cids，则忽略datetime，如果该字段传了，那么可以查询到该时间到现在为止的增量变化 */
    datetime Date `json:"datetime";xml:"datetime"`
    
    /* fields optional需要返回的字段列表，见ItemCat，默认返回：cid,parent_cid,name,is_parent；增量类目信息,根据fields传入的参数返回相应的结果。 features字段： 1、如果存在attr_key=freeze表示该类目被冻结了，attr_value=0,5，value可能存在2个值（也可能只有1个），用逗号分割，0表示禁编辑，5表示禁止发布 */
    fields string `json:"fields";xml:"fields"`
    
    /* parent_cid special父商品类目 id，0表示根节点, 传输该参数返回所有子类目。 (cids、parent_cid至少传一个) */
    parent_cid int64 `json:"parent_cid";xml:"parent_cid"`
    
}

func (req *TaobaoItemcatsGetRequest) GetAPIName() string {
	return "taobao.itemcats.get"
}

// TaobaoItemcatsGetResponse 获取后台供卖家发布商品的标准商品类目。
type TaobaoItemcatsGetResponse struct {
    
    /* item_cats Object Array增量类目信息,根据fields传入的参数返回相应的结果；增量类目信息,根据fields传入的参数返回相应的结果。 features字段： 1、如果存在attr_key=freeze表示该类目被冻结了，attr_value=0,5，value可能存在2个值（也可能只有1个），用逗号分割，0表示禁编辑，5表示禁止发布 */
    item_cats ItemCat `json:"item_cats";xml:"item_cats"`
    
    /* last_modified Basic最近修改时间(如果取增量，会返回该字段)。 */
    last_modified Date `json:"last_modified";xml:"last_modified"`
    
}
