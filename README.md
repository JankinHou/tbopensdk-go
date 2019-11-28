# tbopensdk-go
封装的淘宝Api的签名和请求
#淘宝Api sign算法
请移步[官方说明](https://open.taobao.com/doc.htm?docId=101617&docType=1")。
#Example
## 1 安装
```go
go get github.com/JankinHou/tbopensdk-go

```
##2 创建文件，如example.go
首先需要安装Go语言和设置GoPath，此处不多说明。以获取淘宝系统当前时间api为示例
```go
package main
 import (
 	"encoding/json"
 	"fmt"
 	"github.com/JankinHou/tbopensdk-go"
 )
 //返回参数结构体
 type TimeRespone struct {
 	TimeGetResponse struct {
 		Time      string `json:"time"`
 		RequestID string `json:"request_id"`
 	} `json:"time_get_response"`
 }
 func main() {
    //设置必须的参数
    tbopensdk_go.AppKey = ""
    tbopensdk_go.AppSecret = ""
    tbopensdk_go.ApiUrl = ""
    tbopensdk_go.Session = ""	
 	//以获取淘宝系统当前时间Api为例
 	method := "taobao.time.get"
 	//请求api
 	result, err := tbopensdk_go.RequestApi(method, tbopensdk_go.ApiParams{
        //"key":"value",
        //如果有参数，在此处按照如上格式添加即可
    })
 	if err != nil {
 		fmt.Println(err)
 	}
 	timeres := &TimeRespone{}
 	json.Unmarshal(result,timeres)
 	fmt.Println(timeres.TimeGetResponse.Time)
 }
```
## 3 运行
```go
go run example.go
```
#说明
#####由于官方没有Go语言SDK，为了自己方便实用而写。
#####参考过[taobaogo](https://github.com/nilorg/go-opentaobao)。
#####联系作者：houzhenkai@houzhenkai.com
#####[不去适用、努力改变](http://www.houzhenkai.com)
