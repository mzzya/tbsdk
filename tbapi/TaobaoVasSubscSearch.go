package tbsdk

// TaobaoVasSubscSearchRequest 用于ISV查询自己名下的应用及收费项目的订购记录
type TaobaoVasSubscSearchRequest struct {
    
    /* article_code required应用收费代码，从合作伙伴后台（my.open.taobao.com）-收费管理-收费项目列表 能够获得该应用的收费代码 */
    article_code string `json:"article_code";xml:"article_code"`
    
    /* autosub optional是否自动续费，true=自动续费 false=非自动续费 空=全部 */
    autosub bool `json:"autosub";xml:"autosub"`
    
    /* end_deadline optional到期时间结束值 */
    end_deadline Date `json:"end_deadline";xml:"end_deadline"`
    
    /* expire_notice optional是否到期提醒，true=到期提醒 false=非到期提醒 空=全部 */
    expire_notice bool `json:"expire_notice";xml:"expire_notice"`
    
    /* item_code optional收费项目代码，从合作伙伴后台（my.open.taobao.com）-收费管理-收费项目列表 能够获得收费项目代码 */
    item_code string `json:"item_code";xml:"item_code"`
    
    /* nick optional淘宝会员名 */
    nick string `json:"nick";xml:"nick"`
    
    /* page_no optional页码 */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size optional一页包含的记录数 */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* start_deadline optional到期时间起始值（当start_deadline和end_deadline都不填写时，默认返回最近90天的数据） */
    start_deadline Date `json:"start_deadline";xml:"start_deadline"`
    
    /* status optional订购记录状态，1=有效 2=过期 空=全部 */
    status int64 `json:"status";xml:"status"`
    
}

func (req *TaobaoVasSubscSearchRequest) GetAPIName() string {
	return "taobao.vas.subsc.search"
}

// TaobaoVasSubscSearchResponse 用于ISV查询自己名下的应用及收费项目的订购记录
type TaobaoVasSubscSearchResponse struct {
    
    /* article_subs Object Array订购关系对象 */
    article_subs ArticleSub `json:"article_subs";xml:"article_subs"`
    
    /* total_item Basic总记录数 */
    total_item int64 `json:"total_item";xml:"total_item"`
    
}
