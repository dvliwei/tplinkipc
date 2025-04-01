/**
 * @Title
 * @Author: liwei
 * @Description:  TODO
 * @File:  tplink_fac
 * @Version: 1.0.0
 * @Date: 2025/04/01 17:05
 * @Update liwei 2025/4/1 17:05
 */

package tplinkipc

type TplinkFactory interface {
	MakeTplink() IsTplink
}

type IsTplink interface {
	AuthLogin() (string, error)

	ContinuousMove(ptzParam *PtzRequest) error

	ContinuousStop(ptzParam *PtzRequest) error
}
