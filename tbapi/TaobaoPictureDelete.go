package tbsdk

// TaobaoPictureDeleteRequest 删除图片空间图片
type TaobaoPictureDeleteRequest struct {
    
    /* picture_ids required图片ID字符串,可以一个也可以一组,用英文逗号间隔,如450,120,155.限制数量是100 */
    picture_ids string `json:"picture_ids";xml:"picture_ids"`
    
}

func (req *TaobaoPictureDeleteRequest) GetAPIName() string {
	return "taobao.picture.delete"
}

// TaobaoPictureDeleteResponse 删除图片空间图片
type TaobaoPictureDeleteResponse struct {
    
    /* success Basic是否删除 */
    success string `json:"success";xml:"success"`
    
}
