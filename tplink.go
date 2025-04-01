/**
 * @Title
 * @Author: liwei
 * @Description:  TODO
 * @File:  tplink
 * @Version: 1.0.0
 * @Date: 2025/04/01 17:04
 * @Update liwei 2025/4/1 17:04
 */

package tplinkipc

type TplinkIpc struct {
	Stok     string
	Ip       string
	UserName string
	Password string
}

func NewTplinkIpc(stok string, config *TplinkConfig) *TplinkIpc {
	return &TplinkIpc{Stok: stok, Ip: config.Ip, UserName: config.UserName, Password: config.Password}
}

func (tp *TplinkIpc) MakeTplink() IsTplink {
	return newTplinkRes(tp)
}
