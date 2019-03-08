package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"runtime"
	"time"
)

type Result string
type Search func(query string) Result

func EmulateSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100))*time.Millisecond) // Each search may take upto 100ms
		return Result(fmt.Sprintf("%s result for %q", kind, query))
	}
}

func main() {
	TimeCalc(GoogleV1, "Golang")
	TimeCalc(GoogleV2, "Golang")
	TimeCalc(GoogleV2_1, "Golang")
	TimeCalc(GoogleV3, "Golang")
}

// These functions are written in chronological order

func TimeCalc(F func(query string) (results []Result), query string) {
	fmt.Println(runtime.FuncForPC(reflect.ValueOf(F).Pointer()).Name(),":")
	fmt.Println("--- x ---")
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	results := F(query)
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)
	fmt.Println("--- x ---")
}

var (
	web = EmulateSearch("Web")
	img = EmulateSearch("Image")
	vid = EmulateSearch("Video")
)

// Let's make a Google Search function

func GoogleV1(query string) (results []Result){
	results = append(results, web(query)) // Run
	results = append(results, img(query)) // Wait for previous results, then run
	results = append(results, vid(query)) // Wait for previous results, then run
	return
}

func GoogleV2(query string) (results []Result){
	c := make(chan Result)
	// Here comes the FanIn again!!
	go func() { c<-web(query) }()
	go func() { c<-img(query) }()
	go func() { c<-vid(query) }()
	// Since, each query will only take place once, we know that we need 3 instances of those searches
	for i:=0;i<3;i++ {
		results = append(results, <-c)
	}

	return
}

func GoogleV2_1(query string) (results []Result){
	c := make(chan Result)
	// Here comes the FanIn again!!
	go func() { c<-web(query) }()
	go func() { c<-img(query) }()
	go func() { c<-vid(query) }()
	// Since, each query will only take place once, we know that we need 3 instances of those searches

	// Let's add a timeout to it
	timeout := time.After(80*time.Millisecond)
	for i:=0;i<3;i++ {
		select {
		case r:=<-c:
			results = append(results, r)
		case <-timeout:
			fmt.Println("Connection timed out for", <-c)
			return
		}
	}
	return
}

// Avoid Timeout
// Make the same search to multiple servers, pick the first one ie. with the lowest ms response

func First(query string, replicas... Search) Result {
	c := make(chan Result)
	searchReplica := func(i int) { //since, Search = func(query string) Result
		c <- replicas[i](query)
	}
	for i:=range replicas {
		go searchReplica(i)
	}
	return <-c
}

var (
	// Alternate servers
	web2 = EmulateSearch("Web2")
	img2 = EmulateSearch("Image2")
	vid2= EmulateSearch("Video2")
	web3 = EmulateSearch("Web3")
	img3 = EmulateSearch("Image3")
	vid3= EmulateSearch("Video3")
)

func GoogleV3(query string) (results []Result){
	c := make(chan Result)
	// Here comes the FanIn again!!
	go func() { c<-First(query, web, web2, web3)}()
	go func() { c<-First(query, img, img2, img3) }()
	go func() { c<-First(query, vid, vid2, vid3) }()
	// Since, each query will only take place once, we know that we need 3 instances of those searches

	// Let's add a timeout to it
	timeout := time.After(80*time.Millisecond)
	for i:=0;i<3;i++ {
		select {
		case r:=<-c:
			results = append(results, r)
		case <-timeout:
			fmt.Println("Connection timed out!", <-c)
			return
		}
	}
	return
}