/**
 * @Title
 * @Author: liwei
 * @Description:  TODO
 * @File:  tplink_res
 * @Version: 1.0.0
 * @Date: 2025/04/01 17:05
 * @Update liwei 2025/4/1 17:05
 */

package tplinkipc

import (
	"bytes"
	"crypto/md5"
	"crypto/tls"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
)

type tplinkRes struct {
	Stok     string
	Ip       string
	UserName string
	Password string
}

func newTplinkRes(tp *TplinkIpc) *tplinkRes {
	return &tplinkRes{Stok: tp.Stok, Ip: tp.Ip, UserName: tp.UserName, Password: tp.Password}
}

func (res *tplinkRes) AuthLogin() (string, error) {
	postUrl := fmt.Sprintf("http://%s/", res.Ip)
	data := map[string]interface{}{
		"method": "do",
		"login": map[string]interface{}{
			"username":         res.UserName,
			"password":         res.Password,
			"encrypt_type":     "2",
			"md5_encrypt_type": "1",
		},
	}
	body, err := res.CurlPost(postUrl, data)
	if err != nil {
		return "", err
	}
	var fResult FLoginResponse
	err = json.Unmarshal(body, &fResult)
	if err != nil {
		return "", err
	}
	password := fmt.Sprintf("%s:%s", res.Password, fResult.Data.Nonce)
	hash := md5.Sum([]byte(password))
	// 将哈希结果转换为十六进制字符串
	hashString := hex.EncodeToString(hash[:])
	data = map[string]interface{}{
		"method": "do",
		"login": map[string]interface{}{
			"username":         res.UserName,
			"password":         hashString,
			"encrypt_type":     "2",
			"md5_encrypt_type": "1",
		},
	}
	body, err = res.CurlPost(postUrl, data)
	if err != nil {
		return "", err
	}
	var lResult LoginResponse
	if err = json.Unmarshal(body, &lResult); err != nil {
		return "", err
	}
	if lResult.ErrorCode != 0 {
		return "", fmt.Errorf("tplink auth login failed")
	}
	return lResult.Stok, nil
}

func (res *tplinkRes) ContinuousMove(ptzParam *PtzRequest) error {
	postUrl := fmt.Sprintf("http://%s/stok=%s/ds", res.Ip, res.Stok)
	// 请求数据
	var data map[string]interface{}
	if ptzParam.VelocityTilt != nil { //上下转
		data = map[string]interface{}{
			"ptz": map[string]interface{}{
				"continuous_move": map[string]string{
					"velocity_tilt": *ptzParam.VelocityTilt,
					"channel":       ptzParam.Channel,
				},
			},
			"method": "do",
		}
	}
	if ptzParam.VelocityPan != nil { //左右转
		data = map[string]interface{}{
			"ptz": map[string]interface{}{
				"continuous_move": map[string]string{
					"velocity_pan": *ptzParam.VelocityPan,
					"channel":      ptzParam.Channel,
				},
			},
			"method": "do",
		}
	}
	body, err := res.CurlPost(postUrl, data)
	if err != nil {
		return err
	}
	var result MoveResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return err
	}
	if result.ErrorCode != 0 {
		return fmt.Errorf("tplink move failed")
	}
	return nil
}

func (res *tplinkRes) CurlPost(postUrl string, params interface{}) ([]byte, error) {
	// 请求头
	headers := map[string]string{
		"Accept":           "application/json, text/javascript, */*; q=0.01",
		"Accept-Language":  "zh-CN,zh;q=0.9,en;q=0.8,ar;q=0.7,ja;q=0.6,zh-TW;q=0.5",
		"Connection":       "keep-alive",
		"Content-Type":     "application/json; charset=UTF-8",
		"Origin":           fmt.Sprintf("http://%s", res.Ip),
		"Referer":          fmt.Sprintf("http://%s/", res.Ip),
		"User-Agent":       "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/134.0.0.0 Safari/537.36",
		"X-Requested-With": "XMLHttpRequest",
	}
	// 将数据转换为 JSON 字节切片
	jsonData, err := json.Marshal(params)
	if err != nil {
		fmt.Printf("JSON 编码出错: %v\n", err)
		return nil, err
	}

	// 创建 HTTP 请求
	req, err := http.NewRequest("POST", postUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("创建请求出错: %v\n", err)
		return nil, err
	}
	// 设置请求头
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// 创建一个新的 cookie jar
	jar, err := cookiejar.New(nil)
	if err != nil {
		fmt.Printf("创建 cookie jar 出错: %v\n", err)
		return nil, err
	}

	// 创建 HTTP 客户端
	client := &http.Client{
		Jar: jar,
		// 跳过 SSL 验证
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("发送请求出错: %v\n", err)
		return nil, err
	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("读取响应内容出错: %v\n", err)
		return nil, err
	}

	// 打印响应状态码和内容
	fmt.Printf("响应状态码: %d\n", resp.StatusCode)
	fmt.Printf("响应内容: %s\n", string(body))
	return body, nil
}

func (res *tplinkRes) ContinuousStop(ptzParam *PtzRequest) error {
	postUrl := fmt.Sprintf("http://%s/stok=%s/ds", res.Ip, res.Stok)
	// 请求数据
	data := map[string]interface{}{
		"ptz": map[string]interface{}{
			"stop": map[string]string{
				"pan":     "1",
				"tilt":    "1",
				"zoom":    "1",
				"channel": ptzParam.Channel,
			},
		},
		"method": "do",
	}
	body, err := res.CurlPost(postUrl, data)
	if err != nil {
		return err
	}
	var result MoveResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return err
	}
	if result.ErrorCode != 0 {
		return fmt.Errorf("tplink move failed")
	}
	return nil
}
