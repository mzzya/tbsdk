package tbsdk

// TaobaoOpenuidChangeRequest 将当前应用所属的openUId 转换为对应targetAppkey的openUid
规则：
1.如果两个appkey是应用前后台关系，可以直接转换；
2.如果appkey和targetAppkey都属于同一个开发者，不允许互相转换。
3.如果appkey和targetAppkey不属于同一个开发者，不允许互相转换。
type TaobaoOpenuidChangeRequest struct {
    
    /* open_uid requiredopenUid */
    open_uid string `json:"open_uid";xml:"open_uid"`
    
    /* target_app_key required转换到的appkey */
    target_app_key string `json:"target_app_key";xml:"target_app_key"`
    
}

func (req *TaobaoOpenuidChangeRequest) GetAPIName() string {
	return "taobao.openuid.change"
}

// TaobaoOpenuidChangeResponse 将当前应用所属的openUId 转换为对应targetAppkey的openUid
规则：
1.如果两个appkey是应用前后台关系，可以直接转换；
2.如果appkey和targetAppkey都属于同一个开发者，不允许互相转换。
3.如果appkey和targetAppkey不属于同一个开发者，不允许互相转换。
type TaobaoOpenuidChangeResponse struct {
    
    /* new_open_uid Basic转换到新的openUId */
    new_open_uid string `json:"new_open_uid";xml:"new_open_uid"`
    
}
