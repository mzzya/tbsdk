package tbsdk

// TaobaoWlbStoresBaseinfoGetRequest 通过USERID和仓库类型，获取商家自有仓库或菜鸟仓库或全部仓库
type TaobaoWlbStoresBaseinfoGetRequest struct {
    
    /* _type optional0.商家仓库.1.菜鸟仓库.2全部被划分的仓库 */
    _type int64 `json:"type";xml:"type"`
    
}

func (req *TaobaoWlbStoresBaseinfoGetRequest) GetAPIName() string {
	return "taobao.wlb.stores.baseinfo.get"
}

// TaobaoWlbStoresBaseinfoGetResponse 通过USERID和仓库类型，获取商家自有仓库或菜鸟仓库或全部仓库
type TaobaoWlbStoresBaseinfoGetResponse struct {
    
    /* store_info_list Object Array仓库列表 */
    store_info_list StoreInfo `json:"store_info_list";xml:"store_info_list"`
    
    /* total_count Basic返回的总数 */
    total_count int64 `json:"total_count";xml:"total_count"`
    
}
