package tbsdk

// TaobaoQimenExpressinfoQueryRequest 配送公司信息查询
type TaobaoQimenExpressinfoQueryRequest struct {
    
    /* expressCode optional奇门仓储字段 */
    expressCode string `json:"expressCode";xml:"expressCode"`
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* ownerCode optional奇门仓储字段 */
    ownerCode string `json:"ownerCode";xml:"ownerCode"`
    
}

func (req *TaobaoQimenExpressinfoQueryRequest) GetAPIName() string {
	return "taobao.qimen.expressinfo.query"
}

// TaobaoQimenExpressinfoQueryResponse 配送公司信息查询
type TaobaoQimenExpressinfoQueryResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* expressInfos Object Array奇门仓储字段 */
    expressInfos ExpressInfo `json:"expressInfos";xml:"expressInfos"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
}
