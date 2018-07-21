package main

import (
	"net/http"
	"fmt"
)

func main() {
	//resp,err := http.Get("http://192.168.20.152:7071/report/list")
	//if err!=nil{
	//	fmt.Println(err)
	//}
	//defer resp.Body.Close()
	//io.Copy(os.Stderr,resp.Body)
	//fmt.Println(resp.Body)
	//fmt.Println(http.Head("http://192.168.20.152:7071/report/list"))
	//resp, err := http.PostForm("http://192.168.20.152:7071/report/select", url.Values{"data":
	//{"1"}, "time": {"1"}})
	//if err!=nil{
	//	fmt.Println(err.Error())
	//}
	//fmt.Println(resp.Body)
	req,err := http.NewRequest("GET","http://192.168.20.152:7071/report/list",nil)
	if err!=nil{
		fmt.Println(err)
	}
	req.Header.Add("User-Agent","Gobook Custom User-Agent")
	client := &http.Client{}
	resp,err := client.Do(req)
	fmt.Println(resp)
}
