package tbsdk

// TaobaoItempropsGetRequest 通过设置必要的参数，来获取商品后台标准类目属性，以及这些属性里面详细的属性值prop_values。
type TaobaoItempropsGetRequest struct {
    
    /* attr_keys optional属性的Key，支持多条，以“,”分隔 */
    attr_keys string `json:"attr_keys";xml:"attr_keys"`
    
    /* child_path optional类目子属性路径,由该子属性上层的类目属性和类目属性值组成,格式pid:vid;pid:vid.取类目子属性需要传child_path,cid */
    child_path string `json:"child_path";xml:"child_path"`
    
    /* cid required叶子类目ID，如果只传cid，则只返回一级属性,通过taobao.itemcats.get获得叶子类目ID */
    cid int64 `json:"cid";xml:"cid"`
    
    /* datetime optional增量时间戳。格式:yyyy-MM-dd HH:mm:ss假如传2005-01-01 00:00:00，则取所有的属性和子属性ID(如果传了pid会忽略datetime) */
    datetime Date `json:"datetime";xml:"datetime"`
    
    /* fields optional需要返回的字段列表，见：ItemProp，默认返回：pid, name, must, multi, prop_values */
    fields string `json:"fields";xml:"fields"`
    
    /* is_color_prop optional是否颜色属性。可选值:true(是),false(否) (删除的属性不会匹配和返回这个条件) */
    is_color_prop bool `json:"is_color_prop";xml:"is_color_prop"`
    
    /* is_enum_prop optional是否枚举属性。可选值:true(是),false(否) (删除的属性不会匹配和返回这个条件)。如果返回true，属性值是下拉框选择输入，如果返回false，属性值是用户自行手工输入。 */
    is_enum_prop bool `json:"is_enum_prop";xml:"is_enum_prop"`
    
    /* is_input_prop optional在is_enum_prop是true的前提下，是否是卖家可以自行输入的属性（注：如果is_enum_prop返回false，该参数统一返回false）。可选值:true(是),false(否) (删除的属性不会匹配和返回这个条件) */
    is_input_prop bool `json:"is_input_prop";xml:"is_input_prop"`
    
    /* is_item_prop optional是否商品属性，这个属性只能放于发布商品时使用。可选值:true(是),false(否) */
    is_item_prop bool `json:"is_item_prop";xml:"is_item_prop"`
    
    /* is_key_prop optional是否关键属性。可选值:true(是),false(否) */
    is_key_prop bool `json:"is_key_prop";xml:"is_key_prop"`
    
    /* is_sale_prop optional是否销售属性。可选值:true(是),false(否) */
    is_sale_prop bool `json:"is_sale_prop";xml:"is_sale_prop"`
    
    /* parent_pid optional父属性ID */
    parent_pid int64 `json:"parent_pid";xml:"parent_pid"`
    
    /* pid optional属性id (取类目属性时，传pid，不用同时传PID和parent_pid) */
    pid int64 `json:"pid";xml:"pid"`
    
    /* _type optional获取类目的类型：1代表集市、2代表天猫 */
    _type int64 `json:"type";xml:"type"`
    
}

func (req *TaobaoItempropsGetRequest) GetAPIName() string {
	return "taobao.itemprops.get"
}

// TaobaoItempropsGetResponse 通过设置必要的参数，来获取商品后台标准类目属性，以及这些属性里面详细的属性值prop_values。
type TaobaoItempropsGetResponse struct {
    
    /* item_props Object Array类目属性信息(如果是取全量或者增量，不包括属性值),根据fields传入的参数返回相应的结果 */
    item_props ItemProp `json:"item_props";xml:"item_props"`
    
    /* last_modified Basic最近修改时间(只有取全量或增量的时候会返回该字段)。格式:yyyy-MM-dd HH:mm:ss */
    last_modified Date `json:"last_modified";xml:"last_modified"`
    
}
