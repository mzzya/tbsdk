package tbsdk

// TaobaoPictureUpdateRequest 修改指定图片的图片名
type TaobaoPictureUpdateRequest struct {
    
    /* new_name required新的图片名，最大长度50字符，不能为空 */
    new_name string `json:"new_name";xml:"new_name"`
    
    /* picture_id required要更改名字的图片的id */
    picture_id int64 `json:"picture_id";xml:"picture_id"`
    
}

func (req *TaobaoPictureUpdateRequest) GetAPIName() string {
	return "taobao.picture.update"
}

// TaobaoPictureUpdateResponse 修改指定图片的图片名
type TaobaoPictureUpdateResponse struct {
    
    /* done Basic更新是否成功 */
    done bool `json:"done";xml:"done"`
    
}
