package tbsdk

// TaobaoMediaCategoryUpdateRequest 更新媒体文件分类
type TaobaoMediaCategoryUpdateRequest struct {
    
    /* media_category_id required文件分类ID,不能为空 */
    media_category_id int64 `json:"media_category_id";xml:"media_category_id"`
    
    /* media_category_name required文件分类名，最大长度20字符，中英文都算一字符,不能为空 */
    media_category_name string `json:"media_category_name";xml:"media_category_name"`
    
}

func (req *TaobaoMediaCategoryUpdateRequest) GetAPIName() string {
	return "taobao.media.category.update"
}

// TaobaoMediaCategoryUpdateResponse 更新媒体文件分类
type TaobaoMediaCategoryUpdateResponse struct {
    
    /* success Basic更新是否成功标志 */
    success bool `json:"success";xml:"success"`
    
}
