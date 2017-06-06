package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Server struct {
	ServerName string
	ServerIP   string
}

type Serverslice struct {
	Servers []Server
}

type Response struct {
	data    string
	message string
	status  int8
	version int
}

func main() {
	var s Serverslice
	str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	fmt.Println(str)
	json.Unmarshal([]byte(str), &s)
	fmt.Println(s)

	for item, value := range s.Servers {
		fmt.Println(item, value)
	}

	//res,_ := http.Get("http://api.hopapapa.com/test")
	//fmt.Println(res)
	//fmt.Println(res.Body)
	//defer res.Body.Close()
	//body, err := ioutil.ReadAll(res.Body)
	//fmt.Println(body,err)

	//res, err := http.Get("http://api.hopapapa.com/api/v1/home_page/87")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//robots, err := ioutil.ReadAll(res.Body)
	////res.Body.Close()
	////if err != nil {
	////	log.Fatal(err)
	////}
	////fmt.Printf("%s", robots)
	////t := reflect.TypeOf(robots)
	////fmt.Println(t)
	//
	//var response Response
	//json.Unmarshal([]byte(robots), &response)
	//fmt.Println(response)

	url := "http://api.hopapapa.com/api/v1/home_page/87"

	res := new(Response)
	getJson(url, res)
	fmt.Println(res.message)
	fmt.Println(res.data)

}

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
