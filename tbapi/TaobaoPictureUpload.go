package tbsdk

// TaobaoPictureUploadRequest 图片空间上传接口
type TaobaoPictureUploadRequest struct {
    
    /* client_type optional图片上传的来源，有电脑版本宝贝发布，手机版本宝贝发布client:computer电脑版本宝贝使用，client:phone手机版本宝贝使用。注意：当client:phone时，图片限制为宽度在480-620之间，长度不能超过960，否则会报错。 */
    client_type string `json:"client_type";xml:"client_type"`
    
    /* image_input_title required包括后缀名的图片标题,不能为空，如Bule.jpg,有些卖家希望图片上传后取图片文件的默认名 */
    image_input_title string `json:"image_input_title";xml:"image_input_title"`
    
    /* img required图片二进制文件流,不能为空,允许png、jpg、gif图片格式,3M以内。 */
    img byte[] `json:"img";xml:"img"`
    
    /* is_https optional是否获取https连接 */
    is_https bool `json:"is_https";xml:"is_https"`
    
    /* picture_category_id required图片分类ID，设置具体某个分类ID或设置0上传到默认分类，只能传入一个分类 */
    picture_category_id int64 `json:"picture_category_id";xml:"picture_category_id"`
    
    /* picture_id optional如果此参数大于0，而且在后台能查到对应的图片，则此次上传为原图替换 */
    picture_id int64 `json:"picture_id";xml:"picture_id"`
    
    /* title optional图片标题,如果为空,传的图片标题就取去掉后缀名的image_input_title,超过50字符长度会截取50字符,重名会在标题末尾加"(1)";标题末尾已经有"(数字)"了，则数字加1 */
    title string `json:"title";xml:"title"`
    
}

func (req *TaobaoPictureUploadRequest) GetAPIName() string {
	return "taobao.picture.upload"
}

// TaobaoPictureUploadResponse 图片空间上传接口
type TaobaoPictureUploadResponse struct {
    
    /* picture Object当前上传的一张图片信息 */
    picture Picture `json:"picture";xml:"picture"`
    
}
