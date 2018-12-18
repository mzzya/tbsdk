package tbsdk

// TaobaoPictureIsreferencedGetRequest 查询图片是否被引用，被引用返回true，未被引用返回false
type TaobaoPictureIsreferencedGetRequest struct {
    
    /* picture_id required图片id */
    picture_id int64 `json:"picture_id";xml:"picture_id"`
    
}

func (req *TaobaoPictureIsreferencedGetRequest) GetAPIName() string {
	return "taobao.picture.isreferenced.get"
}

// TaobaoPictureIsreferencedGetResponse 查询图片是否被引用，被引用返回true，未被引用返回false
type TaobaoPictureIsreferencedGetResponse struct {
    
    /* is_referenced Basic图片是否被引用 */
    is_referenced bool `json:"is_referenced";xml:"is_referenced"`
    
}
