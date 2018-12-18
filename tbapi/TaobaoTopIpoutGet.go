package tbsdk

// TaobaoTopIpoutGetRequest 获取开放平台出口IP段
type TaobaoTopIpoutGetRequest struct {
    
}

func (req *TaobaoTopIpoutGetRequest) GetAPIName() string {
	return "taobao.top.ipout.get"
}

// TaobaoTopIpoutGetResponse 获取开放平台出口IP段
type TaobaoTopIpoutGetResponse struct {
    
    /* ip_list BasicTOP网关出口IP列表 */
    ip_list string `json:"ip_list";xml:"ip_list"`
    
}
