package main

import (
	"fmt"
	"html/template"

	//"io/ioutil"
	"net/http"
	"time"
)

/*
type myMux struct{}

func (m *myMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		defaultRoute(w, r)
	case "/route1":
		route1(w, r)
	case "/route2":
		route2(w, r)
	case "/writeResponse":
		writeResponse(w, r)
	case "/login":
		login(w, r)
	default:
		http.NotFound(w, r)
	}
}
*/
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", defaultRoute)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/chat", chat)
	mux.HandleFunc("/actioned", actioned)
	//mux.Handle("/anothermethod", handlerA{})
	//mux.HandleFunc("/lookRequest", lookRequest)
	//mux.HandleFunc("/writeResponse", writeResponse)

	//mux := &myMux{}
	server := http.Server{
		Addr:           ":8080",
		Handler:        mux,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
func chat(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	t := template.Must(template.ParseFiles("doc/chat.html"))
	t.Execute(w, r.Form)
}
func actioned(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	t := template.Must(template.ParseFiles("doc/actioined.html"))
	t.Execute(w, r.Form)
}
func defaultRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "It's defaultroute")
}
func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// 入力フォームを返す
		t, _ := template.ParseFiles("doc/main.html")
		t.Execute(w, nil)
	}
	if r.Method == "POST" {
		// Requestを解析し入力情報を出力する。
		/*r.ParseForm()
		fmt.Println("name:", r.Form)
		http.Redirect(w, r, "/login", 301)
		*/
		r.ParseForm()
		t, _ := template.ParseFiles("doc/main.html")
		t.Execute(w, r.Form)
	}
}

/*
type handlerA struct{}
func (h handlerA) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "it's another method")
}
func lookRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "開始行\n")
	fmt.Fprintln(w, "Method:", r.Method)
	fmt.Fprintln(w, "URL:", r.URL)
	fmt.Fprintln(w, "Proto:", r.Proto)
	fmt.Fprintf(w, "\nHEADER\n")
	for k, v := range r.Header {
		fmt.Fprintln(w, k, v)
	}
	fmt.Fprintf(w, "\nBODY\n")
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	fmt.Fprintln(w, string(b))
}
func writeResponse(w http.ResponseWriter, r *http.Request) {
	// レスポンスヘッダの設定
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	// ステータスコードの設定
	// この後でレスポンスヘッダの設定はできない
	w.WriteHeader(http.StatusBadRequest)

	// レスポンスボディの書き込み
	w.Write([]byte("Bad request!\n"))
	// http.Requestはio.Writerとして扱えるため、これでもOK
	// fmt.Fprintln(w, "Bad request!")
}
*/
