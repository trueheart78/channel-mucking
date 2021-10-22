package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Using Channels Like Sleep()")
	dt := time.Now()
	fmt.Println("=> Current time: ", dt.String())

	waitUntil := dt.Add(time.Second * 5)

	fmt.Println("=> Waiting 5 seconds")
	// when we want to wait till
	// until, _ := time.Parse(time.RFC3339, "2021-10-22T15:12:53-04:00")

	// and now we wait
	<-time.After(time.Until(waitUntil))
	fmt.Println("=> That took forever!")
}
