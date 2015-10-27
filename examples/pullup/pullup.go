package main

import (
	"fmt"
	"github.com/gonutz/gopio"
	"os"
)

var (
	// Use mcu pin 22, corresponds to GPIO3 on the pi
	pin = gopio.Pin(22)
)

func main() {
	// Open and map memory to access gpio, check for errors
	if err := gopio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Unmap gpio memory when done
	defer gopio.Close()

	// Pull up and read value
	pin.PullUp()
	fmt.Printf("PullUp: %d\n", pin.Read())

	// Pull down and read value
	pin.PullDown()
	fmt.Printf("PullDown: %d\n", pin.Read())

}
