package tbsdk

// TaobaoOpenuidGetBymixnickRequest 通过mixnick转换openuid
type TaobaoOpenuidGetBymixnickRequest struct {
    
    /* mix_nick required无线类应用获取到的混淆的nick */
    mix_nick string `json:"mix_nick";xml:"mix_nick"`
    
}

func (req *TaobaoOpenuidGetBymixnickRequest) GetAPIName() string {
	return "taobao.openuid.get.bymixnick"
}

// TaobaoOpenuidGetBymixnickResponse 通过mixnick转换openuid
type TaobaoOpenuidGetBymixnickResponse struct {
    
    /* open_uid BasicOpenUID */
    open_uid string `json:"open_uid";xml:"open_uid"`
    
}
