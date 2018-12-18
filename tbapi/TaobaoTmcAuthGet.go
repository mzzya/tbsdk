package tbsdk

// TaobaoTmcAuthGetRequest TMC连接授权Token
type TaobaoTmcAuthGetRequest struct {
    
    /* group optionaltmc组名 */
    group string `json:"group";xml:"group"`
    
}

func (req *TaobaoTmcAuthGetRequest) GetAPIName() string {
	return "taobao.tmc.auth.get"
}

// TaobaoTmcAuthGetResponse TMC连接授权Token
type TaobaoTmcAuthGetResponse struct {
    
    /* result Basicresult */
    result string `json:"result";xml:"result"`
    
}
