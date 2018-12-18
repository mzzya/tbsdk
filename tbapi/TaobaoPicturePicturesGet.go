package tbsdk

// TaobaoPicturePicturesGetRequest 图片空间对外的图片获取接口，该接口只针对分页获取，获取某一页的图片，该接口不支持总数的查询asd
type TaobaoPicturePicturesGetRequest struct {
    
    /* client_type optional图片使用，如果是pc宝贝detail使用，设置为client:computer，查询出来的图片是符合pc的宝贝detail显示的如果是手机宝贝detail使用，设置为client:phone，查询出来的图片是符合手机的宝贝detail显示的,默认值是全部 */
    client_type string `json:"client_type";xml:"client_type"`
    
    /* current_page optional页码.传入值为1代表第一页,传入值为2代表第二页,依此类推,默认值为1 */
    current_page int64 `json:"current_page";xml:"current_page"`
    
    /* deleted optional是否删除,deleted代表没有删除 */
    deleted string `json:"deleted";xml:"deleted"`
    
    /* end_date optional结束时间 */
    end_date Date `json:"end_date";xml:"end_date"`
    
    /* is_https optional是否获取https的链接 */
    is_https bool `json:"is_https";xml:"is_https"`
    
    /* order_by optional图片查询结果排序,time:desc按上传时间从晚到早(默认), time:asc按上传时间从早到晚,sizes:desc按图片从大到小，sizes:asc按图片从小到大,默认time:desc */
    order_by string `json:"order_by";xml:"order_by"`
    
    /* page_size optional每页条数.每页返回最多返回100条,默认值40 */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* picture_category_id optional图片分类 */
    picture_category_id int64 `json:"picture_category_id";xml:"picture_category_id"`
    
    /* picture_id optional图片ID */
    picture_id int64 `json:"picture_id";xml:"picture_id"`
    
    /* start_date optional查询上传结束时间点,格式:yyyy-MM-dd HH:mm:ss */
    start_date Date `json:"start_date";xml:"start_date"`
    
    /* start_modified_date optional图片被修改的时间点，格式:yyyy-MM-dd HH:mm:ss。查询此修改时间点之后到目前的图片。 */
    start_modified_date Date `json:"start_modified_date";xml:"start_modified_date"`
    
    /* title optional图片标题,最大长度50字符,中英文都算一字符 */
    title string `json:"title";xml:"title"`
    
    /* urls optional图片url查询接口 */
    urls string `json:"urls";xml:"urls"`
    
}

func (req *TaobaoPicturePicturesGetRequest) GetAPIName() string {
	return "taobao.picture.pictures.get"
}

// TaobaoPicturePicturesGetResponse 图片空间对外的图片获取接口，该接口只针对分页获取，获取某一页的图片，该接口不支持总数的查询asd
type TaobaoPicturePicturesGetResponse struct {
    
    /* pictures Object Array图片空间图片数据对象 */
    pictures Picture `json:"pictures";xml:"pictures"`
    
}
