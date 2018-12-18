package tbsdk

// TaobaoJushitaJdpUserDeleteRequest 删除应用的数据推送用户，用户被删除后，重新添加时会重新同步历史数据。
type TaobaoJushitaJdpUserDeleteRequest struct {
    
    /* nick special要删除用户的昵称 */
    nick string `json:"nick";xml:"nick"`
    
    /* user_id special需要删除的用户编号 */
    user_id int64 `json:"user_id";xml:"user_id"`
    
}

func (req *TaobaoJushitaJdpUserDeleteRequest) GetAPIName() string {
	return "taobao.jushita.jdp.user.delete"
}

// TaobaoJushitaJdpUserDeleteResponse 删除应用的数据推送用户，用户被删除后，重新添加时会重新同步历史数据。
type TaobaoJushitaJdpUserDeleteResponse struct {
    
    /* is_success Basic是否删除成功 */
    is_success bool `json:"is_success";xml:"is_success"`
    
}
