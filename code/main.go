package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"sort"
)

const mapUrl string = "https://apis.map.qq.com/ws/geocoder/v1"

var key string
var sk string

//MD5
func md5Crypt(str string) string {
	data := []byte(str)
	hash := md5.Sum(data)
	sign := fmt.Sprintf("%x", hash)
	log.Printf("sign: %s \n", sign)
	return sign
}

func init() {
	if os.Getenv("key") != "" {
		key = os.Getenv("key")
	}
	if os.Getenv("sk") != "" {
		sk = os.Getenv("sk")
	}
}

func apHandler(w http.ResponseWriter, req *http.Request) {
	log.Println(sk)
	log.Println(key)
	values := req.URL.Query()
	address := values.Get("address")
	location := values.Get("location")
	log.Printf("req-params:address= %s ,localtion=%s \n", address, location)

	params := make(map[string]string, 3)
	keys := make([]string, 0, 3)
	if address != "" {
		params["address"] = address
	}
	if location != "" {
		params["location"] = location
	}
	params["key"] = key
	params["get_poi"] = "1"
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	param := ""
	for i, k := range keys {
		if i == 0 {
			param += k + "=" + params[k]
		} else {
			param += "&" + k + "=" + params[k]
		}
	}
	log.Printf("sign-params: %s \n", param)

	res, err := http.Get(mapUrl + "?" + url.QueryEscape(param+"&sig="+md5Crypt("/ws/geocoder/v1?"+param+sk)))
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	b, _ := ioutil.ReadAll(res.Body)
	w.Write(b)
}

/*
	http://localhost:9000/map/qq/ws/geocoder/v1?address=北京市海淀区北四环西路66号
	http://localhost:9000/map/qq/ws/geocoder/v1?location=39.984154,116.307490
*/
func main() {
	http.HandleFunc("/map/qq/ws/geocoder/v1", apHandler)
	port := os.Getenv("FC_SERVER_PORT")
	if port == "" {
		port = "9000"
	}
	log.Println("Listening on localhost:9000")
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}
