package main

import (
	"fmt"
	//"math/rand"
	"time"

	//"fmt"
	//"io/ioutil"
	"io/ioutil"
	"net/http"
)



func main()  {

	httpGet := func () {

		resp, err := http.Get("http://127.0.0.1:8080")
		if err != nil {
			// handle error
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		//ioutil.ReadAll(resp.Body)
		if err != nil {
			// handle error
		}
		fmt.Println(string(body))
	}

	//var wg sync.WaitGroup
	var num = 30000

	i := 1
	for i<=num  {
		println(i)
		go httpGet()
		i++
		if i%100==0 {
			time.Sleep(2*time.Second)
			println("sleep")
		}
		//fmt.Println("hello world-"+i++)
	}
	//wg.Wait()
	//var sleep int = 60
	time.Sleep(6000*1000*time.Second)
	fmt.Println("the end")
	//time.sleep(sleep * time.Second)
}
