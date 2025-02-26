package main

import (
	"fmt"
	"github.com/jbirdvegas/tilltrump.com/internal/timex"
	"time"
)

const duration = time.Hour*5 + time.Minute*10 + time.Second*2

func main() {
	printer()
	time.Sleep(time.Second)
	println("")
	printer()
}

func printer() {
	fmt.Printf("time since: %s\n", timex.SinceLastPresident().String())
	fmt.Printf("time till: %s\n", timex.TillNextPresident().String())
}
