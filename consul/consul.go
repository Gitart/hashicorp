package main

import (
	 // "fmt"
	"strings"
	"net/http"
	// "io/ioutil"
	"log"
	"strconv"
)

const (
	Url = "http://127.0.0.1:8500/v1/kv/"
)


func main() {

	for i:=0; i<100; i++{
         Sending("Сообщение 1",i)		
         
         log.Println("------------------------",i, "\n")	

	} 
}


// Отправка сообщения
func Sending(T string, i int) {

   s:=strconv.Itoa(i)

	payload := strings.NewReader(T)
	req, _ := http.NewRequest("PUT", Url+s, payload)

	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("Postman-Token", "b24cee4f-ee66-4f62-aeb2-3da8bd525677")
	req.Header.Add("X-Consul-Token", "9cd56757-ed5f-9124-fb48-e19329f9d0c5")


	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	
	// Для получения ответа
	// body, _ := ioutil.ReadAll(res.Body)

	// fmt.Println(res)
	// fmt.Println(string(body))

}    
