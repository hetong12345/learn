package router

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"html/template"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func StartHandle() error {
	http.HandleFunc("/", firstPage)
	http.HandleFunc("/router", handle)
	http.HandleFunc("/hello", sayHelloName)
	http.HandleFunc("/login", login)
	http.HandleFunc("/redis", doRedis)
	http.HandleFunc("/upload", upload)
	return nil
}

func firstPage(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>Hello World</h1>")
	conn := Pool.Get()
	defer conn.Close()
	reply, err := redis.Int64(conn.Do("incr", "view"))
	if err != nil {
		panic(err)
	}
	//reply, _ := redis.String(conn.Do("get", "view"))
	io.WriteString(w, "您是这个页面的第"+strconv.FormatInt(reply, 10)+"个访问者")
}
func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path: ", r.URL.Path)
	fmt.Println("scheme: ", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key: ", k)
		fmt.Println("val: ", strings.Join(v, " "))
	}
	fmt.Fprintf(w, "hello chain!")
}

func testPage(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>Test Page</h1>")
	conn := Pool.Get()
	defer conn.Close()
	conn.Do("", "")
}

func handle(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{Name: "pa", Value: "dps"})
	r.Header.Add("server", "go/1.0")
	io.WriteString(w, "myserver is running!")
	fmt.Printf("r.header=%+v,r.usr.string=%+v,r.url.path=%+v\n", r.Header, r.URL.String(), r.URL.Path)
	r.ParseForm()
	d := r.Form
	fmt.Println(d)
}

func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析form
	fmt.Println("method: ", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./web/templates/login.html")
		t.Execute(w, nil)
	} else if r.Method == "POST" {
		if len(r.Form["username"][0]) == 0 {
			io.WriteString(w, fmt.Sprintf("username: null or empty \n"))
		}
		age, err := strconv.Atoi(r.Form.Get("age"))
		if err != nil {
			io.WriteString(w, fmt.Sprintf("age: The format of the input is not correct \n"))
		}
		if age < 18 {
			io.WriteString(w, fmt.Sprintf("age: Minors are not registered \n"))
		}
		if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`,
			r.Form.Get("email")); !m {
			io.WriteString(w, fmt.Sprintf("email: The format of the input is not correct \n"))
		}
	}
}
func doRedis(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析form
	fmt.Println("method: ", r.Method)
	//判断传递方式
	if r.Method == "GET" {
		//加载界面模板
		t, err := template.ParseFiles("./web/templates/redis.html")
		if err != nil {
			panic(err)
		}
		//将解析好的模板应用到data上，这里data为nil
		t.Execute(w, nil)
	} else if r.Method == "POST" {
		conn := Pool.Get()
		defer conn.Close()
		arg := r.Form["arg"][0]
		split := strings.Split(arg, ".")
		reply, err := conn.Do(r.Form["com"][0], split[0], split[1])
		if err != nil {
			io.WriteString(w, err.Error())
		}
		io.WriteString(w, fmt.Sprint(reply))
	}
}
func upload(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == "GET" {
		now := time.Now().Unix()
		h := md5.New()
		h.Write([]byte(strconv.FormatInt(now, 10)))
		token := hex.EncodeToString(h.Sum(nil))
		t, _ := template.ParseFiles("./web/templates/upload.html")
		t.Execute(w, token)
	} else if r.Method == "POST" {
		//把上传的文件存储在内存和临时文件中
		r.ParseMultipartForm(32 << 20)
		//获取文件句柄，然后对文件进行存储等处理
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println("form file err: ", err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		//创建上传的目的文件
		f, err := os.OpenFile("./web/upload/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println("open file err: ", err)
			return
		}
		defer f.Close()
		//拷贝文件
		io.Copy(f, file)
	}
}
