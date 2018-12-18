package tbsdk

// TaobaoShopGetbytitleRequest 根据店铺名称获取店铺信息
type TaobaoShopGetbytitleRequest struct {
    
    /* fields optional代表需要获取的店铺信息：sid,cid,title,nick,desc,bulletin,pic_path,created,modified,shop_score */
    fields string `json:"fields";xml:"fields"`
    
    /* title required店铺名称，必须严格匹配（不支持模糊查询） */
    title string `json:"title";xml:"title"`
    
}

func (req *TaobaoShopGetbytitleRequest) GetAPIName() string {
	return "taobao.shop.getbytitle"
}

// TaobaoShopGetbytitleResponse 根据店铺名称获取店铺信息
type TaobaoShopGetbytitleResponse struct {
    
    /* err_code Basic错误码 */
    err_code string `json:"err_code";xml:"err_code"`
    
    /* err_msg Basic错误信息 */
    err_msg string `json:"err_msg";xml:"err_msg"`
    
    /* is_error Basic有无错误 */
    is_error bool `json:"is_error";xml:"is_error"`
    
    /* result_shop Basictest */
    result_shop map[string]interface{} `json:"result_shop";xml:"result_shop"`
    
}
