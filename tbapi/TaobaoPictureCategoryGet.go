package tbsdk

// TaobaoPictureCategoryGetRequest 获取图片分类信息
type TaobaoPictureCategoryGetRequest struct {
    
    /* modified_time optional图片分类的修改时间点，格式:yyyy-MM-dd HH:mm:ss。查询此修改时间点之后到目前的图片分类。 */
    modified_time Date `json:"modified_time";xml:"modified_time"`
    
    /* parent_id optional取二级分类时设置为对应父分类id
取一级分类时父分类id设为0
取全部分类的时候不设或设为-1 */
    parent_id int64 `json:"parent_id";xml:"parent_id"`
    
    /* picture_category_id optional图片分类ID */
    picture_category_id int64 `json:"picture_category_id";xml:"picture_category_id"`
    
    /* picture_category_name optional图片分类名，不支持模糊查询 */
    picture_category_name string `json:"picture_category_name";xml:"picture_category_name"`
    
    /* _type optional1 */
    _type string `json:"type";xml:"type"`
    
}

func (req *TaobaoPictureCategoryGetRequest) GetAPIName() string {
	return "taobao.picture.category.get"
}

// TaobaoPictureCategoryGetResponse 获取图片分类信息
type TaobaoPictureCategoryGetResponse struct {
    
    /* picture_categories Object Array图片分类 */
    picture_categories PictureCategory `json:"picture_categories";xml:"picture_categories"`
    
}
