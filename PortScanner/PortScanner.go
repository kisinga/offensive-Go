package main

import (
	"fmt"
	"net"
	"sort"
	"sync"
)

func main() {
	// usingwg()
	// usingWorkerPool
	betterWorkerPool()
}

func betterWorkerPool() {
	ports := make(chan int, 100)
	results := make(chan int)
	openports := []int{}
	for i := 0; i < cap(ports); i++ {
		go betterWorker(ports, results)
	}
	go func() {
		for i := 1; i < 1024; i++ {
			ports <- i
		}
	}()
	for i := 1; i < 1024; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}
	close(ports)
	close(results)
	sort.Ints(openports)
	for _, val := range openports {
		fmt.Printf("%d open \n", val)
	}
}
func betterWorker(ports, results chan int) {
	address := ""
	for p := range ports {
		address = fmt.Sprintf("scanme.nmap.org:%d", p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			// port closed
			results <- 0
			return
		}
		conn.Close()
		results <- p
	}
}

func usingWorkerPool() {
	pool := make(chan int, 100)
	var wg sync.WaitGroup
	for i := 0; i < cap(pool); i++ {
		go worker(pool, &wg)
	}
	for i := 1; i <= 65535; i++ {
		wg.Add(1)
		pool <- i
	}

	wg.Wait()
	close(pool)
}

func worker(ports chan int, wg *sync.WaitGroup) {
	address := ""
	// avoid reallocating the var every time, define it outside the loop. minor optimization tho
	for p := range ports {
		address = fmt.Sprintf("scanme.nmap.org:%d", p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			// port closed
			// fmt.Printf("%d CLOSED \n", p)
			return
		}
		conn.Close()
		fmt.Printf("%d OPEN \n", p)
		wg.Done()
	}
}

func usingwg() {
	address := ""
	var wg sync.WaitGroup
	// avoid reallocating the var every time, define it outside the loop. minor optimization tho
	for i := 1; i <= 65535; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			address = fmt.Sprintf("scanme.nmap.org:%d", j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				// port closed
				return
			}
			conn.Close()
			fmt.Printf("%d OPEN \n", j)
		}(i)
	}
	wg.Wait()
}
