package main

import (
	"flag"
	"fmt"
	"net"
	"time"
)

func main() {

	ip := flag.String("ip", "127.0.0.1", "Target IP adress")
	start := flag.Int("start", 1, "Target Port")
	end := flag.Int("end", 1024, "Target Port")
	flag.Parse()

	for i := *start; i <= *end; i++ {
		address := fmt.Sprintf("%s:%d", *ip, i)
		conn, err := net.DialTimeout("tcp", address, 1*time.Second)
		if err == nil {
			fmt.Printf("Port: %d is open.\n", i)
			conn.Close()
		} else {
			fmt.Printf("Port: %d is closed.\n", i)
		}
	}

}
