package tbsdk

// TaobaoOmniorderItemTagOperateRequest 用于对全渠道商品进行打标、去标（门店发货标，门店自提标，前置拆单标）操作。另外还包括增加、删除、修改分单系统，接单系统配置。
type TaobaoOmniorderItemTagOperateRequest struct {
    
    /* item_id required商品ID */
    item_id int64 `json:"item_id";xml:"item_id"`
    
    /* omni_setting optional分单&接单设置 */
    omni_setting OmniSettingDto `json:"omni_setting";xml:"omni_setting"`
    
    /* status required操作状态， 填 1 代表打标，填 -1 代表去标 */
    status int64 `json:"status";xml:"status"`
    
    /* types required商品标,storeDeliver代表门店发货, AllocateByFront代表前置拆单, storeCollect代表门店自提 */
    types string `json:"types";xml:"types"`
    
}

func (req *TaobaoOmniorderItemTagOperateRequest) GetAPIName() string {
	return "taobao.omniorder.item.tag.operate"
}

// TaobaoOmniorderItemTagOperateResponse 用于对全渠道商品进行打标、去标（门店发货标，门店自提标，前置拆单标）操作。另外还包括增加、删除、修改分单系统，接单系统配置。
type TaobaoOmniorderItemTagOperateResponse struct {
    
    /* code Basic0 正常，否则异常 */
    code string `json:"code";xml:"code"`
    
    /* message Basiccode 不为 0时有值，代表异常信息 */
    message string `json:"message";xml:"message"`
    
}
