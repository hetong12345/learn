package old

import (
	"io/ioutil"
	"net/http"
	"os"
)

type appHandler func(w http.ResponseWriter, r *http.Request) error

func errApper(handler appHandler) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := handler(w, r)
		if err != nil {

			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsTimeout(err):
				code = http.StatusRequestTimeout
			default:
				code = http.StatusInternalServerError
			}
			http.Error(w, http.StatusText(code), code)
		}
	}
}
func main() {
	http.HandleFunc("/list/", errApper(HandleList))
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		panic(err)
	}
}
func HandleList(writer http.ResponseWriter, request *http.Request) error {

	path := request.URL.Path[len("/list/"):]
	file, err := os.Open(path)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return err
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	writer.Write(all)

	return nil
	//http.HandleFunc("/list/", func(writer http.ResponseWriter,request *http.Request) {
	//	path:=request.URL.Path[len("/list/"):]
	//	file,err:=os.Open(path)
	//	if err!=nil {
	//		http.Error(writer,err.Error(),http.StatusInternalServerError)
	//		return
	//	}
	//	defer file.Close()
	//
	//	all,err:=ioutil.ReadAll(file)
	//	if err!=nil {
	//		panic(err)
	//	}
	//
	//	writer.Write(all)
	//})
	//
	//http.ListenAndServe(":9999",nil)
}

//func firstPage(w http.ResponseWriter, r *http.Request) {
//	io.WriteString(w, "<h1>Hello World</h1>")
//}
//func reLaunch() {
//
//}
