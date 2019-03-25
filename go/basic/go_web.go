package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	for k, v := range r.Form {
		fmt.Println("key:", k, " val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello...")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method: ", r.Method)
	if r.Method == "GET" {
		// 生成随机token
		h := md5.New()
		crutime := time.Now().Unix()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		io.WriteString(h, "ajggeruag")
		token := fmt.Sprintf("%x", h.Sum(nil))
		t, _ := template.ParseFiles("login.gtpl")
		log.Println(t.Execute(w, token))
	} else {
		// 需要显式调用request.ParseForm()解析表单数据
		r.ParseForm()
		token := r.Form.Get("token")
		if token != "" {
			fmt.Println("username:", r.Form["username"])
			fmt.Println("password:", r.Form["password"])
		} else {
			fmt.Println("token is nil")
		}
	}
}

func main() {
	http.HandleFunc("/", sayhelloName)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe("localhost:9999", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
