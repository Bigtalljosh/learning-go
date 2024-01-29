package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {

	http.Handle("/", myHandler("Customer service"))

	s := http.Server{
		Addr: ":3000",
	}

	go func() {
		log.Fatal(s.ListenAndServeTLS("./cert.pem", "./key.pem"))
	}()

	fmt.Println("Server started, press <Enter> to shutdown")
	fmt.Scanln()
	s.Shutdown(context.Background())
	fmt.Println("Server stopped")
}


type myHandler string

func (mh myHandler) ServeHTTP (w http.ResponseWriter, r *http.Request){
	w.Header().Add("X-Powered-By","Mememe")
	http.SetCookie(w, &http.Cookie{
		Name: "session-id",
		Value: "12345",
		Expires: time.Now().Add(24 * time.Hour * 365),
	})

	w.WriteHeader(http.StatusAccepted)

	fmt.Fprintln(w, r.Cookies())
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, string(mh))
	fmt.Fprintln(w, r.Header)
}