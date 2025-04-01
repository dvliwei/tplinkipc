/**
 * @Title
 * @Author: liwei
 * @Description:  TODO
 * @File:  tplink_test
 * @Version: 1.0.0
 * @Date: 2025/04/01 17:22
 * @Update liwei 2025/4/1 17:22
 */

package tplinkipc

import (
	"fmt"
	"testing"
)

func TestAuthLogin(t *testing.T) {
	var config TplinkConfig
	config.UserName = "admin"
	config.Password = "qwert12345_"
	config.Ip = "192.168.1.100"
	res := NewTplinkIpc("", &config)
	stock, err := res.MakeTplink().AuthLogin()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(stock)
}

func TestContinuousMove(t *testing.T) {
	var config TplinkConfig
	config.UserName = "admin"
	config.Password = "qwert12345_"
	config.Ip = "192.168.1.100"
	res := NewTplinkIpc("6XLrmGNBrFgYkr3gF3ymrmlTVmlT4FEO", &config)
	var request PtzRequest
	velocityPan := "0.571429"
	request.VelocityPan = &velocityPan
	request.Channel = "2"
	err := res.MakeTplink().ContinuousMove(&request)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("OK")
}

func TestContinuousStop(t *testing.T) {
	var config TplinkConfig
	config.UserName = "admin"
	config.Password = "qwert12345_"
	config.Ip = "192.168.1.100"
	res := NewTplinkIpc("6XLrmGNBrFgYkr3gF3ymrmlTVmlT4FEO", &config)
	var request PtzRequest
	request.Channel = "2"
	err := res.MakeTplink().ContinuousStop(&request)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("OK")
}
