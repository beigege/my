package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	_ "net/http/pprof"
)
//new  服务端解析POST参数,json
func server_post_json() {
	http.HandleFunc("/json", func(rw http.ResponseWriter, rq *http.Request) {
		defer rq.Body.Close()
		//读取数据
		arr, err := ioutil.ReadAll(rq.Body)
		if err != nil {
			panic(err)
		}
		str := string(arr)
		//获取表单字段
		fmt.Printf("%v %v %v %v\n", rq.RemoteAddr, rq.Method, rq.Proto, str)
		//返回响应
		msg := fmt.Sprintf(`{"code":%v, "message":%v}`, 1, "success")
		rw.Write([]byte(msg))
	})
	http.ListenAndServe(":3004", nil)

}
//new  服务端解析POST参数
func server_post_application() {
	http.HandleFunc("/application", func(rw http.ResponseWriter, rq *http.Request) {
		defer rq.Body.Close()
		//解析表单
		rq.ParseForm()
		//获取表单字段
		id := rq.PostForm.Get("id")
		pid := rq.PostForm.Get("pid")
		fmt.Printf("%v %v %v id=%v pid=%v\n", rq.RemoteAddr, rq.Method, rq.Proto, id, pid)
		//返回响应
		msg := fmt.Sprintf("%v %v %v", rq.Host, rq.Method, rq.Proto)
		rw.Write([]byte(msg))
	})
	http.ListenAndServe(":3003", nil)

}

//new  带参数的GET请求
func server_get_hasparam() {
	http.HandleFunc("/hasparam", func(rw http.ResponseWriter, rq *http.Request) {
		//延迟关闭请求包体
		defer rq.Body.Close()
		fmt.Printf("client %v %v %v\n", rq.RemoteAddr, rq.Method, rq.URL)
		//获取GET请求参数
		val := rq.URL.Query()
		id := val.Get("id")
		pid := val.Get("pid")
		//返回响应
		msg := fmt.Sprintf("%v: id=%v, pid=%v", rq.Host, id, pid)
		rw.Write([]byte(msg))
	})
	http.ListenAndServe(":3002", nil)
}
//new  不带参数的GET请求
func server_get_noparam() {
	http.HandleFunc("/noparam", func(rw http.ResponseWriter, rq *http.Request) {
		fmt.Printf("client %v %v %v\n", rq.RemoteAddr, rq.Method, rq.URL)
		rw.Write([]byte(rq.RemoteAddr))

		fmt.Printf("protocol: %v\n", rq.Proto)
		fmt.Printf("method: %v\n", rq.Method)
		fmt.Printf("content length: %v\n", rq.ContentLength)
		fmt.Printf("url: %v\n", rq.URL)
		fmt.Printf("uri: %v\n", rq.RequestURI)
		fmt.Printf("remoteAddr: %v\n", rq.RemoteAddr)
		fmt.Printf("host: %v\n", rq.Host)
	})
	http.ListenAndServe(":3001", nil)
}

func main() {
	fmt.Printf("requset1    #################")
	server_get_noparam()
	//fmt.Printf("requset2    #################")
	//server_get_hasparam()
	//fmt.Printf("requset3    #################")
	//server_post_application()
	//fmt.Printf("requset4    #################")
	//server_post_json()
}

//old
//func main() {
//	flag.Set("v", "4")
//	glog.V(2).Info("Starting http server...")
//
//
//	http.HandleFunc("/", rootHandler)
//	http.HandleFunc("/healthz", healthz)
//
//	err := http.ListenAndServe(":80", nil)
//
//	if err != nil {
//		log.Fatal(err)
//	}
//
//}
//
//func healthz(w http.ResponseWriter, r *http.Request){
//	fmt.Println("entering /healthz handler")
//	io.WriteString(w, "200\n")
//}
//
//func rootHandler(w http.ResponseWriter, r *http.Request) {
//	fmt.Println("entering root handler")
//	user := r.URL.Query().Get("user")
//	if user != "" {
//		io.WriteString(w, fmt.Sprintf("hello [%s]\n", user))
//
//	} else {
//		io.WriteString(w, "hello [stranger]\n")
//
//	}
//
//	io.WriteString(w, "===================Details of the http request header:============\n")
//	for k, v := range r.Header {
//		//io.WriteString(w, fmt.Sprintf("%s=%s\n", k, v))
//		w.Write([]byte(fmt.Sprintf("%s=%s\n", k, v)))
//
//	}
//	w.WriteHeader(200)
//	w.Write([]byte(r.RemoteAddr+"\n"))
//	version := os.Getenv("VERSION")
//	w.Write([]byte(fmt.Sprintf("%s=%s\n", "version", version)))
//
//	glog.V(2).Info("client ip:"+r.Host)
//
//
//}
