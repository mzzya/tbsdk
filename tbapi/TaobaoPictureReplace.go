package tbsdk

// TaobaoPictureReplaceRequest 替换一张图片，只替换图片数据，图片名称，图片分类等保持不变。
type TaobaoPictureReplaceRequest struct {
    
    /* image_data required图片二进制文件流,不能为空,允许png、jpg、gif图片格式 */
    image_data byte[] `json:"image_data";xml:"image_data"`
    
    /* picture_id required要替换的图片的id，必须大于0 */
    picture_id int64 `json:"picture_id";xml:"picture_id"`
    
}

func (req *TaobaoPictureReplaceRequest) GetAPIName() string {
	return "taobao.picture.replace"
}

// TaobaoPictureReplaceResponse 替换一张图片，只替换图片数据，图片名称，图片分类等保持不变。
type TaobaoPictureReplaceResponse struct {
    
    /* done Basic图片替换是否成功 */
    done bool `json:"done";xml:"done"`
    
}
