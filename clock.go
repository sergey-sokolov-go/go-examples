package main

import (
	"fmt"
	"os"

	"github.com/beevik/ntp"
)

func main() {
	if time, error := ntp.Time("0.beevik-ntp.pool.ntp.org1"); error != nil {
		fmt.Fprintln(os.Stderr, error)
	} else {
		fmt.Println("Time = ", time)
	}

}
