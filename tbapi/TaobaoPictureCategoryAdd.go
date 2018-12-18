package tbsdk

// TaobaoPictureCategoryAddRequest 同一卖家最多添加500个图片分类，图片分类名称长度最大为20个字符
type TaobaoPictureCategoryAddRequest struct {
    
    /* parent_id optional图片分类的父分类,一级分类的parent_id为0,二级分类的则为其父分类的picture_category_id */
    parent_id int64 `json:"parent_id";xml:"parent_id"`
    
    /* picture_category_name required图片分类名称，最大长度20字符，中文字符算2个字符，不能为空 */
    picture_category_name string `json:"picture_category_name";xml:"picture_category_name"`
    
}

func (req *TaobaoPictureCategoryAddRequest) GetAPIName() string {
	return "taobao.picture.category.add"
}

// TaobaoPictureCategoryAddResponse 同一卖家最多添加500个图片分类，图片分类名称长度最大为20个字符
type TaobaoPictureCategoryAddResponse struct {
    
    /* picture_category Object图片分类信息 */
    picture_category PictureCategory `json:"picture_category";xml:"picture_category"`
    
}
