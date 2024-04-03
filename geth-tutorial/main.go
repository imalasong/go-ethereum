package main

import (
	"fmt"
	"time"
)

func main() {
	offlineLogged := time.Now()
	fmt.Println(time.Since(offlineLogged).Seconds())
}
