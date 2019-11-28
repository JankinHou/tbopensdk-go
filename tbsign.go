package tbopensdk_go

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

//系统公共参数
var (
	//应用的appkey
	AppKey string = ""
	//应用的AppSecret
	AppSecret string = ""
	//Router 环境请求地址
	ApiUrl string = ""
	// Timeout ...
	Timeout time.Duration
	Session string = ""
)

//业务参数
type ApiParams map[string]string

//定义公共参数
var CommParams = map[string]string{
	"timestamp":   GetTime(0),
	"app_key":     AppKey,
	"v":           "2.0",
	"format":      "json",
	"sign_method": "md5",
	"session":     Session,
}

//error 结构体
type ErrRespone struct {
	ErrorResponse struct {
		SubMsg  string `json:"sub_msg"`
		Code    int    `json:"code"`
		SubCode string `json:"sub_code"`
		Msg     string `json:"msg"`
	} `json:"error_response"`
}

//发送请求
func RequestApi(method string, param ApiParams) (res []byte, err error) {
	param["method"] = method
	//合并两个参数
	for k, v := range param {
		CommParams[k] = v
	}
	CommParams["sign"] = sign(CommParams)
	urlParams := url.Values{}
	for k, v := range CommParams {
		urlParams.Set(k, v)
	}
	result := HttpPost(ApiUrl, urlParams.Encode())
	var respone = &ErrRespone{}
	json.Unmarshal(result, respone)
	if respone.ErrorResponse.Code != 0 {
		//出现错误
		err = errors.New(respone.ErrorResponse.SubMsg)
	} else {
		res = result

	}
	return
}

/**
淘宝签名
*/
func sign(params map[string]string) string {
	// 获取Key
	keystr := []string{}
	for k := range params {
		keystr = append(keystr, k)
	}
	//asc排序
	sort.Strings(keystr)
	// 把所有参数名和参数值串在一起
	query := AppSecret
	for _, k := range keystr {
		query += k + params[k]
	}
	//拼接appsecret
	query += AppSecret
	// 使用MD5加密
	signBytes := md5.Sum([]byte(query))
	// 把二进制转化为大写的十六进制
	return strings.ToUpper(hex.EncodeToString(signBytes[:]))
}

/**
发送post请求
传入URL和参数   参数是 name=a  格式
*/
func HttpPost(url string, params string) []byte {
	resp, _ := http.Post(url,
		"application/x-www-form-urlencoded",
		strings.NewReader(params))

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}

/**
获取当前时间
timeStamp 时间戳的差值,传入0是获取当前时间
*/
func GetTime(timeStamp int64) string {
	//当前时间的字符串，2006-01-02 15:04:05据说是golang的诞生时间，固定写法
	var ti int64 = time.Now().Unix() //已知的时间戳
	if timeStamp == 0 {
		//如果是0，则获取默认值
		return time.Unix(ti, 0).Format("2006-01-02 15:04:05")
	} else {
		ti = ti + timeStamp
		formatTimeStr := time.Unix(ti, 0).Format("2006-01-02 15:04:05")
		return formatTimeStr
	}
}
