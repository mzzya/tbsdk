package tbsdk

// TaobaoItempropvaluesGetRequest 获取标准类目属性值
type TaobaoItempropvaluesGetRequest struct {
    
    /* attr_keys optional属性的Key，支持多条，以“,”分隔 */
    attr_keys string `json:"attr_keys";xml:"attr_keys"`
    
    /* cid required叶子类目ID ,通过taobao.itemcats.get获得叶子类目ID */
    cid int64 `json:"cid";xml:"cid"`
    
    /* datetime special假如传2005-01-01 00:00:00，则取所有的属性和子属性(状态为删除的属性值不返回prop_name) */
    datetime Date `json:"datetime";xml:"datetime"`
    
    /* fields required需要返回的字段。目前支持有：cid,pid,prop_name,vid,name,name_alias,status,sort_order */
    fields string `json:"fields";xml:"fields"`
    
    /* pvs special属性和属性值 id串，格式例如(pid1;pid2)或(pid1:vid1;pid2:vid2)或(pid1;pid2:vid2) */
    pvs string `json:"pvs";xml:"pvs"`
    
    /* _type optional获取类目的类型：1代表集市、2代表天猫 */
    _type int64 `json:"type";xml:"type"`
    
}

func (req *TaobaoItempropvaluesGetRequest) GetAPIName() string {
	return "taobao.itempropvalues.get"
}

// TaobaoItempropvaluesGetResponse 获取标准类目属性值
type TaobaoItempropvaluesGetResponse struct {
    
    /* last_modified Basic最近修改时间。格式:yyyy-MM-dd HH:mm:ss */
    last_modified string `json:"last_modified";xml:"last_modified"`
    
    /* prop_values Object Array属性值,根据fields传入的参数返回相应的结果 */
    prop_values PropValue `json:"prop_values";xml:"prop_values"`
    
}
