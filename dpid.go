package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Usage: " + `
	dpid [your mac address] [your next mac address] [...]`)
		os.Exit(0)
	}
	for l := 1; l < len(os.Args); l++ {
		mac := os.Args[l]
		hw, err := net.ParseMAC(mac)
		if err != nil {
			fmt.Println(mac, " : ", err)
			continue
		}
		macArray := strings.Split(hw.String(), ":")
		for i := range macArray {
			val := macArray[i]
			n, err := strconv.ParseUint(val, 16, 32)
			if err != nil {
				log.Println(err.Error())
				break
			}
			if n != 0 && i != 0 {
				n2 := uint32(n)
				fmt.Print(n2)
			}

		}
		fmt.Print("\n")
	}
}
