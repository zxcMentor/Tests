package main

import (
	"log"
	"net/http"
	"net/http/pprof"
	_ "net/http/pprof"
	"sync"
)

func main() {
	// Запуск HTTP-сервера для сбора профилировочных данных
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		mux := http.NewServeMux()
		mux.HandleFunc("/mycustompath/pprof/", pprof.Index)
		mux.HandleFunc("/mycustompath/pprof/cmdline", pprof.Cmdline)
		mux.HandleFunc("/mycustompath/pprof/profile", pprof.Profile)
		mux.HandleFunc("/mycustompath/pprof/symbol", pprof.Symbol)
		mux.HandleFunc("/mycustompath/pprof/trace", pprof.Trace)
		log.Println(http.ListenAndServe("localhost:6060", mux))
	}()
	wg.Wait()
	// Остальной код приложения
	// ...
}
