package tbsdk

// TaobaoPictureUserinfoGetRequest 查询用户的图片空间使用信息，包括：订购量，已使用容量，免费容量，总的可使用容量，订购有效期，剩余容量
type TaobaoPictureUserinfoGetRequest struct {
    
}

func (req *TaobaoPictureUserinfoGetRequest) GetAPIName() string {
	return "taobao.picture.userinfo.get"
}

// TaobaoPictureUserinfoGetResponse 查询用户的图片空间使用信息，包括：订购量，已使用容量，免费容量，总的可使用容量，订购有效期，剩余容量
type TaobaoPictureUserinfoGetResponse struct {
    
    /* user_info Object用户使用图片空间的信息 */
    user_info UserInfo `json:"user_info";xml:"user_info"`
    
}
