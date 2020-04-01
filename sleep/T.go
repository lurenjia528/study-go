package main

import (
	"fmt"
	"github.com/sony/sonyflake"
	"net/http"
	"os"
	"time"
	_ "expvar"
)

var ready = true

func main() {

	getId()

	http.Handle("/", timeMiddleware(http.HandlerFunc(Hello)))
	http.Handle("/health", timeMiddleware(http.HandlerFunc(Health)))
	http.ListenAndServe(":8080", nil)

}

func Hello(resp http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	if len(req.Form["delay"]) > 0 {
		delay := req.Form["delay"][0]
		duration, _ := time.ParseDuration(delay + "s")
		time.Sleep(duration)
		fmt.Fprintln(resp, "ok")
	}
}

func Health(w http.ResponseWriter, req *http.Request) {
	if ready {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("200"))
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500"))
	}
}

func timeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		timeStart := time.Now()

		// next handler
		next.ServeHTTP(wr, r)

		timeElapsed := time.Since(timeStart)
		fmt.Println(timeElapsed)
	})
}

func getId() {
	t, _ := time.Parse("2006-01-02", "2018-01-01")
	settings := sonyflake.Settings{
		StartTime: t,
	}

	sf := sonyflake.NewSonyflake(settings)
	id, err := sf.NextID()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(id)
}
