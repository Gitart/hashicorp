package main

import (
	 // "fmt"
	"strings"
	"net/http"
	"io/ioutil"
	"log"
	"strconv"
	"fmt"
	"encoding/base64"
	"encoding/json"
)

const (
	Url = "http://127.0.0.1:8500/v1/kv/"
)



func main() {

	x:=Getting("10")
	fmt.Println(x)
}


func main_2() {

	for i:=0; i<10000; i++{
         Sending("Сообщение yjdj 1",i)		
         
         log.Println("------------------------",i, "\n")	

	} 
}


// ****************************************
// Отправка сообщения
// ****************************************
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

// type Mst map[string]interface{}
// ****************************************
// Structure for key
// ****************************************
type Mst struct{
	    LockIndex     int64
        Key           string
        Flags         int64
        Value         string
        CreateIndex   int64
        ModifyIndex   int64
}


// ****************************************
// Получение по ключу значение
// ****************************************
func Getting(Key string) string {


	url := Url+Key
	req, _ := http.NewRequest("GET", url, nil)


	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var m []Mst
	json.Unmarshal([]byte(body), &m)

    // sr:=m[0]["Value"].(string)
    sr:=m[0].Value
	
    // Decoding
	decoded, err := base64.StdEncoding.DecodeString(sr)
	if err != nil {
		fmt.Println("decode error:", err.Error())
		return "Error"
	}
	
	// fmt.Println(string(decoded))
    // r:=string(decoded)
    // r:=string(body)
    r:=string(decoded)
	return r
}


// ************************************************
// Decoding is a base64 encoded byte array
// ************************************************
func DecodeValue(kv string) (string) {
	// kv.Value is a base64 encoded byte array
	// encodedBase64 := []byte(kv)
	// s := base64.StdEncoding.EncodeToString(encodedBase64)

	// decode string to bytes
	b, err := base64.StdEncoding.DecodeString(kv)
	if err != nil {
		return ""
	}

	return string(b)
}
