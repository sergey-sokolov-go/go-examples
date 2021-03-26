package main

import (
	"fmt"
	"log"

	"github.com/beevik/ntp"
)

func main() {
	if time, error := ntp.Time("0.beevik-ntp.pool.ntp.org"); error != nil {
		log.Fatalf("fatal error: %s", error)
	} else {
		fmt.Println("Time = ", time)
	}
}
