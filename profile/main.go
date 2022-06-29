package main

import (
    _ "net/http/pprof"
    "log"
	"net/http"
    "runtime"
    "fmt"
)

func main(){
    maxCPUNum := runtime.NumCPU()
    fmt.Println("maxCPUNum: ",maxCPUNum)
    runtime.GOMAXPROCS(maxCPUNum)
    log.Fatal(http.ListenAndServe(":9999", nil))
}