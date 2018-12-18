package tbsdk

// TaobaoFenxiaoDealerRequisitionorderCreateRequest 创建经销采购申请
type TaobaoFenxiaoDealerRequisitionorderCreateRequest struct {
    
    /* address required收货人所在街道地址 */
    address string `json:"address";xml:"address"`
    
    /* buyer_name required买家姓名（自提方式填写提货人姓名） */
    buyer_name string `json:"buyer_name";xml:"buyer_name"`
    
    /* city optional收货人所在市 */
    city string `json:"city";xml:"city"`
    
    /* district optional收货人所在区 */
    district string `json:"district";xml:"district"`
    
    /* id_card_number optional身份证号（自提方式必填，填写提货人身份证号码，提货时用于确认提货人身份） */
    id_card_number string `json:"id_card_number";xml:"id_card_number"`
    
    /* logistics_type optional配送方式。SELF_PICKUP：自提；LOGISTICS：仓库发货 */
    logistics_type string `json:"logistics_type";xml:"logistics_type"`
    
    /* mobile optional买家的手机号码（1、此字段与phone字段至少填写一个。2、自提方式下此字段必填，保存提货人联系电话） */
    mobile string `json:"mobile";xml:"mobile"`
    
    /* order_detail required采购清单，存放多个采购明细，每个采购明细内部以‘:’隔开，多个采购明细之间以‘,’隔开. 例(分销产品id:skuid:购买数量:申请单价,分销产品id:skuid:购买数量:申请单价)，申请单价的单位为分。不存在sku请留空skuid，如（分销产品id::购买数量:申请单价） */
    order_detail string `json:"order_detail";xml:"order_detail"`
    
    /* phone optional买家联系电话（此字段和mobile字段至少填写一个） */
    phone string `json:"phone";xml:"phone"`
    
    /* post_code optional收货人所在地区邮政编码 */
    post_code string `json:"post_code";xml:"post_code"`
    
    /* province required收货人所在省份 */
    province string `json:"province";xml:"province"`
    
}

func (req *TaobaoFenxiaoDealerRequisitionorderCreateRequest) GetAPIName() string {
	return "taobao.fenxiao.dealer.requisitionorder.create"
}

// TaobaoFenxiaoDealerRequisitionorderCreateResponse 创建经销采购申请
type TaobaoFenxiaoDealerRequisitionorderCreateResponse struct {
    
    /* dealer_order_id Basic经销采购申请编号 */
    dealer_order_id int64 `json:"dealer_order_id";xml:"dealer_order_id"`
    
}
