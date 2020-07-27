/**
* @Author:zhoutao
* @Date:2020/7/27 下午2:32
 */

package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

var datas []string

func main() {
	go func() {
		log.Printf("len:%d", add("go-tour-book"))
		time.Sleep(time.Millisecond * 10)
	}()

	_ = http.ListenAndServe(":6060", nil)
}

func add(str string) int {
	data := []byte(str)
	datas = append(datas, string(data))
	return len(datas)

}
