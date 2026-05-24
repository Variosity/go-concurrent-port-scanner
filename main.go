package main

import (
	"flag"
	"fmt"
	"net"
	"time"
	"sync"
)

func main() {

	ip := flag.String("ip", "127.0.0.1", "Target IP adress")
	start := flag.Int("start", 1, "Target Port")
	end := flag.Int("end", 1024, "Target Port")
	flag.Parse()
	var wg sync.WaitGroup
	ch := make(chan int, 65535)

	for i := *start; i <= *end; i++ {
		wg.Add(1)
		go func(port int){
		defer wg.Done()
		address := fmt.Sprintf("%s:%d", *ip, port)
		conn, err := net.DialTimeout("tcp", address, 1*time.Second)
		if err == nil {
			conn.Close()
			ch <- port
		}
	}(i)
}
	wg.Wait()
	close(ch)
	for port := range ch {
		fmt.Printf("Port %d is open.\n", port)
		}
}
