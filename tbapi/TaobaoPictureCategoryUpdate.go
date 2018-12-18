package tbsdk

// TaobaoPictureCategoryUpdateRequest 更新图片分类的名字，或者更新图片分类的父分类（即分类移动）。只能移动2级分类到非2级分类，默认分类和1级分类不可移动。
type TaobaoPictureCategoryUpdateRequest struct {
    
    /* category_id required要更新的图片分类的id */
    category_id int64 `json:"category_id";xml:"category_id"`
    
    /* category_name optional图片分类的新名字，最大长度20字符，不能为空 */
    category_name string `json:"category_name";xml:"category_name"`
    
    /* parent_id optional图片分类的新父分类id */
    parent_id int64 `json:"parent_id";xml:"parent_id"`
    
}

func (req *TaobaoPictureCategoryUpdateRequest) GetAPIName() string {
	return "taobao.picture.category.update"
}

// TaobaoPictureCategoryUpdateResponse 更新图片分类的名字，或者更新图片分类的父分类（即分类移动）。只能移动2级分类到非2级分类，默认分类和1级分类不可移动。
type TaobaoPictureCategoryUpdateResponse struct {
    
    /* done Basic更新图片分类是否成功 */
    done bool `json:"done";xml:"done"`
    
}
