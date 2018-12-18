package tbsdk

// TaobaoFenxiaoLoginUserGetRequest 获取用户登录信息
type TaobaoFenxiaoLoginUserGetRequest struct {
    
}

func (req *TaobaoFenxiaoLoginUserGetRequest) GetAPIName() string {
	return "taobao.fenxiao.login.user.get"
}

// TaobaoFenxiaoLoginUserGetResponse 获取用户登录信息
type TaobaoFenxiaoLoginUserGetResponse struct {
    
    /* login_user Object登录用户信息 */
    login_user LoginUser `json:"login_user";xml:"login_user"`
    
}
