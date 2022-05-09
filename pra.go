package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", defaultRoute)
	mux.HandleFunc("/route1", route1)
	mux.HandleFunc("/route2", route2)
	mux.Handle("/anothermethod", handlerA{})
	mux.HandleFunc("/lookRequest", lookRequest)
	mux.HandleFunc("/writeResponse", writeResponse)
	server := http.Server{
		Addr:           ":8080",
		Handler:        mux,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}

func defaultRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "It's defaultroute")
}
func route1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "It's route1")
}
func route2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "It's route2")
}

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
