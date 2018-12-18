package tbsdk

// TaobaoPicturePicturesCountRequest 图片总数查询
type TaobaoPicturePicturesCountRequest struct {
    
    /* client_type optional图片使用，如果是pc宝贝detail使用，设置为client:computer，查询出来的图片是符合pc的宝贝detail显示的如果是手机宝贝detail使用，设置为client:phone，查询出来的图片是符合手机的宝贝detail显示的,默认值是全部 */
    client_type string `json:"client_type";xml:"client_type"`
    
    /* deleted optional是否删除,undeleted代表没有删除,deleted表示删除 */
    deleted string `json:"deleted";xml:"deleted"`
    
    /* end_date optional查询上传结束时间点,格式:yyyy-MM-dd HH:mm:ss */
    end_date Date `json:"end_date";xml:"end_date"`
    
    /* picture_category_id optional图片分类 */
    picture_category_id int64 `json:"picture_category_id";xml:"picture_category_id"`
    
    /* picture_id optional图片ID */
    picture_id int64 `json:"picture_id";xml:"picture_id"`
    
    /* start_date optional查询上传开始时间点,格式:yyyy-MM-dd HH:mm:ss */
    start_date Date `json:"start_date";xml:"start_date"`
    
    /* start_modified_date optional图片被修改的时间点，格式:yyyy-MM-dd HH:mm:ss。查询此修改时间点之后到目前的图片。 */
    start_modified_date Date `json:"start_modified_date";xml:"start_modified_date"`
    
    /* title optional文件名 */
    title string `json:"title";xml:"title"`
    
}

func (req *TaobaoPicturePicturesCountRequest) GetAPIName() string {
	return "taobao.picture.pictures.count"
}

// TaobaoPicturePicturesCountResponse 图片总数查询
type TaobaoPicturePicturesCountResponse struct {
    
    /* totals Basic查询的文件总数 */
    totals int64 `json:"totals";xml:"totals"`
    
}
