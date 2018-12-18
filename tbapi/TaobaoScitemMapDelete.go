package tbsdk

// TaobaoScitemMapDeleteRequest 根据后端商品Id，失效指定用户的商品与后端商品的映射关系
type TaobaoScitemMapDeleteRequest struct {
    
    /* sc_item_id required后台商品ID */
    sc_item_id int64 `json:"sc_item_id";xml:"sc_item_id"`
    
    /* user_nick optional店铺用户nick。 如果该参数为空则删除后端商品与当前调用人的商品映射关系;如果不为空则删除指定用户与后端商品的映射关系 */
    user_nick string `json:"user_nick";xml:"user_nick"`
    
}

func (req *TaobaoScitemMapDeleteRequest) GetAPIName() string {
	return "taobao.scitem.map.delete"
}

// TaobaoScitemMapDeleteResponse 根据后端商品Id，失效指定用户的商品与后端商品的映射关系
type TaobaoScitemMapDeleteResponse struct {
    
    /* module Basic失效条数 */
    module int64 `json:"module";xml:"module"`
    
}
