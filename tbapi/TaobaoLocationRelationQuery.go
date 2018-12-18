package tbsdk

// TaobaoLocationRelationQueryRequest 地点关联关系查询 
门店和仓库关联关系查询
type TaobaoLocationRelationQueryRequest struct {
    
    /* location_relation required关系查询 */
    location_relation LocationRelationDto `json:"location_relation";xml:"location_relation"`
    
}

func (req *TaobaoLocationRelationQueryRequest) GetAPIName() string {
	return "taobao.location.relation.query"
}

// TaobaoLocationRelationQueryResponse 地点关联关系查询 
门店和仓库关联关系查询
type TaobaoLocationRelationQueryResponse struct {
    
    /* result Objectresult */
    result SingleResult `json:"result";xml:"result"`
    
}
