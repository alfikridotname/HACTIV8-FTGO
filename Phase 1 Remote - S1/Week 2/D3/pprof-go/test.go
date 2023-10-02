package main

//import (
//	"fmt"
//	"math/rand"
//	"net/http"
//	"os"
//	"runtime/pprof"
//	"time"
//)
//
//func doSomething() {
//	for i := 0; i < 100000; i++ {
//		rand.Intn(1000)
//	}
//}
//func main() {
//	f, err := os.Create("cpu.prof")
//	if err != nil {
//		panic(err)
//	}
//	defer f.Close()
//	pprof.StartCPUProfile(f)
//	defer pprof.StopCPUProfile()
//	go func() {
//		fmt.Println(http.ListenAndServe("localhost:8080", nil))
//	}()
//	for i := 0; i < 1000; i++ {
//		go doSomething()
//	}
//	time.Sleep(200 * time.Second)
//
//}
