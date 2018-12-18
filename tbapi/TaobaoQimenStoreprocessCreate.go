package tbsdk

// TaobaoQimenStoreprocessCreateRequest ERP调用奇门的接口,创建仓内加工单
type TaobaoQimenStoreprocessCreateRequest struct {
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* materialitems optional加工商品列表 */
    materialitems MaterialItem `json:"materialitems";xml:"materialitems"`
    
    /* orderCreateTime required加工单创建时间(YYYY-MM-DD HH:MM:SS) */
    orderCreateTime string `json:"orderCreateTime";xml:"orderCreateTime"`
    
    /* orderType required单据类型(CNJG=仓内加工作业单) */
    orderType string `json:"orderType";xml:"orderType"`
    
    /* planQty optional成品计划数量 */
    planQty int64 `json:"planQty";xml:"planQty"`
    
    /* planTime required计划加工时间(YYYY-MM-DD HH:MM:SS) */
    planTime string `json:"planTime";xml:"planTime"`
    
    /* processOrderCode required加工单编码 */
    processOrderCode string `json:"processOrderCode";xml:"processOrderCode"`
    
    /* productitems optional商品列表 */
    productitems ProductItem `json:"productitems";xml:"productitems"`
    
    /* remark optional备注 */
    remark string `json:"remark";xml:"remark"`
    
    /* serviceType required加工类型(1:仓内组合加工 2:仓内组合拆分) */
    serviceType string `json:"serviceType";xml:"serviceType"`
    
    /* warehouseCode required仓库编码(统仓统配等无需ERP指定仓储编码的情况填OTHER) */
    warehouseCode string `json:"warehouseCode";xml:"warehouseCode"`
    
}

func (req *TaobaoQimenStoreprocessCreateRequest) GetAPIName() string {
	return "taobao.qimen.storeprocess.create"
}

// TaobaoQimenStoreprocessCreateResponse ERP调用奇门的接口,创建仓内加工单
type TaobaoQimenStoreprocessCreateResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
    /* processOrderId Basic仓储系统处理单ID */
    processOrderId string `json:"processOrderId";xml:"processOrderId"`
    
}
