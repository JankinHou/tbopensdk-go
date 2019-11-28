package tbopensdk_go

import "fmt"

func main() {
	//获取淘宝系统当前时间
	method := "taobao.time.get"
	result, err := RequestApi(method, ApiParams{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(result))
}
