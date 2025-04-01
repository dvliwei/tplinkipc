/**
 * @Title
 * @Author: liwei
 * @Description:  TODO
 * @File:  types
 * @Version: 1.0.0
 * @Date: 2025/04/01 17:08
 * @Update liwei 2025/4/1 17:08
 */

package tplinkipc

type (
	PtzRequest struct {
		VelocityTilt *string `json:"velocity_tilt"` //上下
		VelocityPan  *string `json:"velocity_pan"`  //左右
		Channel      string  `json:"channel"`       //监控点
	}

	TplinkConfig struct {
		Ip       string
		UserName string
		Password string
	}
)

// 定义外层结构体
type FLoginResponse struct {
	Data      Data `json:"data"`
	ErrorCode int  `json:"error_code"`
}

// 定义 data 字段对应的结构体
type Data struct {
	Code           int      `json:"code"`
	Time           int      `json:"time"`
	MaxTime        int      `json:"max_time"`
	EncryptType    []string `json:"encrypt_type"`
	Key            string   `json:"key"`
	Nonce          string   `json:"nonce"`
	MD5EncryptType int      `json:"md5_encrypt_type"`
}

// Response 定义与 JSON 数据对应的结构体
type LoginResponse struct {
	UserGroup string `json:"user_group"`
	Stok      string `json:"stok"`
	ErrorCode int    `json:"error_code"`
}

type MoveResponse struct {
	ErrorCode int `json:"error_code"`
}
